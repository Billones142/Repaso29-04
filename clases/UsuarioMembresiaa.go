package clases

type UsuarioMembresia struct {
}

func (r *UsuarioMembresia) EsValido(usuario Usuario) Resultado {
	authApi1 := AuthAPI{}
	resultado := authApi1.ValidarCredenciales(usuario)
	return resultado
}
