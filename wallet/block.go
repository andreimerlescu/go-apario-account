package wallet

import (
	`crypto/sha256`
	`fmt`
)

// String pretty prints the Block in a manner that Hash can use
func (b *Block) String() string {
	return fmt.Sprintf("*Block [%T] UUID = %v ; BlockSize = %d ; TotalPoints = %d ; TotalCount = %d ; PositivePoints = %d ; PositiveCount = %d ; NegativePoints = %d ; NegativeCount = %d",
		b, b.AccountUUID, b.BlockSize, b.TotalPoints, b.TotalCount, b.PositivePoints, b.PositiveCount, b.NegativePoints, b.NegativeCount)
}

// Hash assigns the sha256sum of String() to the Checksum and returns true if it succeeded
func (b *Block) Hash() (ok bool) {
	defer func() {
		r := recover()
		if r != nil {
			ok = false
		}
	}()
	ok = true
	cs := sha256.New().Sum([]byte(b.String()))
	b.Checksum = [64]byte{}
	for i := 0; i < len(cs); i++ {
		b.Checksum[i] = cs[i]
	}
	return
}
