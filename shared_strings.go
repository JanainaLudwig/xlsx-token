package xlsx_token

import (
	"encoding/xml"
	"errors"
	"io"
)

type Si struct {
	XMLName xml.Name `xml:"si"`
	Text    string   `xml:",chardata"`
	T       string   `xml:"t"`
}

func (x *XlsxReader) getString(id int) (string, error) {
	val, ok := x.sharedStrings[id]
	if ok {
		return val, nil
	}

	if x.sharedStringsDecoder == nil {
		open, err := x.zipReader.Open(SharedStrings)
		if err != nil {
			return "", err
		}

		x.sharedStringsDecoder = xml.NewDecoder(open)


		found, err := goToSheetElement(x.sharedStringsDecoder, "sst", 0)
		if err != nil {
			return "", err
		}

		if !found {
			return "nil", errors.New("cannot find sst")
		}

	}

	var strVal string

	for x.sharedStringIndex <= id {
		t, tokenErr := x.sharedStringsDecoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			}

			return "", tokenErr
		}

		switch t := t.(type) {
		case xml.StartElement:
			if t.Name.Local != "si" {
				continue
			}

			var si Si
			tokenErr := x.sharedStringsDecoder.DecodeElement(&si, &t)
			if tokenErr != nil {
				return "", tokenErr
			}

			strVal = si.T

			x.sharedStrings[x.sharedStringIndex] = strVal

			x.sharedStringIndex++
		}
	}

	return strVal, nil
}