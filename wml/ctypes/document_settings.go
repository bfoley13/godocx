package ctypes

import (
	"encoding/xml"

	"github.com/bfoley13/godocx/wml/stypes"
)

// DocumentSettings represents document settings (w:settings)
type DocumentSettings struct {
	// View Settings
	View *View `xml:"view,omitempty"`

	// Zoom Settings
	Zoom *Zoom `xml:"zoom,omitempty"`

	// Remove Personal Information from File Properties
	RemovePersonalInfo *OnOff `xml:"removePersonalInfo,omitempty"`

	// Remove Date and Time from Annotations
	RemoveDateAndTime *OnOff `xml:"removeDateAndTime,omitempty"`

	// Do Not Display Visual Boundary For Header/Footer or Between Pages and Text
	DoNotDisplayPageBoundaries *OnOff `xml:"doNotDisplayPageBoundaries,omitempty"`

	// Display Background Objects When Displaying Document
	DisplayBackgroundShape *OnOff `xml:"displayBackgroundShape,omitempty"`

	// Print PostScript Over Text
	PrintPostScriptOverText *OnOff `xml:"printPostScriptOverText,omitempty"`

	// Print Fractional Character Widths
	PrintFractionalWidths *OnOff `xml:"printFractionalWidths,omitempty"`

	// Only Print Form Field Content
	PrintFormsData *OnOff `xml:"printFormsData,omitempty"`

	// Embed TrueType Fonts
	EmbedTrueTypeFonts *OnOff `xml:"embedTrueTypeFonts,omitempty"`

	// Embed Common System Fonts
	EmbedSystemFonts *OnOff `xml:"embedSystemFonts,omitempty"`

	// Subset Fonts When Percent of Characters Used is Greater Than...
	SaveSubsetFonts *OnOff `xml:"saveSubsetFonts,omitempty"`

	// Save Form Data as Delimited Text File
	SaveFormsData *OnOff `xml:"saveFormsData,omitempty"`

	// Embed Linguistic Data
	MirrorMargins *OnOff `xml:"mirrorMargins,omitempty"`

	// Align Border and Edges With Page Border
	AlignBorderAndEdges *OnOff `xml:"alignBorderAndEdges,omitempty"`

	// Page Border Excludes Header
	BordersDoNotSurroundHeader *OnOff `xml:"bordersDoNotSurroundHeader,omitempty"`

	// Page Border Excludes Footer
	BordersDoNotSurroundFooter *OnOff `xml:"bordersDoNotSurroundFooter,omitempty"`

	// Position Gutter At Top of Page
	GutterAtTop *OnOff `xml:"gutterAtTop,omitempty"`

	// Do Not Add Space For Underlines
	HideSpellingErrors *OnOff `xml:"hideSpellingErrors,omitempty"`

	// Do Not Display Grammar Errors
	HideGrammaticalErrors *OnOff `xml:"hideGrammaticalErrors,omitempty"`

	// Writing Style Language Settings
	ActiveWritingStyle []ActiveWritingStyle `xml:"activeWritingStyle,omitempty"`

	// Proof State
	ProofState *ProofState `xml:"proofState,omitempty"`

	// Forms Design Mode
	FormsDesign *OnOff `xml:"formsDesign,omitempty"`

	// Structured Document Tag Placeholder Text Should be Resaved
	AttachedTemplate *CTString `xml:"attachedTemplate,omitempty"`

	// Linked Style Type
	LinkStyles *OnOff `xml:"linkStyles,omitempty"`

	// Stylistic Set Version
	StylePaneFormatFilter *CTString `xml:"stylePaneFormatFilter,omitempty"`

	// Default Tab Stop
	DefaultTabStop *DecimalNum `xml:"defaultTabStop,omitempty"`

	// Automatically Hyphenate Document Contents When Displayed
	AutoHyphenation *OnOff `xml:"autoHyphenation,omitempty"`

	// Maximum Number of Consecutively Hyphenated Lines
	ConsecutiveHyphenLimit *DecimalNum `xml:"consecutiveHyphenLimit,omitempty"`

	// Hyphenation Zone
	HyphenationZone *DecimalNum `xml:"hyphenationZone,omitempty"`

	// Do Not Hyphenate Words In ALL CAPITAL LETTERS
	DoNotHyphenateCaps *OnOff `xml:"doNotHyphenateCaps,omitempty"`

	// Show E-Mail Message Header
	ShowEnvelope *OnOff `xml:"showEnvelope,omitempty"`

	// Summary Length
	SummaryLength *DecimalNum `xml:"summaryLength,omitempty"`

	// Divider Character Between Click and Type Text
	ClickAndTypeStyle *CTString `xml:"clickAndTypeStyle,omitempty"`

	// Default Table Style for Newly Inserted Tables
	DefaultTableStyle *CTString `xml:"defaultTableStyle,omitempty"`

	// Do Not Validate Custom XML Markup Against Schema
	DoNotValidateAgainstSchema *OnOff `xml:"doNotValidateAgainstSchema,omitempty"`

	// Save Invalid Custom XML Markup
	SaveInvalidXML *OnOff `xml:"saveInvalidXML,omitempty"`

	// Ignore Mixed Content When Validating Custom XML Markup
	IgnoreMixedContent *OnOff `xml:"ignoreMixedContent,omitempty"`

	// Allow Storing of Previous Save Locations
	AlwaysShowPlaceholderText *OnOff `xml:"alwaysShowPlaceholderText,omitempty"`

	// Do Not Mark Grammar Errors in this Document
	DoNotDemarcateInvalidXML *OnOff `xml:"doNotDemarcateInvalidXML,omitempty"`

	// Save XML Data Only
	SaveXMLDataOnly *OnOff `xml:"saveXMLDataOnly,omitempty"`

	// Use XSL Transform When Saving
	UseXSLTWhenSaving *OnOff `xml:"useXSLTWhenSaving,omitempty"`

	// Save Document as XML File through Custom XSL Transform
	SaveThroughXSLT *SaveThroughXSLT `xml:"saveThroughXSLT,omitempty"`

	// Do Not Show Custom XML Markup Start/End Locations
	ShowXMLTags *OnOff `xml:"showXMLTags,omitempty"`

	// Always Show Placeholder Text
	AlwaysMergeEmptyNamespace *OnOff `xml:"alwaysMergeEmptyNamespace,omitempty"`

	// Update Fields on Open
	UpdateFields *OnOff `xml:"updateFields,omitempty"`

	// Footnote and Endnote Numbering Settings
	FootnoteDocumentWideProperties *FootnoteProperties `xml:"footnotePr,omitempty"`
	EndnoteDocumentWideProperties  *EndnoteProperties  `xml:"endnotePr,omitempty"`

	// Compatibility Settings
	Compat *Compat `xml:"compat,omitempty"`

	// Document Variables
	DocVars []DocVar `xml:"docVar,omitempty"`

	// Revision Identifiers for Parts of a Document
	Rsids *Rsids `xml:"rsids,omitempty"`

	// Math Properties
	MathPr *MathPr `xml:"mathPr,omitempty"`

	// Settings for Universal Input Method Editor
	UICompat97To2003 *OnOff `xml:"uiCompat97To2003,omitempty"`

	// Character Spacing Control Settings
	CharacterSpacingControl *CharacterSpacingValues `xml:"characterSpacingControl,omitempty"`

	// Do Not Automatically Compress Pictures
	DoNotAutoCompressPictures *OnOff `xml:"doNotAutoCompressPictures,omitempty"`
}

