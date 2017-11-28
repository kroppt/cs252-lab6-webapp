package webapp

import (
	"database/sql"
)

type DataBase struct {
	*sql.DB
}

var DB *DataBase
