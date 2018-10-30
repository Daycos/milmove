package iws

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func buildWkEmaURL(host string, custNum string, workEmail string) (string, error) {
	if !emailRegex.MatchString(workEmail) {
		return "", errors.New("Invalid e-mail address")
	}

	// e-mail addresses are limited to 80 characters
	l := len(workEmail)
	if l > 80 {
		l = 80
	}

	return fmt.Sprintf(
		"https://%s/appj/rbs/rest/op=wkEma/customer=%s/schemaName=get_cac_data/schemaVersion=1.0/EMA_TX=%s",
		host, custNum, workEmail[:l]), nil
}

func parseWkEmaResponse(data []byte) (uint64, *Person, []Personnel, error) {
	rec := Record{}
	unmarshalErr := xml.Unmarshal(data, &rec)
	if unmarshalErr != nil {
		// Couldn't unmarshal as a record, try as an RbsError next
		rbsError := RbsError{}
		unmarshalErr = xml.Unmarshal([]byte(data), &rbsError)
		if unmarshalErr == nil {
			return 0, nil, []Personnel{}, &rbsError
		}
		return 0, nil, []Personnel{}, unmarshalErr
	}

	// Not found
	if rec.AdrRecord.WorkEmail == nil {
		return 0, nil, []Personnel{}, nil
	}

	return rec.AdrRecord.WorkEmail.Edipi, rec.AdrRecord.Person, rec.AdrRecord.Personnel, nil
}

// GetPersonUsingWorkEmail retrieves personal information (including SSN and EDIPI) through the IWS:RBS REST API using a work e-mail address.
// If matched succesfully, it returns the EDIPI, the full name and SSN information, and the personnel information for each of the organizations the person belongs to
func GetPersonUsingWorkEmail(client http.Client, host string, custNum string, workEmail string) (uint64, *Person, []Personnel, error) {
	url, urlErr := buildWkEmaURL(host, custNum, workEmail)
	if urlErr != nil {
		return 0, nil, []Personnel{}, urlErr
	}

	resp, getErr := client.Get(url)
	// Interesting fact: RBS responds 200 OK, not 404 Not Found, if there are no matches
	if getErr != nil {
		return 0, nil, []Personnel{}, getErr
	}

	defer resp.Body.Close()
	data, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return 0, nil, []Personnel{}, readErr
	}

	return parseWkEmaResponse(data)
}