// View represents view settings (w:view)
type View struct {
	// Document View Setting
	Val *GenSingleStrVal[stypes.ViewType] `xml:"val,attr,omitempty"`
}

// Zoom represents zoom settings (w:zoom)
type Zoom struct {
	// Zoom Type
	Val *GenSingleStrVal[stypes.ZoomType] `xml:"val,attr,omitempty"`

	// Zoom Percentage
	Percent *CTString `xml:"percent,attr,omitempty"`
}

// ActiveWritingStyle represents active writing style settings (w:activeWritingStyle)
type ActiveWritingStyle struct {
	// Writing Style Language
	Lang *CTString `xml:"lang,attr,omitempty"`

	// Grammatical Engine ID
	VendorID *DecimalNum `xml:"vendorID,attr,omitempty"`

	// Grammatical Check Engine Version
	DllVersion *DecimalNum `xml:"dllVersion,attr,omitempty"`

	// Natural Language Grammar Check
	NlCheck *OnOff `xml:"nlCheck,attr,omitempty"`

	// Check Stylistic Rules With Grammar
	CheckStyle *OnOff `xml:"checkStyle,attr,omitempty"`

	// Application Defined Writing Style
	AppName *CTString `xml:"appName,attr,omitempty"`
}

// ProofState represents proof state settings (w:proofState)
type ProofState struct {
	// Spell Checking State
	Spelling *GenSingleStrVal[stypes.ProofingState] `xml:"spelling,attr,omitempty"`

	// Grammatical Checking State
	Grammar *GenSingleStrVal[stypes.ProofingState] `xml:"grammar,attr,omitempty"`
}

