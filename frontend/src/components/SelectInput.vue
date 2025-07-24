<script lang="ts" setup>
  const modelValue = defineModel<any>();

  defineProps<{
    options: { name: string }[];
    label: string;
  }>();

  defineEmits(['update:modelValue']);
</script>

<template>
  <div>
    <label class="block text-sm font-medium mb-1">{{ label }}</label>
    <select
      v-model="modelValue"
      class="w-full border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-500"
    >
      <option value="" disabled>-- Выберите --</option>
      <option v-for="opt in options" :key="opt.name" :value="opt">{{ opt.name }}</option>
    </select>
  </div>
</template>
