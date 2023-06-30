// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewRol struct {
	Nombre   string   `json:"nombre"`
	Permisos []string `json:"permisos"`
}

type NewRolPermiso struct {
	Metodo string   `json:"metodo"`
	Roles  []string `json:"roles"`
}

type NewUsuario struct {
	Nombres   string   `json:"nombres"`
	Apellidos string   `json:"apellidos"`
	Username  string   `json:"username"`
	Password  *string  `json:"password,omitempty"`
	FotoURL   *string  `json:"foto_url,omitempty"`
	Telefono  *string  `json:"telefono,omitempty"`
	Correo    *string  `json:"correo,omitempty"`
	Roles     []string `json:"roles"`
}

type Rol struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

type RolPermiso struct {
	Metodo  string `json:"metodo"`
	RolBits int    `json:"rol_bits"`
}

type Tokens struct {
	Username   string `json:"username"`
	Token      string `json:"token"`
	Registrado string `json:"registrado"`
}

type UpdateUsuario struct {
	ID        string   `json:"id"`
	Nombres   string   `json:"nombres"`
	Apellidos string   `json:"apellidos"`
	Username  string   `json:"username"`
	Password  *string  `json:"password,omitempty"`
	FotoURL   *string  `json:"foto_url,omitempty"`
	Telefono  *string  `json:"telefono,omitempty"`
	Correo    *string  `json:"correo,omitempty"`
	Roles     []string `json:"roles"`
}

type Usuario struct {
	ID         string  `json:"id"`
	Nombres    string  `json:"nombres"`
	Apellidos  string  `json:"apellidos"`
	Username   string  `json:"username"`
	FotoURL    *string `json:"foto_url,omitempty"`
	Telefono   *string `json:"telefono,omitempty"`
	Correo     *string `json:"correo,omitempty"`
	Registrado string  `json:"registrado"`
	Estado     bool    `json:"estado"`
}

type UsuarioLogin struct {
	ID         string  `json:"id"`
	Nombres    string  `json:"nombres"`
	Apellidos  string  `json:"apellidos"`
	Username   string  `json:"username"`
	FotoURL    *string `json:"foto_url,omitempty"`
	Telefono   *string `json:"telefono,omitempty"`
	Correo     *string `json:"correo,omitempty"`
	Registrado string  `json:"registrado"`
	Estado     bool    `json:"estado"`
	Dataname   string  `json:"dataname"`
	RolID      string  `json:"rol_id"`
	Rol        string  `json:"rol"`
	Exp        string  `json:"exp"`
}