// SaveThroughXSLT represents save through XSLT settings (w:saveThroughXSLT)
type SaveThroughXSLT struct {
	// XSL Transform Location
	ID *CTString `xml:"id,attr,omitempty"`

	// Local Identifier for XSL Transform
	SolutionID *CTString `xml:"solutionID,attr,omitempty"`
}

// FootnoteProperties represents footnote properties (w:footnotePr)
type FootnoteProperties struct {
	// Footnote Positioning
	Pos *GenSingleStrVal[stypes.FootnoteEndnoteType] `xml:"pos,omitempty"`

	// Footnote Numbering Format
	NumFmt *GenSingleStrVal[stypes.NumFmt] `xml:"numFmt,omitempty"`

	// Footnote and Endnote Numbering Starting Value
	NumStart *DecimalNum `xml:"numStart,omitempty"`

	// Footnote and Endnote Numbering Restart Location
	NumRestart *GenSingleStrVal[stypes.RestartNumber] `xml:"numRestart,omitempty"`
}

// EndnoteProperties represents endnote properties (w:endnotePr)
type EndnoteProperties struct {
	// Endnote Positioning
	Pos *GenSingleStrVal[stypes.FootnoteEndnoteType] `xml:"pos,omitempty"`

	// Endnote Numbering Format
	NumFmt *GenSingleStrVal[stypes.NumFmt] `xml:"numFmt,omitempty"`

	// Footnote and Endnote Numbering Starting Value
	NumStart *DecimalNum `xml:"numStart,omitempty"`

	// Footnote and Endnote Numbering Restart Location
	NumRestart *GenSingleStrVal[stypes.RestartNumber] `xml:"numRestart,omitempty"`
}

