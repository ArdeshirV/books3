package main

//package gocolor

import "fmt"

type Color byte

var (
	colors = []string{
		"\033[0m", // Normal
		"\033[1m", // Bold
		"\033[0;31m",
		"\033[0;36m",
		"\033[0;37m",
		"\033[0;34m",
		"\033[0;32m",
		"\033[0;33m",
		"\033[0;35m",
		"\033[1;31m",
		"\033[1;34m",
		"\033[1;36m",
		"\033[1;32m",
		"\033[1;33m",
		"\033[1;35m",
	}
)

const (
	Normal      Color = 0
	Bold        Color = 1
	WhiteBold   Color = 1
	Red         Color = 2
	Teal        Color = 3
	White       Color = 4
	Blue        Color = 5
	Green       Color = 6
	Yellow      Color = 7
	Magenta     Color = 8
	RedBold     Color = 9
	BlueBold    Color = 10
	TealBold    Color = 11
	GreenBold   Color = 12
	YellowBold  Color = 13
	MagentaBold Color = 14
)

var (
	ColorIn      Color
	ColorOut     Color
	ColorPrompt  Color
	ColorMessage Color
	ColorResult  Color
	ColorError   Color
)

func init() {
	SetColorsToDefault()
}

func SetColorsToDefault() {
	ColorIn = GreenBold
	ColorOut = YellowBold
	ColorPrompt = WhiteBold
	ColorMessage = TealBold
	ColorResult = MagentaBold
	ColorError = RedBold
}

func GetColorValue(color Color) string {
	return colors[color]
}

func ColorizeText(text string, color Color) string {
	return colors[color] + text + colors[Normal]
}

func NormalText(text string) string {
	return ColorizeText(text, Normal)
}

func WhiteText(text string) string {
	return NormalText(text)
}

func NormalBoldText(text string) string {
	return ColorizeText(text, Bold)
}

func WhiteBoldText(text string) string {
	return ColorizeText(text, WhiteBold)
}

func RedText(text string) string {
	return colors[RedBold] + text + colors[Normal]
}

func BlueText(text string) string {
	return colors[Blue] + text + colors[Normal]
}

func TealText(text string) string {
	return colors[Teal] + text + colors[Normal]
}

func GreenText(text string) string {
	return colors[Green] + text + colors[Normal]
}

func YellowText(text string) string {
	return colors[Yellow] + text + colors[Normal]
}

func MagentaText(text string) string {
	return colors[Magenta] + text + colors[Normal]
}

func BoldText(text string) string {
	return colors[Bold] + text + colors[Normal]
}

func RedBoldText(text string) string {
	return colors[RedBold] + text + colors[Normal]
}

func BlueBoldText(text string) string {
	return colors[BlueBold] + text + colors[Normal]
}

func TealBoldText(text string) string {
	return colors[TealBold] + text + colors[Normal]
}

func GreenBoldText(text string) string {
	return colors[GreenBold] + text + colors[Normal]
}

func YellowBoldText(text string) string {
	return colors[YellowBold] + text + colors[Normal]
}

func MagentaBoldText(text string) string {
	return colors[MagentaBold] + text + colors[Normal]
}

func Prompt(args ...any) string {
	return colors[ColorPrompt] + fmt.Sprint(args...) + colors[Normal]
}

func Promptf(format string, args ...any) string {
	return fmt.Sprintf(colors[ColorPrompt]+format+colors[Normal], args...)
}

func Out(args ...any) string {
	return colors[ColorOut] + fmt.Sprint(args...) + colors[Normal]
}

func Outf(format string, args ...any) string {
	return fmt.Sprintf(colors[ColorOut]+format+colors[Normal], args...)
}

func In(args ...any) string {
	return colors[ColorIn] + fmt.Sprint(args...) + colors[Normal]
}

func Inf(format string, args ...any) string {
	return fmt.Sprintf(colors[ColorIn]+format+colors[Normal], args...)
}

func Message(args ...any) string {
	return colors[ColorMessage] + fmt.Sprint(args...) + colors[Normal]
}

func Messagef(format string, args ...any) string {
	return fmt.Sprintf(colors[ColorMessage]+format+colors[Normal], args...)
}

func Result(args ...any) string {
	return colors[ColorResult] + fmt.Sprint(args...) + colors[Normal]
}

func Resultf(format string, args ...any) string {
	return fmt.Sprintf(colors[ColorResult]+format+colors[Normal], args...)
}

func Error(args ...any) string {
	return colors[ColorError] + fmt.Sprint(args...) + colors[Normal]
}

func Errorf(format string, args ...any) string {
	return fmt.Sprintf(colors[ColorError]+format+colors[Normal], args...)
}

func InputColor() string {
	return colors[ColorIn]
}
