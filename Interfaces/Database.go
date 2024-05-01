package interfaces

type IDatabase interface {
	ObtenerUsuario(nombreUsuario string) IUser
}
