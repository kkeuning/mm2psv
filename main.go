package main

import (
	"fmt"
	"os"
)

func main() {
	var inFile string
	if len(os.Args) > 1 {
		inFile = os.Args[1]
	} else {
		fmt.Println("No file specified.")
		fmt.Println("")
		fmt.Println("usage: " + os.Args[0] + " filename")
		fmt.Println("Converts freemind .mm file to pipe separated values output.")
		return
	}
	processXmlString(fileContents(inFile))
}
