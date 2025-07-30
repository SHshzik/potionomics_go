<script setup lang="ts">
  import { computed } from 'vue';
  import { domain } from '../models/index';
  

  const props = defineProps<{
    ingredients: domain.Ingredient[];
    index: number;
  }>();

// Группируем ингредиенты по имени, суммируя энергии и количество
const grouped = computed(() => {
  const map = new Map<string, any>();

  props.ingredients.forEach((ing) => {
    if (!map.has(ing.name)) {
      map.set(ing.name, {
        ...ing,
        count: 1,
      });
    } else {
      const existing = map.get(ing.name);
      map.set(ing.name, {
        name: ing.name,
        A: existing.A + ing.A,
        B: existing.B + ing.B,
        C: existing.C + ing.C,
        D: existing.D + ing.D,
        E: existing.E + ing.E,
        basePrice: existing.basePrice + ing.basePrice,
        translit: ing.translit,
        count: existing.count + 1,
      });
    }
  });

  // Добавляем поле total
  return Array.from(map.values()).map((ing) => ({
    ...ing,
    total: ing.A + ing.B + ing.C + ing.D + ing.E,
    totalPrice: ing.basePrice,
  }));
});

// Общая сумма энергий по рецепту
const totalEnergy = computed(() => {
  return grouped.value.reduce(
    (acc, cur) => {
      acc.A += cur.A;
      acc.B += cur.B;
      acc.C += cur.C;
      acc.D += cur.D;
      acc.E += cur.E;
      acc.total += cur.total;
      acc.totalPrice += cur.totalPrice;
      return acc;
    },
    { A: 0, B: 0, C: 0, D: 0, E: 0, total: 0, totalPrice: 0 }
  );
});
</script>

<template>
  <div class="bg-white border border-gray-200 rounded-xl p-4 shadow">
    <h3 class="font-semibold text-lg mb-4">🧪 Рецепт {{ index + 1 }}</h3>

    <ul class="space-y-2 text-sm">
      <li
        v-for="ing in grouped"
        :key="ing.name"
        class="flex flex-col border-b pb-2 last:border-0"
      >
        <div class="flex justify-between font-medium">
          <span>{{ ing.translit }} ({{ ing.name }})</span>
          <span class="text-gray-500"
            >×{{ ing.count }} — Σ: {{ ing.total }} - $: {{ ing.totalPrice }}</span
          >
        </div>
        <div class="text-xs text-gray-600 flex gap-2 mt-1 flex-wrap">
          <span>A: {{ ing.A }}</span>
          <span>B: {{ ing.B }}</span>
          <span>C: {{ ing.C }}</span>
          <span>D: {{ ing.D }}</span>
          <span>E: {{ ing.E }}</span>
        </div>
      </li>
    </ul>

    <div class="mt-4 pt-2 border-t text-sm text-gray-700">
      <div class="font-semibold mb-1">Общая энергия:</div>
      <div class="flex flex-wrap gap-3 text-xs text-gray-600">
        <span>A: {{ totalEnergy.A }}</span>
        <span>B: {{ totalEnergy.B }}</span>
        <span>C: {{ totalEnergy.C }}</span>
        <span>D: {{ totalEnergy.D }}</span>
        <span>E: {{ totalEnergy.E }}</span>
        <span class="font-bold text-gray-800">Σ: {{ totalEnergy.total }}</span>
        <span class="font-bold text-gray-800">$: {{ totalEnergy.totalPrice }}</span>
      </div>
    </div>
  </div>
</template>
