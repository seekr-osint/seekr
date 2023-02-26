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
curl "localhost:8080/search/google/seekr-osint"
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
### GitHub deep investigation
GitHub deep investigation checks rather git leaked any email adresses.
```sh
curl localhost:8080/deep/github/9glenda
```
```json
{
    "plan9git@proton.me": {
        "mail": "plan9git@proton.me",
        "value": 0,
        "src": "github",
        "services": {
            "github": {
                "name": "GitHub",
                "link": "https://github.com/9glenda",
                "username": "9glenda",
                "icon": "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png"
            }
        },
        "valid": false,
        "gmail": false,
        "validGmail": false,
        "provider": ""
    }
}
```
#### Rate limitation
In case of rate limitation there will be a single email only called `fatal`
There will only be one one item.
I don't care for people using fatal as email.
```
{
	"fatal": "error message"
}
```
```sh
curl localhost:8080/search/whois/google.com
```
```raw
"   Domain Name: GOOGLE.COM\r\n   Registry Domain ID: 2138514_DOMAIN_COM-VRSN\r\n   Registrar WHOIS Server: whois.markmonitor.com\r\n   Registrar URL: http://www.markmonitor.com\r\n   Updated Date: 2019-09-09T15:39:04Z\r\n   Creation Date: 1997-09-15T04:00:00Z\r\n   Registry Expiry Date: 2028-09-14T04:00:00Z\r\n   Registrar: MarkMonitor Inc.\r\n   Registrar IANA ID: 292\r\n   Registrar Abuse Contact Email: abusecomplaints@markmonitor.com\r\n   Registrar Abuse Contact Phone: +1.2086851750\r\n   Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited\r\n   Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited\r\n   Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited\r\n   Domain Status: serverDeleteProhibited https://icann.org/epp#serverDeleteProhibited\r\n   Domain Status: serverTransferProhibited https://icann.org/epp#serverTransferProhibited\r\n   Domain Status: serverUpdateProhibited https://icann.org/epp#serverUpdateProhibited\r\n   Name Server: NS1.GOOGLE.COM\r\n   Name Server: NS2.GOOGLE.COM\r\n   Name Server: NS3.GOOGLE.COM\r\n   Name Server: NS4.GOOGLE.COM\r\n   DNSSEC: unsigned\r\n   URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/\r\n\u003e\u003e\u003e Last update of whois database: 2023-02-09T17:59:03Z \u003c\u003c\u003c\r\n\r\nFor more information on Whois status codes, please visit https://icann.org/epp\r\n\r\nNOTICE: The expiration date displayed in this record is the date the\r\nregistrar's sponsorship of the domain name registration in the registry is\r\ncurrently set to expire. This date does not necessarily reflect the expiration\r\ndate of the domain name registrant's agreement with the sponsoring\r\nregistrar.  Users may consult the sponsoring registrar's Whois database to\r\nview the registrar's reported date of expiration for this registration.\r\n\r\nTERMS OF USE: You are not authorized to access or query our Whois\r\ndatabase through the use of electronic processes that are high-volume and\r\nautomated except as reasonably necessary to register domain names or\r\nmodify existing registrations; the Data in VeriSign Global Registry\r\nServices' (\"VeriSign\") Whois database is provided by VeriSign for\r\ninformation purposes only, and to assist persons in obtaining information\r\nabout or related to a domain name registration record. VeriSign does not\r\nguarantee its accuracy. By submitting a Whois query, you agree to abide\r\nby the following terms of use: You agree that you may use this Data only\r\nfor lawful purposes and that under no circumstances will you use this Data\r\nto: (1) allow, enable, or otherwise support the transmission of mass\r\nunsolicited, commercial advertising or solicitations via e-mail, telephone,\r\nor facsimile; or (2) enable high volume, automated, electronic processes\r\nthat apply to VeriSign (or its computer systems). The compilation,\r\nrepackaging, dissemination or other use of this Data is expressly\r\nprohibited without the prior written consent of VeriSign. You agree not to\r\nuse electronic processes that are automated and high-volume to access or\r\nquery the Whois database except as reasonably necessary to register\r\ndomain names or modify existing registrations. VeriSign reserves the right\r\nto restrict your access to the Whois database in its sole discretion to ensure\r\noperational stability.  VeriSign may restrict or terminate your access to the\r\nWhois database for failure to abide by these terms of use. VeriSign\r\nreserves the right to modify these terms at any time.\r\n\r\nThe Registry database contains ONLY .COM, .NET, .EDU domains and\r\nRegistrars.\r\nDomain Name: google.com\nRegistry Domain ID: 2138514_DOMAIN_COM-VRSN\nRegistrar WHOIS Server: whois.markmonitor.com\nRegistrar URL: http://www.markmonitor.com\nUpdated Date: 2019-09-09T15:39:04+0000\nCreation Date: 1997-09-15T07:00:00+0000\nRegistrar Registration Expiration Date: 2028-09-13T07:00:00+0000\nRegistrar: MarkMonitor, Inc.\nRegistrar IANA ID: 292\nRegistrar Abuse Contact Email: abusecomplaints@markmonitor.com\nRegistrar Abuse Contact Phone: +1.2086851750\nDomain Status: clientUpdateProhibited (https://www.icann.org/epp#clientUpdateProhibited)\nDomain Status: clientTransferProhibited (https://www.icann.org/epp#clientTransferProhibited)\nDomain Status: clientDeleteProhibited (https://www.icann.org/epp#clientDeleteProhibited)\nDomain Status: serverUpdateProhibited (https://www.icann.org/epp#serverUpdateProhibited)\nDomain Status: serverTransferProhibited (https://www.icann.org/epp#serverTransferProhibited)\nDomain Status: serverDeleteProhibited (https://www.icann.org/epp#serverDeleteProhibited)\nRegistrant Organization: Google LLC\nRegistrant State/Province: CA\nRegistrant Country: US\nRegistrant Email: Select Request Email Form at https://domains.markmonitor.com/whois/google.com\nAdmin Organization: Google LLC\nAdmin State/Province: CA\nAdmin Country: US\nAdmin Email: Select Request Email Form at https://domains.markmonitor.com/whois/google.com\nTech Organization: Google LLC\nTech State/Province: CA\nTech Country: US\nTech Email: Select Request Email Form at https://domains.markmonitor.com/whois/google.com\nName Server: ns3.google.com\nName Server: ns4.google.com\nName Server: ns1.google.com\nName Server: ns2.google.com\nDNSSEC: unsigned\nURL of the ICANN WHOIS Data Problem Reporting System: http://wdprs.internic.net/\n\u003e\u003e\u003e Last update of WHOIS database: 2023-02-09T17:51:42+0000 \u003c\u003c\u003c\n\nFor more information on WHOIS status codes, please visit:\n  https://www.icann.org/resources/pages/epp-status-codes\n\nIf you wish to contact this domain’s Registrant, Administrative, or Technical\ncontact, and such email address is not visible above, you may do so via our web\nform, pursuant to ICANN’s Temporary Specification. To verify that you are not a\nrobot, please enter your email address to receive a link to a page that\nfacilitates email communication with the relevant contact(s).\n\nWeb-based WHOIS:\n  https://domains.markmonitor.com/whois\n\nIf you have a legitimate interest in viewing the non-public WHOIS details, send\nyour request and the reasons for your request to whoisrequest@markmonitor.com\nand specify the domain name in the subject line. We will review that request and\nmay ask for supporting documentation and explanation.\n\nThe data in MarkMonitor’s WHOIS database is provided for information purposes,\nand to assist persons in obtaining information about or related to a domain\nname’s registration record. While MarkMonitor believes the data to be accurate,\nthe data is provided \"as is\" with no guarantee or warranties regarding its\naccuracy.\n\nBy submitting a WHOIS query, you agree that you will use this data only for\nlawful purposes and that, under no circumstances will you use this data to:\n  (1) allow, enable, or otherwise support the transmission by email, telephone,\nor facsimile of mass, unsolicited, commercial advertising, or spam; or\n  (2) enable high volume, automated, or electronic processes that send queries,\ndata, or email to MarkMonitor (or its systems) or the domain name contacts (or\nits systems).\n\nMarkMonitor reserves the right to modify these terms at any time.\n\nBy submitting this query, you agree to abide by this policy.\n\nMarkMonitor Domain Management(TM)\nProtecting companies and consumers in a digital world.\n\nVisit MarkMonitor at https://www.markmonitor.com\nContact us at +1.8007459229\nIn Europe, at +44.02032062220\n--\n\n;; Query time: 6912 msec\n;; WHEN: Thu Feb 09 17:59:12 UTC 2023\n"
```
