/* Do not change, this code is generated from Golang structs */


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