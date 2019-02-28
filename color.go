package tc

import "fmt"

const (
	StartSequence = "\033["
	EndSequence   = "\033[0m"
)

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
	TextStyleNoEffect    = 0
	TextStyleNoBold      = 1
	TextStyleNoDim       = 2
	TextStyleNoUnderline = 4
	TextStyleNoBlink     = 5
	TextStyleNoInverse   = 7
	TextStyleNoHidden    = 8
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

func ColorText(text string, textColor int) string {
	return fmt.Sprintf("%s%dm%s%s", StartSequence, textColor, text, EndSequence)
}

func ColorTextEx(text string, textColor, textStyle, backgroundColor int) string {
	return fmt.Sprintf("%s%d;%d;%dm%s%s", StartSequence, textColor, textStyle, backgroundColor, text, EndSequence)
}
