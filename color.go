// Copyright (C) 2022, 2023, 2024 by Blackcat Informatics® Inc.
//
// nolint: varnamelen,revive
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

func Sdump(i interface{}) string {
	return normal.Sdump(i)
}

// SdumpColored dumps a variable to a string with colored syntax highlighting.
func SdumpColored(i ...interface{}) string {
	s := normal.Sdump(i...)

	var b bytes.Buffer

	if err := quick.Highlight(&b, s, "go", "terminal16m", "dracula"); err != nil {
		return s
	}

	return b.String()
}

var simple = &spew.ConfigState{
	Indent:                  "  ",
	SortKeys:                true,
	DisableCapacities:       true,
	DisablePointerAddresses: true,
}

var simpleStripRe = regexp.MustCompile(`\(len=\d+\)\s?`)

func SimpleColorString(language string, s string) string {
	sR := simpleStripRe.ReplaceAllString(s, "")

	var b bytes.Buffer

	if err := quick.Highlight(&b, sR, language, "terminal16m", "dracula"); err != nil {
		return s
	}

	return b.String()
}

// SdumpColoredSimple dumps a variable to a string with colored syntax highlighting.
func SdumpColorSimple(i ...interface{}) string {
	return SimpleColorString("go", simple.Sdump(i...))
}

// SdumpColoredJSON dumps a variable to a string with colored syntax highlighting.
func SdumpColorJSON(m json.Marshaler) string {
	j, err := m.MarshalJSON()
	if err != nil {
		return "ERROR ENCODING JSON"
	}

	return SimpleColorString("json", string(j))
}

var White = color.New(color.FgWhite, color.Bold)
var Red = color.New(color.FgRed, color.Bold)
var Orange = color.New(color.FgYellow)
var Yellow = color.New(color.FgYellow, color.Bold)
var Green = color.New(color.FgGreen, color.Bold)
var Cyan = color.New(color.FgCyan, color.Bold)
var Blue = color.New(color.FgBlue, color.Bold)
var Magenta = color.New(color.FgMagenta)
var Grey = color.New(color.FgWhite)
var WhiteOnMagenta = color.New(color.BgMagenta, color.Bold, color.FgWhite)
var WhiteOnBlue = color.New(color.BgBlue, color.FgWhite, color.Bold)
var WhiteOnGreen = color.New(color.BgGreen, color.Bold, color.FgWhite)
var WhiteOnCyan = color.New(color.BgCyan, color.Bold, color.FgWhite)
var WhiteOnRed = color.New(color.BgRed, color.Bold, color.FgWhite)
var BlackOnYellow = color.New(color.BgHiYellow, color.FgBlack)

func Dump(name string, i ...interface{}) {
	color.NoColor = false
	title := White.SprintfFunc()
	top := Yellow.SprintfFunc()
	bottom := Blue.SprintfFunc()

	fmt.Printf("\n%s %s %s\n",
		top(">>>>>>>>------------------["),
		title(name),
		top("]--------------------------"))
	fmt.Print(SdumpColored(i...))
	fmt.Printf("%s %s %s\n",
		bottom("<<<<<<<<------------------["),
		title(name),
		bottom("]--------------------------"))
}
