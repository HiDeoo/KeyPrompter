package cli

import (
	"fmt"

	"github.com/fatih/color"
)

var Green = color.New(color.FgHiGreen).SprintFunc()
var BoldGreen = color.New(color.FgHiGreen, color.Bold).SprintFunc()
var BoldRed = color.New(color.FgHiRed, color.Bold).SprintFunc()

func PrintServerError(errorOrMessage interface{}) {
	fmt.Printf("%s %s\n", BoldRed("Server error:"), errorOrMessage)
}

func PrintUIError(message string) {
	fmt.Printf("%s %s\n", BoldRed("UI error:"), message)
}
