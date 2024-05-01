package clases

import interfaces "github.com/Billones142/Repaso29-04/Interfaces"

type BaseDeDatos struct {
	usuarios []interfaces.IUser
}

func (BaseDeDatos) New() BaseDeDatos {
	//se crea una base que ya contiene los usuario
	//pero en realidad aqui deberia estar la logica que se comunica con la base de datos real
	return BaseDeDatos{usuarios: []interfaces.IUser{Usuario{}.New("user1", "password1")}}
}

func (r *BaseDeDatos) ObtenerUsuario(nombreUsuario string) interfaces.IUser {
	var usuarioEncontrado interfaces.IUser
	for _, usuario := range r.usuarios {
		if usuario.GetNombreUsuario() == nombreUsuario {
			usuarioEncontrado = usuario
		}
	}
	return usuarioEncontrado
}
