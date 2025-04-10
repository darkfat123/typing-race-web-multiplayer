<template>
    <div v-if="connected" class="typing-container">
        <h3>Players in this room:</h3>
        <ul class="player-list">
            <li v-for="player in playersInRoom" :key="player">
                {{ player }} <span v-if="(readyPlayers || []).includes(player)">✅ Ready</span>
            </li>
        </ul>
        <div class="button-container">
            <button @click="goBack" class="btn-back">Back</button>
            <button v-if="!isReady" @click="sendReadyFlag" class="btn">Ready</button>
        </div>
    </div>

    <div v-if="connected && isGameStarted" class="typing-container">
        <h3>Type this message:</h3>
        <p class="typing-text">
            <span v-for="(char, index) in givenText" :key="index" :class="getCharClass(index)">
                {{ char }}
            </span>
        </p>
        <textarea v-model="inputText" @input="sendText" class="typing-area"></textarea>
        <h3>Live WPM:</h3>
        <ul class="wpm-list">
            <li v-for="(wpm, user) in wpmData" :key="user">{{ user }}: {{ wpm }} WPM</li>
        </ul>
    </div>

    <div v-if="isCountingDown" class="overlay">
        <div class="countdown">{{ countdownValue }}</div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: "",
            roomID: "",
            language: "",
            inputText: "",
            givenText: "",
            playersInRoom: [],
            readyPlayers: [],
            ws: null,
            wpmData: {},
            connected: false,
            isReady: false,
            isGameStarted: false,
            countdown: null,
            countdownValue: 3,
            isCountingDown: false,
            finishSound: (() => {
                const sound = new Audio("/complete.wav");
                sound.volume = 0.45;
                return sound;
            })(),
            startSound: (() => {
                const startsound = new Audio("/countdown.mp3");
                startsound.volume = 0.15;
                return startsound;
            })(),
        };
    },
    created() {
        this.username = sessionStorage.getItem("username") || "";
        this.roomID = sessionStorage.getItem("roomID") || "";
        this.language = sessionStorage.getItem("language") || "";

        if (!this.username) {
            alert("Invalid username or room ID!");
            this.$router.push("/");
            return;
        }

        this.connectWebSocket();
    },
    methods: {
        goBack() {
            if (this.ws && this.connected) {
                const message = { type: "close", username: this.username, roomID: this.roomID, language: this.language };
                this.ws.send(JSON.stringify(message)); // Notify backend before closing
                this.ws.close();
            }
            sessionStorage.removeItem("roomID")
            this.$router.push('/');
        },
        connectWebSocket() {
            if (this.ws) {
                this.ws.close();
            }

            this.ws = new WebSocket(import.meta.env.VITE_WS_URL + "/ws/typing");

            this.ws.onopen = () => {
                this.connected = true;
                this.isReady = false;
                this.isGameStarted = false;
                this.ws.send(JSON.stringify({ username: this.username, roomID: this.roomID, language: this.language }));
            };

            this.ws.onmessage = (event) => {
                const data = JSON.parse(event.data);

                if (data.type === "update_users") {
                    this.playersInRoom = data.users;
                }
                if (data.type === "update_ready") {
                    this.readyPlayers = Array.isArray(data.users) ? data.users : [];
                }
                if (data.type === "start_game") {
                    this.isGameStarted = true;
                    this.startCountdown();
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
                if (event.code !== 1000) {
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
        getCharClass(index) {
            if (!this.inputText[index]) return "default";
            return this.inputText[index] == this.givenText[index] ? "correct" : "incorrect";
        },
        sendReadyFlag() {
            if (this.ws && this.connected) {
                this.isReady = true;
                this.ws.send(JSON.stringify({ status: "ready" }));
            }
        },
        startCountdown() {
            this.isCountingDown = true;
            this.countdownValue = 3;
            this.startSound.play();
            this.countdown = setInterval(() => {
                this.countdownValue--;
                if (this.countdownValue === 0) {
                    clearInterval(this.countdown);
                    this.isCountingDown = false;
                    this.isGameStarted = true;
                }
            }, 1000);
        }


    },
    beforeUnmount() {
        if (this.ws) {
            this.ws.close();
        }
    },
};
</script>

<style>
.overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.7);
    z-index: 999;
    display: flex;
    justify-content: center;
    align-items: center;
}

.countdown {
    font-size: 100px;
    color: white;
    font-weight: bold;
    animation: scaleUp 0.8s ease-in-out;
}

@keyframes scaleUp {
    0% {
        transform: scale(0.5);
        opacity: 0;
    }

    100% {
        transform: scale(1);
        opacity: 1;
    }
}

.input {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
}

.button-container {
    display: flex;
    gap: 10px;
    margin-top: 10px;
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
    background: var(--bg-color);
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 0px 8px var(--shadow-color);
    width: 100%;
    max-width: 1000px;
    text-align: center;
}

.typing-text {
    font-size: 1.4em;

    background: var(--bg-color);
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
    background: var(--bg-color);
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
    background: transparent;
    margin: 5px;
    padding: 5px;
    box-shadow: 0 0px 8px var(--shadow-color);
    border-radius: 5px;
    width: calc(33.33% - 10px);
    box-sizing: border-box;
}

.correct {
    color: #A0C878;
}

.incorrect {
    background-color: #D2665A;
    color: black;
}


@media (max-width: 600px) {
    .typing-container {
        max-width: 100%;
    }
}
</style>