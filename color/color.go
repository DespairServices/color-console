package color

import "fmt"

// https://pkg.go.dev/github.com/shiena/ansicolor#section-readme
// https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	// Format
	f = "%s%s%s"

	// Text attributes
	allOff       = "\x1b[0m"  // All attributes off(color at startup)
	boldOn       = "\x1b[1m"  // Bold on(enable foreground intensity)
	underlineOn  = "\x1b[4m"  // Underline on
	blinkOn      = "\x1b[5m"  // Blink on(enable background intensity)
	boldOff      = "\x1b[21m" // Bold off(disable foreground intensity)
	underlineOff = "\x1b[24m" // Underline off
	blinkOff     = "\x1b[25m" // Blink off(disable background intensity)

	// Foreground colors
	fgBlack        = "\x1b[30m" // Black
	fgRed          = "\x1b[31m" // Red
	fgGreen        = "\x1b[32m" // Green
	fgYellow       = "\x1b[33m" // Yellow
	fgBlue         = "\x1b[34m" // Blue
	fgMagenta      = "\x1b[35m" // Magenta
	fgCyan         = "\x1b[36m" // Cyan
	fgWhite        = "\x1b[37m" // White
	fgDefault      = "\x1b[39m" // Default(foreground color at startup)
	fgLightGray    = "\x1b[90m" // Light Gray
	fgLightRed     = "\x1b[91m" // Light Red
	fgLightGreen   = "\x1b[92m" // Light Green
	fgLightYellow  = "\x1b[93m" // Light Yellow
	fgLightBlue    = "\x1b[94m" // Light Blue
	fgLightMagenta = "\x1b[95m" // Light Magenta
	fgLightCyan    = "\x1b[96m" // Light Cyan
	fgLightWhite   = "\x1b[97m" // Light White

	// Background colors
	bgBlack        = "\x1b[40m"  // Black
	bgRed          = "\x1b[41m"  // Red
	bgGreen        = "\x1b[42m"  // Green
	bgYellow       = "\x1b[43m"  // Yellow
	bgBlue         = "\x1b[44m"  // Blue
	bgMagenta      = "\x1b[45m"  // Magenta
	bgCyan         = "\x1b[46m"  // Cyan
	bgWhite        = "\x1b[47m"  // White
	bgDefault      = "\x1b[49m"  // Default(background color at startup)
	bgLightGray    = "\x1b[100m" // Light Gray
	bgLightRed     = "\x1b[101m" // Light Red
	bgLightGreen   = "\x1b[102m" // Light Green
	bgLightYellow  = "\x1b[103m" // Light Yellow
	bgLightBlue    = "\x1b[104m" // Light Blue
	bgLightMagenta = "\x1b[105m" // Light Magenta
	bgLightCyan    = "\x1b[106m" // Light Cyan
	bgLightWhite   = "\x1b[107m" // Light White
)

// Attributes
func Reset(s any) string {
	return fmt.Sprintf(f, allOff, s, allOff)
}

func Bold(s any) string {
	return fmt.Sprintf(f, boldOn, s, allOff)
}

func Underline(s any) string {
	return fmt.Sprintf(f, underlineOn, s, allOff)
}

func Blink(s any) string {
	return fmt.Sprintf(f, blinkOn, s, allOff)
}

// Foreground
func FgBlack(s any) string {
	return fmt.Sprintf(f, fgBlack, s, fgDefault)
}

func FgRed(s any) string {
	return fmt.Sprintf(f, fgRed, s, fgDefault)
}

func FgGreen(s any) string {
	return fmt.Sprintf(f, fgGreen, s, fgDefault)
}

func FgYellow(s any) string {
	return fmt.Sprintf(f, fgYellow, s, fgDefault)
}

func FgBlue(s any) string {
	return fmt.Sprintf(f, fgBlue, s, fgDefault)
}

func FgMagenta(s any) string {
	return fmt.Sprintf(f, fgMagenta, s, fgDefault)
}

func FgCyan(s any) string {
	return fmt.Sprintf(f, fgCyan, s, fgDefault)
}

func FgWhite(s any) string {
	return fmt.Sprintf(f, fgWhite, s, fgDefault)
}

func FgDefault(s any) string {
	return fmt.Sprintf(f, fgDefault, s, fgDefault)
}

func FgLightGray(s any) string {
	return fmt.Sprintf(f, fgLightGray, s, fgDefault)
}

func FgLightRed(s any) string {
	return fmt.Sprintf(f, fgLightRed, s, fgDefault)
}

func FgLightGreen(s any) string {
	return fmt.Sprintf(f, fgLightGreen, s, fgDefault)
}

func FgLightYellow(s any) string {
	return fmt.Sprintf(f, fgLightYellow, s, fgDefault)
}

func FgLightBlue(s any) string {
	return fmt.Sprintf(f, fgLightBlue, s, fgDefault)
}

func FgLightMagenta(s any) string {
	return fmt.Sprintf(f, fgLightMagenta, s, fgDefault)
}

func FgLightCyan(s any) string {
	return fmt.Sprintf(f, fgLightCyan, s, fgDefault)
}

func FgLightWhite(s any) string {
	return fmt.Sprintf(f, fgLightWhite, s, fgDefault)
}

// Background
func BgBlack(s any) string {
	return fmt.Sprintf(f, bgBlack, s, bgDefault)
}

func BgRed(s any) string {
	return fmt.Sprintf(f, bgRed, s, bgDefault)
}

func BgGreen(s any) string {
	return fmt.Sprintf(f, bgGreen, s, bgDefault)
}

func BgYellow(s any) string {
	return fmt.Sprintf(f, bgYellow, s, bgDefault)
}

func BgBlue(s any) string {
	return fmt.Sprintf(f, bgBlue, s, bgDefault)
}

func BgMagenta(s any) string {
	return fmt.Sprintf(f, bgMagenta, s, bgDefault)
}

func BgCyan(s any) string {
	return fmt.Sprintf(f, bgCyan, s, bgDefault)
}

func BgWhite(s any) string {
	return fmt.Sprintf(f, bgWhite, s, bgDefault)
}

func BgDefault(s any) string {
	return fmt.Sprintf(f, bgDefault, s, bgDefault)
}

func BgLightGray(s any) string {
	return fmt.Sprintf(f, bgLightGray, s, bgDefault)
}

func BgLightRed(s any) string {
	return fmt.Sprintf(f, bgLightRed, s, bgDefault)
}

func BgLightGreen(s any) string {
	return fmt.Sprintf(f, bgLightGreen, s, bgDefault)
}

func BgLightYellow(s any) string {
	return fmt.Sprintf(f, bgLightYellow, s, bgDefault)
}

func BgLightBlue(s any) string {
	return fmt.Sprintf(f, bgLightBlue, s, bgDefault)
}

func BgLightMagenta(s any) string {
	return fmt.Sprintf(f, bgLightMagenta, s, bgDefault)
}

func BgLightCyan(s any) string {
	return fmt.Sprintf(f, bgLightCyan, s, bgDefault)
}

func BgLightWhite(s any) string {
	return fmt.Sprintf(f, bgLightWhite, s, bgDefault)
}
