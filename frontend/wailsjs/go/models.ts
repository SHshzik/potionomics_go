export namespace domain {
	
	export class Ingredient {
	    name: string;
	    A: number;
	    B: number;
	    C: number;
	    D: number;
	    E: number;
	
	    static createFrom(source: any = {}) {
	        return new Ingredient(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.A = source["A"];
	        this.B = source["B"];
	        this.C = source["C"];
	        this.D = source["D"];
	        this.E = source["E"];
	    }
	}
	export class BrewResult {
	    id: string;
	    Receipt: Ingredient[];
	
	    static createFrom(source: any = {}) {
	        return new BrewResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.Receipt = this.convertValues(source["Receipt"], Ingredient);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Cauldron {
	    name: string;
	    capacity: number;
	    magmin: number;
	
	    static createFrom(source: any = {}) {
	        return new Cauldron(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.capacity = source["capacity"];
	        this.magmin = source["magmin"];
	    }
	}
	export class Potion {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Potion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class GenerateRequest {
	    Potion: Potion;
	    Cauldron: Cauldron;
	
	    static createFrom(source: any = {}) {
	        return new GenerateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Potion = this.convertValues(source["Potion"], Potion);
	        this.Cauldron = this.convertValues(source["Cauldron"], Cauldron);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	

}

