<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useCreateCategory } from '../composables/useCategories';

const router = useRouter();
const { mutateAsync: createCategory, isPending } = useCreateCategory();

const form = ref({
    name: '',
    description: ''
});

const handleSubmit = async () => {
    try {
        await createCategory(form.value);
        router.push('/categories');
    } catch (e) {
        console.error("Failed to create category", e);
    }
};

const cancel = () => {
    router.back();
};
</script>

<template>
<div class="flex-1 flex flex-col h-full bg-background-forest p-8 overflow-y-auto custom-scrollbar font-inter">
    <div class="max-w-2xl mx-auto w-full">
        <!-- Header -->
        <div class="flex items-center gap-4 mb-8">
            <button @click="cancel" class="size-10 rounded-full hover:bg-white/10 flex items-center justify-center text-white transition-colors">
                <span class="material-symbols-outlined">arrow_back</span>
            </button>
            <h1 class="text-3xl font-bold text-white">Add New Category</h1>
        </div>

        <div class="bg-surface-forest border border-white/5 rounded-[2.5rem] p-10 shadow-2xl relative overflow-hidden">
            <!-- Decorative Accent -->
            <div class="absolute -top-24 -right-24 size-64 bg-primary/10 blur-[100px] rounded-full"></div>
            
            <form @submit.prevent="handleSubmit" class="space-y-8 relative z-10">
                <div class="space-y-6">
                    <!-- Category Name -->
                    <div class="space-y-2">
                        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Category Name</label>
                        <input 
                            v-model="form.name" 
                            type="text" 
                            required
                            placeholder="e.g. Utilities, Subscriptions, Rent" 
                            class="w-full h-14 bg-white/5 border border-white/5 rounded-2xl px-5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:bg-white/10 transition-all font-medium" 
                        />
                    </div>

                    <!-- Description -->
                    <div class="space-y-2">
                        <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Description (Optional)</label>
                        <textarea 
                            v-model="form.description" 
                            placeholder="What kind of bills belong in this category?" 
                            class="w-full h-32 bg-white/5 border border-white/5 rounded-2xl p-5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:bg-white/10 transition-all font-medium resize-none"
                        ></textarea>
                    </div>
                </div>

                <!-- Actions -->
                <div class="flex items-center justify-end gap-4 pt-4">
                    <button 
                        type="button" 
                        @click="cancel" 
                        class="px-8 py-4 rounded-2xl border border-white/10 text-white font-bold hover:bg-white/5 transition-colors"
                    >
                        Cancel
                    </button>
                    <button 
                        type="submit" 
                        :disabled="isPending"
                        class="px-10 py-4 bg-primary hover:bg-green-400 disabled:opacity-50 text-slate-900 font-bold rounded-2xl shadow-lg shadow-primary/20 transition-all active:scale-95 flex items-center gap-2"
                    >
                        <span v-if="isPending" class="size-5 border-2 border-slate-900 border-t-transparent rounded-full animate-spin"></span>
                        <span v-else class="material-symbols-outlined">check</span>
                        Create Category
                    </button>
                </div>
            </form>
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
