package wallet

import (
	`fmt`
)

// String pretty prints the Reputation to show earned points old and new point values
func (r *Reputation) String() string {
	return fmt.Sprintf("*Reputation [%T] Earned Points = %d ; Old = %d ; New = %d",
		r, r.EarnedPoints(), r.LastPointsCount, r.Points())
}

// Points returns LastPointsCount + EarnedPoints()
func (r *Reputation) Points() int64 {
	return r.LastPointsCount + r.EarnedPoints()
}

// EarnedPoints calculates the points based on the Reason, Multiplier and Action stored in the Reputation
func (r *Reputation) EarnedPoints() int64 {
	return int64(r.Reason) * int64(r.Multiplier) * int64(r.Action)
}
