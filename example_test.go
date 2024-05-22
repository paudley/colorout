// Copyright (C) 2022-2024 by Blackcat Informatics® Inc.
//nolint
package colorout_test

import (
	"fmt"

	"github.com/paudley/colorout"
)

func ExampleSdump() {
	intTest := 10
	fmt.Println(colorout.Sdump(intTest))

	structTest := struct {
		strKey  string
		strNext struct {
			nestKey int
		}
	}{
		strKey: "strValue",
		strNext: struct {
			nestKey int
		}{
			nestKey: 10,
		},
	}
	fmt.Println(colorout.Sdump(structTest))
	// Output:
	// (int) 10
	//
	// (struct { strKey string; strNext struct { nestKey int } }) {
	// strKey: (string) (len=8) "strValue",
	// strNext: (struct { nestKey int }) {
	// nestKey: (int) 10
	// }
	// }
}
