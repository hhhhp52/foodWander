package migrations

import "database/sql"

func CreateMigrations() {
	tx, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")
	CheckErr(err)
	_, err = tx.Exec("CREATE TABLE migrations (id serial PRIMARY KEY, name VARCHAR (50) NOT NULL);")
	CheckErr(err)
}
