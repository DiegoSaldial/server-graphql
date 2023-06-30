package rolpermiso

import (
	"database/sql"
	"errors"
)

func EliminarPermiso(db *sql.DB, metodo string) (bool, error) {
	sql := "delete from rol_permiso where metodo=?"
	res, err := db.Exec(sql, metodo)
	if err != nil {
		return false, nil
	}
	dels, err := res.RowsAffected()
	if err != nil {
		return false, nil
	}
	if dels == 0 {
		return false, errors.New("no se ha eliminado ningun registro")
	}
	return true, nil
}
