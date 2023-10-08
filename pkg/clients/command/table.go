package command

import (
	"fmt"
	"strings"
)

type Table struct {
	columnNames []string
	rows        [][]string
}

func NewTable(columnNames []string) *Table {
	return &Table{
		columnNames: columnNames,
		rows:        make([][]string, 0),
	}
}

func (t *Table) AddRow(row []string) error {
	if len(row) != len(t.columnNames) {
		return fmt.Errorf("expected %d columns, got %d", len(t.columnNames), len(row))
	}
	t.rows = append(t.rows, row)
	return nil
}

func (t *Table) String() string {
	maxLengths := make([]int, len(t.columnNames))
	for i := range t.columnNames {
		maxLengths[i] = len(t.columnNames[i])
		for j := 0; j < len(t.rows); j++ {
			if len(t.rows[j][i]) > maxLengths[i] {
				maxLengths[i] = len(t.rows[j][i])
			}
		}
	}

	tableStr := "|"
	for i, name := range t.columnNames {
		tableStr += fmt.Sprintf("%s%s", name, strings.Repeat(" ", maxLengths[i]-len(name)+1))
		tableStr += "|"
	}
	tableStr += "\n" + strings.Repeat("-", len(tableStr)) + "\n"

	for _, row := range t.rows {
		tableStr += "|"
		for i, col := range row {
			tableStr += fmt.Sprintf("%s%s", col, strings.Repeat(" ", maxLengths[i]-len(col)+1))
			tableStr += "|"
		}
		tableStr += "\n"
	}
	tableStr += strings.Repeat("-", len(tableStr)) + "\n"
	return tableStr
}
