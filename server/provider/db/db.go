package db

import (
	"database/sql"
	"fmt"
	"server/server/config"
)

func Connect(conf *config.Config) (*sql.DB, error) {
	connSTR := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.UserBalanceDB,
		conf.PasswordBalanceDB, conf.NameDB)
	db, err := sql.Open("postgres", connSTR)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
