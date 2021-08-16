package xlsx_token

import (
	"encoding/xml"
	"errors"
	"io"
)

// GetRowColumns returns the row values.
//
// If limit is less than one, all values from the row will be returned
func (x *XlsxReader) GetRowColumns(sheet string, limit int) (cols []string, err error) {
	if sheet == "" {
		return nil, errors.New("the sheet name cannot be empty")
	}

	worksheetFile, err := x.zipReader.Open(worksheetPath(sheet))
	if err != nil {
		return nil, err
	}

	r := WorksheetReader{file: worksheetFile}

	values, err := r.getValues()
	if err != nil {
		return nil, err
	}

	for _, cell := range values {
		if cell.colType == TypeString {
			cell.value, err = x.getString(cell.valueId)
			if err != nil {
				return nil, err
			}
		}

		cols = append(cols, cell.value)
	}
	return cols, nil
}

func goToSheetElement(decoder *xml.Decoder, localName string, position int) (bool, error) {
	currPosition := 0
	for {
		t, tokenErr := decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			}

			return false, tokenErr
		}

		switch t := t.(type) {
		case xml.StartElement:
			if t.Name.Local == localName {
				if currPosition == position {
					return true, nil
				}

				currPosition++
			}
		}
	}

	return false, nil
}
