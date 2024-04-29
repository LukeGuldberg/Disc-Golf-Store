<template>
  <div>
    <h1>Disc Golf Store</h1>
    <div v-if="discs.length">
      <DiscCard v-for="disc in discs" :key="disc.discId" :disc="disc" />
    </div>
    <div v-else>
      <p>Loading discs...</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import DiscCard from './components/DiscCard.vue';

const discs = ref([]);

onMounted(async () => {
  try {
    const response = await fetch('http://localhost:8000/getdiscs');
    if (!response.ok) throw new Error('Network response was not ok');
    discs.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch discs:', error);
  }
});
</script>

<style scoped>
/* Styles specific to Home component */
</style>
