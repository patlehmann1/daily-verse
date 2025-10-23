package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yourusername/daily-verse/verses"
)

const version = "1.0.0"

func main() {
	daily := flag.Bool("daily", false, "Get the same verse for the entire day")
	book := flag.String("book", "", "Filter by book name (case-insensitive)")
	testament := flag.String("testament", "", "Filter by testament (old or new)")
	showVersion := flag.Bool("version", false, "Show version information")

	flag.Parse()

	if *showVersion {
		fmt.Printf("daily-verse v%s\n", version)
		os.Exit(0)
	}

	if *testament != "" && *testament != "old" && *testament != "new" {
		fmt.Fprintf(os.Stderr, "Error: testament must be 'old' or 'new'\n")
		os.Exit(1)
	}

	opts := verses.FilterOptions{
		Testament: *testament,
		Book:      *book,
	}

	verse, err := verses.GetVerse(*daily, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(verse.Text)
	fmt.Println(verse.Reference)
}
