## Leo Bishop Checkout.com (ProcessOut) Technical Assessment Task

To build the docker image wrapping the API run:
```
docker build -t processout .
```
*Note*: This may take a minute as `cgo` is required to run the SQLite3 relational database.

To run an API in a docker container from the image built above run:
```
docker run -p 8080:8080 processout
```

This API is exposed on port 8080 and has two valid endpoints.

### `/transaction/sync`
Makes a synchronous transaction request to an acquiring bank.
Requires a JSON body in the format:
```json
{
    "timestamp_unix": 1726141174,
    "card_no": "1234-5678-1234-5667",
    "expiry_month": 5,
    "expiry_year": 2026,
    "cvv": "344",
    "currency": "GPB",
    "amount": 1234
}
```

It returns the transaction UUID and outcome in the following JSON:
```json
{
    "id": "de1c9648-8bd2-4b4e-b062-47160ccd6618",
    "state": "Successful"
}
```

### `/transaction/:id`
Returns the masked transaction details for a single transaction
Returns the transaction details in following JSON
```json
{
    "id": "de1c9648-8bd2-4b4e-b062-47160ccd6618",
    "timestamp_unix": 1726141174,
    "card_no": "****-****-****-5667",
    "expiry_month": 0,
    "expiry_year": 0,
    "cvv": "***",
    "currency": "GPB",
    "amount": 1234,
    "state": "Successful"
}
```