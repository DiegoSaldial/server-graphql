package roles

import (
	"database/sql"
)

func Eliminar(db *sql.DB, rolname string) (bool, error) {
	if err := verificarExisteRol(db, rolname); err != nil {
		return false, err
	}
	rol, err := GetByName(db, rolname)
	if err != nil {
		return false, err
	}

	tx, err := db.Begin()
	if err != nil {
		return false, err
	}

	sql := "delete from roles where nombre=?"
	_, err = tx.Exec(sql, rol.Nombre)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = quitardePermisos(db, tx, rol)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}
