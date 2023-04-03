## Post Person

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "2"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 0,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {},
	"hobbies": "",
	"id": "2",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Overwrite Person

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "1"
}'
```

**Response:**

```json
{
	"message": "overwritten person"
}
```

**Status Code:** 202


## Get Person by ID

**Curl Request:**

```sh
curl -X GET http://localhost:8080/people/2
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 0,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {},
	"hobbies": "",
	"id": "2",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 200


## Get Person which does not exsist

**Curl Request:**

```sh
curl -X GET http://localhost:8080/people/100
```

**Response:**

```json
null
```

**Status Code:** 404


## Post person with included email

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"fsdfadsfasdfasdf@gmail.com": {
			"mail": "fsdfadsfasdfasdf@gmail.com"
		}
	},
	"hobbies": "",
	"id": "10",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": null,
	"political": "",
	"prevoccupation": "",
	"relations": null,
	"religion": "",
	"sources": null,
	"ssn": "",
	"tags": null
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"fsdfadsfasdfasdf@gmail.com": {
			"gmail": true,
			"mail": "fsdfadsfasdfasdf@gmail.com",
			"provider": "",
			"services": {},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"validGmail": true,
			"value": 0
		}
	},
	"hobbies": "",
	"id": "10",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email detecting only discord as a services

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 10,
	"email": {
		"has_discord_account@gmail.com": {
			"mail": "has_discord_account@gmail.com"
		}
	},
	"id": "11",
	"name": "Email test"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"has_discord_account@gmail.com": {
			"gmail": true,
			"mail": "has_discord_account@gmail.com",
			"provider": "",
			"services": {
				"Discord": {
					"icon": "./images/mail/discord.png",
					"link": "",
					"name": "Discord",
					"username": ""
				}
			},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"validGmail": false,
			"value": 0
		}
	},
	"hobbies": "",
	"id": "11",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email detecting all services

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 10,
	"email": {
		"all@gmail.com": {
			"mail": "all@gmail.com"
		}
	},
	"id": "12",
	"name": "Email test"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 10,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"all@gmail.com": {
			"gmail": true,
			"mail": "all@gmail.com",
			"provider": "",
			"services": {
				"Discord": {
					"icon": "./images/mail/discord.png",
					"link": "",
					"name": "Discord",
					"username": ""
				},
				"Spotify": {
					"icon": "./images/mail/spotify.png",
					"link": "",
					"name": "Spotify",
					"username": ""
				},
				"Twitter": {
					"icon": "./images/mail/twitter.png",
					"link": "",
					"name": "Twitter",
					"username": ""
				}
			},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"validGmail": true,
			"value": 0
		}
	},
	"hobbies": "",
	"id": "12",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


## Post person with included email and discord check failing

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"address": "",
	"age": 13,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"discord_error@gmail.com": {
			"mail": "discord_error@gmail.com"
		}
	},
	"hobbies": "",
	"id": "13",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": null,
	"political": "",
	"prevoccupation": "",
	"relations": null,
	"sources": null,
	"tags": null
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 13,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"education": "",
	"email": {
		"discord_error@gmail.com": {
			"gmail": true,
			"mail": "discord_error@gmail.com",
			"provider": "",
			"services": {},
			"skipped_services": {
				"Discord": true
			},
			"src": "",
			"valid": true,
			"validGmail": false,
			"value": 0
		}
	},
	"hobbies": "",
	"id": "13",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": "",
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"ssn": "",
	"tags": []
}
```

**Status Code:** 201


