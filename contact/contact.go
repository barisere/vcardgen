package contact

// Contact is our model for contact data
type Contact struct {
	FirstName  string
	MiddleName string
	LastName   string
	Telephone  string
}

// NewContact constructs a Contact object
func NewContact(strslc []string) *Contact {
	var (
		fn, mn, ln, tel string
	)
	fn = strslc[0]
	if len(strslc) == 4 {
		mn, ln, tel = strslc[1], strslc[2], strslc[3]
	} else {
		mn, ln, tel = "", strslc[1], strslc[2]
	}
	return &Contact{
		FirstName:  fn,
		MiddleName: mn,
		LastName:   ln,
		Telephone:  tel,
	}
}
