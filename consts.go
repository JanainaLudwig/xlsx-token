package xlsx_token

const (
	WorksheetFolder = "xl/worksheets"
	SharedStrings = "xl/sharedStrings.xml"
)

func worksheetPath(name string) string {
	return WorksheetFolder + "/" + name + ".xml"
}
