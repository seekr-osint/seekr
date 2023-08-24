/* Do not change, this code is generated from Golang structs */


export class General {
    force_port: boolean;
    browser: boolean;
    discord: boolean;
    workers: number;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.force_port = source["force_port"];
        this.browser = source["browser"];
        this.discord = source["discord"];
        this.workers = source["workers"];
    }
}
export class Server {
    port: number;
    ip: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.port = source["port"];
        this.ip = source["ip"];
    }
}
export class Config {
    server: Server;
    general: General;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.server = this.convertValues(source["server"], Server);
        this.general = this.convertValues(source["general"], General);
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