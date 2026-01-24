<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '../composables/useAuth';

const { loginAsync, isLoggingIn } = useAuth();
const email = ref('');
const password = ref('');
const showPassword = ref(false);
const error = ref('');

const handleSubmit = async () => {
  error.value = '';
  try {
    await loginAsync({
      email: email.value,
      password: password.value
    });
  } catch (err: any) {
    if (err.response?.data?.message) {
      error.value = err.response.data.message;
    } else {
      error.value = 'Failed to login. Please check your credentials.';
    }
  }
};
</script>

<template>
  <div class="min-h-screen flex flex-col font-inter text-white bg-register-bg-dark relative overflow-hidden">
    <!-- Background Gradient Effect -->
    <div class="absolute top-0 left-0 w-full h-[600px] bg-gradient-to-b from-[#1e255e] to-transparent opacity-60 pointer-events-none"></div>
    <div class="absolute -top-40 -right-40 w-[800px] h-[800px] bg-blue-900/20 rounded-full blur-3xl pointer-events-none"></div>

    <!-- Top Navigation Bar Component -->
    <header class="w-full relative z-10 px-6 sm:px-12 py-6 flex items-center justify-between">
      <div class="flex items-center gap-2.5">
        <div class="size-9 bg-login-primary rounded-lg flex items-center justify-center text-white shadow-lg shadow-blue-500/20">
          <span class="material-symbols-outlined text-[22px]">account_balance_wallet</span>
        </div>
        <h2 class="text-xl font-bold tracking-tight">BillTrack <span class="text-slate-400 font-medium">Enterprise</span></h2>
      </div>
      <div class="hidden md:flex items-center gap-8 text-sm font-medium text-slate-300">
        <span class="hidden sm:block">Financial Transparency Suite</span>
      </div>
    </header>

    <main class="flex-1 flex items-center justify-center p-8 sm:p-12 relative z-10">
      <!-- Authentication Card -->
      <div class="w-full max-w-[480px] bg-register-card border border-white/5 rounded-2xl shadow-2xl p-8 sm:p-10">
        
        <!-- Branding & Titles -->
        <div class="flex flex-col items-center text-center mb-10">
          <div class="mb-6 p-3 bg-register-primary/10 rounded-full">
            <span class="material-symbols-outlined text-register-primary text-4xl">security</span>
          </div>
          <h1 class="text-white text-3xl font-bold tracking-tight mb-2">Welcome Back</h1>
          <p class="text-slate-400 text-base">Log in to manage your enterprise bills</p>
        </div>

        <!-- Login Form -->
        <form @submit.prevent="handleSubmit" class="space-y-5">
            <!-- Error Alert -->
            <div v-if="error" class="p-4 rounded-lg bg-red-500/10 border border-red-500/20 text-red-400 text-sm flex items-center gap-2">
                <span class="material-symbols-outlined text-lg">error</span>
                {{ error }}
            </div>

          <!-- Email Field -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-slate-200" for="email">
              Email Address
            </label>
            <div class="relative">
              <span class="material-symbols-outlined absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 text-xl">mail</span>
              <input v-model="email" class="w-full pl-12 pr-4 py-3.5 bg-register-input border border-white/5 rounded-lg text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" id="email" placeholder="name@company.com" type="email" required/>
            </div>
          </div>
          <!-- Password Field -->
          <div class="space-y-2">
            <label class="block text-sm font-medium text-slate-200" for="password">
              Password
            </label>
            <div class="relative">
              <span class="material-symbols-outlined absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 text-xl">lock</span>
              <input v-model="password" class="w-full pl-12 pr-12 py-3.5 bg-register-input border border-white/5 rounded-lg text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" id="password" placeholder="••••••••" :type="showPassword ? 'text' : 'password'" required/>
              <button @click="showPassword = !showPassword" class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-200" type="button">
                <span class="material-symbols-outlined text-xl">{{ showPassword ? 'visibility_off' : 'visibility' }}</span>
              </button>
            </div>
          </div>
          
          <!-- Options Row -->
          <div class="flex items-center justify-between py-2">
            <label class="flex items-center gap-2 cursor-pointer group">
              <input class="rounded border-white/10 text-register-primary focus:ring-register-primary bg-register-input" type="checkbox"/>
              <span class="text-sm text-slate-400 group-hover:text-white transition-colors">Remember me</span>
            </label>
            <a class="text-sm font-semibold text-login-accent hover:underline decoration-2 underline-offset-4" href="#">Forgot Password?</a>
          </div>

          <!-- Submit Button -->
          <button :disabled="isLoggingIn" class="w-full py-4 px-6 bg-register-primary hover:bg-[#4049e0] text-white font-bold rounded-xl shadow-lg shadow-indigo-500/20 transition-all flex items-center justify-center gap-2 group mt-2 disabled:opacity-50 disabled:cursor-not-allowed" type="submit">
             <span v-if="isLoggingIn" class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
             <span v-else>Log In</span>
          </button>
        </form>

        <!-- Divider -->
        <div class="relative my-8">
          <div class="absolute inset-0 flex items-center">
            <span class="w-full border-t border-white/10"></span>
          </div>
          <div class="relative flex justify-center text-sm uppercase">
            <span class="bg-register-card px-4 text-slate-500 font-medium">Or continue with</span>
          </div>
        </div>

        <!-- Social/SSO Logins -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <button class="flex items-center justify-center gap-3 px-4 py-3 border border-white/10 rounded-lg bg-register-input hover:bg-white/5 transition-colors text-slate-200 font-medium">
            <svg class="size-5" viewbox="0 0 24 24">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"></path>
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"></path>
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l3.66-2.84z" fill="#FBBC05"></path>
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 12-4.53z" fill="#EA4335"></path>
            </svg>
            Google
          </button>
          <button class="flex items-center justify-center gap-3 px-4 py-3 border border-white/10 rounded-lg bg-register-input hover:bg-white/5 transition-colors text-slate-200 font-medium">
            <span class="material-symbols-outlined text-xl">hub</span>
            SSO
          </button>
        </div>

        <!-- Footer Link -->
        <div class="mt-8 text-center text-sm">
           <p class="text-slate-400">
             Don't have an account? <router-link to="/register" class="text-register-primary font-bold hover:text-white transition-colors">Register</router-link>
           </p>
        </div>
      </div>

    </main>
  </div>
</template>
