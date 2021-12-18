package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "postgres"
const DbName = "postgres"

var db sql.DB

func ObterConexao() *sql.DB {
	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
	conexao, err := sql.Open(PostgresDriver, DataSourceName)
	if err != nil {
		log.Panicln(err.Error())
	}
	return conexao
}
