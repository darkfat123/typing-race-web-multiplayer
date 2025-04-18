<template>
    <div v-if="connected" class="typing-container">
        <h3>Players in this room:</h3>
        <ul class="player-list">
            <li v-for="player in playersInRoom" :key="player"
                :class="{ 'ready': (readyPlayers || []).includes(player) }">
                {{ player }}
            </li>
        </ul>
        <div class="button-container">
            <button @click="goBack" class="btn-back">Back</button>

            <template v-if="!isGameStarted">
                <button v-if="!isReady" @click="sendReadyFlag" class="btn">Ready</button>
                <button v-else @click="sendReadyFlag" class="btn unready">Unready</button>
            </template>

            <button v-if="connected && isGameStarted && !hasVotedRestart" class="btn" @click="voteRestart">
                Vote to Restart
                <span v-if="restartVoteInfo.total > 0" class="vote-info"> {{ restartVoteInfo.votes }}/{{
                    restartVoteInfo.total }}
                </span>
            </button>


        </div>
    </div>

    <div v-if="connected && isGameStarted" class="typing-container" @keydown="handleKeydown" tabindex="0"
        ref="typingBox">
        <div class="row">
            <div class="current-wpm">
                <RoundedIcon label="current wpm:" color="#9FB3DF" />
                <span v-if="currentUser && wpmData[currentUser]">
                    {{ wpmData[currentUser] }} wpm
                </span>
            </div>
            <div class="timer">
                <RoundedIcon label="time:" color="#E38E49" />
                <span>{{ formattedElapsedTime }}</span>
            </div>
        </div>
        <p class="typing-text">
            <span v-for="(char, index) in givenText" :key="index" :class="[
                getCharClass(index),
                { 'has-cursor': index + 1 === inputText.length }
            ]">
                {{ char }}
            </span>
        </p>

        <ul class="wpm-list">
            <li v-for="(wpm, user, index) in wpmDataFinished" :key="user"> No. {{ index + 1 }} - {{ user }}: <span
                    class="wpm">{{ wpm }}</span> WPM</li>
        </ul>
    </div>

    <div v-if="isCountingDown" class="overlay">
        <div class="countdown">{{ countdownValue }}</div>
    </div>
</template>

<script>
import RoundedIcon from '@/components/RoundedIcon.vue';

