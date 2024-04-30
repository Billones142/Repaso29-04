package clases

import (
	"crypto/sha1"
	"encoding/base64"
)

type Usuario struct {
	NombreUsuario string
	password      string
	PasswordHash  string
}

func (Usuario) New(NombreUsuario string, Password string) Usuario {
	hasher := sha1.New()
	hasher.Write([]byte(Password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return Usuario{NombreUsuario: NombreUsuario,
		password:     Password,
		PasswordHash: sha,
	}
}