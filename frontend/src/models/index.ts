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
    translit: string;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.id = source["id"];
      this.name = source["name"];
      this.proportions = source["proportions"] || [];
      this.translit = source["translit"];
    }

    get title(): string {
      return `${this.translit} (${this.name})`
    }
  }

  export class Cauldron {
	  id: string;
    name: string;
    capacity: number;
    magmin: number;
    translit: string;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.id = source["id"];
      this.name = source["name"];
      this.capacity = source["capacity"];
      this.magmin = source["magmin"];
      this.translit = source["translit"];
    }

    get title(): string {
      return `${this.translit} (${this.name})`
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

  export class InventoryCell {
    ingredient: Ingredient;
    cell_number: number;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.ingredient = source["ingredient"];
      this.cell_number = source["cell_number"];
    }
  }
}
