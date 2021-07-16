# Principais comando em GO

* **Ambientes:**
`go env` // lista as variáveis de ambiente internas e seus respectivos valores


* **Build:**
`go build <caminhoArquivo>` // compilar programas em Go
`go run <caminhoArquivo>`   // compilar e executar programas em Go



* **Debug:**
`go doc cmd/vet` // identificar código morto, construções suspeitas



* **Documentação:**
`godoc -http=:<porta>` // acesso a documentação offline, localmente



* **Pacotes:**
`go get <nomePacote>` // faz o download de pacote
`go get -u <endereco/caminho>`  // atualiza ou cria o pacote mais recente criando suas pasta e subpastas



* **Testes:**
`go test` // roda teste dentro do diretório atual
`go test ./...` // roda testes do diretório atual e todos os subdiretórios
`go test -v` // roda o testes em modo verboso, demostrando passos dos testes
`go test -cover` // roda o test, gerando um saída da cobertura no console
`go test -coverprofile=<arquivo.out>` // roda o test, gerando um saída da cobertura detalhada. Para lê-la necessita do comando: `go tool cover -func=<arquivo.out>`