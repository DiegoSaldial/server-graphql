package xauth

import (
	"database/sql"
	"errors"
	"fmt"
)

func CheckPermiso(db *sql.DB, metodo string, rolid int, rol string) error {
	rol_bit := 1 << (rolid - 1)
	rol_bits := 0
	sql := "select rol_bits from rol_permiso where metodo=?"
	db.QueryRow(sql, metodo).Scan(&rol_bits)
	permitido := rol_bits&rol_bit == rol_bit
	if permitido {
		return nil
	}
	texto := fmt.Sprintf("acceso denegado: solicitado: %d, devuelto: %d, para tu rol: %s", rol_bit, rol_bits, rol)
	return errors.New(texto)
}
