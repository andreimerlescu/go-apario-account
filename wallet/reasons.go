package wallet

import (
	`fmt`
	`strings`
)

// String lists the translations for ReasonType that are stored inside of Reasons map
func (rt ReasonType) String() string {
	s := strings.Builder{}
	sm, ok := Reasons[rt]
	if !ok {
		return ""
	}
	s.Write([]byte(fmt.Sprintf("ReasonType [%d]:\n", len(sm))))
	for lang, msg := range sm {
		s.Write([]byte(fmt.Sprintf("- [%v] %v\n", lang, msg)))
	}
	s.Write([]byte("\n"))
	return s.String()
}

var Reasons = map[ReasonType]ReasonDialect{
	ReasonCode001: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Tag Vote"),
	},
	ReasonCode002: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("StumbleInto Read"),
	},
	ReasonCode003: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Proposal Vote"),
	},
	ReasonCode004: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Page Vote"),
	},
	ReasonCode005: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Document Vote"),
	},
	ReasonCode006: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode006"),
	},
	ReasonCode007: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Snippet"),
	},
	ReasonCode008: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Tag"),
	},
	ReasonCode009: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Document Submitted"),
	},
	ReasonCode010: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Document Processing"),
	},
	ReasonCode011: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Document Processed"),
	},
	ReasonCode012: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Document Adoption"),
	},
	ReasonCode013: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Document Verification"),
	},
	ReasonCode014: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("New Document Published"),
	},
	ReasonCode015: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Small Edit OCR Transcription"),
	},
	ReasonCode016: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Large Edit OCR Transcription"),
	},
	ReasonCode017: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("OCR Transcription Replacement"),
	},
	ReasonCode018: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Small Edit Friend's Translation"),
	},
	ReasonCode019: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Large Edit Friend's Translation"),
	},
	ReasonCode020: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Replace Friend's Translation"),
	},
	ReasonCode021: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Translation Vote Casted"),
	},
	ReasonCode022: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Small Translation Proposed"),
	},
	ReasonCode023: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Large Translation Proposed"),
	},
	ReasonCode024: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Small Translation Adoption"),
	},
	ReasonCode025: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Large Translation Adoption"),
	},
	ReasonCode026: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Small Translation Published"),
	},
	ReasonCode027: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Large Translation Published"),
	},
	ReasonCode028: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Translation Vote Casted"),
	},
	ReasonCode029: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode029"),
	},
	ReasonCode030: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode030"),
	},
	ReasonCode031: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode031"),
	},
	ReasonCode032: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode032"),
	},
	ReasonCode033: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Small Translation Received Positive Vote"),
	},
	ReasonCode034: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode034"),
	},
	ReasonCode035: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode035"),
	},
	ReasonCode036: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Large Translation Received Positive Vote"),
	},
	ReasonCode037: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode037"),
	},
	ReasonCode038: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode038"),
	},
	ReasonCode039: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode039"),
	},
	ReasonCode040: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode040"),
	},
	ReasonCode041: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode041"),
	},
	ReasonCode042: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode042"),
	},
	ReasonCode043: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode043"),
	},
	ReasonCode044: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode044"),
	},
	ReasonCode045: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode045"),
	},
	ReasonCode046: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode046"),
	},
	ReasonCode047: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode047"),
	},
	ReasonCode048: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode048"),
	},
	ReasonCode049: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode049"),
	},
	ReasonCode050: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode050"),
	},
	ReasonCode051: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode051"),
	},
	ReasonCode052: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode052"),
	},
	ReasonCode053: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode053"),
	},
	ReasonCode054: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode054"),
	},
	ReasonCode055: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode055"),
	},
	ReasonCode056: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode056"),
	},
	ReasonCode057: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode057"),
	},
	ReasonCode058: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode058"),
	},
	ReasonCode059: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode059"),
	},
	ReasonCode060: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode060"),
	},
	ReasonCode061: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode061"),
	},
	ReasonCode062: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode062"),
	},
	ReasonCode063: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode063"),
	},
	ReasonCode064: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode064"),
	},
	ReasonCode065: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode065"),
	},
	ReasonCode066: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode066"),
	},
	ReasonCode067: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode067"),
	},
	ReasonCode068: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode068"),
	},
	ReasonCode069: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode069"),
	},
	ReasonCode070: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode070"),
	},
	ReasonCode071: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode071"),
	},
	ReasonCode072: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode072"),
	},
	ReasonCode073: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode073"),
	},
	ReasonCode074: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode074"),
	},
	ReasonCode075: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode075"),
	},
	ReasonCode076: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode076"),
	},
	ReasonCode077: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("ReasonCode077"),
	},
	ReasonCode078: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Inappropriate Metadata"),
	},
	ReasonCode079: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Being Military Controlled"),
	},
	ReasonCode080: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Being A Fed"),
	},
	ReasonCode081: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Being A Spook"),
	},
	ReasonCode082: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Bad Words"),
	},
	ReasonCode083: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Harassment"),
	},
	ReasonCode084: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("User Profile Reported For Spam"),
	},
	ReasonCode085: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Comment Censored By Bad Words Filter"),
	},
	ReasonCode086: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Snippet Omitted From Front End Interface"),
	},
	ReasonCode087: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Snippet Omitted From Search Results"),
	},
	ReasonCode088: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Snippet Omitted From StumbleInto"),
	},
	ReasonCode089: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Document Omitted From Front End Interface"),
	},
	ReasonCode090: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Document Omitted From Search Results"),
	},
	ReasonCode091: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Document Omitted From StumbleInto"),
	},
	ReasonCode092: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Page Omitted From Front End Interface"),
	},
	ReasonCode093: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Page Omitted From Search Results"),
	},
	ReasonCode094: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Page Omitted From StumbleInto"),
	},
	ReasonCode095: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Small Translation Received Negative Vote"),
	},
	ReasonCode096: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Your Large Translation Received Negative Vote"),
	},
	ReasonCode097: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Omit From Front End Interface"),
	},
	ReasonCode098: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Omit From Recommendations"),
	},
	ReasonCode099: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Omit From Search Results"),
	},
	ReasonCode100: map[ISO6393CodeType]ReasonMessage{
		ISO6393CodeType([]rune("eng")): ReasonMessage("Omit From StumbleInto"),
	},
}
