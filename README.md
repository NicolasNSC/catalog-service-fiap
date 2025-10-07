# catalog-service-fiap

Microserviço para gerenciamento (cadastro e edição) de veículos. Este projeto integra o Tech Challenge do curso SOAT da Pós-Tech da FIAP, seguindo os princípios da **Clean Architecture**.

## Tecnologias Utilizadas

- **Linguagem:** Go
- **Banco de Dados:** PostgreSQL
- **Infraestrutura:** Docker & Docker Compose
- **Roteador HTTP:** Chi
- **Testes:** Testify & Gomock

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

A API estará disponível em [http://localhost:8080](http://localhost:8080).  
A documentação Swagger estará em [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

## Comandos Úteis (Makefile)

- `make docker-up`: Sobe os containers da aplicação e do banco de dados.
- `make docker-down`: Para e remove os containers e volumes.
- `make test`: Executa todos os testes e exibe a cobertura no terminal.
- `make cov`: Abre o relatório de cobertura de testes em HTML no navegador.
- `make gen`: Gera os mocks para as interfaces (gomock).

## Endpoints da API

A documentação interativa completa está disponível em `/swagger/index.html`.

### Endpoints Públicos

- `POST /vehicles/add`: Cadastra um novo veículo.
- `PUT /vehicles/{id}`: Atualiza os dados de um veículo existente.