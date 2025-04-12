<template>
  <div class="room-container">
    <div class="row">
      <h3 class="available-rooms">Available Rooms</h3>
      <MenuButton label="Create New Room" to="/create-room" class="create-room-btn" />
    </div>
    <input type="text" v-model="searchRoomID" placeholder="Search Room ID" class="input" />
    <LanguageSelector :selectedLang="language" @update:lang="language = $event" class="language-selector" />
    <div v-if="Object.keys(filteredRoomList).length === 0">
      <p>No active rooms.</p>
    </div>

    <ul v-else class="room-list">
      <li v-for="(room, roomID) in filteredRoomList" :key="roomID" class="room-card">
        <h4>
          🔑 Room ID: {{ roomID }}
          <span v-if="room.users.length > 0" class="room-user-count">({{ room.users.length }}/{{ room.limit }} users)</span>
        </h4>
        <p v-if="room.users && room.users.length > 0">{{ room.users.join(", ") }}</p>
        <p class="room-language">🌐 Language: {{ room.language }}</p>
        <div class="room-footer">
          <button class="join-btn" @click="openModal(roomID)">Join This Room</button>
        </div>
      </li>

    </ul>


    <!-- Username Modal -->
    <div v-if="showUsernameModal" class="modal-overlay">
      <div class="modal-content">
        <h3>Enter Your Username</h3>
        <input v-model="username" placeholder="Username" class="input" />
        <div class="modal-buttons">
          <button @click="closeModal" class="btn cancel">Cancel</button>
          <button @click="confirmJoin" class="btn">Join</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import MenuButton from '../components/MenuButton.vue';
import LanguageSelector from '../components/LanguageSelector.vue';
sessionStorage.removeItem("username")
sessionStorage.removeItem("language")
sessionStorage.removeItem("roomID")
export default {
  components: {
    MenuButton,
    LanguageSelector
  },
  data() {
    return {
      username: "",
      roomID: "",
      selectedRoomID: "",
      searchRoomID: "",
      showUsernameModal: false,
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
      for (const [roomID, room] of Object.entries(this.roomList)) {
        if (
          (this.searchRoomID && roomID.includes(this.searchRoomID)) ||
          !this.searchRoomID
        ) {
          filtered[roomID] = room;
        }
      }
      return filtered;
    }
    ,
  },
  methods: {
    openModal(roomID) {
      this.selectedRoomID = roomID;
      this.username = "";
      this.showUsernameModal = true;
    },
    closeModal() {
      this.showUsernameModal = false;
    },
    confirmJoin() {
      if (!this.username) {
        alert("Please enter your username!");
        return;
      }
      sessionStorage.setItem("username", this.username);
      sessionStorage.setItem("roomID", this.selectedRoomID);
      sessionStorage.setItem("language", this.language);
      this.$router.push("/typing-test");
    },
    updateRoomList(newData) {
      const filtered = {};
      for (const [roomID, room] of Object.entries(newData)) {
        if (room.users && room.users.length > 0) {
          filtered[roomID] = room;
        }
      }
      this.roomList = filtered;
    }
    ,
    connectWebSocket() {
      this.socket = new WebSocket(import.meta.env.VITE_WS_URL + "/ws/lobby");

      this.socket.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.type === "room_list") {
          this.updateRoomList(data.roomList);
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
  margin: 1rem 1rem 1rem 1rem;
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
  background-color: var(--main-btn-color);
  color: var(--text-color);
  position: absolute;
  right: 20px;
  width: auto;
}

.room-footer {
  margin-top: auto;
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
  background: linear-gradient(to right, var(--text-color), #6d6d6d);
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
  background: linear-gradient(to right, #757575, var(--text-color));
  transform: scale(1.03);
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.room-user-count {
  font-size: 0.9rem;
  color: #888;
  margin-left: 8px;
  font-weight: bold;
}

.modal-content {
  background: var(--bg-color);
  color: var(--text-color);
  padding: 1.5rem;
  border-radius: 10px;
  width: 90%;
  max-width: 400px;
  text-align: center;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.2);
}

.input {
  width: 100%;
  padding: 10px;
  margin-top: 1rem;
  margin-bottom: 0px;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.modal-buttons {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  gap: 10px;
}

.btn {
  flex: 1;
  padding: 10px;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.btn:hover {
  background: #0056b3;
}

.btn.cancel {
  background: #dc3545;
}

.btn.cancel:hover {
  background: #c82333;
}

.language-selector {
  margin-bottom: 25px;
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