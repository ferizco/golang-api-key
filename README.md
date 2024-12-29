# golang-api-key
Implementasi api key di golang

## Cara menggunakan 
1. Lakukan go mod init terlebih dahulu
2. Run server golang dengan command go run .

## API 
http://host:8000/api/login 

Header: 
1. Content-type: application/json
2. X-API-KEY: key-xxx

Parameter(Body request): 
1. username
2. password
