<template>
  <div class="disc-card">
    <div class="image-container">
      <img :src="imageURL" :alt="disc.Name" />
    </div>
    <div class="content">
      <h2>{{ disc.Name }}</h2>
      <p class="type">{{ disc.Type }}</p>
      <p class="manufacturer">{{ disc.Manufacturer }}</p>
      <div class="specs">
        <span>Speed: {{ disc.Speed }}</span>
        <span>Glide: {{ disc.Glide }}</span>
        <span>Turn: {{ disc.Turn }}</span>
        <span>Fade: {{ disc.Fade }}</span>
      </div>
      <p class="description">{{ disc.Description }}</p>
      <p class="price">Price: ${{ disc.Price.toFixed(2) }}</p>
    </div>
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
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  margin: 10px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.disc-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 12px rgba(0,0,0,0.15);
}

.image-container img {
  width: 100%;
  height: auto;
  display: block; 
}

.content {
  padding: 15px;
  background-color: #fff;
}

.content h2 {
  margin-top: 0;
  color: #333;
}

.specs {
  display: flex;
  justify-content: space-between;
  font-size: 0.9em;
  margin: 10px 0;
}

.specs span {
  background: #f4f4f4;
  padding: 3px 6px;
  border-radius: 4px;
}

.description {
  font-size: 0.9em;
  color: #666;
  margin: 10px 0;
}

.price {
  font-weight: bold;
  color: #2c3e50;
}
</style>
