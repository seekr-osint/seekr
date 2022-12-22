![seekr-logo2](https://user-images.githubusercontent.com/67828948/209133527-88a5eec7-fd24-43a5-bbcb-363777ef5bd9.jpg)
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
```sh
curl http://localhost:8080/getAccounts/9glenda
```
```json
    {
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
```
