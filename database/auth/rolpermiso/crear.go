package rolpermiso

import (
	"database/sql"
	"opentaxi/graph/model"
	"strconv"
)

func CrearNuevoPermiso(db *sql.DB, input model.NewRolPermiso) (*model.RolPermiso, error) {
	if err := verificarPermisoExistente(db, input.Metodo); err != nil {
		return nil, err
	}
	if err := verificarRolesExistentes(db, input.Roles); err != nil {
		return nil, err
	}

	roles, err := obtenerRolesById(db, input.Roles)
	if err != nil {
		return nil, err
	}

	bits := 0
	for _, p := range roles {
		idrol, _ := strconv.Atoi(p.ID)
		rolbit := 1 << (idrol - 1)
		bits += rolbit
	}

	sql := "insert into rol_permiso(metodo,rol_bits) values(?,?)"
	_, err = db.Exec(sql, input.Metodo, bits)
	if err != nil {
		return nil, err
	}

	return GetByMetodo(db, input.Metodo)
}
