package usuarios

import (
	"database/sql"
	"fmt"
	"opentaxi/graph/model"
	"strings"
)

func Update(db *sql.DB, input model.UpdateUsuario) (*model.Usuario, error) {
	if err := verificarRolesRepetidos(input.Roles); err != nil {
		return nil, err
	}
	if err := verificarUsernameYaAsignado(db, input.Username, &input.ID); err != nil {
		return nil, err
	}
	if err := verificarUserIDValido(db, input.ID); err != nil {
		return nil, err
	}
	if err := verificarRolesvalidos(db, input.Roles); err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sql := "update usuarios set nombres=?,apellidos=?,username=?,password=?,foto_url=?,telefono=?,correo=? where id=?"
	_, err = tx.Exec(sql,
		input.Nombres, input.Apellidos, input.Username, input.Password, input.FotoURL, input.Telefono, input.Correo, input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("delete from rol_usuario where usuario_id=?", input.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id := input.ID

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
