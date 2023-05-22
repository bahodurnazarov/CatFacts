package bot

import (
	d "github.com/bahodurnazarov/CatFacts/pkg/db"
	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
)

func InsertToDB(factEN, factRU string) {
	db := d.ConnDB()
	if factEN != "" {
		_, err := db.Exec("INSERT into facts VALUES ($1, $2)", factEN, factRU)
		if err != nil {
			lg.Errl.Fatalf("111An error occured while executing query: %v", err)
		}
	}

}
