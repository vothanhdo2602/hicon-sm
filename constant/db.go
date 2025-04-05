package constant

const (
	DBPostgres  = "postgres"
	DBMysql     = "mysql"
	DBMOracle   = "oracle"
	DBSQLServer = "sqlserver"
)

// RelationType defines the type of relationship
const (
	HasOne        = "has-one"
	BelongsTo     = "belongs-to"
	HasMany       = "has-many"
	HasManyToMany = "has-many-to-many"
)
