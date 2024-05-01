package clases

import (
	"crypto/sha1"
	"encoding/base64"

	interfaces "github.com/Billones142/Repaso29-04/Interfaces"
)

type Usuario struct {
	NombreUsuario string
	password      string
	PasswordHash  string
}

func (Usuario) New(NombreUsuario string, Password string) interfaces.IUser {
	hasher := sha1.New()
	hasher.Write([]byte(Password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return &Usuario{NombreUsuario: NombreUsuario,
		password:     Password,
		PasswordHash: sha,
	}
}

func (r *Usuario) GetNombreUsuario() string {
	return r.NombreUsuario
}

func (r *Usuario) SetNombreUsuario(n string) {
	r.NombreUsuario = n
}

func (r *Usuario) getPassword() string {
	return r.password
}

func (r *Usuario) setPassword(n string) {
	r.password = n
}

func (r *Usuario) GetPasswordHash() string {
	return r.PasswordHash
}

func (r *Usuario) setPasswordHash(n string) {
	r.PasswordHash = n
}
