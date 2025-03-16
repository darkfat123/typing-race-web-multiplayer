<template>
    <div v-if="connected" class="typing-container">
        <h3>Players in this room:</h3>
        <ul class="player-list">
            <li v-for="player in playersInRoom" :key="player">
                {{ player }} <span v-if="readyPlayers.includes(player)">✅ Ready</span>
            </li>
        </ul>
        <div class="button-container">
            <button @click="goBack" class="btn-back">Back</button>
            <button v-if="!isReady" @click="sendReadyFlag" class="btn">Ready</button>
        </div>
    </div>

    <div v-if="connected && isGameStarted" class="typing-container">
        <h3>Type this message:</h3>
        <p class="typing-text">{{ givenText }}</p>
        <textarea v-model="inputText" @input="sendText" class="typing-area"></textarea>
        <h3>Live WPM:</h3>
        <ul class="wpm-list">
            <li v-for="(wpm, user) in wpmData" :key="user">{{ user }}: {{ wpm }} WPM</li>
        </ul>
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
            playersInRoom: [],
            readyPlayers: [],
            ws: null,
            wpmData: {},
            connected: false,
            isReady: false,
            isGameStarted: false,
            finishSound: (() => {
                const sound = new Audio("/complete.wav");
                sound.volume = 0.45;
                return sound;
            })(),
        };
    },
    created() {
        // Retrieve username & roomID from sessionStorage, not from URL
        this.username = sessionStorage.getItem("username") || "";
        this.roomID = sessionStorage.getItem("roomID") || "";

        if (!this.username || !this.roomID) {
            alert("Invalid username or room ID!");
            this.$router.push("/");
            return;
        }

        this.connectWebSocket();
    },
    methods: {
        goBack() {
            // Send a close message and then close the WebSocket connection
            if (this.ws && this.connected) {
                const message = { type: "close", username: this.username, roomID: this.roomID };
                this.ws.send(JSON.stringify(message)); // Notify backend before closing
                this.ws.close();
            }
            this.$router.push('/');
        },
        connectWebSocket() {
            if (this.ws) {
                this.ws.close();
            }

            this.ws = new WebSocket(import.meta.env.VITE_WS_URL + "/ws");

            this.ws.onopen = () => {
                this.connected = true;
                this.isReady = false;
                this.isGameStarted = false;
                this.ws.send(JSON.stringify({ username: this.username, roomID: this.roomID }));
            };

            this.ws.onmessage = (event) => {
                const data = JSON.parse(event.data);

                if (data.type === "update_users") {
                    this.playersInRoom = data.users;
                }
                if (data.type === "update_ready") {
                    this.readyPlayers = data.users;
                }
                if (data.type === "start_game") {
                    this.isGameStarted = true;
                    alert("Game has started! Type the given text as fast as you can.");
                }
                if (data.text) {
                    this.givenText = data.text;
                }
                if (data.wpm !== undefined) {
                    this.wpmData[data.username] = data.wpm.toFixed(2);
                }
                if (data.type === "finished") {
                    this.finishSound.play();
                }
            };

            this.ws.onerror = (err) => {
                console.error("WebSocket error:", err);
            };

            this.ws.onclose = (event) => {
                if (event.code !== 1000) { // Ensure it's a normal closure
                    console.log("WebSocket closed unexpectedly");
                }
                this.connected = false;
            };
        },
        sendText() {
            if (this.ws && this.connected) {
                this.ws.send(JSON.stringify({ text: this.inputText }));
            }
        },
        sendReadyFlag() {
            if (this.ws && this.connected) {
                this.isReady = true;
                this.ws.send(JSON.stringify({ status: "ready" }));
            }
        },
    },
    beforeUnmount() {
        if (this.ws) {
            const message = { type: "close", username: this.username, roomID: this.roomID };
            this.ws.send(JSON.stringify(message)); 
            this.ws.close();
        }
    },
};
</script>




<style>

.input {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
}

.button-container {
    display: flex;
    gap: 10px; /* Adds space between buttons */
}

.btn-back {
    width: 100%;
    padding: 10px;
    background: #960d0d;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
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
}
</style>