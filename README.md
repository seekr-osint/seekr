```sh
curl http://localhost:8080/people
curl http://localhost:8080/people     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"id": "4","name": "hacker","age": 49}'
```
# get names
```sh
curl http://localhost:8080/names
```
```json
["name1","name2"]
```
