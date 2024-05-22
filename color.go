// Copyright (C) 2022, 2023, 2024 by Blackcat Informatics® Inc.
package colorout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/alecthomas/chroma/quick"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

var normal = spew.ConfigState{
	SortKeys: true,
}

func Sdump(i any) string {
	return normal.Sdump(i)
}

// SdumpColored dumps a variable to a string with colored syntax highlighting.
func SdumpColored(i ...any) string {
	s := normal.Sdump(i...)

	var buffer bytes.Buffer

	if err := quick.Highlight(&buffer, s, "go", "terminal16m", "dracula"); err != nil {
		return s
	}

	return buffer.String()
}

var simple = &spew.ConfigState{
	Indent:                  "  ",
	SortKeys:                true,
	DisableCapacities:       true,
	DisablePointerAddresses: true,
}

var simpleStripRe = regexp.MustCompile(`\(len=\d+\)\s?`)

func SimpleColorString(language, str string) string {
	strReplaced := simpleStripRe.ReplaceAllString(str, "")

	var buffer bytes.Buffer

	if err := quick.Highlight(
		&buffer,
		strReplaced,
		language,
		"terminal16m",
		"dracula",
	); err != nil {
		return str
	}

	return buffer.String()
}

// SdumpColoredSimple dumps a variable to a string with colored syntax highlighting.
func SdumpColorSimple(i ...any) string {
	return SimpleColorString("go", simple.Sdump(i...))
}

// SdumpColoredJSON dumps a variable to a string with colored syntax highlighting.
func SdumpColorJSON(object json.Marshaler) string {
	jsonStr, err := object.MarshalJSON()
	if err != nil {
		return "ERROR ENCODING JSON"
	}

	return SimpleColorString("json", string(jsonStr))
}

var (
	White          = color.New(color.FgWhite, color.Bold)
	Red            = color.New(color.FgRed, color.Bold)
	Orange         = color.New(color.FgYellow)
	Yellow         = color.New(color.FgYellow, color.Bold)
	Green          = color.New(color.FgGreen, color.Bold)
	Cyan           = color.New(color.FgCyan, color.Bold)
	Blue           = color.New(color.FgBlue, color.Bold)
	Magenta        = color.New(color.FgMagenta)
	Grey           = color.New(color.FgWhite)
	WhiteOnMagenta = color.New(color.BgMagenta, color.Bold, color.FgWhite)
	WhiteOnBlue    = color.New(color.BgBlue, color.FgWhite, color.Bold)
	WhiteOnGreen   = color.New(color.BgGreen, color.Bold, color.FgWhite)
	WhiteOnCyan    = color.New(color.BgCyan, color.Bold, color.FgWhite)
	WhiteOnRed     = color.New(color.BgRed, color.Bold, color.FgWhite)
	BlackOnYellow  = color.New(color.BgHiYellow, color.FgBlack)
)

func Dump(name string, object ...any) {
	color.NoColor = false
	title := White.SprintfFunc()
	top := Yellow.SprintfFunc()
	bottom := Blue.SprintfFunc()

	fmt.Print(
		top(">>>>>>>>------------------[ ") +
			title(name) +
			top(" ]--------------------------\n") +
			SdumpColored(object...) +
			bottom("\n<<<<<<<<------------------[ ") +
			title(name) +
			bottom(" ]--------------------------\n"))
}
