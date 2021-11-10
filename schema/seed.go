package schema

import (
	"database/sql"
	"fmt"
)

func GetSeed() []string {
	return []string{
		"insert into package_features (id, name) values ('39330b9b-9ad6-4900-a19d-dbb94f1b56e5', 'Custome')",
		"insert into companies (name, code, address, city_id, city, province_id, province, phone, pic, pic_phone, package_feature_id, updated_by) values ('gajahmada', 'GMD', 'jakarta', uuid_generate_v4(), 'Jakarta Timur', uuid_generate_v4(), 'Jakarta', '08174815221', 'Jacky', '08174815221', '39330b9b-9ad6-4900-a19d-dbb94f1b56e5', uuid_generate_v4())",
	}
}

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Using a constant in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *sql.DB, seeds ...string) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, seed := range seeds {
		_, err = tx.Exec(seed)
		if err != nil {
			tx.Rollback()
			fmt.Println("error execute seed")
			return err
		}
	}

	return tx.Commit()
}
