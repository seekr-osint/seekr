![seekr-logo2](https://user-images.githubusercontent.com/67828948/209133527-88a5eec7-fd24-43a5-bbcb-363777ef5bd9.jpg)
# SEEKR
## Introduction
SEEKR is a multi-purpose toolkit for gathering and managing OSINT-data with a sleek web interface. The backend is written in Go and offers a wide range of features for data collection, organization, and analysis. Whether you're a researcher, investigator, or just someone looking to gather information, SEEKR makes it easy to find and manage the data you need. Give it a try and see how it can streamline your OSINT workflow!
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
