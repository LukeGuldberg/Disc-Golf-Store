<template>
  <div>
    <!-- Ensure that the heading updates by using the reactive query variable -->
    <h1>Search Results for {{ query }}</h1>
    <div v-if="disc">
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
const query = ref(''); // Reactive reference for the query string

onMounted(() => {
  // Set the initial query and fetch the disc on mount
  query.value = route.query.name;
  if (query.value) {
    fetchDisc(query.value);
  }
});

// Watch for changes in the route query and update accordingly
watch(() => route.query.name, (newQuery) => {
  query.value = newQuery;
  fetchDisc(newQuery);
}, { immediate: true });

async function fetchDisc(name) {
  try {
    const response = await fetch(`http://localhost:8000/getdiscbyname?name=${name}`);
    if (response.ok) {
      disc.value = await response.json();
    } else {
      throw new Error('Failed to fetch disc');
    }
  } catch (error) {
    console.error('Error fetching disc:', error);
  }
}
</script>
