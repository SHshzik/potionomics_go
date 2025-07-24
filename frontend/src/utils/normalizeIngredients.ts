export type Ingredient = { name: string };

export function normalizeIngredients( input: { name: string }[] ): { name: string; count: number }[] {
  const map = new Map<string, number>();

  input.forEach(({ name }) => {
    map.set(name, (map.get(name) || 0) + 1);
  });

  return Array.from(map.entries()).map(([name, count]) => ({ name, count }));
}
