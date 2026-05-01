package greeting

import "github.com/fatih/color"

var greeting string = (color.CyanString("Golang ")
						+ color.RedString("for")
						+ color.MagentaString(" Brave!"))

func Hello() string {
	return greeting
}
