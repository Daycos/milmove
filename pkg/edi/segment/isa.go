package edisegment

import (
	"fmt"
	"strconv"
)

// ISA represents the ISA EDI segment
type ISA struct {
	AuthorizationInformationQualifier string
	AuthorizationInformation          string
	SecurityInformationQualifier      string
	SecurityInformation               string
	InterchangeSenderIDQualifier      string
	InterchangeSenderID               string
	InterchangeReceiverIDQualifier    string
	InterchangeReceiverID             string
	InterchangeDate                   string
	InterchangeTime                   string
	InterchangeControlStandards       string
	InterchangeControlVersionNumber   string
	InterchangeControlNumber          int64
	AcknowledgementRequested          int
	UsageIndicator                    string
	ComponentElementSeparator         string
}

// StringArray converts ISA to an array of strings
func (s *ISA) StringArray() []string {
	return []string{
		"ISA",
		s.AuthorizationInformationQualifier,
		s.AuthorizationInformation,
		s.SecurityInformationQualifier,
		s.SecurityInformation,
		s.InterchangeSenderIDQualifier,
		s.InterchangeSenderID,
		s.InterchangeReceiverIDQualifier,
		s.InterchangeReceiverID,
		s.InterchangeDate,
		s.InterchangeTime,
		s.InterchangeControlStandards,
		s.InterchangeControlVersionNumber,
		fmt.Sprintf("%09d", s.InterchangeControlNumber),
		strconv.Itoa(s.AcknowledgementRequested),
		s.UsageIndicator,
		s.ComponentElementSeparator,
	}
}

// Parse parses an X12 string that's split into an array into the ISA struct
func (s *ISA) Parse(elements []string) error {
	expectedNumElements := 16
	if len(elements) != expectedNumElements {
		return fmt.Errorf("ISA: Wrong number of elements, expected %d, got %d", expectedNumElements, len(elements))
	}

	var err error
	s.AuthorizationInformationQualifier = elements[0]
	s.AuthorizationInformation = elements[1]
	s.SecurityInformationQualifier = elements[2]
	s.SecurityInformation = elements[3]
	s.InterchangeSenderIDQualifier = elements[4]
	s.InterchangeSenderID = elements[5]
	s.InterchangeReceiverIDQualifier = elements[6]
	s.InterchangeReceiverID = elements[7]
	s.InterchangeDate = elements[8]
	s.InterchangeTime = elements[9]
	s.InterchangeControlStandards = elements[10]
	s.InterchangeControlVersionNumber = elements[11]
	s.InterchangeControlNumber, err = strconv.ParseInt(elements[12], 10, 64)
	if err != nil {
		return err
	}
	s.AcknowledgementRequested, err = strconv.Atoi(elements[13])
	if err != nil {
		return err
	}
	s.UsageIndicator = elements[14]
	s.ComponentElementSeparator = elements[15]

	return nil
}
