package rolpermiso

import (
	"database/sql"
	"errors"
	"fmt"
	"opentaxi/graph/model"
	"strings"
)

func parse(rows *sql.Rows, t *model.RolPermiso) error {
	return rows.Scan(
		&t.Metodo,
		&t.RolBits,
	)
}
func parseRow(rows *sql.Row, t *model.RolPermiso) error {
	return rows.Scan(
		&t.Metodo,
		&t.RolBits,
	)
}

func verificarPermisoExistente(db *sql.DB, metodo string) error {
	count := 0
	db.QueryRow("select count(metodo) from rol_permiso where metodo=?", metodo).Scan(&count)
	if count == 0 {
		return nil
	}
	return errors.New("el metodo ya se encuentra registrado previamente")
}

func verificarPermisoNoExistente(db *sql.DB, metodo string) error {
	count := 0
	db.QueryRow("select count(metodo) from rol_permiso where metodo=?", metodo).Scan(&count)
	if count == 0 {
		return errors.New("el metodo no se encuentra registrado previamente")
	}
	return nil
}

func verificarRolesExistentes(db *sql.DB, roles []string) error {
	existingPermissions := make(map[string]bool)
	rows, err := db.Query("SELECT id FROM roles")
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
	for _, permission := range roles {
		if _, ok := existingPermissions[permission]; !ok {
			ausentes = append(ausentes, permission)
		}
	}

	if len(ausentes) == 0 {
		return nil
	}

	rep := fmt.Sprintf("%+v", ausentes)
	return errors.New("roles no existentes: " + rep)
}

func obtenerRolesById(db *sql.DB, rolesid []string) ([]*model.Rol, error) {
	query := "SELECT id,nombre FROM roles WHERE id IN (%s)"
	placeholders := make([]string, len(rolesid))
	args := make([]interface{}, len(rolesid))
	for i, v := range rolesid {
		placeholders[i] = "?"
		args[i] = v
	}
	query = fmt.Sprintf(query, strings.Join(placeholders, ", "))

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*model.Rol{}
	for rows.Next() {
		rol := model.Rol{}
		if err := rows.Scan(&rol.ID, &rol.Nombre); err != nil {
			return nil, err
		}
		roles = append(roles, &rol)
	}
	return roles, nil
}
