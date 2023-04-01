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
    "relations": null,
    "religion": "",
    "sources": {},
    "ssn": "",
    "tags": null
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


## Overwrite Person

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{"id":"1"}'
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
-d '{"id":"2"}'
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
    "relations": null,
    "religion": "",
    "sources": {},
    "ssn": "",
    "tags": null
}
```

**Status Code:** 201


