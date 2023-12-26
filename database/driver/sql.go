package driver

import (
	"database/sql"
	"fmt"
)

func SqlConnect() (*sql.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", GetDbUser(), GetDbPassword(), GetDbHost(), GetDbPort(), GetDbName())
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}
