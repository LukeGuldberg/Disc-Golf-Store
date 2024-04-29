<script setup>

</script>

<template>
  <div id="app">
    <div class="navbar">
      <router-link to="/">Home</router-link>
      <router-link to="/innova">Innova Discs</router-link>
      <router-link to="/discraft">Discraft Discs</router-link>
      <router-link to="/dynamicdiscs">Dynamic Discs</router-link>
    </div>
    <router-view/>
  </div>
</template>

<style>
/* styles as above */
</style>
