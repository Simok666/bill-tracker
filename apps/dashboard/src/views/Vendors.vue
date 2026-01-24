<script setup lang="ts">
import { useVendors, useDeleteVendor } from '../composables/useVendors';
import { useRouter } from 'vue-router';

const router = useRouter();
const { data: vendors, isLoading } = useVendors();
const { mutateAsync: deleteVendor } = useDeleteVendor();

const handleDelete = async (id: string) => {
    if (confirm('Are you sure you want to delete this vendor?')) {
        try {
            await deleteVendor(id);
        } catch (e) {
            console.error("Failed to delete vendor", e);
        }
    }
};

const navigateToAdd = () => {
    router.push('/vendors/add');
};
</script>

<template>
<div class="flex-1 flex flex-col h-full bg-background-forest p-8 overflow-hidden font-inter">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 mb-8 shrink-0">
        <div>
            <h1 class="text-3xl font-bold text-white mb-2">Vendors</h1>
            <p class="text-slate-400 font-medium">Manage your service providers and utility companies</p>
        </div>
        <button @click="navigateToAdd" class="flex items-center gap-2 bg-primary hover:bg-green-400 text-slate-900 px-6 py-3 rounded-full font-bold shadow-lg shadow-primary/20 transition-all active:scale-95">
            <span class="material-symbols-outlined">add</span>
            Add Vendor
        </button>
    </div>

    <!-- Main Content Card -->
    <div class="flex-1 flex flex-col bg-surface-forest border border-white/5 rounded-[2.5rem] overflow-hidden shadow-2xl shadow-black/40">
        <!-- Table Header -->
        <div class="grid grid-cols-12 gap-4 px-8 py-5 border-b border-white/5 bg-white/[0.02]">
            <div class="col-span-5 text-xs font-bold text-slate-400 uppercase tracking-wider">Vendor Info</div>
            <div class="col-span-3 text-xs font-bold text-slate-400 uppercase tracking-wider">Website</div>
            <div class="col-span-3 text-xs font-bold text-slate-400 uppercase tracking-wider text-center">Contact</div>
            <div class="col-span-1 text-xs font-bold text-slate-400 uppercase tracking-wider text-center">Action</div>
        </div>

        <!-- Table Body -->
        <div class="flex-1 overflow-y-auto custom-scrollbar">
            <!-- Loading State -->
            <div v-if="isLoading" class="flex flex-col items-center justify-center py-20 text-slate-500">
                <div class="size-12 mb-4 rounded-full border-4 border-primary border-t-transparent animate-spin"></div>
                <p>Loading vendors...</p>
            </div>

            <!-- Empty State -->
            <div v-else-if="!vendors || vendors.length === 0" class="flex flex-col items-center justify-center py-20 text-slate-500">
                <span class="material-symbols-outlined text-6xl mb-4 text-slate-700">domain</span>
                <p>No vendors found. Add your first vendor to get started.</p>
            </div>

            <!-- Data Rows -->
            <div v-else v-for="vendor in vendors" :key="vendor.id" class="group grid grid-cols-12 gap-4 px-8 py-5 items-center border-b border-white/5 hover:bg-white/[0.02] transition-colors relative">
                <div class="absolute left-0 top-0 bottom-0 w-1 bg-primary scale-y-0 group-hover:scale-y-100 transition-transform origin-center"></div>
                
                <div class="col-span-5 flex items-center gap-4">
                    <div class="size-12 rounded-xl bg-blue-900/30 flex items-center justify-center text-blue-400 font-bold border border-blue-500/20 overflow-hidden shrink-0">
                        <img v-if="vendor.logo_url" :src="vendor.logo_url" class="w-full h-full object-contain" :alt="vendor.name">
                        <span v-else class="text-xl uppercase">{{ vendor.name.charAt(0) }}</span>
                    </div>
                    <div class="flex flex-col overflow-hidden">
                        <span class="text-white font-bold text-sm truncate">{{ vendor.name }}</span>
                        <span class="text-xs text-slate-500 truncate">{{ vendor.contact_info || 'No contact info' }}</span>
                    </div>
                </div>

                <div class="col-span-3">
                    <a v-if="vendor.website" :href="vendor.website" target="_blank" class="text-primary hover:underline text-sm truncate flex items-center gap-1.5">
                        {{ vendor.website.replace(/^https?:\/\//, '') }}
                        <span class="material-symbols-outlined text-[14px]">open_in_new</span>
                    </a>
                    <span v-else class="text-slate-600 text-sm">N/A</span>
                </div>

                <div class="col-span-3 text-center">
                    <span class="text-slate-300 text-sm italic">{{ vendor.contact_info || '-' }}</span>
                </div>

                <div class="col-span-1 flex justify-center gap-2">
                    <button @click="handleDelete(vendor.id)" class="size-10 rounded-xl bg-red-500/10 border border-red-500/20 text-red-400 hover:bg-red-500 hover:text-white flex items-center justify-center transition-all shadow-lg shadow-red-500/10">
                        <span class="material-symbols-outlined text-xl">delete</span>
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.05);
    border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.1);
}
</style>
