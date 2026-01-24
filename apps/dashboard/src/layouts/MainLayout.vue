<script setup lang="ts">
import Sidebar from '../components/Sidebar.vue';
import Header from '../components/Header.vue';
</script>

<template>
  <div class="flex h-screen w-full overflow-hidden bg-background-light dark:bg-background-dark text-slate-900 dark:text-slate-50 font-display transition-colors duration-300">
    <Sidebar />
    <main class="flex-1 flex flex-col h-full overflow-y-auto bg-background-dark relative">
      <Header />
      <div class="flex-1 p-6 md:p-8 max-w-7xl mx-auto w-full flex flex-col gap-8">
        <router-view />
      </div>
    </main>
  </div>
</template>
