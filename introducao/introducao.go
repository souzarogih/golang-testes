package main

import (
	"fmt"
	"introducaotests/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("avenida epitacio pessoa")
	fmt.Println(tipoEndereco)
}

//go mod tidy
//go run introducao.go
//go mod init introducao-testes