package wallet

const (
	ActionPositive ActionType = 1
	ActionNegative ActionType = -1
)

const (
	Multiplier00 MultiplierType = iota
	Multiplier01
	Multiplier02
	Multiplier03
	Multiplier04
	Multiplier05
	Multiplier06
	Multiplier07
	Multiplier08
	Multiplier09
	Multiplier10
	Multiplier11
	Multiplier12
	Multiplier13
	Multiplier14
	Multiplier15
	Multiplier16
	Multiplier17
)

const (
	ReasonCode001 ReasonType = 0.003 * iota
	ReasonCode002            // 0.003
	ReasonCode003            // 0.006
	ReasonCode004            // special rounding error
	ReasonCode005
	ReasonCode006
	ReasonCode007 // special rounding error
	ReasonCode008
	ReasonCode009
	ReasonCode010
	ReasonCode011 // 11 = 0.03
	ReasonCode012 // 12 == 0.033
	ReasonCode013 // special rounding error
	ReasonCode014
	ReasonCode015
	ReasonCode016
	ReasonCode017
	ReasonCode018 // special rounding error
	ReasonCode019
	ReasonCode020
	ReasonCode021
	ReasonCode022
	ReasonCode023
	ReasonCode024
	ReasonCode025 // special rounding error
	ReasonCode026
	ReasonCode027
	ReasonCode028
	ReasonCode029
	ReasonCode030 // special rounding error
	ReasonCode031
	ReasonCode032
	ReasonCode033
	ReasonCode034
	ReasonCode035 // special rounding error
	ReasonCode036
	ReasonCode037
	ReasonCode038
	ReasonCode039
	ReasonCode040 // 40 = 0.117
	ReasonCode041
	ReasonCode042
	ReasonCode043
	ReasonCode044
	ReasonCode045
	ReasonCode046
	ReasonCode047
	ReasonCode048 // special rounding error
	ReasonCode049 // special rounding error
	ReasonCode050
	ReasonCode051
	ReasonCode052
	ReasonCode053
	ReasonCode054
	ReasonCode055
	ReasonCode056
	ReasonCode057
	ReasonCode058
	ReasonCode059 // special rounding error
	ReasonCode060
	ReasonCode061
	ReasonCode062
	ReasonCode063
	ReasonCode064
	ReasonCode065
	ReasonCode066
	ReasonCode067
	ReasonCode068
	ReasonCode069 // special rounding error
	ReasonCode070 // special rounding error = 70 = .207 w/ a rounding 2 == majestic =D 72 27 = 9 9 = 18 = 9
	ReasonCode071
	ReasonCode072
	ReasonCode073
	ReasonCode074
	ReasonCode075
	ReasonCode076
	ReasonCode077
	ReasonCode078
	ReasonCode079
	ReasonCode080 // special rounding error
	ReasonCode081
	ReasonCode082
	ReasonCode083
	ReasonCode084
	ReasonCode085
	ReasonCode086
	ReasonCode087
	ReasonCode088
	ReasonCode089
	ReasonCode090
	ReasonCode091
	ReasonCode092
	ReasonCode093
	ReasonCode094
	ReasonCode095 // special rounding error
	ReasonCode096 // EXTREMELY special rounding error = .285 with a rounding 3 on 96 = 963 @ .285
	ReasonCode097 // special rounding error
	ReasonCode098
	ReasonCode099
	ReasonCode100
)
