![seekr-logo](https://user-images.githubusercontent.com/67828948/209133320-f742f9dd-e938-4ed7-b5de-7c451c606cd5.jpg)
# SEEKR
## SEEKR explained in fortnite terms
Seekr is like a special weapon that you can use in **Fortnite** to gather information about other players or locations on the map. It's kind of like the "recon scanner" item, which allows you to scan the area and see what's around you. However, Seekr is much more powerful and versatile, because it can search for information about almost anything you can think of, and it has a nice web interface that makes it easy to use.

The backend API of Seekr is like the engine that powers the weapon. It's written in a programming language called Golang, which is kind of like the fuel that makes Seekr work. Just like how you need to have the right type of fuel in your vehicle to make it run, the backend API of Seekr needs to be written in Golang in order to function properly.

Overall, Seekr is a useful tool that can help you gather information and get an advantage over your opponents in **Fortnite**. Just like how a good weapon or piece of equipment can give you an edge in battle, Seekr can help you find valuable information that can help you make better decisions and come out on top.
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

# add accounts
```sh
curl http://localhost:8080/people/3/addAccounts/9glenda
```

# getAccounts
to get the accounts of the user with an api call run this
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
