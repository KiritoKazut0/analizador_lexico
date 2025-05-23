package entities

type User struct {
	clave    int    `json:"clave"`
	nombre   string `json:"nombre"`
	correo   string `json:"correo"`
	telefono string `json:"telefono"`
}
