<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header class="bg-red-100">
    <div class="ml-3.5">
        <p class="text-xl font-bold text-white">Email-Search</p>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>

.logo {
  display: block;
  margin: 0 auto 2rem;
}


@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
  }
}
</style>
