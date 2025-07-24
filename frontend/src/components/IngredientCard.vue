<script lang="ts" setup>
  defineProps<{
    ingredient: { name: string; count: number };
  }>();
</script>

<template>
  <div class="bg-white shadow rounded p-4 flex justify-between items-center">
    <div class="text-lg font-semibold">{{ ingredient.name }}</div>
    <div class="text-sm text-gray-600">x{{ ingredient.count }}</div>
  </div>
</template>
