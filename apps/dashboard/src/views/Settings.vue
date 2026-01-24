<script setup lang="ts">
import { ref, watchEffect } from 'vue';
import { useUserProfile, useUpdateProfile, useChangePassword } from '../composables/useUsers';

const { data: user, isLoading } = useUserProfile();
const { mutateAsync: updateProfile, isPending: isUpdating } = useUpdateProfile();
const { mutateAsync: changePassword, isPending: isChangingPassword } = useChangePassword();

const profileForm = ref({
    name: ''
});

const passwordForm = ref({
    current_password: '',
    new_password: '',
    confirm_password: ''
});

const passwordMessage = ref('');
const passwordError = ref('');

watchEffect(() => {
    if (user.value) {
        profileForm.value.name = user.value.name;
    }
});

const handleUpdateProfile = async () => {
    try {
        await updateProfile({ name: profileForm.value.name });
        // Success feedback handled by UI state or toast
    } catch (e) {
        console.error("Failed to update profile", e);
    }
};

const handleChangePassword = async () => {
    passwordMessage.value = '';
    passwordError.value = '';
    
    if (passwordForm.value.new_password !== passwordForm.value.confirm_password) {
        passwordError.value = "New passwords don't match";
        return;
    }

    try {
        await changePassword({ 
            current_password: passwordForm.value.current_password,
            new_password: passwordForm.value.new_password
        });
        passwordMessage.value = 'Password changed successfully';
        passwordForm.value = { current_password: '', new_password: '', confirm_password: '' };
    } catch (e) {
        passwordError.value = 'Failed to change password. Check old password.';
    }
};
</script>

<template>
<div class="flex-1 flex flex-col h-full relative bg-background-forest">
    <!-- Subtle Gradient -->
    <div class="absolute top-0 left-0 w-full h-96 bg-primary/5 blur-3xl pointer-events-none rounded-full -translate-y-1/2"></div>
    
    <!-- Header -->
    <header class="px-8 py-6 z-10 shrink-0">
        <h2 class="text-4xl font-extrabold text-white tracking-tight">Settings</h2>
        <p class="text-slate-400 text-sm mt-1">Manage your account preferences</p>
    </header>

    <!-- Content -->
    <div class="flex-1 px-8 pb-8 overflow-y-auto no-scrollbar z-10">
        <div class="max-w-4xl space-y-8">
            
            <!-- Profile Section -->
            <div class="bg-surface-forest border border-white/5 rounded-[2rem] p-8">
                <h3 class="text-xl font-bold text-white mb-6">Profile Information</h3>
                
                <div v-if="isLoading" class="flex items-center gap-2 text-slate-500">
                    <div class="size-4 border-2 border-primary border-t-transparent rounded-full animate-spin"></div>
                    Loading profile...
                </div>

                <div v-else class="space-y-6 max-w-md">
                    <div class="space-y-2">
                         <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Email Address</label>
                         <input type="text" :value="user?.email" disabled class="w-full h-12 bg-white/5 border border-white/5 rounded-xl px-4 text-slate-400 cursor-not-allowed font-medium" />
                         <p class="text-[10px] text-slate-500">Email cannot be changed.</p>
                    </div>

                    <div class="space-y-2">
                         <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Full Name</label>
                         <input type="text" v-model="profileForm.name" class="w-full h-12 bg-white/5 border border-white/5 rounded-xl px-4 text-white focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium" />
                    </div>

                    <button @click="handleUpdateProfile" :disabled="isUpdating" class="h-12 px-6 bg-primary hover:bg-primary/90 text-background-dark font-bold rounded-xl shadow-lg shadow-primary/20 transition-transform active:scale-95 flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
                        <span v-if="isUpdating" class="size-4 border-2 border-background-dark border-t-transparent rounded-full animate-spin"></span>
                        <span v-else>Save Changes</span>
                    </button>
                </div>
            </div>

            <!-- Password Section -->
            <div class="bg-surface-forest border border-white/5 rounded-[2rem] p-8">
                <h3 class="text-xl font-bold text-white mb-6">Change Password</h3>
                
                <div class="space-y-6 max-w-md">
                    <div class="space-y-2">
                         <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Current Password</label>
                         <input type="password" v-model="passwordForm.current_password" class="w-full h-12 bg-white/5 border border-white/5 rounded-xl px-4 text-white focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium" />
                    </div>
                    
                    <div class="space-y-2">
                         <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">New Password</label>
                         <input type="password" v-model="passwordForm.new_password" class="w-full h-12 bg-white/5 border border-white/5 rounded-xl px-4 text-white focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium" />
                    </div>

                    <div class="space-y-2">
                         <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Confirm New Password</label>
                         <input type="password" v-model="passwordForm.confirm_password" class="w-full h-12 bg-white/5 border border-white/5 rounded-xl px-4 text-white focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium" />
                    </div>

                    <div v-if="passwordError" class="text-red-400 text-sm font-bold">{{ passwordError }}</div>
                    <div v-if="passwordMessage" class="text-green-400 text-sm font-bold">{{ passwordMessage }}</div>

                    <button @click="handleChangePassword" :disabled="isChangingPassword" class="h-12 px-6 bg-white/10 hover:bg-white/20 text-white font-bold rounded-xl border border-white/10 transition-colors flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
                        <span v-if="isChangingPassword" class="size-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
                        <span v-else>Update Password</span>
                    </button>
                </div>
            </div>

        </div>
    </div>
</div>
</template>

<style scoped>
.no-scrollbar::-webkit-scrollbar {
    display: none;
}
.no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>
