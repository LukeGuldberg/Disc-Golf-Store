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
        <span class="spec">Speed: {{ disc.Speed }}</span>
        <span class="spec">Glide: {{ disc.Glide }}</span>
        <span class="spec">Turn: {{ disc.Turn }}</span>
        <span class="spec">Fade: {{ disc.Fade }}</span>
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
  return `/images/${props.disc.Name.toLowerCase()}.png`;
});
</script>

<style scoped>
.disc-card {
  display: flex;
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

.image-container {
  flex: 0 0 auto;
}

.image-container img {
  width: 100%;
  height: auto; 
  display: block;
  max-width: 150px;
}

.content {
  padding: 15px;
  background-color: #fff;
  flex-grow: 1;
}

.spec {
  margin-right: 10px;
}
</style>
