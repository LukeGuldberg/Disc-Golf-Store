<template>
  <div id="app">
    <Navbar />
    <div class="content">
      <router-view/>
    </div>
  </div>
</template>

<script setup>
import Navbar from './components/Navbar.vue';
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background: #f0f2f5;
}

#app {
  display: flex;
  min-height: 100vh;
}

/* Ensure the content div starts to the right of the navbar */
.content {
  flex-grow: 1;
  padding: 20px;
  margin-left: 250px; /* Adjust this value based on the actual width of your navbar */
  background: #fff;
  box-shadow: -1px 0 10px rgba(0,0,0,0.1);
  overflow-y: auto;
}

a {
  color: #007BFF;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}

*, *::before, *::after {
  box-sizing: border-box;
}
</style>
