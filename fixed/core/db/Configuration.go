package YamiDB

import (
	"Yami/core/models/Json"

	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var SQL *sql.DB

func Connection() error {
	Conn, err := sql.Open("mysql", JsonParse.ConfigSyncs.SQL.SQLUsername+":"+JsonParse.ConfigSyncs.SQL.SQLPassword+"@tcp("+JsonParse.ConfigSyncs.SQL.SQLHost+")/"+JsonParse.ConfigSyncs.SQL.SQLName); if err != nil {
		return err
	}

	SQL = Conn
	return nil
}