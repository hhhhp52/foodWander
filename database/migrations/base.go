package migrations

import (
	"database/sql"
)

func DB() {
	_, err := sql.Open("postgres", "user=astaxie password=astaxie dbname=test sslmode=disable")
	CheckErr(err)

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
