package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/devlights/gomy/convert"
)

// Convert -- 指定された文字列を10進数文字列として16進数文字列へ変換します.
func Convert(val string) (string, error) {
	return convert.Dec2Hex(val, "0x", 0)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: dec2hex dec-value(e.g: 255)")
		os.Exit(1)
	}

	v := args[0]
	h, err := Convert(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\t-->\t%s\n", v, h)
	os.Exit(0)
}
