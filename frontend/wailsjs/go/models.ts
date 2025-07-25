export namespace main {
	
	export class ytData {
	
	
	    static createFrom(source: any = {}) {
	        return new ytData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

