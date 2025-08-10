package stypes

// FldCharType represents field character types
type FldCharType string

const (
	FldCharTypeBegin    FldCharType = "begin"
	FldCharTypeEnd      FldCharType = "end"
	FldCharTypeSeparate FldCharType = "separate"
)

// TextFormFieldType represents text form field types
type TextFormFieldType string

const (
	TextFormFieldTypeRegular  TextFormFieldType = "regular"
	TextFormFieldTypeNumber   TextFormFieldType = "number"
	TextFormFieldTypeDate     TextFormFieldType = "date"
	TextFormFieldTypeCurrentTime TextFormFieldType = "currentTime"
	TextFormFieldTypeCurrentDate TextFormFieldType = "currentDate"
	TextFormFieldTypeCalculated  TextFormFieldType = "calculated"
)

// InfoTextType represents form field info text types
type InfoTextType string

const (
	InfoTextTypeText     InfoTextType = "text"
	InfoTextTypeAutoText InfoTextType = "autoText"
)

// DisplacedByCustomXml represents displacement by custom XML values
type DisplacedByCustomXml string

const (
	DisplacedByCustomXmlNext DisplacedByCustomXml = "next"
	DisplacedByCustomXmlPrev DisplacedByCustomXml = "prev"
)

// Space represents xml:space values
type Space string

const (
	SpaceDefault  Space = "default"
	SpacePreserve Space = "preserve"
)