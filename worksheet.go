package xlsx_token

import (
	"encoding/xml"
	"errors"
	"io"
	"io/fs"
	"strconv"
	"sync"
)

type WorksheetReader struct {
	file fs.File
	decoder *xml.Decoder
	cols []ColValue
	colsMutex sync.Mutex
}

type ColType int8

const (
	TypeString = iota
	TypeNumber
)

type ColValue struct {
	colType ColType
	value string
	valueId int
}

type Cell struct {
	XMLName xml.Name `xml:"c"`
	Text    string   `xml:",chardata"`
	R       string   `xml:"r,attr"`
	T       string   `xml:"t,attr"`
	S       string   `xml:"s,attr"`
	V       string   `xml:"v"`
}

func (w *WorksheetReader) NewColValue(element *xml.StartElement) (*ColValue, error) {
	cell := Cell{}
	err := w.decoder.DecodeElement(&cell, element)
	if err != nil {
		return nil, err
	}

	val := ColValue{ colType: TypeNumber }

	if cell.T == "s" && cell.V != ""{
		val.colType = TypeString

		id, err := strconv.Atoi(cell.V)
		if err != nil {
			return nil, err
		}

		val.valueId = id
	}

	return &val, err
}

func (w *WorksheetReader) goToSheetElement(localName string, position int) (bool, error) {
	currPosition := 0
	for {
		t, tokenErr := w.decoder.Token()
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

func (w *WorksheetReader) getValues() ([]ColValue, error) {
	w.decoder = xml.NewDecoder(w.file)

	found, err := w.goToSheetElement("sheetData", 0)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("cannot find sheetData")
	}

	found, err = w.goToSheetElement("row", 0)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("cannot find row")
	}

	rowCount := 0
	for {
		t, tokenErr := w.decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			}

			return nil, tokenErr
		}

		switch t := t.(type) {
		case xml.StartElement:
			if t.Name.Local == "row" {
				rowCount++
				if rowCount > 1 {
					break
				}
				continue
			}

			if t.Name.Local == "c" {
				value, tokenErr := w.NewColValue(&t)
				if tokenErr != nil {
					return nil, tokenErr
				}

				w.cols = append(w.cols, *value)
			}
		}
	}

	return w.cols, nil
}