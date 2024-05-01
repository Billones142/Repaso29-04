package clases

import interfaces "github.com/Billones142/Repaso29-04/Interfaces"

type UsuarioMembresia struct {
}

func (r *UsuarioMembresia) EsValido(usuario interfaces.IUser) Resultado {
	authApi1 := AuthAPI{}
	resultado := authApi1.ValidarCredenciales(usuario)
	return resultado
}
