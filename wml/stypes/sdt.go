package stypes

// SdtLock represents content control lock settings
type SdtLock string

const (
	SdtLockUnlocked        SdtLock = "unlocked"        // Content control can be deleted and edited
	SdtLockSdtLocked       SdtLock = "sdtLocked"       // Content control cannot be deleted but can be edited
	SdtLockContentLocked   SdtLock = "contentLocked"   // Content control can be deleted but content cannot be edited
	SdtLockSdtContentLocked SdtLock = "sdtContentLocked" // Content control cannot be deleted or edited
)

// CalendarType represents calendar types for date content controls
type CalendarType string

const (
	CalendarTypeGregorian                CalendarType = "gregorian"
	CalendarTypeGregorianUs              CalendarType = "gregorianUs"
	CalendarTypeJapanese                 CalendarType = "japanese"
	CalendarTypeTaiwan                   CalendarType = "taiwan"
	CalendarTypeKorea                    CalendarType = "korea"
	CalendarTypeHijri                    CalendarType = "hijri"
	CalendarTypeThai                     CalendarType = "thai"
	CalendarTypeHebrew                   CalendarType = "hebrew"
	CalendarTypeGregorianMeFrench        CalendarType = "gregorianMeFrench"
	CalendarTypeGregorianArabic          CalendarType = "gregorianArabic"
	CalendarTypeGregorianXlitEnglish     CalendarType = "gregorianXlitEnglish"
	CalendarTypeGregorianXlitFrench      CalendarType = "gregorianXlitFrench"
	CalendarTypeNone                     CalendarType = "none"
)

// HexChar represents a hexadecimal character value
type HexChar string