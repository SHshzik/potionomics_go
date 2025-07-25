import { domain } from '../models/index'

export type Ingredient = { name: string };

export function normalizeIngredients( input: { name: string }[] ): { name: string; count: number }[] {
  const map = new Map<string, number>();

  input.forEach(({ name }) => {
    map.set(name, (map.get(name) || 0) + 1);
  });

  return Array.from(map.entries()).map(([name, count]) => ({ name, count }));
}

export function normalizeInventory(input: domain.InventoryCell[]): { name: string, count: number }[] {
  const map = new Map<string, number>();

  input.forEach(({ ingredient: { name, translit } }) => {
    const fullName = `${name} - ${translit}`
    map.set(fullName, (map.get(fullName) || 0) + 1);
  });

  return Array.from(map.entries()).map(([name, count]) => ({ name, count }));
}
