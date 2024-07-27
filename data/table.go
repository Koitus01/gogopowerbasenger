package data

// Column TODO: default values for some properties
type Column struct {
	name         string
	dataType     string
	isPrimaryKey bool
	isNullable   bool
	defaultValue interface{}
}

type ForeignKey struct {
	name               string
	referencedTable    Table
	referencedColumns  []Column
	referencingTable   []Table
	referencingColumns []Column
}

type Table struct {
	dataBaseName string
	tableName    string
	columns      []Column
	indexes      []Index
	foreignKeys  []ForeignKey
}

type Index struct {
	name    string
	columns []Column
}

func Create(databaseName string, tableName string, columns []Column, indexes []Index) (table Table, err error) {

	return Table{
		dataBaseName: databaseName,
		tableName:    tableName,
		columns:      columns,
		indexes:      indexes,
	}, nil
}
