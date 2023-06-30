package roles

import (
	"database/sql"
	"opentaxi/graph/model"
)

func Crear(db *sql.DB, input model.NewRol) (*model.Rol, error) {
	if err := verificarNombreExiste(db, input.Nombre); err != nil {
		return nil, err
	}
	if err := verificarPermisosExistentes(db, input.Permisos); err != nil {
		return nil, err
	}
	if err := verificarMaxRegistros(db); err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := "insert into roles(nombre) values(?)"
	res, err := tx.Exec(sql, input.Nombre)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	idrol, _ := res.LastInsertId()

	_, err = asignarAPermisos(tx, db, input, idrol)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return GetByName(db, input.Nombre)
}
