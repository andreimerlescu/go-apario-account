package wallet

import (
	`fmt`
	`strings`
)

// String lists the translations for MultiplierType that are stored inside of Multipliers map
func (mt MultiplierType) String() string {
	s := strings.Builder{}
	sm, ok := Multipliers[mt]
	if !ok {
		return ""
	}
	s.Write([]byte(fmt.Sprintf("MultiplierType [%d]:\n", len(sm))))
	for lang, msg := range sm {
		s.Write([]byte(fmt.Sprintf("- [%v] %v\n", lang, msg)))
	}
	s.Write([]byte("\n"))
	return s.String()
}

var Multipliers = map[MultiplierType]MultiplierDialect{
	Multiplier00: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("Canceled"),
	},
	Multiplier01: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("No Change"),
	},
	Multiplier02: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("1x Boost"),
	},
	Multiplier03: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("2x Boost"),
	},
	Multiplier04: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("3x Boost"),
	},
	Multiplier05: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("4x Boost"),
	},
	Multiplier06: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("5x Boost"),
	},
	Multiplier07: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("6x Boost"),
	},
	Multiplier08: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("7x Boost"),
	},
	Multiplier09: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("8x Boost"),
	},
	Multiplier10: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("9x Boost"),
	},
	Multiplier11: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("10x Boost"),
	},
	Multiplier12: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("11x Boost"),
	},
	Multiplier13: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("MAJESTIC BONUS!"),
	},
	Multiplier14: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("13x Boost"),
	},
	Multiplier15: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("14x Boost"),
	},
	Multiplier16: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("15x Boost"),
	},
	Multiplier17: {
		ISO6393CodeType([]rune("eng")): MultiplierMessage("Q Booster Administered"),
	},
}
