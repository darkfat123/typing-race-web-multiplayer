import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faMoon, faSun } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faMoon, faSun)


const app = createApp(App);
app.component('FontAwesomeIcon', FontAwesomeIcon)
// Register FontAwesomeIcon as global component

app.use(router).mount('#app');
