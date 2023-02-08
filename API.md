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

### Delete person by id
There are 2 ways to delete a person by it's id both not giving any response back.
#### DELETE request
```sh
curl --request "DELETE" localhost:8080/people/:id
``` 
```sh
curl --request "DELETE" localhost:8080/people/1
```
#### GET request
In some cases it's simpler to use a `GET` request.
```sh
curl localhost:8080/people/:id/delete
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
### Example 
```sh
curl --request "POST" localhost:8080/person --data '
{
    "id": "10",
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
    "email": { "hacker@gmail.com" : { "mail": "hacker@gmail.com" } },
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
'
```
```json
{
    "id": "10",
    "name": "test",
    "pictures": {},
    "maidenname": "",
    "age": 0,
    "bday": "",
    "address": "",
    "phone": "",
    "ssn": "",
    "civilstatus": "",
    "kids": "",
    "hobbies": "",
    "email": {
        "hacker@gmail.com": {
            "mail": "hacker@gmail.com",
            "value": 0,
            "src": "",
            "services": {
                "Discord": {
                    "name": "Discord",
                    "link": "",
                    "username": "",
                    "icon": "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6cc3c481a15a141738_icon_clyde_white_RGB.png"
                }
            },
            "valid": true,
            "gmail": true,
            "validGmail": true,
            "provider": ""
        }
    },
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
    "sources": {},
    "accounts": {},
    "tags": null,
    "notaccounts": null
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
### Google

To get results from google:
```sh
curl "localhost:8080/google/seekr-osint"
```
```json
[
    {
        "url": "https://github.com/seekr-osint/seekr",
        "description": "Seekr is a multi-purpose toolkit for gathering and managing OSINT-data with a sleek web interface. The backend is written in Go and offers a wide range of ...A multi-purpose toolkit for gathering and managing OSINT-Data with a neat web-interface. JavaScript 2. Repositories. Type. Select type. All Public Sources",
        "title": "seekr-osint/seekr: A multi-purpose toolkit for ... - GitHubseekr - GitHub"
    },
    {
        "url": "https://securityonline.info/seekr-multi-purpose-toolkit-for-gathering-and-managing-osint-data/",
        "description": "1 day ago —",
        "title": "multi-purpose toolkit for gathering and managing OSINT Data"
    },
    {
        "url": "https://news.ycombinator.com/from?site=github.com/seekr-osint",
        "description": "All in one OSINT tool with web interface (github.com/seekr-osint). 2 points by Reusable2862 20 minutes ago | past | 1 comment ...Hello guy. I'm new here. I recently started to get into OSINT and therefore wrote this little tool. Now I am searching for feedback on my tool ...",
        "title": "in one OSINT tool with web interface - Hacker NewsAll in one OSINT tool with web interface - Hacker News"
    },
    {
        "url": "https://forensic-architecture.org/methodology/osint",
        "description": "Asylum seekers crossing the Aegean Sea are intercepted within Greek waters or ... Open source intelligence or OSINT is information collected from publicly ...",
        "title": "Osint - Forensic Architecture"
    },
    {
        "url": "https://haxf4rall.com/2023/02/07/seekr-multi-purpose-toolkit-for-gathering-and-managing-osint-data/",
        "description": "22 hours ago —",
        "title": "Seekr: multi-purpose toolkit for gathering and managing ..."
    },
    {
        "url": "https://www.liferaftinc.com/blog/7-osint-podcasts-every-analyst-should-follow",
        "description": "Here are our top seven OSINT podcasts for analysts, investigators, and researchers. ... OSINT for Job Seekers · Everyday OSINT with Rae Baker ...",
        "title": "7 OSINT Podcasts Every Analyst Should Follow - LifeRaft"
    },
    {
        "url": "https://jobs.baesystems.com/global/en/job/86783BR/OSINT-Trainer",
        "description": "These dedicated email and telephonic options are only for job seekers with disabilities to request an accommodation. Please do not use these services to check ...",
        "title": "OSINT Trainer in Grafenwoehr, APO, Germany"
    },
    {
        "url": "https://books.google.de/books?id=4aH0DwAAQBAJ\u0026pg=PA149\u0026lpg=PA149\u0026dq=seekr-osint\u0026source=bl\u0026ots=F75I3cMb9l\u0026sig=ACfU3U0f0yY5ekeF9BX6QxK_NMYxiO4HlQ\u0026hl=en\u0026sa=X\u0026ved=2ahUKEwiOoay_1ob9AhU4QvEDHaA_Di8Q6AF6BAgmEAM",
        "description": "Mohammad A. Tayebi",
        "title": "Open Source Intelligence and Cyber Crime: Social Media Analytics"
    }
]
```
