package labsession

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/6lab/6lib/labfile"
)

type SQLite struct {
	client *elastic.Client
}

func newSQlite() (sl *sql.DB) {
	dbName := xtpath.GetWorkingPath() + ""
	log.Println("DB Path :", dbName)

	// Open Database
	sl, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	// Tables & Indexes Structure
	sl.createTables()

	return sl
}

func (sl *sql.DB) createTables() {

}
