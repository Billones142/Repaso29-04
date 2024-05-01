package main_test

import (
	"testing"

	interfaces "github.com/Billones142/Repaso29-04/Interfaces"
	clases "github.com/Billones142/Repaso29-04/clases"
)

func Test_Valido(t *testing.T) {
	membresia1 := clases.UsuarioMembresia{}
	var usuario1 interfaces.IUser = clases.Usuario{}.New("user1",
		"password1")

	resultado := membresia1.EsValido(usuario1)
	if !resultado.Exito {
		t.Errorf("deberia dar verdadero")
	}
}

func Test_Invalido(t *testing.T) {
	membresia1 := clases.UsuarioMembresia{}
	var usuario1 interfaces.IUser = clases.Usuario{}.New("user1",
		"passwordInvalido")

	resultado := membresia1.EsValido(usuario1)
	if resultado.Exito {
		t.Errorf("deberia dar falso")
	}
}

func init() {

}
