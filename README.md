var membresia1 = new UsuarioMembresia();
 
 var usuario1 = new Usuario("user1", "password1")
 
 var resultado = membresia.esValido(usuario1);
 
 assert.True(resultado.Exito);
 
 ----------------------------------------------------------
 
 var membresia1 = new UsuarioMembresia();
 
 var usuario1 = new Usuario("user1", "passwordInvalido")
 
 var resultado = membresia.esValido(usuario1);
 
 assert.False(resultado.Exito);
 assert.Equual("Datos invalidos", resultado.Mensaje);
 
 ----------------------------------------------------------
 
 //Clase UsuarioMembresia > Metodo: esValido(Usuario usuario){
 
 var authApi1 = new AuthAPI();
 
 var resultado = authApi1.ValidarCredenciales(usuario);
 
 return resultado;
 
 }
 
 
 ----------------------------------------------------------
 
 //Clase AuthAPI > Metodo: ValidarCredenciales(Usuario usuario){
 
 var db1 = new BaseDeDatos();
 
 var usuarioDb1 = db1.ObtenerUsuario(usuario.NombreUsuario);
 
 if(usuario.PasswordHash == usuarioDb1.PasswordHash){
 return new Resultado(true);
 }else{
 return new Resultado(false, "Credenciales invalidas");
 }
 
 }
