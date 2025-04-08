<template>
  <div :class="{ 'dark-mode': isDarkMode }" class="main-container">
    <Header @toggleDarkMode="toggleDarkMode" />
    <router-view />
    <Footer />
  </div>
</template>

<script>
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';

export default {
  components: {
    Header,
    Footer,
  },
  data() {
    return {
      isDarkMode: localStorage.getItem('darkMode') === 'true'
    };
  },
  methods: {
    toggleDarkMode() {
      this.isDarkMode = !this.isDarkMode;
      localStorage.setItem('darkMode', this.isDarkMode);
    }
  }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Sans+Thai+Looped:wght@100;200;300;400;500;600;700&display=swap');

:root {
  --bg-color: white;
  --text-color: black;
  --shadow-color: rgba(0, 0, 0, 0.2);
  --main-btn-color: rgb(255, 183, 0);
}

.dark-mode {
  --bg-color: #1a1a1a;
  --text-color: #e0e0e0;
  --shadow-color: rgba(255, 255, 255, 0.2);
  --main-btn-color: rgb(172, 123, 0);
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-family: "IBM Plex Sans Thai Looped", sans-serif;
  transition: background 0.2s, color 0.2s;
}

.main-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  min-height: 100vh;
  background: var(--bg-color);
  color: var(--text-color);
}
</style>
