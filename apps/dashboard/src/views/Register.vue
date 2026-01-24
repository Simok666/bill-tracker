<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '../composables/useAuth';

const { registerAsync, isRegistering } = useAuth();
const name = ref('');
const companyName = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const error = ref('');

const handleSubmit = async () => {
    error.value = '';
    
    if (password.value !== confirmPassword.value) {
        error.value = 'Passwords do not match';
        return;
    }

    try {
        await registerAsync({
            name: name.value,
            company_name: companyName.value,
            email: email.value,
            password: password.value
        });
    } catch (err: any) {
        if (err.response?.data?.error) {
            error.value = err.response.data.error;
        } else if (err.response?.data?.message) {
            error.value = err.response.data.message;
        } else {
            error.value = 'Failed to create account. Please try again.';
        }
    }
};
</script>

<template>
  <div class="min-h-screen flex flex-col font-inter text-white bg-register-bg-dark relative overflow-hidden">
    <!-- Background Gradient Effect -->
    <div class="absolute top-0 left-0 w-full h-[600px] bg-gradient-to-b from-[#1e255e] to-transparent opacity-60 pointer-events-none"></div>
    <div class="absolute -top-40 -right-40 w-[800px] h-[800px] bg-blue-900/20 rounded-full blur-3xl pointer-events-none"></div>

    <!-- Header -->
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

    <main class="flex-1 flex flex-col items-center justify-center p-6 relative z-10">
      
      <!-- Headline -->
      <div class="text-center mb-10 max-w-2xl mx-auto">
        <h1 class="text-4xl sm:text-5xl font-bold tracking-tight mb-4">Create your enterprise <br/> account</h1>
        <p class="text-slate-400 text-lg">Financial transparency and late fee prevention at scale.</p>
      </div>

      <!-- Register Card -->
      <div class="w-full max-w-[520px] bg-register-card border border-white/5 rounded-2xl p-8 sm:p-10 shadow-2xl">
        <form @submit.prevent="handleSubmit" class="space-y-5">
            
            <!-- Error Alert -->
            <div v-if="error" class="p-4 rounded-lg bg-red-500/10 border border-red-500/20 text-red-400 text-sm flex items-center gap-2">
                <span class="material-symbols-outlined text-lg">error</span>
                {{ error }}
            </div>

          <!-- Full Name -->
          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-slate-200 flex items-center gap-2">
              <span class="material-symbols-outlined text-xs text-slate-400">person</span>
              Full Name
            </label>
            <input v-model="name" type="text" placeholder="John Doe" class="w-full bg-register-input border border-white/5 rounded-lg px-4 py-3.5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" required/>
          </div>

          <!-- Company Name -->
          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-slate-200 flex items-center gap-2">
              <span class="material-symbols-outlined text-xs text-slate-400">business</span>
              Company Name
            </label>
            <input v-model="companyName" type="text" placeholder="Acme Corp" class="w-full bg-register-input border border-white/5 rounded-lg px-4 py-3.5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" required/>
          </div>

          <!-- Work Email -->
          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-slate-200 flex items-center gap-2">
              <span class="material-symbols-outlined text-xs text-slate-400">mail</span>
              Work Email
            </label>
            <input v-model="email" type="email" placeholder="john.doe@company.com" class="w-full bg-register-input border border-white/5 rounded-lg px-4 py-3.5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" required/>
          </div>

          <!-- Password Row -->
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label class="text-sm font-semibold text-slate-200 flex items-center gap-2">
                <span class="material-symbols-outlined text-xs text-slate-400">lock</span>
                Password
              </label>
              <input v-model="password" type="password" placeholder="••••••••" class="w-full bg-register-input border border-white/5 rounded-lg px-4 py-3.5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" required minlength="8"/>
            </div>
            <div class="space-y-1.5">
              <label class="text-sm font-semibold text-slate-200 flex items-center gap-2">
                <span class="material-symbols-outlined text-xs text-slate-400">shield</span>
                Confirm
              </label>
              <input v-model="confirmPassword" type="password" placeholder="••••••••" class="w-full bg-register-input border border-white/5 rounded-lg px-4 py-3.5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-register-primary focus:border-transparent transition-all" required minlength="8"/>
            </div>
          </div>

          <!-- Helper Text -->
          <div class="flex items-start gap-2 text-xs text-slate-400 px-1">
             <span class="material-symbols-outlined text-sm pt-0.5">info</span>
             <p>Use 8 or more characters with a mix of letters, numbers & symbols.</p>
          </div>

          <!-- Submit Button -->
          <button :disabled="isRegistering" type="submit" class="w-full bg-register-primary hover:bg-[#4049e0] text-white font-bold py-4 rounded-xl shadow-lg shadow-indigo-500/20 transition-all flex items-center justify-center gap-2 group mt-2 disabled:opacity-50 disabled:cursor-not-allowed">
            <span v-if="isRegistering" class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
            <span v-else class="flex items-center gap-2">
                Create Account
                <span class="material-symbols-outlined transition-transform group-hover:translate-x-1">arrow_forward</span>
            </span>
          </button>

          <!-- Security Badges -->
          <div class="flex items-center justify-center gap-6 pt-2 text-[10px] font-bold text-slate-500 tracking-wider uppercase">
            <div class="flex items-center gap-1.5">
              <span class="material-symbols-outlined text-sm">verified_user</span>
              Bank-Grade Security
            </div>
            <div class="flex items-center gap-1.5">
              <span class="material-symbols-outlined text-sm">security</span>
              GDPR Compliant
            </div>
          </div>

        </form>
      </div>

      <!-- Footer Action -->
      <div class="mt-8 text-center">
        <p class="text-slate-400 text-sm">
          Already have an account? <router-link to="/login" class="text-register-primary font-bold hover:text-white transition-colors">Log In</router-link>
        </p>
      </div>

    </main>

  </div>
</template>
