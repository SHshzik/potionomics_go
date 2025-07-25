const BASE_URL = 'http://localhost:8080';
import { domain } from '../models';

export async function GetPotions(): Promise<domain.Potion[]> {
  const res = await fetch(`${BASE_URL}/get_potions`);
  if (!res.ok) throw new Error('Ошибка при получении зелий');
  const data = await res.json();
  return data.map((item: any) => new domain.Potion(item));
}

export async function GetCauldrons(): Promise<domain.Cauldron[]> {
  const res = await fetch(`${BASE_URL}/get_cauldrons`);
  if (!res.ok) throw new Error('Ошибка при получении котлов');
  const data = await res.json();
  return data.map((item: any) => new domain.Cauldron(item));
}

export async function Generate(potionId: string, cauldronId: string): Promise<domain.BrewResult[]> {
  const res = await fetch(`${BASE_URL}/generate`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ potion_id: potionId, cauldron_id: cauldronId })
  });
  if (!res.ok) throw new Error('Ошибка при генерации рецепта');
  const data = await res.json();
  return data.map((item: any) => new domain.BrewResult(item));
}
