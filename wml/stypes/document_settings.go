package stypes

// ViewType represents document view types
type ViewType string

const (
	ViewTypeNone       ViewType = "none"
	ViewTypePrint      ViewType = "print"
	ViewTypeOutline    ViewType = "outline"
	ViewTypeMasterPage ViewType = "masterPage"
	ViewTypeNormal     ViewType = "normal"
	ViewTypeWeb        ViewType = "web"
)

// ZoomType represents zoom types
type ZoomType string

const (
	ZoomTypeNone        ZoomType = "none"
	ZoomTypeFullPage    ZoomType = "fullPage"
	ZoomTypeBestFit     ZoomType = "bestFit"
	ZoomTypeTextFit     ZoomType = "textFit"
)

// ProofingState represents proofing states
type ProofingState string

const (
	ProofingStateClean ProofingState = "clean"
	ProofingStateDirty ProofingState = "dirty"
)

// FootnoteEndnoteType represents footnote/endnote positioning
type FootnoteEndnoteType string

const (
	FootnoteEndnoteTypeBeneathText FootnoteEndnoteType = "beneathText"
	FootnoteEndnoteTypePageBottom  FootnoteEndnoteType = "pageBottom"
	FootnoteEndnoteTypeDocEnd      FootnoteEndnoteType = "docEnd"
	FootnoteEndnoteTypeSectEnd     FootnoteEndnoteType = "sectEnd"
)

// RestartNumber represents numbering restart locations
type RestartNumber string

const (
	RestartNumberContinuous RestartNumber = "continuous"
	RestartNumberEachSect   RestartNumber = "eachSect"
	RestartNumberEachPage   RestartNumber = "eachPage"
)

// BreakBinaryOperatorType represents break binary operator types
type BreakBinaryOperatorType string

const (
	BreakBinaryOperatorTypeBefore BreakBinaryOperatorType = "before"
	BreakBinaryOperatorTypeAfter  BreakBinaryOperatorType = "after"
	BreakBinaryOperatorTypeRepeat BreakBinaryOperatorType = "repeat"
)

// BreakBinarySubtractionType represents break binary subtraction types
type BreakBinarySubtractionType string

const (
	BreakBinarySubtractionTypeMinus  BreakBinarySubtractionType = "minus"
	BreakBinarySubtractionTypePlus   BreakBinarySubtractionType = "plus"
	BreakBinarySubtractionTypeRepeat BreakBinarySubtractionType = "repeat"
)

// JustificationType represents justification types
type JustificationType string

const (
	JustificationTypeLeft         JustificationType = "left"
	JustificationTypeRight        JustificationType = "right"
	JustificationTypeCenter       JustificationType = "center"
	JustificationTypeCenterGroup  JustificationType = "centerGroup"
)

// LimitLocationType represents limit location types
type LimitLocationType string

const (
	LimitLocationTypeSubSup LimitLocationType = "subSup"
	LimitLocationTypeUndOvr LimitLocationType = "undOvr"
)

// CharacterSpacing represents character spacing types
type CharacterSpacing string

const (
	CharacterSpacingDoNotCompress    CharacterSpacing = "doNotCompress"
	CharacterSpacingCompressPunctuation CharacterSpacing = "compressPunctuation"
	CharacterSpacingCompressPunctuationAndKana CharacterSpacing = "compressPunctuationAndKana"
)