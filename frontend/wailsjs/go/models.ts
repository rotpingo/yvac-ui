export namespace main {
	
	export class ytData {
	    Url: string;
	    StartHH: string;
	    StartMM: string;
	    StartSS: string;
	    EndHH: string;
	    EndMM: string;
	    EndSS: string;
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new ytData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Url = source["Url"];
	        this.StartHH = source["StartHH"];
	        this.StartMM = source["StartMM"];
	        this.StartSS = source["StartSS"];
	        this.EndHH = source["EndHH"];
	        this.EndMM = source["EndMM"];
	        this.EndSS = source["EndSS"];
	        this.Name = source["Name"];
	    }
	}

}

