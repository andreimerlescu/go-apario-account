package wallet

import (
	`fmt`
	`strings`
)

// String lists the translations for ActionType that are stored inside of Actions map
func (at ActionType) String() string {
	s := strings.Builder{}
	sm, ok := Actions[at]
	if !ok {
		return ""
	}
	s.Write([]byte(fmt.Sprintf("Actions [%d]:\n", len(sm))))
	for lang, msg := range sm {
		s.Write([]byte(fmt.Sprintf("- [%v] %v\n", lang, msg)))
	}
	s.Write([]byte("\n"))
	return s.String()
}

var Actions = map[ActionType]ActionsDialect{
	ActionPositive: map[ISO6393CodeType]MessageType{
		ISO6393CodeType([]rune("eng")): MessageType("Constructive"),
	},
	ActionNegative: map[ISO6393CodeType]MessageType{
		ISO6393CodeType([]rune("eng")): MessageType("Destructive"),
	},
}
