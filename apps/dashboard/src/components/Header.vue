<script setup lang="ts">
</script>

<template>
<header class="sticky top-0 z-20 flex items-center justify-between px-8 py-6 bg-background-dark/90 backdrop-blur-md border-b border-border-dark">
<div class="flex flex-col">
<h2 class="text-2xl font-bold text-white tracking-tight">Welcome, Admin</h2>
<p class="text-slate-400 text-sm">Here's your financial overview for today.</p>
</div>
<div class="flex items-center gap-4">
<div class="hidden md:flex items-center gap-2 px-4 py-2 rounded-full bg-slate-800 border border-border-dark">
<span class="material-symbols-outlined text-slate-400 text-sm">calendar_month</span>
<span class="text-white text-sm font-bold">Jan 2026</span>
</div>
<button class="relative p-2.5 rounded-full bg-slate-800 border border-border-dark text-slate-300 hover:text-white hover:border-slate-600 transition-colors">
<span class="material-symbols-outlined text-[20px]">notifications</span>
<span class="absolute top-2 right-2.5 size-2 rounded-full bg-red-500 border-2 border-slate-800"></span>
</button>
<button class="md:hidden p-2 rounded-full bg-slate-800 text-white">
<span class="material-symbols-outlined">menu</span>
</button>
</div>
</header>
</template>
