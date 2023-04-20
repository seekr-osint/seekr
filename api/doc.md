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
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
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
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
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
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
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
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
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
	"age": 10,
	"email": {
		"fsdfadsfasdfasdf@gmail.com": {
			"mail": "fsdfadsfasdfasdf@gmail.com"
		}
	},
	"id": "10",
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
	"custom": null,
	"education": "",
	"email": {
		"fsdfadsfasdfasdf@gmail.com": {
			"mail": "fsdfadsfasdfasdf@gmail.com",
			"provider": "gmail",
			"services": {},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
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
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
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
	"custom": null,
	"education": "",
	"email": {
		"has_discord_account@gmail.com": {
			"mail": "has_discord_account@gmail.com",
			"provider": "fake_mail",
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
			"value": 0
		}
	},
	"gender": "",
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
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
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
	"custom": null,
	"education": "",
	"email": {
		"all@gmail.com": {
			"mail": "all@gmail.com",
			"provider": "gmail",
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
				},
				"Ubuntu GPG": {
					"icon": "./images/mail/ubuntu.png",
					"link": "https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com\u0026op=index",
					"name": "Ubuntu GPG",
					"username": ""
				},
				"keys.gnupg.net": {
					"icon": "./images/mail/gnupg.ico",
					"link": "https://keys.gnupg.net/pks/lookup?search=all@gmail.com\u0026op=index",
					"name": "keys.gnupg.net",
					"username": ""
				}
			},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
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
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
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
	"age": 13,
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
	"phone": {},
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
	"custom": null,
	"education": "",
	"email": {
		"discord_error@gmail.com": {
			"mail": "discord_error@gmail.com",
			"provider": "fake_mail",
			"services": {},
			"skipped_services": {
				"Discord": true
			},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
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
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post person with included email detected as a fake email by seekr


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"accounts": null,
	"age": 10,
	"email": {
		"fake_mail@gmail.com": {
			"mail": "fake_mail@gmail.com"
		}
	},
	"id": "14",
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
	"custom": null,
	"education": "",
	"email": {
		"fake_mail@gmail.com": {
			"mail": "fake_mail@gmail.com",
			"provider": "fake_mail",
			"services": {},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "14",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Email test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post Person (civil status)


**Curl Request:**

```sh
curl -X GET http://localhost:8080/getAccounts/snapchat-exsists
```

**Response:**

```json
{
	"Snapchat-snapchat-exsists": {
		"bio": null,
		"blog": "",
		"created": "",
		"firstname": "",
		"followers": 0,
		"following": 0,
		"id": "",
		"lastname": "",
		"location": "",
		"profilePicture": null,
		"service": "Snapchat",
		"updated": "",
		"url": "",
		"username": "snapchat-exsists"
	}
}
```

**Status Code:** 200


## Post Person (civil status)
Possible values are: Single,Married,Widowed,Divorced,Separated

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"civilstatus": "Single",
	"id": "15"
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 0,
	"bday": "",
	"civilstatus": "Single",
	"club": "",
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "15",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post Person (invalid civil status)
Possible values are: Single,Married,Widowed,Divorced,Separated

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"civilstatus": "Invalid",
	"id": "16"
}'
```

**Response:**

```json
{
	"message": "Invalid civil status"
}
```

**Status Code:** 400


## Post Person (missing id)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{}'
```

**Response:**

```json
{
	"message": "Missing ID"
}
```

**Status Code:** 400


## Post Person (invalid religion)
Check [surce code](https://github.com/seekr-osint/seekr/blob/main/api/religion_type.go) for valid religions 

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "17",
	"religion": "invalid"
}'
```

**Response:**

```json
{
	"message": "Invalid religion"
}
```

**Status Code:** 400


## Post Person (invalid Gender)
Possible values are: Male,Female,Other

**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"gender": "Invalid",
	"id": "19"
}'
```

**Response:**

```json
{
	"message": "Invalid gender"
}
```

**Status Code:** 400


## Post Person (Email key missmatch)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"age": 10,
	"email": {
		"key@gmail.com": {
			"mail": "missmatched_value@gmail.com"
		}
	},
	"id": "10",
	"name": "Email test"
}'
```

**Response:**

```json
{
	"message": "Key missmatch: Email[key@gmail.com] = missmatched_value@gmail.com"
}
```

**Status Code:** 400


## Post Person (Email key missmatch)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "20",
	"name": "Phone test",
	"phone": {
		"6502530000": {
			"number": "missmatched_value"
		}
	}
}'
```

**Response:**

```json
{
	"message": "Key missmatch: Phone[6502530000] = missmatched_value"
}
```

**Status Code:** 400


## Post Person (Email key missmatch already taken ID)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "20",
	"name": "Phone test",
	"phone": {
		"6502530000": {
			"number": "missmatched_value"
		}
	}
}'
```

**Response:**

```json
{
	"message": "Key missmatch: Phone[6502530000] = missmatched_value"
}
```

**Status Code:** 400


## Post Person (Phone number formatting)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "21",
	"name": "Phone test",
	"phone": {
		"+1(234) 567-8901": {
			"number": "+1(234) 567-8901"
		}
	}
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
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "21",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Phone test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {
		"+1 234-567-8901": {
			"national_format": "(234) 567-8901",
			"number": "+1 234-567-8901",
			"phoneinfoga": {
				"Carrier": "",
				"Country": "US",
				"CountryCode": 1,
				"E164": "+12345678901",
				"International": "12345678901",
				"Local": "(234) 567-8901",
				"RawLocal": "2345678901",
				"Valid": true
			},
			"tag": "",
			"valid": true
		}
	},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post Person (Phone number formatting missing +)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "30",
	"name": "Phone test",
	"phone": {
		"1-234-567-8901": {
			"number": "1-234-567-8901"
		}
	}
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
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "30",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Phone test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {
		"+1 234-567-8901": {
			"national_format": "(234) 567-8901",
			"number": "+1 234-567-8901",
			"phoneinfoga": {
				"Carrier": "",
				"Country": "US",
				"CountryCode": 1,
				"E164": "+12345678901",
				"International": "12345678901",
				"Local": "(234) 567-8901",
				"RawLocal": "2345678901",
				"Valid": true
			},
			"tag": "",
			"valid": true
		}
	},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post Person (Invalid_number)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "22",
	"name": "Phone test",
	"phone": {
		"Invalid_number": {
			"number": "Invalid_number"
		}
	}
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
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "22",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Phone test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {
		"Invalid_number": {
			"national_format": "",
			"number": "Invalid_number",
			"phoneinfoga": {
				"Carrier": "",
				"Country": "",
				"CountryCode": 0,
				"E164": "",
				"International": "",
				"Local": "",
				"RawLocal": "",
				"Valid": false
			},
			"tag": "",
			"valid": false
		}
	},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post Person (Empty phone number)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"id": "23",
	"name": "Phone test",
	"phone": {
		"": {
			"number": ""
		}
	}
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
	"custom": null,
	"education": "",
	"email": {},
	"gender": "",
	"hobbies": "",
	"id": "23",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Phone test",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## Post Person (Lot of fields)


**Curl Request:**

```sh
curl -X POST http://localhost:8080/person \
-H 'Content-Type: application/json' \
-d '{
	"age": 23,
	"email": {
		"all@gmail.com": {
			"mail": "all@gmail.com"
		}
	},
	"id": "24",
	"name": "Many fields",
	"phone": {
		"+13183442908": {
			"number": "+13183442908"
		}
	}
}'
```

**Response:**

```json
{
	"accounts": {},
	"address": "",
	"age": 23,
	"bday": "",
	"civilstatus": "",
	"club": "",
	"custom": null,
	"education": "",
	"email": {
		"all@gmail.com": {
			"mail": "all@gmail.com",
			"provider": "gmail",
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
				},
				"Ubuntu GPG": {
					"icon": "./images/mail/ubuntu.png",
					"link": "https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com\u0026op=index",
					"name": "Ubuntu GPG",
					"username": ""
				},
				"keys.gnupg.net": {
					"icon": "./images/mail/gnupg.ico",
					"link": "https://keys.gnupg.net/pks/lookup?search=all@gmail.com\u0026op=index",
					"name": "keys.gnupg.net",
					"username": ""
				}
			},
			"skipped_services": {},
			"src": "",
			"valid": true,
			"value": 0
		}
	},
	"gender": "",
	"hobbies": "",
	"id": "24",
	"kids": "",
	"legal": "",
	"maidenname": "",
	"military": "",
	"name": "Many fields",
	"notaccounts": null,
	"notes": "",
	"occupation": "",
	"pets": "",
	"phone": {
		"+1 318-344-2908": {
			"national_format": "(318) 344-2908",
			"number": "+1 318-344-2908",
			"phoneinfoga": {
				"Carrier": "",
				"Country": "US",
				"CountryCode": 1,
				"E164": "+13183442908",
				"International": "13183442908",
				"Local": "(318) 344-2908",
				"RawLocal": "3183442908",
				"Valid": true
			},
			"tag": "",
			"valid": true
		}
	},
	"pictures": {},
	"political": "",
	"prevoccupation": "",
	"relations": {},
	"religion": "",
	"sources": {},
	"tags": []
}
```

**Status Code:** 201


## GET Person Markdown


**Curl Request:**

```sh
curl -X GET http://localhost:8080/people/24/markdown
```

**Response:**

```json
{
	"markdown": "# Many fields\n- Age: `23`\n- Phone: `+1 318-344-2908`\n## Email\n### all@gmail.com\n- Mail: `all@gmail.com`\n- Provider: `gmail`\n#### Services\n##### Discord\n- Name: `Discord`\n- Icon: `./images/mail/discord.png`\n##### Spotify\n- Name: `Spotify`\n- Icon: `./images/mail/spotify.png`\n##### Twitter\n- Name: `Twitter`\n- Icon: `./images/mail/twitter.png`\n##### Ubuntu GPG\n- Name: `Ubuntu GPG`\n- Link: `https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com\u0026op=index`\n- Icon: `./images/mail/ubuntu.png`\n##### keys.gnupg.net\n- Name: `keys.gnupg.net`\n- Link: `https://keys.gnupg.net/pks/lookup?search=all@gmail.com\u0026op=index`\n- Icon: `./images/mail/gnupg.ico`\n\n\n"
}
```

**Status Code:** 200


