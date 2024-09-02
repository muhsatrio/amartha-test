# amartha-test
Repository for Technical Test Amartha

# How to Run

1. Adjust configuration in `config.yaml`
2. Dump the table and user to local database with `dump.sql`
3. Run the API with `go run main.go`
4. Can use `API Postman.json` to import related API to Postman

# API Specification
## Login
### Endpoint
POST `/auth/login`
### Request
```
{
    "username": "guest",
    "password": "123456"
}
```
### Response
```
{
    "token": "<token>"
}
```
### Note
For login can use username `guest` and password `123456` after dumping data process done.

## Reconciliation
### Endpoint
POST `/reconciliation`
### Header
Authorization: `Bearer <token>`
### Request
|Field|Type|
|---|---|
|transaction|File (.csv)|
|bank_statement|File (.csv)|
|start_date|Text (YYYY-mm-dd)|
|end_date|Text (YYYY-mm-dd)|
### Response
```
{
    "transactions_processed": 10,
    "transactions_matched": 9,
    "transactions_unmatch": {
        "transactions": [
            {
                "trxId": "BNI009",
                "amount": 240000,
                "type": "CREDIT",
                "transaction_time": "2024-09-02T17:00:00Z"
            }
        ],
        "bank_statements": [
            {
                "unique_identifier": "BNI005",
                "amount": 240000,
                "transaction_date": "2024-09-02T00:00:00Z"
            }
        ]
    },
    "total_discrepancies": 1710000
}
```
### Note
For format file transaction and bank_statement can use file `transaction.csv` and `bank_statement.csv` in this project folder
