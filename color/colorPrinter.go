// Package color has various Print functions that can be called to change the color of the text in standard out
package color

import (
	"log"

	"github.com/fatih/color"

	"ethical-developer/cli/gitctl/config"
)

// PrintInfo prints yellow colored text to standard out
func PrintInfo(msg interface{}) {
	printColored(msg, color.FgYellow, config.IsQuiet())
}

// PrintSubtleInfo prints magenta colored text to standard out
func PrintSubtleInfo(msg interface{}) {
	printColored(msg, color.FgHiMagenta, config.IsQuiet())
}

// PrintSuccess prints green colored text to standard out
func PrintSuccess(msg interface{}) {
	printColored(msg, color.FgGreen, false)
}

// PrintError prints red colored text to standard out
func PrintError(msg interface{}) {
	printColored(msg, color.FgRed, false)
}

func printColored(msg interface{}, colorOutput color.Attribute, quiet bool) {
	if quiet {
		return
	}
	if config.IsColored() {
		_, err := color.New(colorOutput).Println(msg)
		if err != nil {
			return
		}
	} else {
		log.Println(msg)
	}
}
