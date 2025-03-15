<template>
    <div class="room-container">
      <h2>Typing Test Room: {{ roomID || 'Not Joined' }}</h2>
      <input v-model="username" placeholder="Enter username" class="input" />
      <input v-model="roomID" placeholder="Enter room ID" class="input" />
      <button @click="joinRoom" class="btn">Join Room</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        username: "",
        roomID: "",
      };
    },
    methods: {
      joinRoom() {
        if (!this.username || !this.roomID) {
          alert("Enter username and room ID!");
          return;
        }
  
        // Save username and roomID to sessionStorage for better security
        sessionStorage.setItem("username", this.username);
        sessionStorage.setItem("roomID", this.roomID);
  
        // Navigate to TypingTest.vue
        this.$router.push("/typing-test");
      },
    },
  };
  </script>
  
  
  

<style>
@import url('https://fonts.googleapis.com/css2?family=IBM+Plex+Sans+Thai+Looped:wght@100;200;300;400;500;600;700&display=swap');

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-family: "IBM Plex Sans Thai Looped", sans-serif;
}

.main-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  min-height: 100vh;
  background: #f4f4f4;
}

h2 {
  margin-bottom: 10px;
}

.room-container {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
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

.typing-container {
  margin-top: 20px;
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 1000px;
  text-align: center;
}

.typing-text {
  font-size: 1.4em;

  background: #eef;
  padding: 10px;
  border-radius: 5px;
  margin: 10px 0;
}

.typing-area {
  width: 100%;
  height: 100px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.wpm-list {
  list-style: none;
  padding: 0;
}

.wpm-list li {
  background: #ddd;
  margin: 5px 0;
  padding: 5px;
  border-radius: 5px;
}

.player-list {
  list-style: none;
  padding: 0;
  display: flex;
  flex-wrap: wrap;
}

.player-list li {
  background: #ddd;
  margin: 5px;
  padding: 5px;
  border-radius: 5px;
  width: calc(33.33% - 10px);
  /* Ensures 3 items per row */
  box-sizing: border-box;
}


@media (max-width: 600px) {
  .typing-container {
    max-width: 100%;
  }

  .room-container {
    max-width: 100%;
  }
}
</style>
  