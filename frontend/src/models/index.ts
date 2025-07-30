export namespace domain {
  export class Ingredient {
    name: string;
    translit: string;
    A: number;
    B: number;
    C: number;
    D: number;
    E: number;
    basePrice: number;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.name = source["name"];
      this.translit = source["translit"];
      this.A = source["A"];
      this.B = source["B"];
      this.C = source["C"];
      this.D = source["D"];
      this.E = source["E"];
      this.basePrice = source["base_price"];
    }

    get title(): string {
      return `${this.translit} (${this.name})`
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
