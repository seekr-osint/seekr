```sh
curl http://localhost:8080/people
curl http://localhost:8080/people     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"id": "4","name": "hacker","age": 49}'
```
# get names
```sh
curl http://localhost:8080/names
```
```json
```

# post names
```sh
curl http://localhost:8080/people/3/addAccounts/9glenda
```
```sh
curl http://localhost:8080/people/3/getAccounts/9glenda
```
```json
{
    "id": "3",
    "name": "Max Bulk",
    "age": 22,
    "bday": "1.1.2000",
    "address": "That Street",
    "phone": "123 456789",
    "civilstatus": "married",
    "kids": "none",
    "hobbies": "walking",
    "email": "that@mail.com",
    "occupation": "That job",
    "prevoccupation": "the other job",
    "military": "none",
    "club": "none",
    "legal": "parking ticket",
    "political": "right",
    "notes": "looks goofy",
    "accounts": {
        "github": {
            "service": "github",
            "id": "",
            "username": "9glenda",
            "url": "",
            "profilePicture": null,
            "bio": null
        },
        "slideshare": {
            "service": "slideshare",
            "id": "",
            "username": "9glenda",
            "url": "",
            "profilePicture": null,
            "bio": null
        }
    }
}
```
