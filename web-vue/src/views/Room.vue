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
  
        sessionStorage.setItem("username", this.username);
        sessionStorage.setItem("roomID", this.roomID);
  
        this.$router.push("/typing-test");
      },
    },
  };
  </script>
  
  
  

<style>
.room-container {
  background: var(--bg-color);
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0px 8px var(--shadow-color); /* ใช้ตัวแปรสีเงา */
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

@media (max-width: 600px) {
  .room-container {
    max-width: 100%;
  }
}
</style>
  