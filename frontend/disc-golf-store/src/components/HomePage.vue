<template>
  <div class="home">
    <h1>Disc Golf Store</h1>
    <div v-if="discs.length">
      <DiscItem v-for="disc in discs" :key="disc.discId" :disc="disc" />
    </div>
    <div v-else>
      Loading discs or none available...
    </div>
  </div>
</template>

<script>
import DiscItem from '@/components/DiscItem.vue';

export default {
  name: 'HomePage',
  components: {
    DiscItem
  },
  data() {
    return {
      discs: [] // Array to store fetched discs
    };
  },
  mounted() {
    this.fetchDiscs();
  },
  methods: {
    fetchDiscs() {
      fetch('http://localhost:8080/getdiscs') // Update URL as per your backend
        .then(response => response.json())
        .then(data => {
          this.discs = data; // Assign fetched data to discs array
        })
        .catch(error => {
          console.error('Error fetching discs:', error);
        });
    }
  }
}
</script>

<style scoped>
.home {
  max-width: 800px;
  margin: auto;
  padding: 20px;
}
</style>
