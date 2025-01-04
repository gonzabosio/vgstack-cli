<script setup>
</script>

<template>
  <div id="leading-container">
    <p>Go+Vue App</p>
    <p>Based on <span><a href="https://vite.dev" target="_blank">Vite</a></span>⚡ scaffolding</p>
  </div>
</template>

<style scoped>
#leading-container {
  border: 16px dashed;
  border-radius: 1em;
  padding: 2em;
  transition: 0.5s;
  border-left-color: rgb(98, 159, 163);
  border-top-color: rgb(98, 159, 163);
  border-right-color: rgb(34, 114, 45);
  border-bottom-color: rgb(34, 114, 45);

  &:hover {
    transform: scale(1.1);
    border-right-color: rgb(98, 159, 163);
    border-bottom-color: rgb(98, 159, 163);
    border-left-color: rgb(34, 114, 45);
    border-top-color: rgb(34, 114, 45);
  }
}

a {
  color: rgb(168, 111, 221);
  text-decoration: none;
  transition: 0.3s;

  &:hover {
    text-shadow: rgb(198, 166, 228) 1px 0 15px;
    ;
  }
}
</style>
