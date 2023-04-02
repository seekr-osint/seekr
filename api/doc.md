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


## Get Person by ID

**Curl Request:**

```sh
curl -X GET http://localhost:8080/people/1
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
	"id": "1",
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


