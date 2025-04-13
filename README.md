<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExYnBpNWtwNGpxOXZwODVqZnc3cDE2Nm1ycHBiaThkeHRkYXZqZDA5eCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/zOx4kKZLsfuqShoh2t/giphy.gif" alt="Logo" height="100">
  </a>

<h3 align="center">Typing Race Multiplayer</h3>

  <p align="center">
    A real-time multiplayer typing race web application where users can compete in Thai or English. The platform emphasizes speed and real-time interaction, making it perfect for fun challenges or improving typing skills. <br> <br> <br>⚠️ Note: The app may be slow after being idle for a while, as it’s hosted on a free cloud service.
    <br />
    <br />
    <a href="https://github.com/darkfat123/typing-race-web-multiplayer/issues">🚨 Report Bug</a>
    ·
    <a href="https://github.com/darkfat123/typing-race-web-multiplayer/issues">✉️ Request Feature</a>
    .
    <a href="https://github.com/darkfat123/typing-race-web-multiplayer?tab=readme-ov-file#-getting-started-for-development-only">🚀 Getting Started</a>
  </p>
</div>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">✨ Features:</h3>

  * All components and users are connected through real-time communication via WebSocket.
  * Supports both Thai and English languages.
  * No database is required — in-memory storage is used instead.
  * Players can create rooms (up to 10 users per room) and choose the language for that room.
  * Players can filter available rooms by language or room ID when joining.
  * Word Per Minute (WPM) is calculated live while typing.
  * Light and Dark modes are available and can be switched anytime.
  * After finishing a race, players can vote to restart the game. If more than 60% of the players vote to restart, the game will begin again with new randomized text.
  * A countdown is displayed before the game starts.
  * If all players leave a room, the room is automatically deleted.

</br>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">🖥️ Programming languages and tools:</h3>

- Backend
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" />
  </a>
</p>

- Frontend
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=vue,js,npm" />
  </a>
</p>

- Tools
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=git,github,vscode,docker" />
  </a>
</p>

<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">🖲️ Demo:</h3>
<p align="center">
  <img src="https://github.com/user-attachments/assets/dfefdc27-a388-4e44-ba88-c4f4bfdaf79d" />
</p>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">
<h3 align="left"> 📃 All pages:</h3>

- Home
<p align="center">
  <img src="https://github.com/user-attachments/assets/b8e7ed18-7dc4-4d8a-94f8-cc61fc6921b3" />
</p>
<br>

- Create room
<p align="center">
  <img src="https://github.com/user-attachments/assets/7382f643-0be5-4397-92b1-3bc75060d1ca" />
</p>
<br>

- Active Room List
<p align="center">
  <img src="https://github.com/user-attachments/assets/5b01d09b-b547-4b49-9bb1-a31b3a8e0aab" />
</p>
<br>

- Typing Test Room
<p align="center">
  <img src="https://github.com/user-attachments/assets/5cea22fd-651e-497b-bf79-1dd9be23b900" />
</p>
<br>

- Display Scoreboard After All Players Have Finished
<p align="center">
  <img src="https://github.com/user-attachments/assets/3879cf79-c412-421f-b5d1-e76fc94102e3" />
</p>
<br>

- Vote to Restart Room and Randomize the Text
<p align="center">
  <img src="https://github.com/user-attachments/assets/ea075f64-848b-4604-8b87-d5597ee9b4db" />
</p>

</br>


<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

### 🚀 Getting Started (for development only)

#### 1. Clone the project
```bash
git clone https://github.com/darkfat123/typing-race-web-multiplayer.git
cd typing-race-web-multiplayer
```
#### 2. Frontend
```bash
cd web-vue
npm install
npm run dev
```

#### 3. Backend
```bash
cd server
go mod tidy
go run cmd/main.go
```

#### 4. Environment Variables
```bash
# Backend - In `server/.env`
ALLOWED_ORIGIN=http://localhost:5173

# Frontend - In `web-vue/.env`
VITE_WS_URL=ws://localhost:8080
```



<h3> Connect with me 🎊: <h3>
  <a href="https://www.linkedin.com/in/supakorn-yookack-39a730289/">
   <img align="left" alt="Supakorn Yookack | Linkedin" width="30px" src="https://www.vectorlogo.zone/logos/linkedin/linkedin-icon.svg" />
  </a>
  <a href="mailto:supakorn.yookack@gmail.com">
    <img align="left" alt="Supakorn Yookack | Gmail" width="32px" src="https://www.vectorlogo.zone/logos/gmail/gmail-icon.svg" />
  </a>
  <a href="https://medium.com/@yookack_s">
    <img align="left" alt="Supakorn Yookack | Medium" width="32px" src="https://www.vectorlogo.zone/logos/medium/medium-tile.svg" />
  </a>
   <a href="https://www.facebook.com/supakorn.yookaek/">
    <img align="left" alt="Supakorn Yookack | Facebook" width="32px" src="https://www.vectorlogo.zone/logos/facebook/facebook-tile.svg" />
  </a>
   <a href="https://github.com/darkfat123">
    <img align="left" alt="Supakorn Yookack | Github" width="32px" src="https://www.vectorlogo.zone/logos/github/github-tile.svg" />
  </a>
    <p align="right" > Created by <a href="https://github.com/darkfat123">darkfat</a></p> <p align="right" > <img src="https://komarev.com/ghpvc/?username=darkfat123&label=Profile%20views&color=0e75b6&style=flat" alt="darkfat123" /> </p>
