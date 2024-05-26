package data

import "testing"

func TestCreateTable(t *testing.T) {
	columns := []Column{
		{
			name:         "jij",
			dataType:     "juj",
			isPrimaryKey: false,
			defaultValue: nil,
		},
		{
			name:         "pis",
			dataType:     "pus",
			isPrimaryKey: false,
			defaultValue: nil,
		},
	}

	t.Log(columns)
}
