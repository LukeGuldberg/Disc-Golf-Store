<template>
  <div class="disc-card">
    <img :src="imageURL" :alt="disc.Name" />
    <h2>{{ disc.Name }}</h2>
    <p>Type: {{ disc.Type }}</p>
    <p>Manufacturer: {{ disc.Manufacturer }}</p>
    <p>Speed: {{ disc.Speed }}</p>
    <p>Glide: {{ disc.Glide }}</p>
    <p>Turn: {{ disc.Turn }}</p>
    <p>Fade: {{ disc.Fade }}</p>
    <p>Description: {{ disc.Description }}</p>
    <p>Price: ${{ disc.Price.toFixed(2) }}</p>
  </div>
</template>

<script setup>
import { defineProps, computed } from 'vue';

const props = defineProps({
  disc: {
    type: Object,
    required: true
  }
});

const imageURL = computed(() => {
  return new URL(`../assets/${props.disc.ImageFileName}`, import.meta.url).href;
});
</script>

<style scoped>
.disc-card {
  border: 1px solid #ccc;
  margin: 10px;
  padding: 10px;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.disc-card img {
  max-width: 100%;
  height: auto;
}
</style>
