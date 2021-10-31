package main

import "database/sql"

type db struct {
}

func (db) connect() *sql.DB {
	sql, err := sql.Open("sqlite3", "./kcl.db")
	if err != nil {
		return nil
	}
	return sql
}
