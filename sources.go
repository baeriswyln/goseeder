package goseeder

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

var dataPath = "db/seeds/data"

//SetDataPath this will allow you to specify a custom path where your seed data is located
func SetDataPath(path string) {
	dataPath = path
}

//FromJson inserts into a database table with name same as the filename all the json entries
func (s Seeder) FromJson(filename string) {
	var folder = ""
	if s.context.env != "" {
		folder = fmt.Sprintf("%s/", s.context.env)
	}
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s%s.json", dataPath, folder, filename))
	if err != nil {
		log.Fatal(err)
	}

	m := []map[string]interface{}{}
	err = json.Unmarshal(content, &m)
	if err != nil {
		panic(err)
	}

	for _, e := range m {
		stmQuery, args := prepareStatement(filename, e)
		stmt, _ := s.DB.Prepare(stmQuery.String())
		_, err := stmt.Exec(args...)
		if err != nil {
			panic(err)
		}
	}
}
