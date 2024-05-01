package clases

import interfaces "github.com/Billones142/Repaso29-04/Interfaces"

type AuthAPI struct {
}

func (r *AuthAPI) ValidarCredenciales(usuario interfaces.IUser) Resultado {
	db1 := BaseDeDatos{}.New()

	usuarioDb1 := db1.ObtenerUsuario(usuario.GetNombreUsuario())
	if usuario.GetPasswordHash() == usuarioDb1.GetPasswordHash() {
		return Resultado{Exito: true}
	} else {
		return Resultado{Exito: false,
			Mensaje: "Credenciales invalidas"}
	}
}
