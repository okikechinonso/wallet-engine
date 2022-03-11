# Wallet-Egnine
This demonstrate how wallet a works.  

Run the command to download all dependencies.  
```
go mod tidy
```

Before you run the code make sure you add a ``.env`` file  in the root directory of project that contains the following:
```
DB_PORT=<your db port>
DB_USER=<your db username>
DB_PASS=<your db password>
JWT_SECRET=<your secrete>
PORT=<your server port>
```


Run the application using:
```
go run main.go
```


## Use postman to consume the following endpoints and send request using json

The following endpoints are available for the project:

### createWallet 
endpoint 
```
localhost:<PORT>/api/v1/create
```
payload
```bigquery
{
   "first_name":"john",
   "last_name":"doe",
   "email":"jonndoe@gmail.com",
   "password":"password123",
   "phone": "070317485678"
}
```

### login
This ensures that the user is a member in the system, the user is authenticated using JWT. To perform other transaction copy the access_token without the quotation marks
endpoint `localhost:<PORT>/api/v1/login`
payload
```bigquery
{
    "email":"jonndoe@gmail.com",
   "password":"password123"
}
```

### CreditWallet
Before testing this endpoint, copy the login access_token provided during login and use it to set the Bearer token of your Authorization in postman with the access_token provided.  
Endpoint:
```
localhost:<PORT>/api/v1/credit
```
Payload
```
{
"amount": <enter amount(intergers or decimals)> 
}
```

### DebitWallet 
Endpoint 
```
localhost:<PORT>/api/v1/debit
```
payload
```
{
"amount": <enter amount(intergers or decimals)> 
}
```

### ActivateWallet
This endpoint either activate or deactivates a wallet depending on the status of the wallet.  
endpoint 
```
localhost:<PORT>/api/v1/activate?wallet_address=<PHONE>
```
## Tests
Testing is done using the GoMock framework. The ``gomock`` package and the ``mockgen``code generation tool are used for this purpose.
If you installed the dependencies using the command given above, then the packages would have been installed. Otherwise, installation can be done using the following commands:
```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
```
run all the test files using:
```bigquery
go test -v ./...
```





