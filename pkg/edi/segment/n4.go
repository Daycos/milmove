package edisegment

import (
	"fmt"
)

// N4 represents the N4 EDI segment
type N4 struct {
	CityName            string
	StateOrProvinceCode string
	PostalCode          string
	CountryCode         string
	LocationQualifier   string
	LocationIdentifier  string
}

// StringArray converts N4 to an array of strings
func (s *N4) StringArray() []string {
	return []string{
		"N4",
		s.CityName,
		s.StateOrProvinceCode,
		s.PostalCode,
		s.CountryCode,
		s.LocationQualifier,
		s.LocationIdentifier,
	}
}

// Parse parses an X12 string that's split into an array into the N4 struct
func (s *N4) Parse(parts []string) error {
	expectedNumElements := 6
	if len(parts) != expectedNumElements {
		return fmt.Errorf("N4: Wrong number of fields, expected %d, got %d", expectedNumElements, len(parts))
	}

	s.CityName = parts[0]
	s.StateOrProvinceCode = parts[1]
	s.PostalCode = parts[2]
	s.CountryCode = parts[3]
	s.LocationQualifier = parts[4]
	s.LocationIdentifier = parts[5]
	return nil
}
