<template>
  <div class="checkout-page">
    <h1>Checkout</h1>
    <div v-if="cartItems.length">
      <div class="cart-item" v-for="(item, index) in cartItems" :key="item.DiscId">
        <img :src="`/images/${item.ImageFileName}`" :alt="item.Name" class="cart-image">
        <div class="item-details">
          <h2>{{ item.Name }}</h2>
          <p>{{ item.Description }}</p>
          <p>Price: ${{ item.Price.toFixed(2) }}</p>
          <button @click="removeItem(index)">Remove</button>
        </div>
      </div>
      <button @click="checkout">Checkout</button>
    </div>
    <div v-else-if="showSuccessMessage">
      <p class="success-message">Checkout successful! Thank you for your purchase.</p>
    </div>
    <div v-else>
      <p>Your cart is empty.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const cartItems = ref([]);
const showSuccessMessage = ref(false);

onMounted(() => {
  cartItems.value = JSON.parse(localStorage.getItem('cart')) || [];
});

function removeItem(index) {
  cartItems.value.splice(index, 1);
  localStorage.setItem('cart', JSON.stringify(cartItems.value));
}

function checkout() {
  localStorage.removeItem('cart');
  cartItems.value = [];
  showSuccessMessage.value = true;
  setTimeout(() => showSuccessMessage.value = false, 5000); // Hide the message after 5 seconds
}
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

button {
  padding: 10px;
  margin: 10px 0;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
}

button:hover {
  background-color: #45a049;
}

.success-message {
  color: green;
  font-size: 1.2em;
  font-weight: bold;
  padding: 20px;
  margin-top: 20px;
  background-color: #ccffcc;
  border-radius: 5px;
}
</style>
