// +build !windows

package main

import (
	"fmt"

	"gitlab.com/poldi1405/go-ansi"
)

var (
	progressStyle = "block"
	r, l          = "▕", "▏"
)

func color(content ...interface{}) string {
	return ansi.Blue(fmt.Sprint(content...))
}
