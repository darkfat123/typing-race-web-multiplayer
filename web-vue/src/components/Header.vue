<template>
    <header>
        <h1>{{ displayedText }}</h1>
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

<script setup>
import { ref, onMounted } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const fullText = 'Typing Race Multiplayer'
const displayedText = ref('')
let isDeleting = false
let index = 0

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms))
}

async function typeLoop() {
    while (true) {
        if (!isDeleting) {
            if (index < fullText.length) {
                displayedText.value += fullText[index]
                index++
                await sleep(100)
            } else {
                await sleep(3000)
                isDeleting = true
            }
        } else {
            if (index > 0) {
                displayedText.value = fullText.slice(0, index - 1)
                index--
                await sleep(50)
            } else {
                await sleep(1000)
                isDeleting = false
            }
        }
    }
}

onMounted(() => {
    typeLoop()
})

const emit = defineEmits(['toggleDarkMode'])
function toggleDarkMode() {
    emit('toggleDarkMode')
}
</script>


<style>
header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    height: 4rem;
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

h1 {
    font-size: 28px;
    white-space: nowrap;
    overflow: hidden;
    border-right: 2px solid #ccc;
    animation: blinkCursor 0.7s step-end infinite;
}

@keyframes blinkCursor {

    from,
    to {
        border-color: transparent
    }

    50% {
        border-color: #ccc
    }
}
</style>