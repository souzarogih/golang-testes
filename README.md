# Testes Automatizados

## Exemplos de testes para lingugagem Golang

## Comandos para executar testes

<h2>Teste</h2>

`go test`

<h2>Testa na pasta raiz os arquivos de dentro de um diretório</h2>

`go test ./...`

<h2>Testa mostrando detalhes</h2>

`go test -v`

<h2>Testa mostrando a cobertura de testes</h2>

`go test --cover`

<h2>Para gerar um arquivo com toda cobertura dos testes</h2>

`go test --coverprofile cobertura.txt`

<h2>Ler o arquivo de cobertura</h2>

`go tool cover --func=cobertura.txt`

<h2>Ler o arquivo de cobertura e exibe html</h2>

`go tool cover --html=cobertura.txt`
