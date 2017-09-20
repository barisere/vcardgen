package helpers

import (
	"bufio"
	"log"
	"os"

	"github.com/emersion/go-vcard"
	"github.com/latentgenius/vcardgen/contact"
)

// ReadFromFile reads a line from the given file and returns it as a string
func ReadFromFile(file *os.File) func() string {
	scanner := bufio.NewScanner(file)
	var line string
	var success bool
	return func() string {
		line = ""
		if success = scanner.Scan(); success != false {
			line += scanner.Text()
		} else if err := scanner.Err(); err == nil {
			log.Println("EOF reached.")
		} else {
			panic(err)
		}
		return line
	}
}

// ReadName formats a Contact object into a []*vcard.Field object containing the
// name of the contact with its components
func ReadName(c *contact.Contact) []*vcard.Field {
	return []*vcard.Field{
		&vcard.Field{
			Value: c.FirstName + ";" + c.LastName + ";" + c.MiddleName,
			Params: map[string][]string{
				vcard.ParamSortAs: []string{
					c.FirstName + " " + c.LastName,
				},
			},
		},
	}
}

// ReadFormattedName formats a Contact object into a []*vcard.Field object containing
// the formatted name of the contact
func ReadFormattedName(c *contact.Contact) []*vcard.Field {
	return []*vcard.Field{
		&vcard.Field{
			Value: c.FirstName + " " + c.LastName + " " + c.MiddleName,
		},
	}
}

// ReadTelephone formats a telephone number string into a vcard.Field object
func ReadTelephone(c *contact.Contact) []*vcard.Field {
	return []*vcard.Field{
		&vcard.Field{
			Value: c.Telephone,
		},
	}
}
