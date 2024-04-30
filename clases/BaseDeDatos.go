package clases

type BaseDeDatos struct {
	usuarios []Usuario
}

func (BaseDeDatos) New() BaseDeDatos {
	//se crea una base que ya contiene los usuario
	//pero en realidad aqui deberia estar la logica que se comunica con la base de datos real
	return BaseDeDatos{usuarios: []Usuario{Usuario{}.New("user1", "password1")}}
}

func (r *BaseDeDatos) ObtenerUsuario(nombreUsuario string) Usuario {
	var usuarioEncontrado Usuario
	for _, usuario := range r.usuarios {
		if usuario.NombreUsuario == nombreUsuario {
			usuarioEncontrado = usuario
		}
	}
	return usuarioEncontrado
}
