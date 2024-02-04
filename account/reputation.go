package account

import (
	`crypto/sha256`
	`errors`
	`fmt`
	`log`
	`sync`
	`time`

	ai `github.com/andreimerlescu/go-apario-identifier`

	`github.com/andreimerlescu/go-apario-account/wallet`
)

type Reputation struct {
	Account                      *Account
	Blocks                       []wallet.Block
	BlockSize                    int64
	Pending                      []wallet.Reputation
	BlocksTotalPoints            int64
	BlocksTotalPointsVerifiedAt  time.Time
	PendingTotalPoints           int64
	PendingTotalPointsVerifiedAt time.Time
	PositiveActionCount          int64
	NegativeActionCount          int64
	LastBlockChecksum            [64]byte
	LastBlockTotalPoints         int64
	LastBlockTotalCount          int64
	LastBlockPositivePoints      int64
	LastBlockPositiveActionCount int64
	LastBlockNegativePoints      int64
	LastBlockNegativeActionCount int64
	Checksum                     [64]byte
	Version                      ai.Version
	mu                           *sync.RWMutex
	e                            error
	eat                          time.Time
}

// Copy function duplicates the data inside the Reputation struct into a new Reputation struct and logs the err
func (r *Reputation) Copy() Reputation {
	if r.e != nil {
		log.Printf("*Reputation.Copy() omitted from Copy() an error %v which was triggered on %v", r.e, r.eat.Format(time.RFC3339))
	}
	return Reputation{
		Account:                      r.Account,
		Blocks:                       r.Blocks,
		BlockSize:                    r.BlockSize,
		Pending:                      r.Pending,
		BlocksTotalPoints:            r.BlocksTotalPoints,
		BlocksTotalPointsVerifiedAt:  r.BlocksTotalPointsVerifiedAt,
		PendingTotalPoints:           r.PendingTotalPoints,
		PendingTotalPointsVerifiedAt: r.PendingTotalPointsVerifiedAt,
		PositiveActionCount:          r.PositiveActionCount,
		NegativeActionCount:          r.NegativeActionCount,
		LastBlockTotalPoints:         r.LastBlockTotalPoints,
		LastBlockTotalCount:          r.LastBlockTotalCount,
		LastBlockPositiveActionCount: r.LastBlockPositiveActionCount,
		LastBlockPositivePoints:      r.LastBlockPositivePoints,
		LastBlockNegativeActionCount: r.LastBlockNegativeActionCount,
		LastBlockNegativePoints:      r.LastBlockNegativePoints,
		LastBlockChecksum:            r.LastBlockChecksum,
		Checksum:                     r.Checksum,
		Version:                      r.Version,
		mu:                           &sync.RWMutex{},
	}
}

// ActionCount combines Positive and Negative Action Counts
func (r *Reputation) ActionCount() int64 {
	return r.PositiveActionCount + r.NegativeActionCount
}

// Watchable returns if Positive < Negative in Action Counts
func (r *Reputation) Watchable() bool {
	return r.PositiveActionCount < r.NegativeActionCount
}

// PositiveScaleWatchable multiplies the PositiveActionCount by the scale and ensure its greater than PositiveActionCount
func (r *Reputation) PositiveScaleWatchable(scale int64) bool {
	return scale*r.PositiveActionCount < r.NegativeActionCount
}

// NegativeScaleWatchable multiplies the NegativeActionCount by the scale and ensures its greater than PositiveActionCount
func (r *Reputation) NegativeScaleWatchable(scale int64) bool {
	return r.PositiveActionCount < scale*r.NegativeActionCount
}

// String cleanly prints the Reputation data that can be used for diagnostics and checksum hashing
func (r *Reputation) String() string {
	return fmt.Sprintf("Reputation [%T] Account UUID %v ; Positive = %d ; Negative = %d ; Block Points = %d ; Total Pending = %d ; LastBlock = %d ; LastBlockChecksum = %v ; ",
		r, r.Account.Identifier.UUID(), r.PositiveActionCount, r.NegativeActionCount, r.BlocksTotalPoints, len(r.Pending), len(r.Blocks), r.LastBlockChecksum)
}

// NegativeStreak returns true when the Pending []wallet.Reputation contains a sequence (newest first) of Action==ActionNegative & > 17 streaks hit
func (r *Reputation) NegativeStreak() bool {
	neg, pos := 0, 0
	streak := 0
	lrp := len(r.Pending)
	for i := lrp; i >= 0; i-- { // newest first
		if i < 0 || i > lrp {
			break
		}
		reputation := r.Pending[i]
		if reputation.Action == wallet.ActionPositive {
			pos += 1
			streak += 1
		} else if reputation.Action == wallet.ActionNegative {
			neg += 1
			if neg >= 17 {
				return true
			}
			// good behavior fast lane
			if streak >= 17 {
				return false
			}
			streak = 0
		} else {
			if streak >= 17 {
				return false
			}
			continue
		}
	}
	return false
}

