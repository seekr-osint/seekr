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
### Delete account of person
There are 2 ways to delete a person by it's id both not giving any response back.
#### DELETE request
```sh
curl --request "DELETE" localhost:8080/people/:id/accounts/:account
``` 
```sh
curl --request "DELETE" localhost:8080/people/1/accounts/GitHub-9glenda
```
#### GET request
In some cases it's simpler to use a `GET` request.
```sh
curl localhost:8080/people/:id/accounts/:account/delete
```

### Post person
If the json is invalid there will be no response.

To post a person:
```sh
curl --request "POST" --data '{"id": "1","name": "test"}' localhost:8080/person
```
```json
{
    "id": "1",
    "name": "test",
    "pictures": null,
    "maidenname": "",
    "age": 0,
    "bday": "",
    "address": "",
    "phone": "",
    "ssn": "",
    "civilstatus": "",
    "kids": "",
    "hobbies": "",
    "email": null,
    "occupation": "",
    "prevoccupation": "",
    "education": "",
    "military": "",
    "religion": "",
    "pets": "",
    "club": "",
    "legal": "",
    "political": "",
    "notes": "",
    "relations": null,
    "sources": null,
    "accounts": null,
    "tags": null,
    "notaccounts": null
}
```
To overwrite a person simply do the same:
```sh
curl --request "POST" --data '{"id": "1","name": "test"}' localhost:8080/person
```
```json
{
    "message": "overwritten person"
}
```
#### Get accounts
To get all accounts of a Username:
```sh
curl localhost:8080/getAccounts/:username
```
##### Example
```sh
curl localhost:8080/getAccounts/9glenda
```
```json
{
    "GitHub-9glenda": {
        "service": "GitHub",
        "id": "69043370",
        "username": "9glenda",
        "url": "https://github.com/9glenda",
        "profilePicture": {
            "1": {
                "img": "base64",
                "img_hash": 164
            }
        },
        "bio": {
            "1": {
                "bio": "18yo Russian linux enthusiast"
            }
        },
        "firstname": "",
        "lastname": "",
        "location": "bell labs",
        "created": "2020-07-31T13:04:48Z",
        "updated": "2023-01-28T20:34:05Z",
        "blog": "9glenda.github.io",
        "followers": 0,
        "following": 0
    },
    "Linktree-9glenda": {
        "service": "Linktree",
        "id": "",
        "username": "9glenda",
        "url": "https://linktr.ee/9glenda",
        "profilePicture": null,
        "bio": null,
        "firstname": "",
        "lastname": "",
        "location": "",
        "created": "",
        "updated": "",
        "blog": "",
        "followers": 0,
        "following": 0
    },
    "SlideShare-9glenda": {
        "service": "SlideShare",
        "id": "",
        "username": "9glenda",
        "url": "https://slideshare.net/9glenda",
        "profilePicture": {
            "1": {
                "img": "base64",
                "img_hash": 164
            }
        },
        "bio": null,
        "firstname": "",
        "lastname": "",
        "location": "",
        "created": "",
        "updated": "",
        "blog": "",
        "followers": 0,
        "following": 0
    },
    "Twitter-9glenda": {
        "service": "Twitter",
        "id": "",
        "username": "9glenda",
        "url": "https://twitter.com/9glenda",
        "profilePicture": null,
        "bio": null,
        "firstname": "",
        "lastname": "",
        "location": "",
        "created": "",
        "updated": "",
        "blog": "",
        "followers": 0,
        "following": 0
    }
}
```
