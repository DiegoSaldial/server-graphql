package roles

import (
	"database/sql"
	"opentaxi/graph/model"
)

func GetByName(db *sql.DB, id string) (*model.Rol, error) {
	sql := "select id,nombre from roles where nombre=?"
	row := db.QueryRow(sql, id)
	rol := model.Rol{}
	er := parse(row, &rol)
	if er != nil {
		return nil, er
	}
	return &rol, nil
}
