package usuarios

import (
	"database/sql"
	"errors"
	"opentaxi/graph/model"
)

func UsuarioByUsername(db *sql.DB, username string) (*model.Usuario, error) {
	sql := "select id,nombres,apellidos,username,foto_url,telefono,correo,registrado,estado from usuarios where username=?"
	row := db.QueryRow(sql, username)
	usuario := model.Usuario{}
	parse(row, &usuario)
	if usuario.ID == "" {
		return nil, errors.New("usuario no encontrado")
	}
	return &usuario, nil
}
