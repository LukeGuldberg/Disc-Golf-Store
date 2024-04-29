<template>
  <div>
    <h1>Disc Golf Store</h1>
    <section v-if="putters.length">
      <h2>Putters</h2>
      <div class="discs-container">
        <DiscCard v-for="disc in putters" :key="disc.DiscId" :disc="disc" />
      </div>
    </section>
    <section v-if="midranges.length">
      <h2>Midranges</h2>
      <div class="discs-container">
        <DiscCard v-for="disc in midranges" :key="disc.DiscId" :disc="disc" />
      </div>
    </section>
    <section v-if="fairwayDrivers.length">
      <h2>Fairway Drivers</h2>
      <div class="discs-container">
        <DiscCard v-for="disc in fairwayDrivers" :key="disc.DiscId" :disc="disc" />
      </div>
    </section>
    <section v-if="distanceDrivers.length">
      <h2>Distance Drivers</h2>
      <div class="discs-container">
        <DiscCard v-for="disc in distanceDrivers" :key="disc.DiscId" :disc="disc" />
      </div>
    </section>
    <div v-if="!putters.length && !midranges.length && !fairwayDrivers.length && !distanceDrivers.length">
      <p>Loading discs...</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import DiscCard from './components/DiscCard.vue';

const putters = ref([]);
const midranges = ref([]);
const fairwayDrivers = ref([]);
const distanceDrivers = ref([]);

onMounted(() => {
  fetchDiscsByType('Putter', putters);
  fetchDiscsByType('Midrange', midranges);
  fetchDiscsByType('FairwayDriver', fairwayDrivers);
  fetchDiscsByType('DistanceDriver', distanceDrivers);
});

async function fetchDiscsByType(type, targetArray) {
  try {
    const response = await fetch(`http://localhost:8000/get${type.toLowerCase()}s`);
    if (!response.ok) throw new Error('Network response was not ok');
    targetArray.value = await response.json();
  } catch (error) {
    console.error(`Failed to fetch ${type}:`, error);
  }
}
</script>

<style scoped>
.discs-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

section {
  margin-bottom: 20px;
}
</style>
