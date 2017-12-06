package webapp

import (
	"database/sql"
)

// DataBase stores a pointer to the SQL Database.
var DataBase struct {
	DB *sql.DB
}
