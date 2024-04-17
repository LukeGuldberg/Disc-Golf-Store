<template>
  <div class="disc-item">
    <h3>{{ disc.name }}</h3>
    <p>Type: {{ disc.type }}</p>
    <p>Speed: {{ disc.speed }}, Glide: {{ disc.glide }}, Turn: {{ disc.turn }}, Fade: {{ disc.fade }}</p>
    <p>Manufacturer: {{ disc.manufacturer }}</p>
  </div>
</template>

<script>
export default {
  name: 'DiscItem',
  props: {
    disc: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>
.disc-item {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 10px;
}
</style>
