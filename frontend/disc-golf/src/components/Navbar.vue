<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const searchQuery = ref('');
const router = useRouter();

function searchDisc() {
  if (searchQuery.value.trim() !== '') {
    router.push({ name: 'search', query: { name: searchQuery.value.trim() } })
      .catch(err => {
        if (err.name !== 'NavigationDuplicated') {
          console.error(err);
        }
      });
    searchQuery.value = '';
  }
}
</script>

<template>
  <div class="navbar">
    <div class="nav-items">
      <router-link to="/" class="nav-item">Home</router-link>
      <router-link to="/innova" class="nav-item">Innova Discs</router-link>
      <router-link to="/discraft" class="nav-item">Discraft Discs</router-link>
      <router-link to="/dynamicdiscs" class="nav-item">Dynamic Discs</router-link>
      <router-link to="/checkout" class="nav-item">Checkout</router-link>
    </div>
    <form @submit.prevent="searchDisc" class="search-form">
      <input type="text" v-model="searchQuery" placeholder="Search Discs" />
      <button type="submit">Search</button>
    </form>
  </div>
</template>

<style scoped>
.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 20px;
  background-color: #333;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.3);
}

.nav-items {
  display: flex;
  align-items: center;
}

.nav-item {
  color: #fff;
  text-decoration: none;
  margin-right: 20px;
  transition: color 0.3s;
}

.nav-item:last-child {
  margin-right: 0;
}

.nav-item:hover {
  color: #aaa;
}

.search-form {
  display: flex;
  align-items: center;
}

.search-form input {
  padding: 8px;
  margin-right: 10px;
  border: 1px solid #555;
  border-radius: 4px;
  background-color: #fff;
  color: #333;
}

.search-form button {
  padding: 8px 16px;
  background-color: #555;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.search-form button:hover {
  background-color: #777;
}
</style>
