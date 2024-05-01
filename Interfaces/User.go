package interfaces

type IUser interface {
	New(NombreUsuario string, Password string) IUser
	GetNombreUsuario() string
	SetNombreUsuario(n string)
	GetPasswordHash() string
}
