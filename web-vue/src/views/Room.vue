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

  <div class="room-list">
    <h3>Available Rooms</h3>
    <div v-if="Object.keys(filteredRoomList).length === 0">
      <p>No active rooms.</p>
    </div>
    <div v-else>
      <div v-for="(users, roomID) in filteredRoomList" :key="roomID" class="room-card">
        <h4>Room ID: {{ roomID }}</h4>
        <p v-if="users && users.length > 0">Users: {{ users.join(", ") }}</p>
        <button @click="joinRoom">Join This Room</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      roomID: "",
      language: "th",
      roomList: {}, // เพิ่ม state เพื่อเก็บ room list
      socket: null,
    };
  },
  mounted() {
    this.connectWebSocket();
  },
  computed: {
    // กรองเฉพาะห้องที่มีผู้ใช้อยู่
    filteredRoomList() {
      const filtered = {};
      for (const [roomID, users] of Object.entries(this.roomList)) {
        if (users && users.length > 0) {
          filtered[roomID] = users;
        }
      }
      return filtered;
    },
  },
  methods: {
    selectLanguage(lang) {
      this.language = lang;
    },
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

    connectWebSocket() {
      this.socket = new WebSocket(import.meta.env.VITE_WS_URL + "/ws/lobby");

      this.socket.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.type === "room_list") {
          this.roomList = data.roomList;
        }
      };

      this.socket.onerror = (err) => {
        console.error("WebSocket error:", err);
      };
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

.room-card {
  border: 1px solid #ccc;
  padding: 12px;
  margin: 10px 0;
  border-radius: 8px;
  background: #f9f9f9;
}


@media (max-width: 600px) {
  .room-container {
    max-width: 100%;
  }
}
</style>
