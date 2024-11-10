# Desafio: Client-Server-API

Neste desafio, vamos aplicar conceitos de servidor HTTP, contextos, banco de dados e manipulação de arquivos com Go.

## Descrição

Você deverá desenvolver dois sistemas em Go:

- `client.go`
- `server.go`

### Requisitos

#### `client.go`

O `client.go` deve:

- Realizar uma requisição HTTP para o `server.go` solicitando a cotação do dólar.
- Receber do `server.go` apenas o valor atual do câmbio (campo `"bid"` do JSON).
- Utilizar o package `context` com um timeout máximo de **300ms** para receber o resultado do `server.go`.
- Registrar nos logs um erro caso o tempo de execução exceda o timeout.
- Salvar a cotação em um arquivo `cotacao.txt` no formato: `Dólar: {valor}`.

#### `server.go`

O `server.go` deve:

- Consumir a API [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL) para obter o câmbio de Dólar e Real.
- Retornar ao `client.go` o valor da cotação em formato JSON.
- Utilizar o package `context` para:
  - Registrar no banco de dados SQLite cada cotação recebida.
  - Ter um timeout máximo de **200ms** para obter a cotação da API.
  - Ter um timeout máximo de **10ms** para persistir os dados no banco.
  - Registrar nos logs um erro caso o tempo de execução seja insuficiente para alguma operação.

#### Endpoint e Porta

- Endpoint: `/cotacao`
- Porta: `8080`