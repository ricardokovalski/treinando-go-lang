# treinando-go-lang
Listando algumas dicas

- Iniciando um módulo
```
go mod init <NOME-DO-MODULO>
```

- Editando o nome do módulo no arquivo go.mod via comando do terminal
```
go mod edit --module=<NOVO-NOME>
```

- Instalando um package
```
go get <NOME-DO-PACOTE>
Ex.: go get github.com/google/uuid
```

- Instalando uma versão especifica de algum pacote
```
go get <NOME-DO-PACOTE>@<VERSION>
Ex.: go get github.com/google/uuid@1.2.3
```

- Atualiando uma dependência específica
```
go get -u github.com/google/uuid
```

- Sincronizar as dependências
```
go mod tidy
```
