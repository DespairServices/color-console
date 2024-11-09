package console

import (
	"fmt"
	"os"

	"github.com/DespairServices/asset-store-price-backend/internal/color"
)

// Attributes
func Bold(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.Bold(s)
	fmt.Println(s)
}

func Underline(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.Underline(s)
	fmt.Println(s)
}

func Blink(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.Blink(s)
	fmt.Println(s)
}

// Foreground
func FgBlack(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgBlack(s)
	fmt.Println(s)
}

func FgRed(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgRed(s)
	fmt.Println(s)
}

func FgGreen(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgGreen(s)
	fmt.Println(s)
}

func FgYellow(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgYellow(s)
	fmt.Println(s)
}

func FgBlue(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgBlue(s)
	fmt.Println(s)
}

func FgMagenta(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgMagenta(s)
	fmt.Println(s)
}

func FgCyan(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgCyan(s)
	fmt.Println(s)
}

func FgWhite(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgWhite(s)
	fmt.Println(s)
}

func FgDefault(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgDefault(s)
	fmt.Println(s)
}

func FgLightGray(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightGray(s)
	fmt.Println(s)
}

func FgLightRed(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightRed(s)
	fmt.Println(s)
}

func FgLightGreen(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightGreen(s)
	fmt.Println(s)
}

func FgLightYellow(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightYellow(s)
	fmt.Println(s)
}

func FgLightBlue(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightBlue(s)
	fmt.Println(s)
}

func FgLightMagenta(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightMagenta(s)
	fmt.Println(s)
}

func FgLightCyan(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightCyan(s)
	fmt.Println(s)
}

func FgLightWhite(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.FgLightWhite(s)
	fmt.Println(s)
}

// Background
func BgBlack(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgBlack(s)
	fmt.Println(s)
}

func BgRed(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgRed(s)
	fmt.Println(s)
}

func BgGreen(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgGreen(s)
	fmt.Println(s)
}

func BgYellow(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgYellow(s)
	fmt.Println(s)
}

func BgBlue(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgBlue(s)
	fmt.Println(s)
}

func BgMagenta(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgMagenta(s)
	fmt.Println(s)
}

func BgCyan(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgCyan(s)
	fmt.Println(s)
}

func BgWhite(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgWhite(s)
	fmt.Println(s)
}

func BgDefault(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgDefault(s)
	fmt.Println(s)
}

func BgLightGray(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightGray(s)
	fmt.Println(s)
}

func BgLightRed(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightRed(s)
	fmt.Println(s)
}

func BgLightGreen(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightGreen(s)
	fmt.Println(s)
}

func BgLightYellow(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightYellow(s)
	fmt.Println(s)
}

func BgLightBlue(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightBlue(s)
	fmt.Println(s)
}

func BgLightMagenta(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightMagenta(s)
	fmt.Println(s)
}

func BgLightCyan(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightCyan(s)
	fmt.Println(s)
}

func BgLightWhite(m string, v ...any) {
	s := fmt.Sprintf(m, v...)
	s = color.BgLightWhite(s)
	fmt.Println(s)
}

// Types
func Error(m string, v ...any) {
	s := fmt.Sprintf("[ERROR] %s", m)
	FgRed(s, v...)
	os.Exit(1)
}

func ErrorLite(m string, v ...any) {
	s := fmt.Sprintf("[ERROR] %s", m)
	FgLightRed(s, v...)
}

func Info(m string, v ...any) {
	FgCyan(m, v...)
}

func InfoLite(m string, v ...any) {
	FgLightCyan(m, v...)
}

func Log(m string, v ...any) {
	FgDefault(m, v...)
}

func Success(m string, v ...any) {
	FgGreen(m, v...)
}

func SuccessLite(m string, v ...any) {
	FgLightGreen(m, v...)
}

func Warn(m string, v ...any) {
	s := fmt.Sprintf("[WARNING] %s", m)
	FgMagenta(s, v...)
}

func WarnLite(m string, v ...any) {
	s := fmt.Sprintf("[WARNING] %s", m)
	FgLightMagenta(s, v...)
}
