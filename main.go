package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/twinj/uuid"
	"gopkg.in/gorp.v1"
	"log"
	"net/http"
)

type Record struct {
	UUID    string
	Created int64
	URLs    string
}

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	uuid.SwitchFormat(uuid.CleanHyphen)

	r := mux.NewRouter()
	r.HandleFunc("/{id:[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}}/", SplitHandler).Methods("POST")
	r.HandleFunc("/edit/{id:[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}}/", EditHandler)
	r.HandleFunc("/", AddHandler)

	http.ListenAndServe(":8080", r)
}

func SplitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
}

// func other() {
// 	var records []Record
// 	_, err := dbmap.Select(&records, "select * from records order by created")
// 	checkErr(err, "select * failed")

// 	fmt.Println(records)

// 	fmt.Println(uuid.NewV4().String())
// }

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "/tmp/db.sqlite3")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Record{}, "records").SetKeys(false, "UUID")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
