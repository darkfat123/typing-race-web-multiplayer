<template>
  <div class="room-container">
    <h2>Typing Test Room: {{ roomID || 'Not Joined' }}</h2>
    
    <input v-model="username" placeholder="Enter username" class="input" />
    <input v-model="roomID" placeholder="Enter room ID" class="input" />
    
    <div class="language-selector">
      <button :class="['lang-btn', { selected: language === 'th' }]" @click="selectLanguage('th')">
        TH
      </button>
      <button :class="['lang-btn', { selected: language === 'en' }]" @click="selectLanguage('en')">
        EN
      </button>
    </div>

    <button @click="joinRoom" class="btn">Join Room</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      roomID: "",
      language: "th",
    };
  },
  methods: {
    selectLanguage(lang) {
      this.language = lang;
    },

    // ฟังก์ชันที่ใช้ในการเข้าห้อง
    joinRoom() {
      if (!this.username || !this.roomID) {
        alert("Enter username and room ID!");
        return;
      }

      sessionStorage.setItem("username", this.username);
      sessionStorage.setItem("roomID", this.roomID);
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
  margin-left: 10px;
  margin-right: 10px;
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

@media (max-width: 600px) {
  .room-container {
    max-width: 100%;
  }
}
</style>
