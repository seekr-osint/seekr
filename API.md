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
### Delete person by it's id
There are 2 ways to delete a person by it's id both not giving any response back.
#### DELETE request
```sh
curl --request "DELETE" localhost:8080/people/1
``` 
#### GET request
In some cases it's simpler to use a `GET` request.
```sh
curl localhost:8080/people/1/delete
```
### Post person
If the json is invalid there will be no response.
