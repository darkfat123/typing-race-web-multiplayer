<template>
  <div class="room-container">
    <div class="row">
      <MenuButton label="Back" to="/" class="back-btn" @click="handleBack" />
      <h3 class="new-rooms">Create New Room</h3>
    </div>
    <input type="text" v-model="username" placeholder="Enter username" class="input" />

    <h4>Number of players</h4>
    <div class="number-control">
      <button class="control-btn" @click="decreaseValue">-</button>
      <div class="number-box">{{ value }}</div>
      <button class="control-btn" @click="increaseValue">+</button>
    </div>

    <div class="language-selector">
      <button :class="['lang-btn', { selected: language === 'th' }]" @click="selectLanguage('th')">
        Thai
      </button>
      <button :class="['lang-btn', { selected: language === 'en' }]" @click="selectLanguage('en')">
        English
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
      value: 1,
    };
  },
  methods: {
    selectLanguage(lang) {
      this.language = lang;
    },
    handleBack() {
      sessionStorage.removeItem("username")
      sessionStorage.removeItem("language")
      sessionStorage.removeItem("roomID")
    },
    joinRoom() {
      if (!this.username) {
        alert("Enter username!");
        return;
      }
      console.log("Value of max_players before setting:", this.value);
      sessionStorage.setItem("username", this.username);
      sessionStorage.setItem("language", this.language);
      localStorage.setItem("max_players", this.value);
      this.$router.push("/typing-test");
    },
    increaseValue() {
      if (this.value < 10) {
        this.value++;
      }
    },
    decreaseValue() {
      if (this.value > 1) {
        this.value--;
      }
    }
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

.number-control {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 10px 0;
}


.control-btn {
  width: 50px;
  height: 50px;
  font-size: 24px;
  background-color: var(--text-color);
  color: var(--bg-color);
  border: none;
  cursor: pointer;
  border-radius: 5px;
  margin: 0 10px;
  transition: 0.2s;
}

.control-btn:hover {
  background-color: #ccc;
}

.number-box {
  width: 100px;
  height: 48px;
  font-size: 18px;
  border: 2px solid #ccc;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
}


.btn {
  width: 100%;
  padding: 10px;
  font-weight: bold;
  background-color: var(--main-btn-color);
  color: var(--text-color);
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  box-shadow: 0 0px 8px var(--shadow-color);
  transition: background-color 0.3s, color 0.3s;
  text-align: center;
  transition: background 0.3s ease, transform 0.2s ease;
}

.btn:hover {
  background: linear-gradient(to right, var(--bg-color), goldenrod);
  transform: scale(1.01);
}

.language-selector {
  gap: 10px;
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 10px;
  margin-top: 20px;
}

.lang-btn {
  flex: 1;
  padding: 10px 0;
  border-radius: 5px;
  font-size: 16px;
  border: 1px solid #ccc;
  color: var(--text-color);
  background-color: transparent;
  cursor: pointer;
  transition: background 0.3s ease, transform 0.2s ease;
}

.lang-btn.selected {
  background-color: var(--main-btn-color);
  color: var(--text-color);
}

.lang-btn:hover {
  color: black;
  background-color: #f1f1f1;
  transform: scale(1.01);
}

.back-btn {
  background-color: rgb(153, 33, 33);
  position: absolute;
  width: 100px;
  color: white;
}

.back-btn:hover {
  background: none;
  background-color: var(--text-color);
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
