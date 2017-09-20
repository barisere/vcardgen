package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/latentgenius/vcardgen/contact"
	"github.com/latentgenius/vcardgen/helpers"

	vcard "github.com/emersion/go-vcard"
)

func main() {
	var (
		inFile  string
		outFile string
		outExt  = vcard.Extension
	)

	// parse command line arguments
	flag.Parse()
	args := flag.Args()
	argLen := len(args)
	if argLen < 1 || argLen > 2 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Usage: %s inputfile [outputfile]", flag.Arg(0)))
		os.Exit(2)
	}
	inFile = args[0]
	if argLen == 2 {
		outFile = args[1]
	} else {
		outFile = inFile + "." + outExt
	}

	// open input file for reading
	readFrom, err := os.Open(inFile)
	if err != nil {
		log.Fatal("Could not open file" + inFile)
	}
	defer readFrom.Close()

	// open output file for writing, or create one if it does not exist
	writeTo, err := os.Create(outFile)
	if err != nil {
		log.Fatal("Could not create file" + outFile)
	}
	defer writeTo.Close()

	// get a generator for the file content
	fileReader := helpers.ReadFromFile(readFrom)

	// buffer the writes to save disk I/O time
	writeBuffer := bufio.NewWriter(writeTo)

	// destination where the vcard will be encoded to
	writer := vcard.NewEncoder(writeBuffer)
	var (
		text         string
		textSlc      []string
		card         = make(vcard.Card)
		contactModel *contact.Contact
	)
	text = fileReader()
	for text != "" {
		textSlc = strings.Split(text, " ")
		contactModel = contact.NewContact(textSlc)
		card[vcard.FieldFormattedName] = helpers.ReadFormattedName(contactModel)
		card[vcard.FieldName] = helpers.ReadName(contactModel)
		card[vcard.FieldTelephone] = helpers.ReadTelephone(contactModel)
		vcard.ToV4(card)
		err := writer.Encode(card)
		if err != nil {
			panic(err)
		}
		text = fileReader()
	}
	err = writeBuffer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
