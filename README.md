# Leo Bishop Checkout.com (ProcessOut) Technical Assessment Task

## Design
### Core Design
- Single API service designed to wrap all transaction behaviors.
- Wrapped in a docker image to ensure consistent builds across assessor machines.
- Directory structure:
    - `api`: Root directory
        - `ports`: Entry points to system functionality, contain http handlers.
        - `core`: Shared functionality which could be used across multiple endpoints
            - `bank`: Wraps calls to the bank API (mock) and associated functionality.
            - `db`: Database interaction functionality.
            - `domain`: Shared class definitions which are used widely across codebase.
            - `transaction`: Core transaction functionality, i.e. business logic.

#### Ports
The entry points into system functionality.
In this case only http handlers but enables clean addition of additional connection methods.
I consider handlers to be the interface between the user interface and internal system functionality.
As such no functional requirements should be met by code in ports.
All it should do it extract relevant data from the user request, pass it to internal functions and parse the result to the correct response format and return it to the user.

#### Bank
The Bank API mock. 
Went with simple mock functions (as opposed to an additional API service) for implementation and testing simplicity.
Does not support switching bank implementations/interfaces during run-time.
Considering this API as a microservice providing transaction operations I prefer to transfer incoming traffic to a new instance using a different implementation then to take on the implementation complexity of supporting run-time implementation switching without transaction loss.
(This issue was primarily considered based on a possible future implementation of asynchronous transactions.)
Implementing using a `Go` `interface` object to enable the (theoretical future) provision of a real implementation and a naive, in-memory mock for testing support.

#### DB
The relational database functionality, containing functionality to create and interact with the `transactions` table.
Similar to the design approach for `ports`, `db` is the interface between internal system and the database, enabling looser coupling of the system's transaction model and the database's transaction model.
Provides an in memory version to enable easier integration & end-to-end testing.

#### Domain
Common class implementations and isolated methods required across the codebase.
Whilst domain packages risk recreating the issues of a global `util` packages if allowed to bloat.
I still like using them as the ability to prevent cyclical imports and encourage the use of common objects (as opposed to multiple fiat implementations in different areas that can arise in large project) is helpful.

#### Transaction
The core internal functionality which implement the specification requirements, i.e. where business logic is implemented.

#### SQLite3
For persistent storage went with a standard SQL relational database for simplicity.
Whilst for a proper implementation I would use a separate Postgres instance, for the purposes of this demonstration work only running a SQLite3 instance within the same docker image was deemed okay.

### Assumptions
- A single database.
- Consistent card number & cvv formats.
- Only supporting a defined set of currencies (GBP, USD & EUR).

## Build & Run
**Note**: Requires docker

To build the docker image wrapping the API run:
```
docker build -t processout .
```
*Note*: This may take a minute as `cgo` is required to build the SQLite3 relational database package.

To run the API in a docker container from the image built above, run:
```
docker run -p 8080:8080 processout
```

## Endpoints

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