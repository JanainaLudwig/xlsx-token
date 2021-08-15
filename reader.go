package xlsx_token

import "errors"

// GetRowColumns returns the row values.
//
// If limit is less than one, all values from the row will be returned
func (x *XlsxReader) GetRowColumns(sheet string, limit int) (cols []ColValue, err error) {
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

	return values, nil
}