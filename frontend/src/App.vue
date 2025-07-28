<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import PotionSelect from './components/PotionSelect.vue';
  import CauldronSelect from './components/CauldronSelect.vue';
  import RecipeBlock from './components/ReceiptBlock.vue';
  import { Generate, GetInventory, GetShop } from './client/app';
  import { domain } from './models';
  import { potions } from './models/potions';
  import { cauldrons } from './models/cauldrons';
  import { normalizeInventory } from './utils/normalizeIngredients'

  const selectedPotion = ref<potions.Potion>();
  const selectedCauldron = ref<cauldrons.Cauldron>();
  const withShop = ref<boolean>(false);
  const isStrict = ref<boolean>(false);
  const load = ref<boolean>(false);

  const receipts = ref<domain.BrewResult[]>([]);
  const inventory = ref<{ name: string, count: number }[]>([]);
  const shop = ref<{ name: string, count: number }[]>([]);

  function handleBrew() {
    if (!selectedPotion.value || !selectedCauldron.value) {
      return;
    }

    load.value = true
    receipts.value = []

    Generate(selectedPotion.value.id, selectedCauldron.value.id, withShop.value, isStrict.value).then((result: domain.BrewResult[]) => {
      load.value = false
      receipts.value = result;
    });
  }

  function showInventory() {
    GetInventory().then((result: domain.InventoryCell[]) => {
      inventory.value = normalizeInventory(result)
    })
  }

  function showShop() {
    GetShop().then((result: domain.InventoryCell[]) => {
      shop.value = normalizeInventory(result)
    })
  }
</script>

<template>
  <div class="min-h-screen bg-gray-100 p-6">
    <div class="max-w-3xl mx-auto space-y-6">
      <h1 class="text-3xl font-bold text-center">🧪 Potionomics Helper</h1>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <PotionSelect v-model="selectedPotion" />
        <CauldronSelect v-model="selectedCauldron" />

        <div class="flex items-center space-x-4">
          <label class="inline-flex items-center">
            <input type="checkbox" v-model="withShop" class="form-checkbox text-indigo-600 rounded" />
            <span class="ml-2">с магазином</span>
          </label>
          <label class="inline-flex items-center">
            <input type="checkbox" v-model="isStrict" class="form-checkbox text-indigo-600 rounded" />
            <span class="ml-2">Строго</span>
          </label>
        </div>
      </div>

      <button
        @click="handleBrew"
        class="bg-purple-600 text-white py-2 px-4 rounded hover:bg-purple-700 transition"
      >
        Собрать рецепт
      </button>

      <div v-if="load" class="flex justify-center mt-6">
        <svg class="animate-spin h-6 w-6 text-indigo-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v4l3-3-3-3v4a8 8 0 000 16v-4l-3 3 3 3v-4a8 8 0 01-8-8z" />
        </svg>
      </div>

      <div v-if="receipts.length" class="grid gap-4 pt-6">
        <RecipeBlock
          v-for="(recipe, i) in receipts"
          :key="recipe.id"
          :ingredients="recipe.Receipt"
          :index="i"
        />
      </div>

      <br />

      <button
        @click="showInventory"
        class="bg-purple-600 text-white py-2 px-4 rounded hover:bg-purple-700 transition"
      >
        Показать инвентарь
      </button>

      <ul class="space-y-1 text-sm">
        <li v-for="ing in inventory" :key="ing.name" class="flex justify-between">
          <span>{{ ing.name }} x {{ ing.count }}</span>
        </li>
      </ul>

      <br />

      <button
        @click="showShop"
        class="bg-purple-600 text-white py-2 px-4 rounded hover:bg-purple-700 transition"
      >
        Показать магазин
      </button>

      <ul class="space-y-1 text-sm">
        <li v-for="ing in shop" :key="ing.name" class="flex justify-between">
          <span>{{ ing.name }} x {{ ing.count }}</span>
        </li>
      </ul>
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
