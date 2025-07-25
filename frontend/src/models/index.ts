export namespace domain {
  export class Ingredient {
    name: string;
    translit: string;
    A: number;
    B: number;
    C: number;
    D: number;
    E: number;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.name = source["name"];
      this.translit = source["translit"];
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

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.id = source["id"];
      this.Receipt = (source["Receipt"] || []).map((item: any) => new Ingredient(item));
    }
  }

  export class Potion {
    id: string;
    name: string;
    proportions: number[];

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.id = source["id"];
      this.name = source["name"];
      this.proportions = source["proportions"] || [];
    }
  }

  export class Cauldron {
	id: string;
    name: string;
    capacity: number;
    magmin: number;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.id = source["id"];
      this.name = source["name"];
      this.capacity = source["capacity"];
      this.magmin = source["magmin"];
    }
  }

  export class GenerateRequest {
	potionId: string;
	cauldronId: string

    constructor(source: any = {}) {
		if (typeof source === 'string') source = JSON.parse(source);
		this.potionId = source["potionId"];
		this.cauldronId = source["cauldronId"];
	  }
  }
}
