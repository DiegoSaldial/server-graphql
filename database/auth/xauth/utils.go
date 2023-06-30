package xauth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"opentaxi/graph/model"
	"os"
)

func parse(rows *sql.Rows, t *model.UsuarioLogin) error {
	return rows.Scan(
		&t.ID,
		&t.Nombres,
		&t.Apellidos,
		&t.Username,
		&t.FotoURL,
		&t.Telefono,
		&t.Correo,
		&t.Registrado,
		&t.Estado,
		&t.Rol,
		&t.RolID,
	)
}
func parseRoles(rows *sql.Rows, t *model.Rol) error {
	return rows.Scan(
		&t.ID,
		&t.Nombre,
	)
}

func encrypt(message string) (string, error) {
	aes_key := os.Getenv("AES_KEY")
	key := []byte(aes_key)
	byteMsg := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decrypt(message string) (string, error) {
	aes_key := os.Getenv("AES_KEY")
	key := []byte(aes_key)
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

// func parseUs(rows *sql.Rows, t *model.Usuario) error {
// 	return rows.Scan(
// 		&t.ID,
// 		&t.Nombres,
// 		&t.Apellidos,
// 		&t.Username,
// 		&t.FotoURL,
// 		&t.Telefono,
// 		&t.Correo,
// 		&t.Registrado,
// 		&t.Estado,
// 	)
// }
