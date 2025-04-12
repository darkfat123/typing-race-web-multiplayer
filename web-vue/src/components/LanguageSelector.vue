<template>
    <div class="language-selector">
        <button :class="['lang-btn', { selected: selectedLang === 'th' }]" @click="selectLanguage('th')">
            Thai
        </button>
        <button :class="['lang-btn', { selected: selectedLang === 'en' }]" @click="selectLanguage('en')">
            English
        </button>
    </div>
</template>

<script>
export default {
    name: 'LanguageSelector',
    props: {
        selectedLang: {
            type: String,
            required: true,
        },
    },
    emits: ['update:lang'],
    methods: {
        selectLanguage(lang) {
            this.$emit('update:lang', lang);
        },
    },
};
</script>

<style scoped>
.language-selector {
    gap: 10px;
    display: flex;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 10px;
    margin-top: 20px;
}

.lang-btn {
    flex: 1;
    padding: 10px 0;
    border-radius: 5px;
    font-size: 16px;
    border: 1px solid #ccc;
    color: var(--text-color);
    background-color: transparent;
    cursor: pointer;
    transition: background 0.3s ease, transform 0.2s ease;
}

.lang-btn.selected {
    background-color: var(--main-btn-color);
    color: var(--text-color);
}

.lang-btn:hover {
    color: black;
    background-color: #f1f1f1;
    transform: scale(1.01);
}
</style>