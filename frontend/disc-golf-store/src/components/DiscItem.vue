<template>
  <div class="disc-item">
    <h3>{{ disc.Name }}</h3> <!-- Capitalize the first letter -->
    <p>Type: {{ disc.Type }}</p> <!-- Capitalize the first letter -->
    <p>Speed: {{ disc.Speed }}, Glide: {{ disc.Glide }}, Turn: {{ disc.Turn }}, Fade: {{ disc.Fade }}</p> <!-- Capitalize the first letter -->
    <p>Manufacturer: {{ disc.Manufacturer }}</p> <!-- Capitalize the first letter -->
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
  },
  mounted() {
    console.log("Disc received in component:", this.disc); // This will now display properly formatted disc data
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
