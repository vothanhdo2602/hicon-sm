package constant

// DBType defines the database type
type DBType string

const (
	DBPostgres  DBType = "postgres"
	DBMysql     DBType = "mysql"
	DBMOracle   DBType = "oracle"
	DBSQLServer DBType = "sqlserver"
)

// RelationType defines the type of relationship
type RelationType string

const (
	HasOne        RelationType = "has-one"
	BelongsTo     RelationType = "belongs-to"
	HasMany       RelationType = "has-many"
	HasManyToMany RelationType = "m2m"
)

type ColumnType string

const (
	ColumnTypeString    ColumnType = "string" // "text", "varchar", "char"
	ColumnTypeTimestamp ColumnType = "timestamp"
	ColumnTypeInteger   ColumnType = "int"
	ColumnTypeFloat     ColumnType = "float"
	ColumnTypeBoolean   ColumnType = "boolean"
	ColumnTypeJSON      ColumnType = "json"
)
