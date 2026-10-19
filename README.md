# Integração de Aplicações com RabbitMQ (Go e C#)

Comunicação assíncrona entre duas aplicações independentes utilizando RabbitMQ e Docker Compose:
- **Aplicação 1 (`app-1/`)**: Producer desenvolvido em **Go**, que publica pedidos de pizzas na fila `fila-pizzas` a cada 10 segundos.
- **Aplicação 2 (`app-2/`)**: Consumer desenvolvido em **C# (.NET 8)**, que consome e processa as mensagens da fila `fila-pizzas` em tempo real.

## Integrantes

- **Nome:** Caio Sousa — **RA:** 04251101
- **Nome:** Bruno Oliveira — **RA:** 04251037

## Estrutura do Projeto

```text
projeto/
├── app-1/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── app-2/
│   ├── Dockerfile
│   ├── Program.cs
│   └── consumer.csproj
├── compose.yml
└── README.md
```

## Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/) e [Docker Compose](https://docs.docker.com/compose/) instalados.

## Como executar e testar a comunicação

1. No diretório raiz do projeto (onde está localizado o arquivo `compose.yml`), execute o comando:

   ```bash
   docker compose up --build
   ```

2. Aguarde a inicialização do container `rabbitmq` (verificada automaticamente via *healthcheck*) e a subida dos serviços `app-1` e `app-2`.

3. Acompanhe os logs diretamente no terminal para verificar a comunicação assíncrona:
   - O **app-1 (Producer em Go)** exibirá:
     ```text
     conectado ao rabbitmq com sucesso.
     pizza enviada: Calabresa
     ```
   - O **app-2 (Consumer em C#)** exibirá em seguida:
     ```text
     conectado ao rabbitmq com sucesso.
     pizza recebida: Calabresa
     ```

4. *(Opcional)* Você também pode acessar o painel de gerenciamento do RabbitMQ no navegador:
   - **URL:** `http://localhost:15672`
   - **Usuário:** `admin`
   - **Senha:** `admin`

5. Para encerrar a execução dos containers, pressione `Ctrl + C` ou execute:

   ```bash
   docker compose down
   ```
