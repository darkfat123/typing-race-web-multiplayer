<template>
  <div>
    <h2>Typing Test Room: {{ roomID }}</h2>
    <input v-model="username" placeholder="Enter username" />
    <input v-model="roomID" placeholder="Enter room ID" />
    <button @click="joinRoom">Join Room</button>

    <div v-if="connected">
      <h3>พิมพ์ข้อความนี้:</h3>
      <p>{{ givenText }}</p>
      <textarea v-model="inputText" @input="sendText"></textarea>
      <h3>Live WPM:</h3>
      <ul>
        <li v-for="(wpm, user) in wpmData" :key="user">{{ user }}: {{ wpm }} WPM</li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      roomID: "",
      inputText: "",
      givenText: "",
      ws: null,
      wpmData: {},
      connected: false,
      finishSound: (() => {
        const sound = new Audio("/complete.wav");
        sound.volume = 0.45;
        return sound;
      })(),
    };
  },
  methods: {
    joinRoom() {
      if (!this.username || !this.roomID) {
        alert("Enter username and room ID!");
        return;
      }
      this.ws = new WebSocket(import.meta.env.VITE_WS_URL + "/ws");

      this.ws.onopen = () => {
        this.connected = true;
        this.ws.send(JSON.stringify({ username: this.username, roomID: this.roomID }));
      };
      this.ws.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.text) {
          this.givenText = data.text;
        } else {
          this.wpmData[data.username] = data.wpm;
        }

        if (data.status == "finished") {
          this.finishSound.play();
        }
      };
    },
    sendText() {
      if (this.ws && this.connected) {
        this.ws.send(JSON.stringify({ text: this.inputText }));
      }
    },
  },
};
</script>

<style>
textarea {
  width: 100%;
  height: 100px;
  margin-bottom: 10px;
}
</style>
