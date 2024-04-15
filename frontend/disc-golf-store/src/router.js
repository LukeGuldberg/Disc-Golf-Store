import { createRouter, createWebHistory } from 'vue-router';
import Home from './components/Home.vue';
import DiscCategory from './components/DiscCategory.vue';
import DiscDetails from './components/DiscDetails.vue';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
        },
        {
            path: '/category/:id',
            name: 'disc-category',
            component: DiscCategory,
            props: true,
        },
        {
            path: '/disc/:id',
            name: 'disc-details',
            component: DiscDetails,
            props: true,
        },
    ],
});

export default router;
