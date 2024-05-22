/* Copyright (C) 2022, 2023, 2024 by Blackcat Informatics® Inc.
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

/*
Package colorout implements a colorizing and formatting wrapper
around Dave Collin's Spew Go data structure dumping utility.

Changes and additions above and beyond Spew:

  - syntax highlight of strings via https://github.com/chroma/quick
  - a specific colored dumper for JSON strings
  - a console output version of spew.Dump that supports
    more readable and visual segmented output
*/
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

// Sdump provides a slightly console optimized version of spew.Sdump; returning
// a string representation of a given Go variable.
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

// SimpleColorString performs syntax highlight on a string and returns it.
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

// Dump produces a formatted and colorized data-structure dump to the console.
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
