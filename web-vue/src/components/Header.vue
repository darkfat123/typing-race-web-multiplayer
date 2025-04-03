<template>
    <header>
        <h1>Typing Race Multiplayer</h1>
        <button @click="$emit('toggleDarkMode')">
            Toggle Dark Mode
        </button>
    </header>
</template>

<script>
export default {
    name: 'Header',
};
</script>

<style>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 10px 20px;
  background: var(--bg-color);
  color: var(--text-color);
  box-shadow: 0 0px 8px var(--shadow-color);
  margin-bottom: 20px;
}

button {
    padding: 8px 12px;
    border: none;
    cursor: pointer;
    background: var(--text-color);
    color: var(--bg-color);
    border-radius: 5px;
    transition: 0.3s;
}

button:hover {
    opacity: 0.8;
}
</style>