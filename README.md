# Go GraphQL Project

Este projeto é uma aplicação de exemplo que utiliza Go e GraphQL para criar uma API eficiente e flexível. O objetivo é demonstrar como configurar e usar GraphQL com Go para construir APIs modernas.

## Pré-requisitos

Antes de começar, certifique-se de ter o Go instalado na sua máquina. Você pode baixar e instalar o Go a partir do [site oficial](https://golang.org/dl/).

## Instalação

1. Clone o repositório para a sua máquina local:
    ```sh
    git clone https://github.com/dyhalmeida/go-graphql.git
    cd go-graphql
    ```

2. Instale as dependências do projeto:
    ```sh
    go mod tidy
    ```

2. Crie as tabelas no SQLite:
    ```sh
    sqlite3 database.db < ./internal/database/create_table.sql 
    ```

## Executando o Projeto

1. Inicie o servidor GraphQL:
    ```sh
    go run main.go
    ```

2. Acesse a interface GraphQL no seu navegador:
    Abra o navegador e vá para `http://localhost:8080/graphql` para acessar a interface GraphQL onde você pode testar suas [queries e mutations](./query.api.graphql).
