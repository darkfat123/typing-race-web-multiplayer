import { createRouter, createWebHistory } from 'vue-router';
import Room from '@/views/Room.vue';
import TypingTest from '@/views/TypingTest.vue';

const routes = [
  { path: '/', name: 'Home', component: Room },
  { path: '/typing-test', name: 'TypingTest', component: TypingTest },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