export default {
    components: {
        RoundedIcon
    },
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
            wpmDataFinished: {},
            connected: false,
            isReady: false,
            restartVoteInfo: { votes: 0, total: 0 },
            hasVotedRestart: false,
            isGameStarted: false,
            elapsedTime: 0,
            startTime: null,
            timerInterval: null,
            countdown: null,
            countdownValue: 3,
            isCountingDown: false,
            timer: null,
            finishSound: new Audio("/complete.wav"),
            startSound: new Audio("/countdown.mp3"),
        };
    },
    mounted() {
        this.finishSound.volume = 0.45;
        this.startSound.volume = 0.15;
        this.$refs.typingBox?.focus();
    },
    computed: {
        currentUser() {
            return sessionStorage.getItem("username");
        },
        filteredWpmData() {
            const currentUser = currentUser();
            if (!currentUser) return {};
            return {
                [currentUser]: this.wpmData[currentUser]
            };
        },
        formattedElapsedTime() {
            const totalSeconds = Math.floor(this.elapsedTime);
            const minutes = Math.floor(totalSeconds / 60);
            const seconds = totalSeconds % 60;
            return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
        }
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
                this.ws.send(JSON.stringify(message));
                this.ws.close();
            }
            sessionStorage.removeItem("username");
            sessionStorage.removeItem("language");
            sessionStorage.removeItem("roomID");
            this.$router.push('/');
        },
        connectWebSocket() {
            if (this.ws) this.ws.close();

            this.ws = new WebSocket(import.meta.env.VITE_WS_URL + "/ws/typing");

            this.ws.onopen = () => {
                this.connected = true;
                this.isReady = false;
                this.isGameStarted = false;

                this.ws.send(JSON.stringify({
                    username: this.username,
                    roomID: this.roomID,
                    language: this.language,
                    ...(localStorage.getItem("max_players") && { limit: localStorage.getItem("max_players") })
                }));

                localStorage.removeItem("max_players");
            };

            this.ws.onmessage = (event) => {
                const data = JSON.parse(event.data);
                if (data.error) {
                    alert(data.error);
                    sessionStorage.removeItem("username");
                    sessionStorage.removeItem("roomID");
                    this.ws.close();
                    this.$router.push("/");
                    return;
                }
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
                    this.wpmDataFinished = JSON.parse(JSON.stringify(this.wpmData));
                    this.finishSound.play();
                }
                if (data.type === "update_votes") {
                    this.restartVoteInfo.votes = data.votes;
                    this.restartVoteInfo.total = data.total;
                }
                if (data.type === "restart_game") {
                    this.inputText = "";
                    this.givenText = data.text;
                    this.isGameStarted = false;
                    this.hasVotedRestart = false;
                    this.wpmData = {}
                    this.wpmDataFinished = {}
                    this.restartVoteInfo = { votes: 0, total: 0 };
                    this.startCountdown();
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
        handleKeydown(event) {
            if (event.key === 'Backspace') {
                this.inputText = this.inputText.slice(0, -1);
                this.sendText();
                return;
            }
            if (!event.key || event.key.length !== 1) return;

            this.inputText += event.key;
            this.sendText();
        },
        getCharClass(index) {
            if (!this.inputText[index]) return "default";
            return this.inputText[index] === this.givenText[index] ? "correct" : "incorrect";
        },
        sendReadyFlag() {
            if (this.ws && this.connected) {
                this.isReady = !this.isReady;
                this.ws.send(JSON.stringify({ status: this.isReady ? "ready" : "not_ready" }));
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

                    this.startTime = performance.now();
                    this.timerInterval = setInterval(() => {
                        const now = performance.now();
                        this.elapsedTime = ((now - this.startTime) / 1000).toFixed(2);
                    }, 100);

                    this.$nextTick(() => {
                        this.$refs.typingBox?.focus();
                    });
                }
            }, 1000);
        },

        voteRestart() {
            if (this.ws && this.connected && !this.hasVotedRestart) {
                this.hasVotedRestart = true;
                this.ws.send(JSON.stringify({ type: "vote_restart" }));
            }
        },

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

.ready {
    background-color: #4CAF50;
    color: var(--bg-color);
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

.btn-back:hover {
    transform: scale(1.02);
    transition: background 0.3s ease, transform 0.2s ease;
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
    transform: scale(1.02);
    transition: background 0.3s ease, transform 0.2s ease;
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

.typing-container h3 {
    margin-bottom: 20px;
}

.typing-text {
    font-size: 1.4em;
    position: relative;
    display: inline-block;
    background: var(--bg-color);
    padding: 10px;
    border-radius: 5px;
    margin: 10px 0;
}

.typing-text span {
  display: inline;
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

.unready {
    background-color: #8b8b8b;
    color: white;
}

.unready:hover {
    background-color: #2b2b2b;
    color: white;
}

.vote-info {
    background: var(--bg-color);
    color: var(--text-color);
    padding: 1px 8px 1px 8px;
    margin-left: 10px;
    border-radius: 5px;
}

@media (max-width: 600px) {
    .typing-container {
        max-width: 100%;
    }
}

.row {
    display: flex;
    justify-content: flex-end;
    text-align: left;
    gap: 15px;
}

.current-wpm,
.timer {
    font-size: 12px;
    display: flex;
    letter-spacing: 1px;
}

.timer span,
.current-wpm span {
    min-width: 80px;
}

.wpm {
    color: #FF8343;
}

.typing-text .has-cursor {
    position: relative;
}

.typing-text .has-cursor::after {
    content: "";
    position: absolute;
    right: 0;
    top: 0;
    height: 100%;
    width: 1px;
    background: var(--text-color);
    animation: blink 1s step-start infinite;
}

@keyframes blink {
    50% {
        opacity: 0;
    }
}
</style>