package config

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/golang-sql/civil"
)

func GetMySQLDB() (db *sql.DB, err error) {

	//dbDriver := "sqlserver"
	dbUser := "agoraopsportaluser"
	dbPassword := "#evd@Ops21portal"
	DbServe := "agoraopsportaldbserver-dev.database.windows.net"
	DbPort := "1433"
	DbName := "agoraopsportaldbdev"

	constring := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", DbServe, dbUser, dbPassword, DbPort, DbName)
	db, err = sql.Open("sqlserver", constring)
	if err != nil {
		panic(err.Error())
	}
	return
}
