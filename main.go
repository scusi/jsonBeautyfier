// xmlBeautyfier - prints an xml file indented
//
package main

import (
	"flag"
	"fmt"
	"github.com/clbanning/mxj"
	"io/ioutil"
	"log"
)

var prefix string
var indent string

func init() {
	flag.StringVar(&prefix, "p", "", "prefix to be used; when set to FILENAME it uses the current filename as prefix")
	flag.StringVar(&indent, "i", "  ", "indent to be used")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	jsonFile := flag.Arg(0)
	if prefix == "FILENAME" {
		prefix = jsonFile
	}

	data, err := readJSONFile(jsonFile)
	check(err)

	data, err = remarshalIndentJSON(data)
	check(err)

	// printout
	fmt.Printf("%s\n", data)
}

func readJSONFile(fileName string) (data []byte, err error) {
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		return data, err
	}
	return data, err
}

func remarshalIndentJSON(in []byte) (out []byte, err error) {
	// Unmarshall
	mv, err := mxj.NewMapJson(in)
	if err != nil {
		return out, err
	}
	// Marshall with Indent
	out, err = mv.JsonIndent(prefix, indent)
	if err != nil {
		return out, err
	}
	return out, nil
}
