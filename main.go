package main

import (
	"crypto/sha1"
	"encoding/base64"
)

func main() {
	mensaje := "hola"
	hasher := sha1.New()
	hasher.Write([]byte(mensaje))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	print("SHA64:" + mensaje + "= " + sha)
}
