<template>
  <div class="home">
    <h1>Disc Golf Store</h1>
    <div v-if="discs && discs.length">
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
      discs: []
    };
  },
  mounted() {
    this.fetchDiscs();
  },
  methods: {
  fetchDiscs() {
    fetch('http://localhost:8000/getdiscs', {
      method: 'POST'
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(data => {
      console.log("Data received:", data); // This will show us the structure of the received data
      this.discs = data;
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
