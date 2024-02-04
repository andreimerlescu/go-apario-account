package wallet

type MessageType string
type ActionsDialect map[ISO6393CodeType]MessageType

type Block struct {
	AccountUUID    string
	BlockSize      int64
	TotalPoints    int64
	TotalCount     int64
	PositivePoints int64
	PositiveCount  int64
	NegativePoints int64
	NegativeCount  int64
	Checksum       [64]byte
}

type MultiplierMessage string
type MultiplierDialect map[ISO6393CodeType]MultiplierMessage

type ISO6393CodeType [3]rune
type ReasonMessage string
type ReasonDialect map[ISO6393CodeType]ReasonMessage

type ActionType float32
type MultiplierType float32
type ReasonType float32

type Reputation struct {
	AccountUUID     string
	LastPointsCount int64
	Action          ActionType
	Reason          ReasonType
	Multiplier      MultiplierType
}
