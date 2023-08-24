/* Do not change, this code is generated from Golang structs */


export class Bio {
    bio: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.bio = source["bio"];
    }
}
export class Account {
    service: string;
    id: string;
    username: string;
    url: string;
    profilePicture: {[key: string]: Picture};
    bio: {[key: string]: Bio};
    firstname: string;
    lastname: string;
    location: string;
    created: string;
    updated: string;
    blog: string;
    followers: number;
    following: number;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.service = source["service"];
        this.id = source["id"];
        this.username = source["username"];
        this.url = source["url"];
        this.profilePicture = this.convertValues(source["profilePicture"], Picture, true);
        this.bio = this.convertValues(source["bio"], Bio, true);
        this.firstname = source["firstname"];
        this.lastname = source["lastname"];
        this.location = source["location"];
        this.created = source["created"];
        this.updated = source["updated"];
        this.blog = source["blog"];
        this.followers = source["followers"];
        this.following = source["following"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class Tag {
    name: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.name = source["name"];
    }
}
export class Errors {
    exists: any;
    info: any;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.exists = source["exists"];
        this.info = source["info"];
    }
}
export class AccountInfo {
    url: string;
    profile_picture: { latest: { data: string } } ;
    bio: { latest: { data: {bio: string} } };

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.url = source["url"];
        this.profile_picture = source["profile_picture"];
        this.bio = source["bio"];
    }
}
export class Service {
    name: string;
    domain: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.name = source["name"];
        this.domain = source["domain"];
    }
}
export class User {
    Username: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.Username = source["Username"];
    }
}
export class InputData {
    user: User;
    service: Service;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.user = this.convertValues(source["user"], User);
        this.service = this.convertValues(source["service"], Service);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class ServiceCheckResult {
    input_data: InputData;
    exists: boolean;
    info: AccountInfo;
    errors: Errors;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.input_data = this.convertValues(source["input_data"], InputData);
        this.exists = source["exists"];
        this.info = this.convertValues(source["info"], AccountInfo);
        this.errors = this.convertValues(source["errors"], Errors);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class Source {
    url: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.url = source["url"];
    }
}
export class Club {
    club: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.club = source["club"];
    }
}
export class EmailService {
    name: string;
    link: string;
    username: string;
    icon: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.name = source["name"];
        this.link = source["link"];
        this.username = source["username"];
        this.icon = source["icon"];
    }
}
export class Email {
    mail: string;
    value: number;
    src: string;
    services: {[key: string]: EmailService};
    skipped_services: {[key: string]: boolean};
    valid: boolean;
    provider: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.mail = source["mail"];
        this.value = source["value"];
        this.src = source["src"];
        this.services = this.convertValues(source["services"], EmailService, true);
        this.skipped_services = source["skipped_services"];
        this.valid = source["valid"];
        this.provider = source["provider"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class Hobby {
    hobby: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.hobby = source["hobby"];
    }
}
export class Ip {
    ip: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.ip = source["ip"];
    }
}
export class Number {
    Valid: boolean;
    RawLocal: string;
    Local: string;
    E164: string;
    International: string;
    CountryCode: number;
    Country: string;
    Carrier: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.Valid = source["Valid"];
        this.RawLocal = source["RawLocal"];
        this.Local = source["Local"];
        this.E164 = source["E164"];
        this.International = source["International"];
        this.CountryCode = source["CountryCode"];
        this.Country = source["Country"];
        this.Carrier = source["Carrier"];
    }
}
export class PhoneNumber {
    number: string;
    valid: boolean;
    national_format: string;
    tag: string;
    phoneinfoga: Number;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.number = source["number"];
        this.valid = source["valid"];
        this.national_format = source["national_format"];
        this.tag = source["tag"];
        this.phoneinfoga = this.convertValues(source["phoneinfoga"], Number);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class Picture {
    img: string;
    img_hash: number;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.img = source["img"];
        this.img_hash = source["img_hash"];
    }
}
export class Person {
    id: string;
    name: string;
    gender: string;
    ethnicity: string;
    pictures: {[key: string]: Picture};
    maidenname: string;
    age: number;
    bday: string;
    address: string;
    phone: {[key: string]: PhoneNumber};
    ips: {[key: string]: Ip};
    civilstatus: string;
    kids: string;
    hobbies: {[key: string]: Hobby};
    email: {[key: string]: Email};
    occupation: string;
    prevoccupation: string;
    education: string;
    military: string;
    religion: string;
    pets: string;
    clubs: {[key: string]: Club};
    legal: string;
    political: string;
    notes: string;
    relations: {[key: string]: string[]};
    sources: {[key: string]: Source};
    accounts: {[key: string]: ServiceCheckResult};
    tags: Tag[];
    notaccounts: {[key: string]: Account};
    custom: any;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"] || '';
        this.name = source["name"] || '';
        this.gender = source["gender"] || '';
        this.ethnicity = source["ethnicity"] || '';
        this.pictures = this.convertValues(source["pictures"], Picture, true);
        this.maidenname = source["maidenname"] || '';
        this.age = source["age"];
        this.bday = source["bday"] || '';
        this.address = source["address"] || '';
        this.phone = this.convertValues(source["phone"], PhoneNumber, true);
        this.ips = this.convertValues(source["ips"], Ip, true);
        this.civilstatus = source["civilstatus"] || '';
        this.kids = source["kids"] || '';
        this.hobbies = this.convertValues(source["hobbies"], Hobby, true);
        this.email = this.convertValues(source["email"], Email, true);
        this.occupation = source["occupation"] || '';
        this.prevoccupation = source["prevoccupation"] || '';
        this.education = source["education"] || '';
        this.military = source["military"] || '';
        this.religion = source["religion"] || '';
        this.pets = source["pets"] || '';
        this.clubs = this.convertValues(source["clubs"], Club, true);
        this.legal = source["legal"] || '';
        this.political = source["political"] || '';
        this.notes = source["notes"] || '';
        this.relations = source["relations"];
        this.sources = this.convertValues(source["sources"], Source, true);
        this.accounts = this.convertValues(source["accounts"], ServiceCheckResult, true);
        this.tags = this.convertValues(source["tags"], Tag);
        this.notaccounts = this.convertValues(source["notaccounts"], Account, true);
        this.custom = source["custom"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}