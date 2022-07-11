package db_mysql

import "database/sql"

var DB *sql.DB

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1)/db_amosblog?charset=utf8mb4&parseTime=True"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
