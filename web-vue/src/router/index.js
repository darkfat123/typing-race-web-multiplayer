import { createRouter, createWebHistory } from 'vue-router';
import Room from '@/views/Room.vue';
import TypingTest from '@/views/TypingTest.vue';
import CreateRoom from '@/views/CreateRoom.vue';

const routes = [
  { path: '/create-room', name: 'CreateRoom', component: CreateRoom },
  { path: '/', name: 'Room', component: Room },
  { path: '/typing-test', name: 'TypingTest', component: TypingTest },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