// Compat represents compatibility settings (w:compat)
type Compat struct {
	// Use Simplified Rules For Table Border Conflicts
	UseSingleBorderforContiguousCells *OnOff `xml:"useSingleBorderforContiguousCells,omitempty"`

	// Emulate Word 6.x/Word 95/Word 97 Text Wrapping Around Objects
	WpJustification *OnOff `xml:"wpJustification,omitempty"`

	// Do Not Create Custom Tab Stop for Hanging Indent
	NoTabHangInd *OnOff `xml:"noTabHangInd,omitempty"`

	// Do Not Add Leading Between Lines of Text
	NoLeading *OnOff `xml:"noLeading,omitempty"`

	// Do Not Add Space For Underlines
	SpaceForUL *OnOff `xml:"spaceForUL,omitempty"`

	// Balance SBCS Characters and DBCS Characters
	NoColumnBalance *OnOff `xml:"noColumnBalance,omitempty"`

	// Do Not Balance Text Columns within a Section
	BalanceSingleByteDoubleByteWidth *OnOff `xml:"balanceSingleByteDoubleByteWidth,omitempty"`

	// Do Not Snap to Grid in Table Cells with Objects
	NoExtraLineSpacing *OnOff `xml:"noExtraLineSpacing,omitempty"`

	// Do Not Allow Hanging Punctuation With Character Grid
	DoNotLeaveBackslashAlone *OnOff `xml:"doNotLeaveBackslashAlone,omitempty"`

	// Convert Backslash To Yen Sign When Entered
	UlTrailSpace *OnOff `xml:"ulTrailSpace,omitempty"`

	// Add Space for Underlines
	DoNotExpandShiftReturn *OnOff `xml:"doNotExpandShiftReturn,omitempty"`

	// Don't Justify Lines Ending in Soft Line Break
	SpacingInWholePoints *OnOff `xml:"spacingInWholePoints,omitempty"`

	// Line Wrap Trailing Spaces
	LineWrapLikeWord6 *OnOff `xml:"lineWrapLikeWord6,omitempty"`

	// Emulate Word 6.0 Line Wrapping for East Asian Text
	PrintBodyTextBeforeHeader *OnOff `xml:"printBodyTextBeforeHeader,omitempty"`

	// Print Colors as Black And White without Dithering
	PrintColBlack *OnOff `xml:"printColBlack,omitempty"`

	// Space Width Like WordPerfect 5.x
	WpSpaceWidth *OnOff `xml:"wpSpaceWidth,omitempty"`

	// Display Page/Column Breaks As Printed
	ShowBreaksInFrames *OnOff `xml:"showBreaksInFrames,omitempty"`

	// Increase Priority of Font Size During Font Substitution
	SubFontBySize *OnOff `xml:"subFontBySize,omitempty"`

	// Ignore Exact Line Height for Last Line on Page
	SuppressBottomSpacing *OnOff `xml:"suppressBottomSpacing,omitempty"`

	// Ignore Minimum and Exact Line Height for First Line on Page
	SuppressTopSpacing *OnOff `xml:"suppressTopSpacing,omitempty"`

	// Ignore Minimum Line Height for First Line on Page
	SuppressSpacingAtTopOfPage *OnOff `xml:"suppressSpacingAtTopOfPage,omitempty"`

	// Emulate WordPerfect 6.x Font Height Calculation
	SuppressTopSpacingWP *OnOff `xml:"suppressTopSpacingWP,omitempty"`

	// Emulate Word 5.x Line Spacing
	SuppressSpBfAfterPgBrk *OnOff `xml:"suppressSpBfAfterPgBrk,omitempty"`

	// Use Printer Metrics to Display Documents
	SwapBordersFacingPages *OnOff `xml:"swapBordersFacingPages,omitempty"`

	// Treat Backslash Quote Delimiter as Two Backslash Characters
	ConvMailMergeEsc *OnOff `xml:"convMailMergeEsc,omitempty"`

	// Truncate Font Height
	TruncateFontHeightsLikeWP6 *OnOff `xml:"truncateFontHeightsLikeWP6,omitempty"`

	// Emulate Word 6.x/Word 95/Word 97 Text Wrapping Around Objects
	MwSmallCaps *OnOff `xml:"mwSmallCaps,omitempty"`

	// Use ANSI Kerning Pairs from Fonts
	UsePrinterMetrics *OnOff `xml:"usePrinterMetrics,omitempty"`

	// Do Not Use HTML Paragraph Auto Spacing
	DoNotSuppressParagraphBorders *OnOff `xml:"doNotSuppressParagraphBorders,omitempty"`

	// Ignore Width of Last Tab Stop When Aligning Paragraph If It Is Not Left Aligned
	WrapTrailSpaces *OnOff `xml:"wrapTrailSpaces,omitempty"`

	// Do Not Use East Asian Break Rules
	FootnoteLayoutLikeWW8 *OnOff `xml:"footnoteLayoutLikeWW8,omitempty"`

	// Do Not Automatically Apply List Paragraph Style To Bulleted/Numbered Text
	ShapeLayoutLikeWW8 *OnOff `xml:"shapeLayoutLikeWW8,omitempty"`

	// Align Table Rows Independently
	AlignTablesRowByRow *OnOff `xml:"alignTablesRowByRow,omitempty"`

	// Ignore Hanging Indent When Creating Tab Stop After Numbering
	ForgetLastTabAlignment *OnOff `xml:"forgetLastTabAlignment,omitempty"`

	// Use Word 2002 Table Style Rules
	DoNotUseHTMLParagraphAutoSpacing *OnOff `xml:"doNotUseHTMLParagraphAutoSpacing,omitempty"`

	// Emulate Word 6.x/Word 95 Full-Width Character Spacing
	LayoutRawTableWidth *OnOff `xml:"layoutRawTableWidth,omitempty"`

	// Emulate Word 97 Text Wrapping Around Floating Objects
	LayoutTableRowsApart *OnOff `xml:"layoutTableRowsApart,omitempty"`

	// Allow Tables to AutoFit Into Page Margins
	UseWord97LineBreakRules *OnOff `xml:"useWord97LineBreakRules,omitempty"`

	// Emulate Microsoft Word Version
	DoNotBreakWrappedTables *OnOff `xml:"doNotBreakWrappedTables,omitempty"`

	// Don't Allow Rows to Break Across Pages in Tables
	DoNotSnapToGridInCell *OnOff `xml:"doNotSnapToGridInCell,omitempty"`

	// Select Field When First or Last Character is Selected
	SelectFldWithFirstOrLastChar *OnOff `xml:"selectFldWithFirstOrLastChar,omitempty"`

	// Use Legacy Ethiopic and Amharic Line Breaking Rules
	ApplyBreakingRules *OnOff `xml:"applyBreakingRules,omitempty"`

	// Do Not Allow Floating Tables To Break Across Pages
	DoNotAllowInsOfMoveOfMoveIntoTextbox *OnOff `xml:"doNotAllowInsOfMoveOfMoveIntoTextbox,omitempty"`

	// Use Word 2003 Kashida Justification Method
	UseWord2002TableStyleRules *OnOff `xml:"useWord2002TableStyleRules,omitempty"`

	// Grow AutoFit Tables Into Page Margins
	GrowAutofit *OnOff `xml:"growAutofit,omitempty"`

	// Don't Bypass East Asian/Complex Script Layout Code
	UseFELayout *OnOff `xml:"useFELayout,omitempty"`

	// Don't Automatically Apply List Paragraph Style To Bulleted/Numbered Text
	UseNormalStyleForList *OnOff `xml:"useNormalStyleForList,omitempty"`

	// Ignore Hanging Indent When Creating Tab Stop After Numbering
	DoNotUseIndentAsNumberingTabStop *OnOff `xml:"doNotUseIndentAsNumberingTabStop,omitempty"`

	// Use Alternate Set of East Asian Line Breaking Rules
	UseAltKinsokuLineBreakRules *OnOff `xml:"useAltKinsokuLineBreakRules,omitempty"`

	// Allow Contextual Spacing of Paragraphs in Tables
	AllowSpaceOfSameStyleInTable *OnOff `xml:"allowSpaceOfSameStyleInTable,omitempty"`

	// Do Not Ignore Floating Objects When Calculating Paragraph Indentation
	DoNotSuppressIndentation *OnOff `xml:"doNotSuppressIndentation,omitempty"`

	// Do Not AutoFit Tables To Fit Next To Wrapped Objects
	DoNotAutofitConstrainedTables *OnOff `xml:"doNotAutofitConstrainedTables,omitempty"`

	// Allow Table Columns To Exceed Preferred Widths of Constituent Cells
	AutofitToFirstFixedWidthCell *OnOff `xml:"autofitToFirstFixedWidthCell,omitempty"`

	// Underline Following Character Following Numbering
	UnderlineTabInNumList *OnOff `xml:"underlineTabInNumList,omitempty"`

	// Always Use Fixed Width for Hangul Characters
	DisplayHangulFixedWidth *OnOff `xml:"displayHangulFixedWidth,omitempty"`

	// Always Move Paragraph Mark to Page after a Page Break
	SplitPgBreakAndParaMark *OnOff `xml:"splitPgBreakAndParaMark,omitempty"`

	// Don't Vertically Align Cells Containing Floating Objects
	DoNotVertAlignCellWithSp *OnOff `xml:"doNotVertAlignCellWithSp,omitempty"`

	// Don't Break Table Rows Around Floating Tables
	DoNotBreakConstrainedForcedTable *OnOff `xml:"doNotBreakConstrainedForcedTable,omitempty"`

	// Ignore Vertical Alignment in Textboxes
	DoNotVertAlignInTxbx *OnOff `xml:"doNotVertAlignInTxbx,omitempty"`

	// Use ANSI Kerning Pairs from Fonts
	UseAnsiKerningPairs *OnOff `xml:"useAnsiKerningPairs,omitempty"`

	// Use Cached Paragraph Information for Column Balancing
	CachedColBalance *OnOff `xml:"cachedColBalance,omitempty"`
}

