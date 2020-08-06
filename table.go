package prettytable

import (
	"github.com/bingoohuang/prettytable/pkg/table"
	"reflect"
	"regexp"
)

// TablePrinter ...
type TablePrinter struct {
	DittoMark          string
	SingleRowTranspose bool
	NoPrintRowSeq      bool
	TagName            string
}

type tableInfo struct {
	option       TablePrinter
	structType   reflect.Type
	structFields []reflect.StructField
	header       table.Row
	rows         []table.Row
	singleRow    bool
}

func (p TablePrinter) Print(value interface{}) string {
	if value == nil {
		return ""
	}

	if p.TagName == "" {
		p.TagName = "table"
	}

	v := reflect.ValueOf(value)
	tableInfo := tableInfo{
		option: p,
	}

	if !p.NoPrintRowSeq {
		tableInfo.header = append(tableInfo.header, "#")
	}

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}

		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		tableInfo.structType = v.Type()
		tableInfo.singleRow = true
		tableInfo.createHeader()
		tableInfo.createRow(0, v)
	case reflect.Slice:
		if v.Len() == 0 {
			return ""
		}

		tableInfo.structType = v.Type().Elem()
		tableInfo.singleRow = v.Len() == 1
		tableInfo.createHeader()

		for i := 0; i < v.Len(); i++ {
			tableInfo.createRow(i, v.Index(i))
		}
	default:
		return ""
	}

	return tableInfo.tableRender(p)
}

func (t *tableInfo) createRow(rowIndex int, v reflect.Value) {
	row := make(table.Row, 0)

	if !t.option.NoPrintRowSeq {
		row = append(row, rowIndex+1)
	}

	for _, f := range t.structFields {
		row = append(row, v.FieldByIndex(f.Index).Interface())
	}

	t.rows = append(t.rows, row)
}

func (t *tableInfo) createHeader() {
	for i := 0; i < t.structType.NumField(); i++ {
		f := t.structType.Field(i)
		if f.PkgPath != "" {
			continue
		}

		tag := f.Tag.Get(t.option.TagName)
		if tag == "-" {
			continue
		}

		title := tag
		if title == "" {
			title = BlankCamel(f.Name)
		}

		t.header = append(t.header, title)
		t.structFields = append(t.structFields, f)
	}
}

func (t *tableInfo) tableRender(p TablePrinter) string {
	if t.option.SingleRowTranspose && t.singleRow {
		return t.tableRenderSingleRow()
	}

	return t.tableRenderNormal(p)
}

func (t *tableInfo) tableRenderSingleRow() string {
	w := table.NewWriter()
	w.SetStyle(table.StyleLight)

	if !t.option.NoPrintRowSeq {
		w.AppendHeader(table.Row{"#", "Key", "Value"})
		row := t.rows[0][1:]
		head := t.header[1:]

		for i := 0; i < len(head); i++ {
			w.AppendRow(table.Row{i + 1, head[i], row[i]})
		}
	} else {
		w.AppendHeader(table.Row{"Key", "Value"})
		row := t.rows[0]
		head := t.header

		for i := 0; i < len(head); i++ {
			w.AppendRow(table.Row{head[i], row[i]})
		}
	}

	return w.Render()
}

func (t *tableInfo) tableRenderNormal(p TablePrinter) string {
	w := table.NewWriter()
	w.SetStyle(table.StyleLight)
	w.AppendHeader(t.header)

	if p.DittoMark != "" {
		w.AppendRows(p.dittoMarkRows(t.rows))
	} else {
		w.AppendRows(t.rows)
	}

	return w.Render()
}

func (p TablePrinter) dittoMarkRows(rows []table.Row) []table.Row {
	mark := make(map[int]interface{})

	for i, row := range rows {
		for j, cell := range row {
			v, ok := mark[j]
			if ok && v != "" && v == cell {
				rows[i][j] = p.DittoMark
			} else {
				mark[j] = cell
			}
		}
	}

	return rows
}

// BlankCamel ...
func BlankCamel(str string) string {
	blank := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(str, "${1} ${2}")
	return regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(blank, "${1} ${2}")
}
