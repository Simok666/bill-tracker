<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '../composables/useAuth';
import { useRouter } from 'vue-router';

const isProfileOpen = ref(false);
const { currentUser: user, logout } = useAuth();
const router = useRouter();

const handleLogout = () => {
    logout();
    router.push('/login');
};
</script>

<template>
<aside class="hidden lg:flex w-72 flex-col justify-between border-r border-border-dark bg-[#0b1120] p-6 h-full">
<div class="flex flex-col gap-8">
<!-- Brand -->
<div class="flex items-center gap-3 px-2">
<div class="flex items-center justify-center size-10 rounded-xl bg-gradient-to-br from-primary to-green-700 shadow-lg shadow-green-900/50 text-white">
<span class="material-symbols-outlined">account_balance_wallet</span>
</div>
<div class="flex flex-col">
<h1 class="text-white text-lg font-bold tracking-tight">BillTracker</h1>
<p class="text-slate-400 text-xs font-medium uppercase tracking-wider">Enterprise</p>
</div>
</div>
<!-- Navigation -->
    <nav class="flex flex-col gap-2">
        <router-link to="/" class="group" custom v-slot="{ href, navigate, isExactActive }">
            <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isExactActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
                <span class="material-symbols-outlined" style="font-variation-settings: 'FILL' 1;">dashboard</span>
                <span class="text-sm font-bold">Dashboard</span>
            </a>
        </router-link>
        <router-link to="/bills" class="group" custom v-slot="{ href, navigate, isActive }">
            <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
                <span class="material-symbols-outlined">receipt_long</span>
                <span class="text-sm font-medium">Bills</span>
            </a>
        </router-link>
        <router-link to="/recurring" class="group" custom v-slot="{ href, navigate, isActive }">
            <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
                <span class="material-symbols-outlined">update</span>
                <span class="text-sm font-medium">Recurring</span>
            </a>
        </router-link>
        <router-link to="/categories" class="group" custom v-slot="{ href, navigate, isActive }">
            <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
                <span class="material-symbols-outlined">folder_open</span>
                <span class="text-sm font-medium">Categories</span>
            </a>
        </router-link>
        <router-link to="/vendors" class="group" custom v-slot="{ href, navigate, isActive }">
            <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
                <span class="material-symbols-outlined">domain</span>
                <span class="text-sm font-medium">Vendors</span>
            </a>
        </router-link>
        <router-link to="/reports" class="group" custom v-slot="{ href, navigate, isActive }">
            <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
                <span class="material-symbols-outlined">bar_chart</span>
                <span class="text-sm font-medium">Reports</span>
            </a>
        </router-link>
</nav>
</div>
<div class="flex flex-col gap-2">
<router-link to="/settings" class="group" custom v-slot="{ href, navigate, isExactActive }">
    <a :href="href" @click="navigate" class="flex items-center gap-3 px-4 py-3 rounded-full transition-all duration-200" :class="isExactActive ? 'bg-primary/20 text-primary' : 'text-slate-400 hover:text-white hover:bg-slate-800'">
        <span class="material-symbols-outlined">settings</span>
        <span class="text-sm font-medium">Settings</span>
    </a>
</router-link>

    <div class="mt-4 relative" v-if="user">
        <button @click="isProfileOpen = !isProfileOpen" class="w-full flex items-center gap-3 px-4 py-3 rounded-xl bg-slate-800/50 border border-slate-800 hover:bg-slate-800 transition-colors text-left group">
            <div class="size-8 rounded-full bg-slate-700 bg-cover bg-center shrink-0 flex items-center justify-center text-white font-bold text-xs" :style="user.avatar_url ? `background-image: url('${user.avatar_url}')` : ''">
                {{ user.avatar_url ? '' : user.name.charAt(0).toUpperCase() }}
            </div>
            <div class="flex flex-col flex-1 overflow-hidden">
                <p class="text-white text-xs font-bold group-hover:text-primary transition-colors truncate">{{ user.name }}</p>
                <p class="text-slate-400 text-[10px] truncate">{{ user.email }}</p>
            </div>
            <span class="material-symbols-outlined text-slate-500 text-lg transition-transform duration-200 shrink-0" :class="{ 'rotate-180': isProfileOpen }">expand_less</span>
        </button>

        <!-- Dropdown Menu -->
        <div v-if="isProfileOpen" class="absolute bottom-full left-0 w-full mb-2 bg-[#161f32] border border-white/10 rounded-xl shadow-xl overflow-hidden animate-in fade-in slide-in-from-bottom-2 duration-200">
            <div class="p-1">
                <router-link to="/settings" class="flex items-center gap-2 px-3 py-2 text-sm text-slate-300 hover:bg-white/5 hover:text-white rounded-lg transition-colors">
                    <span class="material-symbols-outlined text-lg">edit</span>
                    Edit Profile
                </router-link>
                <div class="h-px bg-white/5 my-1"></div>
                <button @click="handleLogout" class="w-full flex items-center gap-2 px-3 py-2 text-sm text-red-400 hover:bg-red-500/10 hover:text-red-300 rounded-lg transition-colors text-left">
                    <span class="material-symbols-outlined text-lg">logout</span>
                    Logout
                </button>
            </div>
        </div>
    </div>
    
    <div v-else class="mt-4">
        <router-link to="/login" class="flex items-center justify-center gap-2 w-full py-3 bg-primary/10 hover:bg-primary/20 text-primary font-bold rounded-xl transition-colors text-sm border border-primary/20">
            Login / Register
        </router-link>
    </div>
</div>
</aside>
</template>
