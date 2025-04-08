<template>
    <header>
        <h1>Typing Race Multiplayer</h1>
        <div>
            <input type="checkbox" class="checkbox" id="checkbox" @change="toggleDarkMode" />
            <label for="checkbox" class="checkbox-label">
                <FontAwesomeIcon :icon="['fas', 'moon']" />
                <FontAwesomeIcon :icon="['fas', 'sun']" />
                <span class="ball"></span>
            </label>
        </div>
    </header>
</template>

<script>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
export default {
    name: 'Header',
    components: {
        FontAwesomeIcon
    },
    methods: {
        toggleDarkMode() {
            this.$emit('toggleDarkMode');
        }
    },
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
}
.checkbox {
    opacity: 0;
    position: absolute;
}

.checkbox-label {
    background-color: rgb(56, 56, 56);
    width: 80px;
    height: 36px;
    border-radius: 50px;
    position: relative;
    padding: 10px;
    cursor: pointer;
    display: flex;
    justify-content: space-between;
    align-items: center;
    box-shadow: 0 0px 8px var(--shadow-color);
}

.fa-moon {
    color: #f1c40f;
}

.fa-sun {
    color: #f39c12;
}

.checkbox-label .ball {
    background-color: #fff;
    width: 28px;
    height: 28px;
    position: absolute;
    left: 4px;
    top: 4px;
    border-radius: 50%;
    transition: transform 0.2s linear;
}

.checkbox:checked+.checkbox-label .ball {
    transform: translateX(42px);
}
</style>