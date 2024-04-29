<template>
  <div class="disc-card">
    <h2>{{ disc.name }}</h2>
    <p>Type: {{ disc.typ }}</p>
    <p>Manufacturer: {{ disc.manufacturer }}</p>
    <p>Speed: {{ disc.speed }}</p>
    <p>Glide: {{ disc.glide }}</p>
    <p>Turn: {{ disc.turn }}</p>
    <p>Fade: {{ disc.fade }}</p>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  disc: {
    type: Object,
    required: true
  }
});
</script>

<style scoped>
.disc-card {
  border: 1px solid #ccc;
  margin: 10px;
  padding: 10px;
  border-radius: 5px;
}
</style>
