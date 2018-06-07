package main

import (
	"fmt"
	// "log"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: growth [filename]")
		os.Exit(1)
	}

	file := os.Args[1]

	first := GetSize(file)

	fmt.Printf("Current size: %s\n", humanize.Bytes(uint64(first)))
	fmt.Printf("Checking again in 5 seconds...\n")

	time.Sleep(5 * time.Second)

	last := GetSize(file)

	fmt.Printf("Size is now: %s\n", humanize.Bytes(uint64(last)))
	fmt.Println()

	if first == last {
		fmt.Println("File did not grow.")
		os.Exit(0)
	}

	if first > last {
		fmt.Println("File did actually shrink.")
		os.Exit(0)
	}

	diff := last - first

	fmt.Printf("File grew %s in 5 seconds.\n", humanize.Bytes(uint64(diff)))

	fmt.Println()

	fmt.Printf("Growth prediction (linear):\n")
	fmt.Printf("In 1 minute:\t%s total\n", humanize.Bytes(uint64(last+12*diff)))
	fmt.Printf("In 1 hour:\t%s total\n", humanize.Bytes(uint64(last+60*12*diff)))
	fmt.Printf("In 1 day:\t%s total\n", humanize.Bytes(uint64(last+24*60*12*diff)))
	fmt.Printf("In 1 week:\t%s total\n", humanize.Bytes(uint64(last+7*24*60*12*diff)))
	fmt.Printf("In 1 month:\t%s total\n", humanize.Bytes(uint64(last+30*24*60*12*diff)))
	fmt.Printf("In 1 year:\t%s total\n", humanize.Bytes(uint64(last+365*24*60*12*diff)))
}

func GetSize(file string) int64 {
	stats, err := os.Stat(file)
	if err != nil {
		fmt.Printf("growth: %v\n", err)
		os.Exit(1)
	}

	return stats.Size()
}
