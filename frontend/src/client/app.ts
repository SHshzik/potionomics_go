import { BASE_URL } from '../config';
import { domain } from '../models';

export async function GetInventory(): Promise<domain.InventoryCell[]> {
  const res = await fetch(`${BASE_URL}/get_inventory`);
  if (!res.ok) throw new Error('Ошибка при получении инвентаря');
  const data = await res.json();
  return data.map((item: { ingredient: any, cell_number: number }) => {
    return new domain.InventoryCell({ ingredient: new domain.Ingredient(item.ingredient), cell_number: item.cell_number })
  });
}

export async function GetShop(): Promise<domain.InventoryCell[]> {
  const res = await fetch(`${BASE_URL}/get_shop`);
  if (!res.ok) throw new Error('Ошибка при получении магазина');
  const data = await res.json();
  return data.map((item: { ingredient: any, cell_number: number }) => {
    return new domain.InventoryCell({ ingredient: new domain.Ingredient(item.ingredient), cell_number: item.cell_number })
  });
}

export async function Generate(potionId: number, cauldronId: number): Promise<domain.BrewResult[]> {
  const res = await fetch(`${BASE_URL}/generate`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ potion_id: potionId, cauldron_id: cauldronId })
  });
  if (!res.ok) throw new Error('Ошибка при генерации рецепта');
  const data = await res.json();
  return data.map((item: any) => new domain.BrewResult(item));
}
