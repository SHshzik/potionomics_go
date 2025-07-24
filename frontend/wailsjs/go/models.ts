export namespace domain {
	
	export class Cauldron {
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new Cauldron(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	    }
	}
	export class Potion {
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new Potion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	    }
	}

}

