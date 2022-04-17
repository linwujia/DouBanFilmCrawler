package main

import (
	"flag"
	"fmt"
)

func main() {
	start := flag.Uint("s", 1, "start page index")
	end := flag.Uint("e", 10, "end page index")

	if *start > *end {
		fmt.Errorf("start page over than end page")
		return
	}

	manager := NewDouBanManager(*start, *end)
	manager.Run()
}
