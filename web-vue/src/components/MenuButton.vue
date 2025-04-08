<template>
    <button @click="navigate" class="menu-btn">{{ label }}</button>
</template>

<script>
export default {
    name: 'MenuButton',
    props: {
        label: {
            type: String,
            required: true
        },
        to: {
            type: String,
            required: true
        }
    },
    methods: {
        navigate() {
            this.$router.push(this.to);
        }
    }
};
</script>
<style scoped>
.menu-btn {
    width: 100%;
    padding: 10px;
    font-weight: bold;
    background-color: goldenrod;
    color: var(--bg-color);
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 16px;
    box-shadow: 0 0px 8px var(--shadow-color);
    transition: background-color 0.3s, color 0.3s;
    text-align: center;
    transition: background 0.3s ease, transform 0.2s ease;
}

.menu-btn:hover {
    background: linear-gradient(to right,  var(--bg-color), goldenrod);
    transform: scale(1.03);
}
</style>