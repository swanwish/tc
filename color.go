package tc

import (
	"fmt"
	"strings"
)

const (
	StartSequence = "\x1b["
	EndSequence   = "\x1b[0m"
)

const Reset = 0

const (
	TextColorBlack        = 30
	TextColorRed          = 31
	TextColorGreen        = 32
	TextColorYellow       = 33
	TextColorBlue         = 34
	TextColorPurple       = 35
	TextColorCyan         = 36
	TextColorLightGrey    = 37
	TextColorDarkGrey     = 90
	TextColorLightRed     = 91
	TextColorLightGreen   = 92
	TextColorLightYellow  = 93
	TextColorLightBlue    = 94
	TextColorLightMagenta = 95
	TextColorLightCyan    = 96
	TextColorWhite        = 97
)

const (
	TextStyleNoEffect  = 0
	TextStyleBold      = 1
	TextStyleDim       = 2
	TextStyleUnderline = 4
	TextStyleBlink     = 5
	TextStyleInverse   = 7
	TextStyleHidden    = 8
)

const (
	BackgroundColorBlack        = 40
	BackgroundColorRed          = 41
	BackgroundColorGreen        = 42
	BackgroundColorYellow       = 43
	BackgroundColorBlue         = 44
	BackgroundColorPurple       = 45
	BackgroundColorCyan         = 46
	BackgroundColorLightGrey    = 47
	BackgroundColorDarkGrey     = 100
	BackgroundColorLightRed     = 101
	BackgroundColorLightGreen   = 102
	BackgroundColorLightYellow  = 103
	BackgroundColorLightBlue    = 104
	BackgroundColorLightMagenta = 105
	BackgroundColorLightCyan    = 106
	BackgroundColorWhite        = 107
)

func Text(text string, codes ...int) string {
	var colorCodes []string
	for _, code := range codes {
		colorCodes = append(colorCodes, fmt.Sprintf("%d", code))
	}
	if len(colorCodes) == 0 {
		return text
	}
	return fmt.Sprintf("%s%sm%s%s", StartSequence, strings.Join(colorCodes, ";"), text, EndSequence)
}
