import { BASE_URL } from "../config";
import { potions } from "../models/potions";

export async function GetPotions(): Promise<potions.BDPotions> {
  const res = await fetch(`${BASE_URL}/get_potions`);
  if (!res.ok) throw new Error('Ошибка при получении зелий');
  const data = await res.json();

  const result: potions.BDPotions = {};

  for (const key in data) {
    const numericKey = Number(key);
    result[numericKey] = new potions.Potion(data[key]);
  }

  return result;
}
