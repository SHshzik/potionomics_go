import { BASE_URL } from "../config";
import { cauldrons } from "../models/cauldrons";

export async function GetCauldrons(): Promise<cauldrons.BDCauldrons> {
  const res = await fetch(`${BASE_URL}/get_cauldrons`);
  if (!res.ok) throw new Error('Ошибка при получении котлов');
  const data = await res.json();
  const result: cauldrons.BDCauldrons = {};

  for (const key in data) {
    const numericKey = Number(key);
    result[numericKey] = new cauldrons.Cauldron(data[key]);
  }
  
  return result;
}
