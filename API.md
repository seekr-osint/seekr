# Api
## Api calls
### Get whole DataBase
To get the whole DataBase simply run:
```sh
curl localhost:8080/
```
### Get a person Object by it's id
```sh
curl localhost:8080/people/:id
```
To get person `1`:
```sh
curl localhost:8080/people/1
```
If the person is not found the status code will be `404` and the response `nil`/`null`.
Therefore in case `2` doesn't exist:
```sh
curl localhost:8080/people/1
```
```json
null
```
