<script setup lang="ts">
  import { normalizeIngredients } from '../utils/normalizeIngredients';

  const props = defineProps<{
    ingredients: { name: string }[];
    index: number;
  }>();

  const grouped = normalizeIngredients(props.ingredients);
</script>

<template>
  <div class="bg-white border border-gray-200 rounded-xl p-4 shadow">
    <h3 class="font-semibold text-lg mb-2">🧪 Рецепт {{ index + 1 }}</h3>
    <ul class="space-y-1 text-sm">
      <li v-for="ing in grouped" :key="ing.name" class="flex justify-between">
        <span>{{ ing.name }}</span>
        <span class="text-gray-500">×{{ ing.count }}</span>
      </li>
    </ul>
  </div>
</template>
