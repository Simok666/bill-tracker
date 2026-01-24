<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useVendors } from '../composables/useVendors';
import { useCategories } from '../composables/useCategories';
import { useCreateBill } from '../composables/useBills';
import { useQueryClient } from '@tanstack/vue-query';

const router = useRouter();
const queryClient = useQueryClient();

const { data: vendors } = useVendors();
const { data: categories } = useCategories();
const { mutateAsync: createBill, isPending: isCreating } = useCreateBill();

const formData = ref({
    title: '',
    vendor_id: '',
    category_id: '',
    amount: '',
    due_date: new Date().toISOString().split('T')[0],
    description: '',
    is_recurring: false,
    recurring_frequency: 'monthly' as 'monthly' | 'weekly' | 'yearly',
    invoice_number: `INV-${new Date().getFullYear()}-${Math.floor(Math.random() * 1000)}`
});

const handleSubmit = async () => {
    try {
        await createBill({
            title: formData.value.title,
            amount: Number(formData.value.amount),
            due_date: new Date(formData.value.due_date || '').toISOString(),
            vendor_id: formData.value.vendor_id || undefined,
            category_id: formData.value.category_id || undefined,
            notes: formData.value.description, // Map description to notes
            is_recurring: formData.value.is_recurring,
            recurring_frequency: formData.value.is_recurring ? formData.value.recurring_frequency : undefined,
            invoice_number: formData.value.invoice_number,
            currency: 'USD',
            status: 'unpaid'
        });
        
        // Invalidate queries to refresh lists
        queryClient.invalidateQueries({ queryKey: ['bills'] });
        queryClient.invalidateQueries({ queryKey: ['dashboard-stats'] });
        
        router.push('/bills');
    } catch (e) {
        console.error("Failed to create bill", e); // Show toast in real app
    }
};

const handleSaveDraft = async () => {
     try {
        await createBill({
            title: formData.value.title,
            invoice_number: formData.value.invoice_number,
            amount: Number(formData.value.amount),
            due_date: new Date(formData.value.due_date || '').toISOString(),
            vendor_id: formData.value.vendor_id || undefined,
            category_id: formData.value.category_id || undefined,
            notes: formData.value.description,
            status: 'draft',
            currency: 'USD'
        });
        router.push('/bills');
    } catch (e) {
        console.error("Failed to save draft", e);
    }
};
</script>

