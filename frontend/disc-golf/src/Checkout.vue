<template>
  <div class="checkout-page">
    <h1>Checkout</h1>
    <div v-if="cartItems.length">
      <div class="cart-item" v-for="item in cartItems" :key="item.DiscId">
        <img :src="`/images/${item.ImageFileName}`" :alt="item.Name" class="cart-image">
        <div class="item-details">
          <h2>{{ item.Name }}</h2>
          <p>{{ item.Description }}</p>
          <p>Price: ${{ item.Price.toFixed(2) }}</p>
        </div>
      </div>
    </div>
    <div v-else>
      <p>Your cart is empty.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const cartItems = ref([]);

onMounted(() => {
  cartItems.value = JSON.parse(localStorage.getItem('cart')) || [];
});
</script>

<style scoped>
.checkout-page {
  padding: 20px;
}

.cart-item {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.cart-image {
  width: 100px;
  height: 100px;
  object-fit: cover;
  margin-right: 20px;
}

.item-details {
  flex-grow: 1;
}
</style>
