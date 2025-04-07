<template>
  <div class="room-container">
    <div class="row">
      <h3 class="available-rooms">Available Rooms</h3>
      <MenuButton label="Create New Room" to="/create-room" class="create-room-btn" />
    </div>

    <div v-if="Object.keys(filteredRoomList).length === 0">
      <p>No active rooms.</p>
    </div>

    <ul v-else class="room-list">
      <li v-for="(users, roomID) in filteredRoomList" :key="roomID" class="room-card">
        <h4>🔑 Room ID: {{ roomID }}</h4>
        <p v-if="users && users.length > 0">Users: {{ users.join(", ") }}</p>
        <div class="room-footer">
          <button class="join-btn" @click="joinRoom(roomID)">Join This Room</button>
        </div>
      </li>
    </ul>
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
      roomID: "",
      language: "th",
      roomList: {},
      socket: null,
    };
  },
  mounted() {
    this.connectWebSocket();
  },
  computed: {
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
  margin: 1rem 1rem 2rem 1rem;
  position: relative;
}

.row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding: 20px;
}

.available-rooms {
  margin: 0;
  flex-grow: 1;
  text-align: center;
}

.create-room-btn {
  position: absolute;
  right: 20px;
  width: auto;
}

.room-footer {
  margin-top: auto;
  /* ให้ปุ่มอยู่ติดขอบล่าง */
}

.room-list {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 2rem;
  padding: 0;
  list-style: none;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.room-card {
  background-color: var(--bg-color);
  color: var(--text-color);
  padding: 1rem;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  transform: translateY(-4px);
  box-shadow: 0 8px 24px var(--shadow-color);


}

.join-btn {
  background: linear-gradient(to right, var(--text-color), #333333);
  color: var(--bg-color);
  padding: 0.6rem 1.2rem;
  border: none;
  width: 100%;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.3s ease, transform 0.2s ease;
}

.join-btn:hover {
  background: linear-gradient(to right, #333333, var(--text-color));
  transform: scale(1.03);
}


@media (max-width: 768px) {
  .room-container {
    max-width: 100%;
  }

  .create-room-btn {
    position: static;
    margin: 10px 0;
  }

  .room-list {
    grid-template-columns: 1fr;
  }

  .available-rooms {
    margin: 0;
    flex-grow: 1;
    text-align: left;
  }
}
</style>