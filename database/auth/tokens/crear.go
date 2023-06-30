package tokens

import "database/sql"

func Guardar(db *sql.DB, token string, username string) error {
	sql := "replace into tokens(token,username) values(?,?)"
	_, err := db.Exec(sql, token, username)
	return err
}
