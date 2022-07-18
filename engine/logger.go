package engine

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

type LogLevel uint8

const (
	Err LogLevel = iota
	Warn
	Info
	Trace
)

func Print(message string) {
	fmt.Printf("%v\n", message)
}

func Log(message string) {
	log.Printf("%v\n", message)
}

func LogTrace(message string) {
	log.Printf("%v %v\n", color.MagentaString("Trace:"), message)
}

func LogErr(message string) {
	log.Printf("%v %v\n", color.RedString("Error:"), message)
}

func LogWarn(message string) {
	log.Printf("%v %v\n", color.YellowString("Warning:"), message)
}

func LogInfo(message string) {
	log.Printf("%v %v\n", color.GreenString("Info:"), message)
}
