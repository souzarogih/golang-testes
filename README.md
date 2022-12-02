# Testes Automatizados

## Exemplos de testes para linguagem Golang

## Comandos para executar testes

<h3>Teste</h3>

`go test`

<h3>Testa na pasta raiz os arquivos de dentro de um diretório</h3>

`go test ./...`

<h3>Testa mostrando detalhes</h3>

`go test -v`

<h3>Testa mostrando a cobertura de testes</h3>

`go test --cover`

<h3>Para gerar um arquivo com toda cobertura dos testes</h3>

`go test --coverprofile cobertura.txt`

<h3>Ler o arquivo de cobertura</h3>

`go tool cover --func=cobertura.txt`

<h3>Ler o arquivo de cobertura e exibe html</h3>

`go tool cover --html=cobertura.txt`
