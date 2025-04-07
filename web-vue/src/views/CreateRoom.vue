<template>
  <div class="room-container">
    <div class="row">
      <MenuButton label="Back" to="/" class="back-btn" />
      <h3 class="new-rooms">Create New Room</h3>
    </div>
    <input v-model="username" placeholder="Enter username" class="input" />

    <div class="language-selector">
      <button :class="['lang-btn', { selected: language === 'th' }]" @click="selectLanguage('th')">
        TH
      </button>
      <button :class="['lang-btn', { selected: language === 'en' }]" @click="selectLanguage('en')">
        EN
      </button>
    </div>

    <button @click="joinRoom" class="btn">Create Room</button>
  </div>
</template>

<script>
import MenuButton from '../components/MenuButton.vue';

export default {
  components: {
    MenuButton
  },
  data() {
    return {
      username: "",
      language: "th",
      roomList: {},
      socket: null,
    };
  },
  methods: {
    selectLanguage(lang) {
      this.language = lang;
    },
    joinRoom() {
      if (!this.username) {
        alert("Enter username and room ID!");
        return;
      }
      sessionStorage.setItem("username", this.username);
      sessionStorage.setItem("language", this.language);
      this.$router.push("/typing-test");
    },
  },
};
</script>

<style scoped>
.room-container {
  background: var(--bg-color);
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0px 8px var(--shadow-color);
  text-align: center;
  width: 100%;
  max-width: 1000px;
  margin: 1rem 1rem 2rem 1rem;
}

.input {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.btn {
  width: 100%;
  padding: 10px;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.btn:hover {
  background: #218838;
}

.language-selector {
  gap: 10px;
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 10px;
}

.lang-btn {
  flex: 1;
  padding: 10px 0;
  font-size: 16px;
  border: 1px solid #ccc;
  color: var(--text-color);
  background-color: transparent;
  cursor: pointer;
  transition: background-color 0.1s ease;
}

.lang-btn.selected {
  background-color: #4CAF50;
  color: white;
}

.lang-btn:hover {
  color: black;
  background-color: #f1f1f1;
}

.back-btn {
  background-color: rgb(153, 33, 33);
  position: absolute;
  width: 100px;
}

.row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;
  margin-top: 20px;
  position: relative;
}

.new-rooms {
  flex: 1;
  text-align: center;
}

@media (max-width: 600px) {
  .room-container {
    max-width: 100%;
  }

  .new-rooms {
    margin: 0;
    flex-grow: 1;
    text-align: right;
  }
}
</style>
