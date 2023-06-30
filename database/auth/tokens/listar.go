package tokens

import (
	"database/sql"
	"opentaxi/graph/model"
)

func GetToken(db *sql.DB, username string) *model.Tokens {
	r := model.Tokens{}
	sql := "select token,username from tokens where username=?"
	row := db.QueryRow(sql, username)
	row.Scan(
		&r.Token,
		&r.Username,
	)
	return &r
}
