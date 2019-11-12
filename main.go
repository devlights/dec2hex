package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Do は、指定された10進数文字列を16進数に変換します。
func Do(v string) (string, error) {
	if len(v) == 0 {
		return "", nil
	}

	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("0x%X", i), nil
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: dec2hex dec-value(e.g: 255)")
		os.Exit(1)
	}

	v := args[0]
	h, err := Do(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\t-->\t%s\n", v, h)
	os.Exit(0)
}
