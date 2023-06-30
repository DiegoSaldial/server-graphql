package rolpermiso

import (
	"database/sql"
	"opentaxi/graph/model"
	"strconv"
	"strings"
)

func ActualizarOnRolCreated(db *sql.Tx, permissions []*model.RolPermiso) error {
	query := "UPDATE rol_permiso SET rol_bits = CASE "
	values := []interface{}{}

	// Construir los valores y las cláusulas WHERE
	for _, permiso := range permissions {
		query += "WHEN metodo = ? THEN ? "
		values = append(values, permiso.Metodo, permiso.RolBits)
	}

	query += "END WHERE metodo IN ("
	placeholders := make([]string, len(permissions))
	for i, permiso := range permissions {
		placeholders[i] = "?"
		values = append(values, permiso.Metodo)
	}
	query += strings.Join(placeholders, ", ") + ")"

	// Ejecutar la consulta de actualización masiva
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

func ActualizarPermiso(db *sql.DB, input model.NewRolPermiso) (*model.RolPermiso, error) {
	if err := verificarPermisoNoExistente(db, input.Metodo); err != nil {
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

	sql := "update rol_permiso set rol_bits=? where metodo=?"
	_, err = db.Exec(sql, bits, input.Metodo)
	if err != nil {
		return nil, err
	}

	return GetByMetodo(db, input.Metodo)
}
