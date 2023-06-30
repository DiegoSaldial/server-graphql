package rolpermiso

import (
	"database/sql"
	"fmt"
	"opentaxi/graph/model"
	"strings"
)

func Listar(db *sql.DB) ([]*model.RolPermiso, error) {
	permisos := []*model.RolPermiso{}
	sql := "select metodo,rol_bits from rol_permiso"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		perm := model.RolPermiso{}
		er := parse(rows, &perm)
		if er == nil {
			permisos = append(permisos, &perm)
		} else {
			return nil, er
		}
	}
	return permisos, nil
}

func GetByMetodo(db *sql.DB, metodo string) (*model.RolPermiso, error) {
	rp := model.RolPermiso{}
	sql := "select metodo,rol_bits from rol_permiso where metodo=?"
	row := db.QueryRow(sql, metodo)
	err := parseRow(row, &rp)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}

func ListarByMetodos(db *sql.DB, metodos []string) ([]*model.RolPermiso, error) {
	permisos := []*model.RolPermiso{}
	sql := fmt.Sprintf("SELECT metodo,rol_bits FROM rol_permiso WHERE metodo IN ('%s')", strings.Join(metodos, "','"))
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		perm := model.RolPermiso{}
		er := parse(rows, &perm)
		if er == nil {
			permisos = append(permisos, &perm)
		} else {
			return nil, er
		}
	}
	return permisos, nil
}
