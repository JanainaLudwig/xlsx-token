package xlsx_token

import (
	"log"
	"testing"
)

func TestXlsxReader_GetRowColumns(t *testing.T) {
	type fields struct {
		file string
	}
	type args struct {
		sheet string
		limit int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCols []string
		wantErr  bool
	}{
		{
			name:     "",
			fields:   fields{
				file: "test_files/test_one.xlsx",
			},
			args:     args{},
			wantCols: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err2 := Open(tt.fields.file)
			if err2 != nil {
				t.Fatal(err2)
			}

			gotCols, err := reader.GetRowColumns(tt.args.sheet, tt.args.limit)
			log.Println(gotCols, err)
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("GetRowColumns() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//
			//if !reflect.DeepEqual(gotCols, tt.wantCols) {
			//	t.Errorf("GetRowColumns() gotCols = %v, want %v", gotCols, tt.wantCols)
			//}
		})
	}
}
