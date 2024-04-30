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

[![](https://mermaid.ink/img/pako:eNqtVEuP2jAQ_ivWnBYpIJwAoTmstOqqpz4O7fZQ5WLi0WI1sakfu90i_nsdnEBIglCl5javb77PM5k9FIojZGDwl0NZ4KNgz5pVuST-Y4UVL8wiKUqB0tLgbQwyvb8nFVYbjUYwmhGJr3eTXuE5HgJnm0zr-gYrI0NoZxzTQjXAOTiDmuYQ5bBj5lVp7o1-u7YmuFvrZquuCjTfWSm4umurJ0PmvuTB2e3DTlyR3UaDu7UCj26zG9BHJky_18j9bAQr0bS0Jn1oX8g3lIzz8ZHgqVOml03GgTLyZWNRon4K_dq-s8-qJvt0wWIEtUl_PDXGETJ4ZXCstOSjMmTLzJYYJYl4drX4Ear-3VxpGVdB-DftsC_-lBL8J_PWQ_SHdS4c0O-1wNLgWYBU_6bhA_PlEcmh6MydCPly3AYz3Pr_LnCg73KfUfIbf_N1pOuR_qnohBpciKBCXTHB_cHa11k52C1WmIMXB_5P-ZlDLg8-jzmrvr7JAjLrNyICt-MeqDlurXPH5A-luiZke_gNGZ2vZ0n8bpmk8TpO5ukqjeANstViliQruqKUput0saaHCP4cAeYzby-TuE5fxotFuowAubBKfwrn9XhlD38BWjO-Vg?type=png)](https://mermaid.live/edit#pako:eNqtVEuP2jAQ_ivWnBYpIJwAoTmstOqqpz4O7fZQ5WLi0WI1sakfu90i_nsdnEBIglCl5javb77PM5k9FIojZGDwl0NZ4KNgz5pVuST-Y4UVL8wiKUqB0tLgbQwyvb8nFVYbjUYwmhGJr3eTXuE5HgJnm0zr-gYrI0NoZxzTQjXAOTiDmuYQ5bBj5lVp7o1-u7YmuFvrZquuCjTfWSm4umurJ0PmvuTB2e3DTlyR3UaDu7UCj26zG9BHJky_18j9bAQr0bS0Jn1oX8g3lIzz8ZHgqVOml03GgTLyZWNRon4K_dq-s8-qJvt0wWIEtUl_PDXGETJ4ZXCstOSjMmTLzJYYJYl4drX4Ear-3VxpGVdB-DftsC_-lBL8J_PWQ_SHdS4c0O-1wNLgWYBU_6bhA_PlEcmh6MydCPly3AYz3Pr_LnCg73KfUfIbf_N1pOuR_qnohBpciKBCXTHB_cHa11k52C1WmIMXB_5P-ZlDLg8-jzmrvr7JAjLrNyICt-MeqDlurXPH5A-luiZke_gNGZ2vZ0n8bpmk8TpO5ukqjeANstViliQruqKUput0saaHCP4cAeYzby-TuE5fxotFuowAubBKfwrn9XhlD38BWjO-Vg)
