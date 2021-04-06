# rbwallet-go

rbwallet although it has a wallet in its name is just an application to create transactions for users who have an account.

## Features

- Create and retrieve account information
- Create transactions

## Tech

- [Gorilla mux](https://pkg.go.dev/github.com/gorilla/mux@v1.8.0) - request router and dispatcher
- [Negroni](https://pkg.go.dev/github.com/urfave/negroni@v1.0.0) - negroni is an idiomatic approach to web middleware
- [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10@v10.4.1) - validations for structs and individual fields based on tags.
- [Mysql](https://www.mysql.com/) - The world's most popular open source database

## Installation

rbwallet requires [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/).

Install the dependencies and devDependencies and start the server.

```sh
git clone https://github.com/devbenatti/rbwallet-go.git
cd rbwallet-go
docker-compose up -d
```
## Create account

To create a account, you must send an HTTP POST request for route `/accounts`.
#### request body example
```json
{
    "documentIdentifier": "01145609745",
}
```
#### response body example
```json
{
    "id": "A54C182B-7E22-4B1D-6ABB-0DF3A02FEE4E"
}
```
## Find account

To find a account, you must send an HTTP GET request for route `/accounts/{id}`.
#### example
`/accounts/A54C182B-7E22-4B1D-6ABB-0DF3A02FEE4E`

#### response body example
```json
{
    "id": "A54C182B-7E22-4B1D-6ABB-0DF3A02FEE4E",
    "documentIdentifier": "01145609745"
}
```
## Transaction

To create a transaction, you must send an HTTP POST request for route `/transactions`.

#### request body example
```json
{
   "accountId": "A54C182B-7E22-4B1D-6ABB-0DF3A02FEE4E",
    "operationType": 2,
    "total": 110.10
}
```

#### response body example
```json
{
   "code": "7DC61C43-15EA-4AE1-6D4F-30D30C7084FA"
}
```
### Transaction Operations
| Code | Description |
| ------ | ------ |
| 1 | Cash Purchase |
| 2 | Installment Purchase |
|3 | Withdraw |
|4 | Payment |

## Roadmap
- anti-corruption layer to validate input data