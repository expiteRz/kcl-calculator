package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

type DB struct {
	mode     DataMode
	effectId int
}

type ColTypes struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type EffList struct {
	ef0 string `db:"ef_0"`
	ef1 string `db:"ef_1"`
	ef2 string `db:"ef_2"`
	ef3 string `db:"ef_3"`
	ef4 string `db:"ef_4"`
	ef5 string `db:"ef_5"`
	ef6 string `db:"ef_6"`
	ef7 string `db:"ef_7"`
}

// Connect to Database
func Connect() (*sql.DB, error) {
	_, err := ioutil.ReadFile("./kcl.db")
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("sqlite3", "./kcl.db")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// GetIndexes Get some data per refresh
func GetIndexes(d DB) ([]string, error) {
	var (
		q *sql.Rows
		r *sql.Row
	)

	conn, err := Connect()
	if err != nil {
		return nil, err
	}

	switch d.mode {
	case DataMode(colType):
		q, _ = conn.QueryContext(context.Background(), "select * from col_types")
	case DataMode(effType):
		r = conn.QueryRowContext(context.Background(), fmt.Sprintf("select ef_0, ef_1, ef_2, ef_3, ef_4, ef_5, ef_6, ef_7 from eff_list where id = %d", d.effectId))
	default:
		log.Fatal("Error: DataMode is out of range.")
	}

	if err != nil {
		return nil, err
	}

	var (
		c  ColTypes
		ar []ColTypes
		e  EffList
		cr []string
	)

	switch d.mode {
	case DataMode(colType):
		for q.Next() {
			err := q.Scan(&c.Id, &c.Name)
			if err != nil {
				return nil, err
			}
			ar = append(ar, c)
		}

		for _, colTypes := range ar {
			cr = append(cr, colTypes.Name)
		}
	case DataMode(effType):
		err := r.Scan(&e.ef0, &e.ef1, &e.ef2, &e.ef3, &e.ef4, &e.ef5, &e.ef6, &e.ef7)
		if err != nil {
			return nil, err
		}

		cr = append(cr, e.ef0, e.ef1, e.ef2, e.ef3, e.ef4, e.ef5, e.ef6, e.ef7)
	}

	return cr, nil
}
