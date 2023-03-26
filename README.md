# desafio-idealctvm

1. Wire utilizado para gerar as dependências, para gerar usar o comando wire ./application
2. PostgreSQL para banco de dados
3. GORM para ORM
4. Redis para cache
5. Para carregar as variáveis na aplicação utilizar aquivo .env conforme exemplo ex.env

# Requisições (Primeiro, criar usuário para usar as outras rotas, o ID está fixo como 1)
```
/api/signup
{
    "name":"3",
    "password":"123",
    "confirm_password":"123"
}
```
```
/api/user/quotes?asset=AAPL,BTC-USD - os assets deverão ser separados por vírgula
```
```
/api/user/favorite
[
    {
        "symbol": "AAPL",
        "region": "US",
        "favorite": true,
        "position": 1
    },
    {
        "symbol": "BTC-USD",
        "region": "US",
        "favorite": true,
        "position": 3
    },
    {
        "symbol": "GOGL",
        "region": "US",
        "favorite": true,
        "position": 2
    }
]
```
```
/api/user/favorite?order=value - possiveis valores (alphabetical, user, value)
```
