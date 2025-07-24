export type Ingredient = { name: string };

export function normalizeIngredients(
  input: Ingredient[][]
): { name: string; count: number }[] {
  const map = new Map<string, number>();

  input.flat().forEach(({ name }) => {
    map.set(name, (map.get(name) || 0) + 1);
  });

  return Array.from(map.entries()).map(([name, count]) => ({ name, count }));
}
