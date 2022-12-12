curl http://localhost:8080/people --include --header "Content-Type: application/json" --request "POST" --data '{"id": "1","name": "hacker1","age": 49}'
curl http://localhost:8080/people --include --header "Content-Type: application/json" --request "POST" --data '{"id": "2","name": "hacker2","age": 49}'
curl http://localhost:8080/people --include --header "Content-Type: application/json" --request "POST" --data '{"id": "3","name": "hacker3","age": 49}'
curl http://localhost:8080/people/1/addAccounts/9glenda
curl http://localhost:8080/people/2/addAccounts/9glenda
curl http://localhost:8080/people/3/getAccounts/9glenda
