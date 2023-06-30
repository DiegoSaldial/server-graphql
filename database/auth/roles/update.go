package roles

import (
	"database/sql"
	"opentaxi/graph/model"
	"strconv"
)

func Modificar(db *sql.DB, input model.NewRol) (*model.Rol, error) {
	rol, err := GetByName(db, input.Nombre)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = quitardePermisos(db, tx, rol)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	tx, err = db.Begin()
	if err != nil {
		return nil, err
	}

	idrol, _ := strconv.ParseInt(rol.ID, 10, 32)

	_, err = asignarAPermisos(tx, db, input, idrol)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return rol, nil
}
