<script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import SelectInput from './SelectInput.vue';
  import { GetCauldrons } from '../client/cauldrons';
  import { cauldrons } from '../models/cauldrons';
  
  const selectedPotion = defineModel<cauldrons.Cauldron>();
  const allPCauldrons = ref<cauldrons.Cauldron[]>([]);

  onMounted(() => {
    GetCauldrons().then((result: cauldrons.BDCauldrons) => {
      allPCauldrons.value = Object.values(result);
    })
  });
</script>

<template>
  <SelectInput
    v-model="selectedPotion"
    :options="allPCauldrons"
    label="Выберите котёл"
  />
</template>
