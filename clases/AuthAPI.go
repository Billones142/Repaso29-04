package clases

type AuthAPI struct {
}

func (r *AuthAPI) ValidarCredenciales(usuario Usuario) Resultado {
	db1 := BaseDeDatos{}.New()

	usuarioDb1 := db1.ObtenerUsuario(usuario.NombreUsuario)
	if usuario.PasswordHash == usuarioDb1.PasswordHash {
		return Resultado{Exito: true}
	} else {
		return Resultado{Exito: false,
			Mensaje: "Credenciales invalidas"}
	}
}
