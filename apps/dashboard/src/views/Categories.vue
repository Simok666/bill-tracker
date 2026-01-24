<script setup lang="ts">
import { useCategories, useDeleteCategory } from '../composables/useCategories';
import { useRouter } from 'vue-router';

const router = useRouter();
const { data: categories, isLoading } = useCategories();
const { mutateAsync: deleteCategory } = useDeleteCategory();

const handleDelete = async (id: string) => {
    if (confirm('Are you sure you want to delete this category?')) {
        try {
            await deleteCategory(id);
        } catch (e) {
            console.error("Failed to delete category", e);
        }
    }
};

const navigateToAdd = () => {
    router.push('/categories/add');
};
</script>

<template>
<div class="flex-1 flex flex-col h-full bg-background-forest p-8 overflow-hidden font-inter">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 mb-8 shrink-0">
        <div>
            <h1 class="text-3xl font-bold text-white mb-2">Categories</h1>
            <p class="text-slate-400 font-medium">Organize your expenses into meaningful groups</p>
        </div>
        <button @click="navigateToAdd" class="flex items-center gap-2 bg-primary hover:bg-green-400 text-slate-900 px-6 py-3 rounded-full font-bold shadow-lg shadow-primary/20 transition-all active:scale-95">
            <span class="material-symbols-outlined">add</span>
            Add Category
        </button>
    </div>

    <!-- Main Content Card -->
    <div class="flex-1 flex flex-col bg-surface-forest border border-white/5 rounded-[2.5rem] overflow-hidden shadow-2xl shadow-black/40">
        <!-- Table Header -->
        <div class="grid grid-cols-12 gap-4 px-8 py-5 border-b border-white/5 bg-white/[0.02]">
            <div class="col-span-1 text-xs font-bold text-slate-400 uppercase tracking-wider text-center">#</div>
            <div class="col-span-4 text-xs font-bold text-slate-400 uppercase tracking-wider">Category Name</div>
            <div class="col-span-6 text-xs font-bold text-slate-400 uppercase tracking-wider">Description</div>
            <div class="col-span-1 text-xs font-bold text-slate-400 uppercase tracking-wider text-center">Action</div>
        </div>

        <!-- Table Body -->
        <div class="flex-1 overflow-y-auto custom-scrollbar">
            <!-- Loading State -->
            <div v-if="isLoading" class="flex flex-col items-center justify-center py-20 text-slate-500">
                <div class="size-12 mb-4 rounded-full border-4 border-primary border-t-transparent animate-spin"></div>
                <p>Loading categories...</p>
            </div>

            <!-- Empty State -->
            <div v-else-if="!categories || categories.length === 0" class="flex flex-col items-center justify-center py-20 text-slate-500">
                <span class="material-symbols-outlined text-6xl mb-4 text-slate-700">folder_open</span>
                <p>No categories found. Add your first category to get started.</p>
            </div>

            <!-- Data Rows -->
            <div v-else v-for="(cat, index) in categories" :key="cat.id" class="group grid grid-cols-12 gap-4 px-8 py-5 items-center border-b border-white/5 hover:bg-white/[0.02] transition-colors relative">
                <div class="absolute left-0 top-0 bottom-0 w-1 bg-primary scale-y-0 group-hover:scale-y-100 transition-transform origin-center"></div>
                
                <div class="col-span-1 text-center text-slate-500 font-mono text-sm">{{ index + 1 }}</div>
                
                <div class="col-span-4 flex items-center gap-3">
                    <div class="size-2 rounded-full bg-primary"></div>
                    <span class="text-white font-bold text-sm">{{ cat.name }}</span>
                </div>

                <div class="col-span-6">
                    <span class="text-slate-400 text-sm italic">{{ cat.description || 'No description provided' }}</span>
                </div>

                <div class="col-span-1 flex justify-center">
                    <button @click="handleDelete(cat.id)" class="size-9 rounded-lg bg-red-500/10 border border-red-500/20 text-red-400 hover:bg-red-500 hover:text-white flex items-center justify-center transition-all">
                        <span class="material-symbols-outlined text-lg">delete</span>
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
</style>