// DocVar represents document variable (w:docVar)
type DocVar struct {
	// Document Variable Name
	Name *CTString `xml:"name,attr,omitempty"`

	// Document Variable Value
	Val *CTString `xml:"val,attr,omitempty"`
}

// Rsids represents revision save IDs (w:rsids)
type Rsids struct {
	// Original Document Revision Save ID
	RsidRoot *stypes.LongHexNum `xml:"rsidRoot,omitempty"`

	// Revision Save ID Values
	Rsid []stypes.LongHexNum `xml:"rsid,omitempty"`
}

// MathPr represents math properties (w:mathPr)
type MathPr struct {
	// Math Font
	MathFont *CTString `xml:"mathFont,omitempty"`

	// Break Binary Operators
	BrkBin *GenSingleStrVal[stypes.BreakBinaryOperatorType] `xml:"brkBin,omitempty"`

	// Break Binary Subtraction
	BrkBinSub *GenSingleStrVal[stypes.BreakBinarySubtractionType] `xml:"brkBinSub,omitempty"`

	// Small Fraction
	SmallFrac *OnOff `xml:"smallFrac,omitempty"`

	// Display Math
	DispDef *OnOff `xml:"dispDef,omitempty"`

	// Left Margin
	LMargin *DecimalNum `xml:"lMargin,omitempty"`

	// Right Margin
	RMargin *DecimalNum `xml:"rMargin,omitempty"`

	// Default Justification
	DefJc *GenSingleStrVal[stypes.JustificationType] `xml:"defJc,omitempty"`

	// Pre-Equation Spacing
	PreSpacing *DecimalNum `xml:"preSpacing,omitempty"`

	// Post-Equation Spacing
	PostSpacing *DecimalNum `xml:"postSpacing,omitempty"`

	// Inter-Equation Spacing
	InterSpacing *DecimalNum `xml:"interSpacing,omitempty"`

	// Intra-Equation Spacing
	IntraSpacing *DecimalNum `xml:"intraSpacing,omitempty"`

	// Wrap Indent
	WrapIndent *DecimalNum `xml:"wrapIndent,omitempty"`

	// Wrap Right
	WrapRight *OnOff `xml:"wrapRight,omitempty"`

	// Integral Limit Locations
	IntLim *GenSingleStrVal[stypes.LimitLocationType] `xml:"intLim,omitempty"`

	// n-ary Limit Locations
	NaryLim *GenSingleStrVal[stypes.LimitLocationType] `xml:"naryLim,omitempty"`
}

