package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTable(t *testing.T) {
	assert := assert.New(t)
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
	var indexes []Index

	cock, _ := Create("pipi", "pupu", columns, indexes)

	assert.Equal("pipi", cock.dataBaseName)
	assert.Equal("pupu", cock.tableName)
	assert.Len(cock.columns, len(columns))
	assert.Len(cock.indexes, len(indexes))
}
