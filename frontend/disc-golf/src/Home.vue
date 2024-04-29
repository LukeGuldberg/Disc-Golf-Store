<template>
  <div>
    <h1>Disc Golf Store</h1>
    <div v-if="discs.length">
      <DiscCard v-for="disc in discs" :key="disc.discId" :disc="disc" @checkout="handleCheckout" />
    </div>
    <div v-else>
      <p>Loading discs...</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import DiscCard from './components/DiscCard.vue';

const discs = ref([]);
const router = useRouter();

onMounted(async () => {
  try {
    const response = await fetch('http://localhost:8000/getdiscs');
    if (!response.ok) throw new Error('Network response was not ok');
    discs.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch discs:', error);
  }
});

const handleCheckout = (disc) => {
  router.push({ name: 'checkout', params: { discId: disc.DiscId } });
};
</script>
