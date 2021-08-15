package xlsx_token

const (
	WorksheetFolder = "xl/worksheets"
)

func worksheetPath(name string) string {
	return WorksheetFolder + "/" + name + ".xml"
}
