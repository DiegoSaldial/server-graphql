package usuarios

import (
	"database/sql"
	"fmt"
	"opentaxi/graph/model"
	"strings"
)

func Crear(db *sql.DB, input model.NewUsuario) (*model.Usuario, error) {
	if err := verificarRolesRepetidos(input.Roles); err != nil {
		return nil, err
	}
	if err := verificarUsernameYaAsignado(db, input.Username, nil); err != nil {
		return nil, err
	}
	if err := verificarRolesvalidos(db, input.Roles); err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := "insert into usuarios(nombres,apellidos,username,password,foto_url,telefono,correo) values(?,?,?,?,?,?,?)"
	res, err := tx.Exec(sql,
		input.Nombres, input.Apellidos, input.Username, input.Password, input.FotoURL, input.Telefono, input.Correo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := res.LastInsertId()
	sql = "insert into rol_usuario(usuario_id,rol_id) values %s"
	places := make([]string, len(input.Roles))
	args := make([]interface{}, len(input.Roles)*2)
	for i, v := range input.Roles {
		places[i] = "(?,?)"
		args[i*2] = id
		args[i*2+1] = v
	}

	sql = fmt.Sprintf(sql, strings.Join(places, ", "))
	_, err = tx.Exec(sql, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return UsuarioByUsername(db, input.Username)
}
