export namespace cauldrons {
  export interface BDCauldrons {
    [key: number]: Cauldron
  }

  export class Cauldron {
	id: number;
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
      return `${this.translit} (${this.name}) - ${this.capacity} - ${this.magmin}`
    }
  }
}
