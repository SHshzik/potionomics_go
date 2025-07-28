<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import SelectInput from './SelectInput.vue';
  import { GetPotions } from '../client/potions';
  import { potions } from '../models/potions';
  
  const selectedPotion = defineModel<potions.Potion>();
  const allPotions = ref<potions.Potion[]>([]);

  onMounted(() => {
    GetPotions().then((result: potions.BDPotions) => {
      allPotions.value = Object.values(result);
    })
  });
</script>

<template>
  <SelectInput
    v-model="selectedPotion"
    :options="allPotions"
    label="Выберите зелье"
  />
</template>
