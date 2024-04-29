<template>
  <div class="navbar">
    <router-link to="/">Home</router-link>
    <router-link to="/innova">Innova Discs</router-link>
    <router-link to="/discraft">Discraft Discs</router-link>
    <router-link to="/dynamicdiscs">Dynamic Discs</router-link>
  </div>
</template>

<style scoped>
.navbar {
  display: flex;
  flex-direction: column;
  padding: 20px;
  background-color: #f0f0f0;
  height: 100vh; /* full height of the viewport */
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
.navbar a {
  padding: 10px;
  margin: 5px;
  text-decoration: none;
  color: #333;
}
.navbar a:hover {
  background-color: #ddd;
}
</style>
