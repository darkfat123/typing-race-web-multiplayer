<template>
    <div>
      <div class="dot" :style="{ backgroundColor: color }"></div>
      <span>{{ label }}</span>
    </div>
  </template>
  
  <script>
  export default {
    name: 'RoundedIcon',
    props: {
      label: {
        type: String,
        required: true
      },
      color: {
        type: String,
        required: true
      }
    },
  };
  </script>
  
  <style scoped>
  .dot {
    display: inline-block;
    width: 10px;
    height: 10px;
    border-radius: 50%;
    margin-right: 8px;
  }
  span {
    font-size: 12px;
    margin-right: 8px;
    letter-spacing: 1px;
  }
  </style>
  