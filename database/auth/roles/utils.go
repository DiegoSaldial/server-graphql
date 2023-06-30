package roles

import (
	"database/sql"
	"errors"
	"fmt"
	"opentaxi/database/auth/rolpermiso"
	"opentaxi/graph/model"
	"strconv"
)

func parse(row *sql.Row, t *model.Rol) error {
	return row.Scan(
		&t.ID,
		&t.Nombre,
	)
}

func verificarNombreExiste(db *sql.DB, nombre string) error {
	count := 0
	sql := "select count(nombre) from roles where nombre=?"
	db.QueryRow(sql, nombre).Scan(&count)
	if count > 0 {
		return errors.New("nombre de rol ya existe")
	}
	return nil
}

func verificarExisteRol(db *sql.DB, nombre string) error {
	count := 0
	sql := "select count(nombre) from roles where nombre=?"
	db.QueryRow(sql, nombre).Scan(&count)
	if count == 0 {
		return errors.New("el rol no existe")
	}
	return nil
}

func verificarPermisosExistentes(db *sql.DB, permisos []string) error {
	existingPermissions := make(map[string]bool)
	rows, err := db.Query("SELECT metodo FROM rol_permiso")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Almacenar los permisos existentes en el mapa
	for rows.Next() {
		var metodo string
		if err := rows.Scan(&metodo); err != nil {
			return err
		}
		existingPermissions[metodo] = true
	}

	// Verificar los permisos en el array
	ausentes := []string{}
	for _, permission := range permisos {
		if _, ok := existingPermissions[permission]; !ok {
			ausentes = append(ausentes, permission)
		}
	}

	if len(ausentes) == 0 {
		return nil
	}

	rep := fmt.Sprintf("%+v", ausentes)
	return errors.New("permisos no existentes: " + rep)
}

func verificarMaxRegistros(db *sql.DB) error {
	count := 0
	sql := "select count(nombre) from roles"
	db.QueryRow(sql).Scan(&count)
	max := 14
	if count >= max {
		return errors.New("se ha alzanzado el maximo de registros de roles permitido")
	}
	return nil
}

func asignarAPermisos(tx *sql.Tx, db *sql.DB, input model.NewRol, idrol int64) (*model.Rol, error) {
	permisos, err := rolpermiso.ListarByMetodos(db, input.Permisos)
	if err != nil {
		return nil, err
	}
	for _, p := range permisos {
		rolbit := 1 << (idrol - 1)
		p.RolBits += rolbit
	}

	er := rolpermiso.ActualizarOnRolCreated(tx, permisos)
	if er != nil {
		return nil, er
	}
	return nil, nil
}

func quitardePermisos(db *sql.DB, tx *sql.Tx, rol *model.Rol) (bool, error) {
	permisos, err := rolpermiso.Listar(db)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	for _, p := range permisos {
		idrol, _ := strconv.Atoi(rol.ID)
		rolbit := 1 << (idrol - 1)
		permitido := p.RolBits&rolbit == rolbit
		if permitido {
			p.RolBits -= rolbit
		}

	}

	er := rolpermiso.ActualizarOnRolCreated(tx, permisos)
	if er != nil {
		tx.Rollback()
		return false, er
	}
	return true, nil
}