<template>
<div class="flex-1 flex flex-col h-full relative bg-background-forest">
    <!-- Subtle Gradient from top -->
    <div class="absolute top-0 left-0 w-full h-96 bg-primary/5 blur-3xl pointer-events-none rounded-full -translate-y-1/2"></div>
    
    <!-- Header -->
    <header class="px-8 py-6 z-10 shrink-0">
        <div class="flex items-center gap-2 mb-1">
             <span class="material-symbols-outlined text-primary text-sm font-bold">add_circle</span>
             <span class="text-primary text-xs font-bold tracking-wider uppercase">New Entry</span>
        </div>
        <h2 class="text-4xl font-extrabold text-white tracking-tight">Add New Enterprise Bill</h2>
        <p class="text-slate-400 text-sm mt-1 max-w-2xl">Register a new payable for tracking. Ensure vendor details and amounts are accurate to prevent late fees.</p>
    </header>

    <!-- Main Content Area -->
    <div class="flex-1 px-8 pb-8 overflow-y-auto no-scrollbar z-10">
        <div class="max-w-6xl w-full grid grid-cols-1 lg:grid-cols-12 gap-8">
            
            <!-- Left Column: Forms -->
            <div class="lg:col-span-8 space-y-6">
                
                <!-- Bill Details Card -->
                <div class="bg-surface-forest border border-white/5 rounded-[2rem] p-8 relative overflow-hidden group">
                    <div class="absolute top-0 right-0 p-8 opacity-20 group-hover:opacity-40 transition-opacity">
                         <!-- Decorative Icon -->
                         <span class="material-symbols-outlined text-[120px] text-primary/10 rotate-12">description</span>
                    </div>

                    <div class="flex items-center gap-3 mb-6 relative z-10">
                        <div class="size-10 rounded-full bg-primary/10 flex items-center justify-center text-primary">
                            <span class="material-symbols-outlined">description</span>
                        </div>
                        <h3 class="text-xl font-bold text-white">Bill Details</h3>
                    </div>

                    <div class="space-y-6 relative z-10">
                        <div class="space-y-2">
                             <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Bill Title</label>
                             <input type="text" v-model="formData.title" placeholder="e.g. Q3 Server Maintenance" class="w-full h-14 bg-white/5 border border-white/5 rounded-2xl px-5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:bg-white/10 transition-all font-medium" />
                        </div>
                        
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                             <div class="space-y-2">
                                 <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Vendor</label>
                                 <div class="relative">
                                     <select v-model="formData.vendor_id" class="w-full h-14 bg-white/5 border border-white/5 rounded-2xl px-5 text-white appearance-none focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium cursor-pointer">
                                         <option value="" disabled>Select Vendor</option>
                                         <option v-for="vendor in vendors" :key="vendor.id" :value="vendor.id">{{ vendor.name }}</option>
                                     </select>
                                     <span class="material-symbols-outlined absolute right-4 top-1/2 -translate-y-1/2 text-slate-500 pointer-events-none">expand_more</span>
                                 </div>
                             </div>
                             <div class="space-y-2">
                                 <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Category</label>
                                 <div class="relative">
                                     <select v-model="formData.category_id" class="w-full h-14 bg-white/5 border border-white/5 rounded-2xl px-5 text-white appearance-none focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium cursor-pointer">
                                         <option value="" disabled>Select Category</option>
                                         <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
                                     </select>
                                     <span class="material-symbols-outlined absolute right-4 top-1/2 -translate-y-1/2 text-slate-500 pointer-events-none">expand_more</span>
                                 </div>
                             </div>
                        </div>

                        <div class="space-y-2">
                             <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Description/Notes</label>
                             <textarea v-model="formData.description" placeholder="Additional details..." class="w-full h-24 bg-white/5 border border-white/5 rounded-2xl p-5 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:bg-white/10 transition-all font-medium resize-none"></textarea>
                        </div>
                    </div>
                </div>

                <!-- Payment Information Card -->
                <div class="bg-surface-forest border border-white/5 rounded-[2rem] p-8">
                    <div class="flex items-center gap-3 mb-6">
                        <div class="size-10 rounded-full bg-primary/10 flex items-center justify-center text-primary">
                            <span class="material-symbols-outlined">payments</span>
                        </div>
                        <h3 class="text-xl font-bold text-white">Payment Information</h3>
                    </div>
                    
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                        <div class="space-y-2">
                             <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Total Amount</label>
                             <div class="relative">
                                 <span class="absolute left-5 top-1/2 -translate-y-1/2 text-primary font-bold">$</span>
                                 <input type="number" step="0.01" v-model="formData.amount" placeholder="0.00" class="w-full h-14 bg-white/5 border border-white/5 rounded-2xl pl-12 pr-5 text-white placeholder-slate-500 font-mono text-lg font-medium focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all" />
                             </div>
                        </div>
                        <div class="space-y-2">
                             <label class="text-xs font-bold text-slate-400 uppercase tracking-wider">Due Date</label>
                             <div class="relative">
                                 <input type="date" v-model="formData.due_date" class="w-full h-14 bg-white/5 border border-white/5 rounded-2xl px-5 text-white focus:outline-none focus:ring-2 focus:ring-primary/50 transition-all font-medium" />
                             </div>
                        </div>
                    </div>
                </div>

            </div>

            <!-- Right Column: Sidebar -->
            <div class="lg:col-span-4 space-y-6">
                
                <!-- Recurring Bill Card -->
                <div class="bg-surface-forest border border-white/5 rounded-[2rem] p-6">
                    <div class="flex items-center justify-between mb-6">
                        <h3 class="text-lg font-bold text-white">Recurring Bill</h3>
                        <!-- Custom Toggle -->
                        <button class="w-12 h-6 rounded-full transition-colors relative" :class="formData.is_recurring ? 'bg-primary' : 'bg-white/10'" @click="formData.is_recurring = !formData.is_recurring">
                             <div class="absolute top-1 left-1 bg-white size-4 rounded-full transition-transform" :class="formData.is_recurring ? 'translate-x-6 shadow-sm' : ''"></div>
                        </button>
                    </div>
                    
                    <div class="space-y-4" :class="{ 'opacity-50 pointer-events-none': !formData.is_recurring }">
                        <div class="space-y-2">
                             <label class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">Frequency</label>
                             <div class="relative">
                                 <select v-model="formData.recurring_frequency" class="w-full h-12 bg-white/5 border border-white/5 rounded-xl px-4 text-white text-sm appearance-none focus:outline-none focus:ring-2 focus:ring-primary/50 font-medium">
                                     <option value="monthly">Monthly</option>
                                     <option value="weekly">Weekly</option>
                                     <option value="yearly">Yearly</option>
                                 </select>
                                 <span class="material-symbols-outlined absolute right-4 top-1/2 -translate-y-1/2 text-slate-500 text-sm pointer-events-none">expand_more</span>
                             </div>
                        </div>

                         <div class="flex items-start gap-2 bg-primary/10 p-3 rounded-xl border border-primary/10">
                              <span class="material-symbols-outlined text-primary text-sm mt-0.5">info</span>
                              <p class="text-[10px] text-slate-300 leading-relaxed">This will automatically generate a draft bill on the selected date.</p>
                         </div>
                    </div>
                </div>

                <!-- Summary Preview -->
                <div class="bg-surface-forest border border-white/5 rounded-[2rem] p-6">
                    <h3 class="text-xs font-bold text-slate-400 uppercase tracking-wider mb-4">Summary Preview</h3>
                    
                    <div class="space-y-3 mb-6">
                        <div class="flex items-center justify-between">
                            <span class="text-slate-400 text-sm">Status</span>
                            <span class="bg-amber-500/10 text-amber-500 text-xs font-bold px-2 py-1 rounded-full border border-amber-500/20">Unpaid</span>
                        </div>
                        <div class="flex items-center justify-between border-b border-white/5 pb-3">
                            <span class="text-slate-400 text-sm">Currency</span>
                            <span class="text-white font-bold text-sm">USD ($)</span>
                        </div>
                    </div>

                    <div class="flex gap-3 bg-white/5 p-3 rounded-xl mb-6">
                         <span class="material-symbols-outlined text-slate-400">receipt</span>
                         <p class="text-[10px] text-slate-400 leading-relaxed">By saving, you agree to the internal financial compliance policy.</p>
                    </div>

                    <div class="grid grid-cols-2 gap-3 mb-3">
                         <button @click="handleSaveDraft" :disabled="isCreating" class="flex items-center justify-center h-12 rounded-full border border-white/10 text-white font-bold hover:bg-white/5 transition-colors text-sm disabled:opacity-50">Save Draft</button>
                         <router-link to="/bills" class="flex items-center justify-center h-12 rounded-full border border-white/10 text-white font-bold hover:bg-white/5 transition-colors text-sm">Cancel</router-link>
                    </div>
                    <button @click="handleSubmit" :disabled="isCreating || !formData.amount" class="w-full h-14 bg-primary hover:bg-primary/90 text-background-dark font-bold rounded-full shadow-lg shadow-primary/20 transition-transform active:scale-95 flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
                        <span v-if="isCreating" class="size-5 border-2 border-background-dark border-t-transparent rounded-full animate-spin"></span>
                        <span v-else class="material-symbols-outlined">check</span>
                        <span v-if="!isCreating">Save Bill</span>
                    </button>

                </div>

            </div>
        
        </div>
    </div>
</div>
</template>

<style scoped>
/* Hide scrollbar for clean UI */
.no-scrollbar::-webkit-scrollbar {
    display: none;
}
.no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>
