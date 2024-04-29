<template>
  <div>
    <h1>Search Results for {{ query }}</h1>
    <div v-if="error">
      <p>Error: {{ error }}</p>
    </div>
    <div v-else-if="disc">
      <DiscCard :disc="disc" />
    </div>
    <div v-else>
      <p>No results found.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import DiscCard from './components/DiscCard.vue';

const route = useRoute();
const disc = ref(null);
const query = ref(''); 
const error = ref(null); 

onMounted(() => {
  query.value = route.query.name;
  if (query.value) {
    fetchDisc(query.value);
  }
});

watch(() => route.query.name, (newQuery) => {
  query.value = newQuery;
  fetchDisc(newQuery);
}, { immediate: true });

async function fetchDisc(name) {
  try {
    const response = await fetch(`http://localhost:8000/getdiscsbyname?name=${name}`);
    if (response.ok) {
      disc.value = await response.json();
      error.value = null; 
    } else {
      throw new Error('Failed to fetch disc');
    }
  } catch (error) {
    console.error('Error fetching disc:', error);
    disc.value = null; 
    error.value = error.message;
  }
}
</script>