// CharacterSpacingValues represents character spacing control values
type CharacterSpacingValues struct {
	// Value
	Val *GenSingleStrVal[stypes.CharacterSpacing] `xml:"val,attr,omitempty"`
}

// MarshalXML implements xml.Marshaler for DocumentSettings
func (d DocumentSettings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:settings"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Marshal all settings properties
	if d.View != nil {
		if err := d.View.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if d.Zoom != nil {
		if err := d.Zoom.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if d.RemovePersonalInfo != nil {
		if err := e.EncodeElement(d.RemovePersonalInfo, xml.StartElement{Name: xml.Name{Local: "w:removePersonalInfo"}}); err != nil {
			return err
		}
	}

	if d.DefaultTabStop != nil {
		if err := d.DefaultTabStop.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:defaultTabStop"}}); err != nil {
			return err
		}
	}

	if d.AutoHyphenation != nil {
		if err := e.EncodeElement(d.AutoHyphenation, xml.StartElement{Name: xml.Name{Local: "w:autoHyphenation"}}); err != nil {
			return err
		}
	}

	// Add other fields as needed...
	for _, docVar := range d.DocVars {
		if err := docVar.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	if d.Compat != nil {
		if err := d.Compat.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for DocumentSettings
func (d *DocumentSettings) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "view":
				d.View = &View{}
				if err := d.View.UnmarshalXML(decoder, elem); err != nil {
					return err
				}
			case "zoom":
				d.Zoom = &Zoom{}
				if err := d.Zoom.UnmarshalXML(decoder, elem); err != nil {
					return err
				}
			case "removePersonalInfo":
				d.RemovePersonalInfo = &OnOff{}
				if err := decoder.DecodeElement(d.RemovePersonalInfo, &elem); err != nil {
					return err
				}
			case "defaultTabStop":
				d.DefaultTabStop = &DecimalNum{}
				if err := decoder.DecodeElement(d.DefaultTabStop, &elem); err != nil {
					return err
				}
			case "autoHyphenation":
				d.AutoHyphenation = &OnOff{}
				if err := decoder.DecodeElement(d.AutoHyphenation, &elem); err != nil {
					return err
				}
			case "docVar":
				var docVar DocVar
				if err := docVar.UnmarshalXML(decoder, elem); err != nil {
					return err
				}
				d.DocVars = append(d.DocVars, docVar)
			case "compat":
				d.Compat = &Compat{}
				if err := d.Compat.UnmarshalXML(decoder, elem); err != nil {
					return err
				}
			default:
				if err := decoder.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}

// MarshalXML implements xml.Marshaler for View
func (v View) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:view"
	if v.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(v.Val.Val)})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for View
func (v *View) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			v.Val = &GenSingleStrVal[stypes.ViewType]{Val: stypes.ViewType(attr.Value)}
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		if _, ok := token.(xml.EndElement); ok {
			break
		}
	}
	return nil
}

