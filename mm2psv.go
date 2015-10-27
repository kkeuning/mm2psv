package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"os"
)

func processNode (node xml.Node, row string) {
	row = row + node.Attr("TEXT") + "|"
	kids, err := node.Search("node")
	if err != nil {
		log.Println("Error searching for node:", err)
		return
	}
	if len(kids) > 0 {  // has children, not a leaf node
		for i := range kids {
			processNode(kids[i], row)
		}
	} else {
		fmt.Println(row) // print leaf node
	}
}

func fileContents(inFile string) []byte {
	// Read entire file contents into memory, ioutil.ReadFile() closes file after reading.
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Println("Error reading file:", err)
		return contents
	}
	return contents
}

func processXmlString(unparsedXml []byte) {
	doc, err := gokogiri.ParseXml(unparsedXml)
	if err != nil {
		log.Println("Error parsing file:", err)
		return
	}
	firstNode, err := doc.Node.Search("//node")
	row := "|" // empty pipe separated row, starter
	processNode(firstNode[0], row)
}

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
