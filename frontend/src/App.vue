<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import SelectInput from './components/SelectInput.vue';
  import IngredientCard from './components/IngredientCard.vue';
  import { normalizeIngredients } from './utils/normalizeIngredients';
  import { GetPotions, GetCauldrons } from '../wailsjs/go/service/App';

  type Ingredient = { name: string };

  const potions = ref<string[]>([]);
  const cauldrons = ref<string[]>([]);
  const selectedPotion = ref('');
  const selectedCauldron = ref('');
  const result = ref<{ name: string; count: number }[]>([]);

  onMounted(() => {
    GetPotions().then((result: { Name: string }[]) => {
      potions.value = result.map((p: { Name: string }) => p.Name);
    })
    GetCauldrons().then((result: { Name: string }[]) => {
      cauldrons.value = result.map((p: { Name: string }) => p.Name);
    })
  });

  function handleBrew() {
    const input: Ingredient[][] = [
      [{ name: 'Слизь' }, { name: 'Коготь' }],
      [{ name: 'Слизь' }, { name: 'Корень' }]
    ];
    result.value = normalizeIngredients(input);
  }
</script>

<template>
  <div class="min-h-screen bg-gray-100 p-6">
    <div class="max-w-3xl mx-auto space-y-6">
      <h1 class="text-3xl font-bold text-center">🧪 Potionomics Helper</h1>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <SelectInput
          v-model="selectedPotion"
          :options="potions"
          label="Выберите зелье"
        />
        <SelectInput
          v-model="selectedCauldron"
          :options="cauldrons"
          label="Выберите котёл"
        />
      </div>

      <p>{{ selectedPotion }}</p>
      <p>{{ selectedCauldron }}</p>

      <button
        @click="handleBrew"
        class="bg-purple-600 text-white py-2 px-4 rounded hover:bg-purple-700 transition"
      >
        Собрать рецепт
      </button>

      <div v-if="result.length" class="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-4">
        <IngredientCard
          v-for="ingredient in result"
          :key="ingredient.name"
          :ingredient="ingredient"
        />
      </div>
    </div>
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
