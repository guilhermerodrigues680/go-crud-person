# GO CRUD Pessoa

GO CRUD Pessoa é uma aplicação simples com a implementação de um CRUD de Pessoa com e uma REST API.

## Configuração do projeto

### Run and Debug Vscode

No [Visual Studio Code](https://code.visualstudio.com/) você pode debugar a aplicação usando o `Launch Main` disponivel na janela `Run and Debug`

### Compilação
```sh
# Constroe apenas o binário do Linux
make build

# Constroe todos os binários para todas as plataformas
make cross
```

### Deploy Binário
```sh
# Execute o binário contruido
./bin/main-linux-amd64
./bin/main-darwin-amd64
./bin/bin/main-windows-amd64.exe
```

### Deploy Docker
```sh
# Multi-stage build
docker build -f build/package/Dockerfile -t gocrudperson .
docker run -i --rm -p 8080:8080 gocrudperson

# Single-stage build
docker build -f build/package/Dockerfile.single-stage -t gocrudperson/single-stage .
docker run -i --rm -p 8080:8080 gocrudperson/single-stage
```

### Deploy Docker Compose
```sh
docker-compose -f deployments/docker-compose.yml up
```

## Layout do Projeto

Baseado no [github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Documentação da API Rest

A API Rest do projeto foi documentada no padrão OpenAPI 3.0, e está disponivel no arquivo `api/openapi-spec/openapi.yaml` ou na página `http://localhost:8080/api/docs` da aplicação.