// MarshalXML implements xml.Marshaler for Zoom
func (z Zoom) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:zoom"
	if z.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(z.Val.Val)})
	}
	if z.Percent != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:percent"}, Value: z.Percent.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for Zoom
func (z *Zoom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "val":
			z.Val = &GenSingleStrVal[stypes.ZoomType]{Val: stypes.ZoomType(attr.Value)}
		case "percent":
			z.Percent = NewCTString(attr.Value)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		if _, ok := token.(xml.EndElement); ok {
			break
		}
	}
	return nil
}

// MarshalXML implements xml.Marshaler for DocVar
func (d DocVar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:docVar"
	if d.Name != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:name"}, Value: d.Name.Val})
	}
	if d.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: d.Val.Val})
	}
	return e.EncodeElement("", start)
}

// UnmarshalXML implements xml.Unmarshaler for DocVar
func (d *DocVar) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			d.Name = NewCTString(attr.Value)
		case "val":
			d.Val = NewCTString(attr.Value)
		}
	}
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if _, ok := token.(xml.EndElement); ok {
			break
		}
	}
	return nil
}

// MarshalXML implements xml.Marshaler for Compat
func (c Compat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:compat"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Only marshal a few key compatibility settings for brevity
	if c.UseSingleBorderforContiguousCells != nil {
		if err := e.EncodeElement(c.UseSingleBorderforContiguousCells, xml.StartElement{Name: xml.Name{Local: "w:useSingleBorderforContiguousCells"}}); err != nil {
			return err
		}
	}

	if c.WpJustification != nil {
		if err := e.EncodeElement(c.WpJustification, xml.StartElement{Name: xml.Name{Local: "w:wpJustification"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for Compat
func (c *Compat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "useSingleBorderforContiguousCells":
				c.UseSingleBorderforContiguousCells = &OnOff{}
				if err := d.DecodeElement(c.UseSingleBorderforContiguousCells, &elem); err != nil {
					return err
				}
			case "wpJustification":
				c.WpJustification = &OnOff{}
				if err := d.DecodeElement(c.WpJustification, &elem); err != nil {
					return err
				}
			// Add other compatibility settings as needed
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}
}