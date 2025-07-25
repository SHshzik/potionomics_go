<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import SelectInput from './components/SelectInput.vue';
  import RecipeBlock from './components/ReceiptBlock.vue';
  import { GetPotions, GetCauldrons, Generate } from './client/app';
  import { domain } from './models';

  const potions = ref<domain.Potion[]>([]);
  const cauldrons = ref<domain.Cauldron[]>([]);
  const selectedPotion = ref<domain.Potion>();
  const selectedCauldron = ref<domain.Cauldron>();
  const receipts = ref<domain.BrewResult[]>([]);

  onMounted(() => {
    GetPotions().then((result: domain.Potion[]) => {
      potions.value = result;
    })
    GetCauldrons().then((result: domain.Cauldron[]) => {
      cauldrons.value = result;
    })
  });

  function handleBrew() {
    if (!selectedPotion.value || !selectedCauldron.value) {
      return;
    }

    Generate(selectedPotion.value.id, selectedCauldron.value.id).then((result: domain.BrewResult[]) => {
      receipts.value = result;
    });
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

      <div v-if="receipts.length" class="grid gap-4 pt-6">
        <RecipeBlock
          v-for="(recipe, i) in receipts"
          :key="recipe.id"
          :ingredients="recipe.Receipt"
          :index="i"
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
