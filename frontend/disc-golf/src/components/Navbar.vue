<template>
  <div class="navbar">
    <router-link to="/" class="nav-item">Home</router-link>
    <router-link to="/innova" class="nav-item">Innova Discs</router-link>
    <router-link to="/discraft" class="nav-item">Discraft Discs</router-link>
    <router-link to="/dynamicdiscs" class="nav-item">Dynamic Discs</router-link>
  </div>
</template>

<style scoped>
.navbar {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 30px 20px;
  background-color: #333; 
  height: 100vh;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
  width: 250px; 
  position: fixed; 
  top: 0;
  left: 0;
  z-index: 1000;
}

.nav-item {
  display: block;
  padding: 12px 20px;
  margin: 5px 0;
  width: 100%;
  text-decoration: none;
  color: #fff;
  font-size: 16px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: transparent;
  text-align: left;
  transition: background-color 0.3s, color 0.3s, padding-left 0.3s;
}

.nav-item:hover, .nav-item:active {
  background-color: #555;
  color: #fff;
  padding-left: 30px; 
}

@media (max-width: 768px) {
  .navbar {
    width: 100%;
    height: auto;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    padding: 10px 0;
  }

  .nav-item {
    padding: 10px;
    text-align: center;
    flex-grow: 1;
  }
}
</style>