// PositiveStreak returns true when the Pending []wallet.Reputation contains a sequence (newest first) of Action==ActionPositive & > 17 streaks hit
func (r *Reputation) PositiveStreak() bool {
	neg, pos := 0, 0
	streak := 0
	lrp := len(r.Pending)
	for i := lrp; i >= 0; i-- { // newest first
		if i < 0 || i > lrp {
			break
		}
		reputation := r.Pending[i]
		if reputation.Action == wallet.ActionPositive {
			pos += 1
			streak += 1
		} else if reputation.Action == wallet.ActionNegative {
			neg += 1
			if neg >= 17 {
				return false
			}
			// good behavior fast lane
			if streak >= 17 {
				return true
			}
			streak = 0
		} else {
			if streak >= 17 {
				return true
			}
			continue
		}
	}
	return false
}

// Multiplier enforces no points for bad behavior
func (r *Reputation) Multiplier() int64 {
	if r.Watchable() {
		if !r.NegativeStreak() && r.PositiveStreak() {
			return 1 // when you've been bad and are now good, you resume earning points
		}
		return 0 // default == no points rewarded and positive action granted
	}
	return 1 // default when not watchable means no zeroing out of points earned on new Reputation issues
}

// Hash returns the sha256sum of the String() output of Reputation
func (r *Reputation) Hash() (ok bool) {
	defer func() {
		r := recover()
		if r != nil {
			ok = false
		}
	}()
	ok = true
	cs := sha256.New().Sum([]byte(r.String()))
	r.Checksum = [64]byte{}
	for i := 0; i < len(cs); i++ {
		r.Checksum[i] = cs[i]
	}
	return
}

// CheckPendingBlockSize wraps up Pending in []wallet.Reputation into a new wallet.Block when at BlockSize in length
func (r *Reputation) CheckPendingBlockSize() *Reputation {
	if r.mu == nil {
		r.mu = &sync.RWMutex{}
	}
	if len(r.Pending) != int(r.BlockSize) {
		return r
	}
	if errors.Is(r.e, ErrBadChecksum) {
		log.Printf("[*Reputation] error stuck in CheckingPendingBlockSize() triggering ErrBadChecksum: %v", r.e)
		return nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	block := wallet.Block{
		AccountUUID:    r.Account.Identifier.UUID(),
		BlockSize:      r.BlockSize,
		TotalPoints:    0,
		TotalCount:     0,
		PositivePoints: 0,
		PositiveCount:  0,
		NegativePoints: 0,
		NegativeCount:  0,
		Checksum:       [64]byte{},
	}
	nPos, nNeg := 0, 0
	for _, pending := range r.Pending {
		block.TotalPoints += pending.EarnedPoints()
		block.TotalCount += 1
		if pending.Action == wallet.ActionPositive {
			nPos += 1
		}
		if pending.Action == wallet.ActionNegative {
			nNeg += 1
		}
	}
	if !block.Hash() {
		r.e = ErrBadChecksum // technically should never happen
		r.eat = time.Now().UTC()
		return nil
	}
	r.Blocks = append(r.Blocks, block)
	r.BlocksTotalPoints += block.TotalPoints
	r.BlocksTotalPointsVerifiedAt = time.Now().UTC()
	r.PendingTotalPoints = 0
	r.PendingTotalPointsVerifiedAt = time.Now().UTC()
	r.PositiveActionCount += int64(nPos)
	r.NegativeActionCount += int64(nNeg)
	r.LastBlockChecksum = block.Checksum
	r.LastBlockNegativePoints = block.NegativePoints
	r.LastBlockNegativeActionCount = block.NegativeCount
	r.LastBlockPositivePoints = block.PositivePoints
	r.LastBlockPositiveActionCount = block.PositiveCount
	r.LastBlockTotalPoints = block.TotalPoints
	r.LastBlockTotalCount = block.TotalCount
	r.Version.BumpMajor()
	r.Account.Reputation = *r
	return r
}

// CountBlocks uses the BlocksTotalPoints and adds each Pending []wallet.Reputation to the result
func (r *Reputation) CountBlocks() int64 {
	if r.mu == nil {
		log.Printf("[NOTICE] Reputation.CountBlocks r.mu == nil invoked r.mu = &sync.RWMutex{}")
		r.mu = &sync.RWMutex{}
	}
	r.CheckPendingBlockSize()
	r.mu.Lock()
	defer r.mu.Unlock()
	var total int64 = 0
	var counter int64 = 0
	for _, b := range r.Blocks {
		total += b.TotalPoints
		counter += 1
	}
	if counter != int64(len(r.Blocks)) {
		log.Printf("[NOTICE] Reputation.CounterBlocks counter != int64(len(r.Blocks))")
	}
	r.BlocksTotalPoints = total
	r.BlocksTotalPointsVerifiedAt = time.Now().UTC()
	return r.BlocksTotalPoints
}

// Points returns the BlockTotalPoints + Pending EarnedPoints() or initializes the BlocksTotalPoints if 0 with wallet.Block present in Blocks
func (r *Reputation) Points() int64 {
	if r.BlocksTotalPoints == 0 && len(r.Blocks) != 0 {
		return r.CountBlocks()
	}
	var pending int64 = 0
	for _, reputation := range r.Pending {
		pending += reputation.EarnedPoints()
	}
	return r.BlocksTotalPoints + pending
}
