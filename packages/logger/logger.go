package informer

import (
	"fmt"

	"github.com/fatih/color"
)

type MsgType int

const (
	Error MsgType = iota
	Warning
	Good
	Info
)

type Message struct {
	Type    MsgType
	Message string
	NoBreak bool
}

type finalMessage struct {
	Color       color.Attribute
	Placeholder string
}

func getPrintFunc(typeOf MsgType) *finalMessage {
	colorMap := make(map[int]*finalMessage)

	colorMap[0] = &finalMessage{Color: color.FgRed, Placeholder: "[HALT] "}
	colorMap[1] = &finalMessage{Color: color.FgYellow, Placeholder: "[WARN] "}
	colorMap[2] = &finalMessage{Color: color.FgGreen, Placeholder: "[GOOD] "}
	colorMap[3] = &finalMessage{Color: color.FgCyan, Placeholder: "[INFO] "}

	return colorMap[int(typeOf)]
}

func (m Message) Log() {
	printInfo := getPrintFunc(m.Type)
	colorized := color.New(printInfo.Color).SprintFunc()

	if !m.NoBreak {
		m.Message += "\n"
	}

	fmt.Print(colorized(printInfo.Placeholder + m.Message))
}
