# catalog-service-fiap

Microserviço para gerenciamento (cadastro e edição) de veículos. Este projeto integra o Tech Challenge do curso SOAT da Pós-Tech da FIAP, seguindo os princípios da **Clean Architecture**.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Banco de Dados:** PostgreSQL
- **Infraestrutura:** Docker & Docker Compose
- **Roteador HTTP:** Chi
- **Testes:** Testify & Gomock
- **Documentação da API:** Swagger (OpenAPI)

## Como Executar o Projeto

O projeto é totalmente containerizado. É necessário ter **Docker** e **Docker Compose** instalados.

### 1. Clone o repositório

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto, baseado no `.env-sample`.

### 3. Suba os containers

Utilize o Makefile para construir e iniciar a aplicação e o banco de dados:

```bash
make docker-up
```

## Comandos Úteis (Makefile)

- `make docker-up`: Sobe os containers da aplicação e do banco de dados.
- `make docker-down`: Para e remove os containers e volumes.
- `make test`: Executa todos os testes e exibe a cobertura no terminal.
- `make cov`: Abre o relatório de cobertura de testes em HTML no navegador.
- `make gen`: Gera os mocks para as interfaces (gomock).

## Documentação da API (Swagger)

A documentação interativa da API está disponível via Swagger após subir os containers. Acesse:

```
http://localhost:8080/swagger/index.html
```

> O endpoint pode variar conforme a configuração da aplicação.

## Endpoints da API

- `POST /vehicles/add`: Cadastra um novo veículo.
- `PUT /vehicles/{id}`: Atualiza os dados de um veículo existente.

