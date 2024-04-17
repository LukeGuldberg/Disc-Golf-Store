import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '@/components/HomePage.vue';
import DiscItem from '@/components/DiscItem.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home-page',
      component: HomePage
    },
    {
      path: '/item',
      name: 'disc-item',
      component: DiscItem
    }
  ]
});

export default router;
