export namespace potions {
  export interface BDPotions {
    [key: number]: Potion
  }

  export class Potion {
    id: number;
    name: string;
    tier: number;
    baseValue: number;
    a: number;
    b: number;
    c: number;
    d: number;
    e: number;
    balance: number[];
    translit: string;

    constructor(source: any = {}) {
      if (typeof source === 'string') source = JSON.parse(source);
      this.id = source["id"];
      this.name = source["name"];
      this.tier = source["tier"];
      this.baseValue = source["base_value"]
      this.a = source["a"]
      this.b = source["b"]
      this.c = source["c"]
      this.d = source["d"]
      this.e = source["e"]
      this.balance = source["balance"];
      this.translit = source["translit"];
    }

    get title(): string {
      return `${this.translit} (${this.name})`
    }
  }
}
