<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { useBill, useBillActivities, usePayBill, useDeleteBill } from '../composables/useBills';
import { computed } from 'vue';

const route = useRoute();
const router = useRouter();
const billId = computed(() => route.params.id as string);

const { data: bill, isLoading: isLoadingBill } = useBill(billId);
const { data: activities, isLoading: isLoadingActivities } = useBillActivities(billId);

const { mutateAsync: payBill, isPending: isPaying } = usePayBill();
const { mutateAsync: deleteBill, isPending: isDeleting } = useDeleteBill();

const goBack = () => {
    router.back();
};

const handlePay = async () => {
    if (!bill.value) return;
    try {
        await payBill({ id: bill.value.id, date: new Date() });
        // Status will auto-update via query invalidation
    } catch (e) {
        console.error("Failed to pay bill", e); // In a real app show toast
    }
};

const handleDelete = async () => {
    if (!bill.value) return;
    try {
        await deleteBill(bill.value.id);
        router.push('/bills');
    } catch (e) {
        console.error("Failed to delete bill", e);
    }
};

const formatCurrency = (amount: string | number) => {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
    }).format(Number(amount));
};

const formatDate = (date: string) => {
    return new Date(date).toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric'
    });
};

const formatDateTime = (date: string) => {
    return new Date(date).toLocaleString('en-US', {
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    });
};

const getStatusColor = (status: string) => {
    switch (status) {
        case 'paid': return 'text-primary bg-primary/20 border-primary/20';
        case 'unpaid': return 'text-amber-500 bg-amber-500/20 border-amber-500/20';
        case 'overdue': return 'text-red-500 bg-red-500/20 border-red-500/20';
        default: return 'text-slate-400 bg-slate-700/20 border-slate-600/20';
    }
};

const isOverdue = computed(() => {
    if (!bill.value || bill.value.status === 'paid') return false;
    const dueDate = new Date(bill.value.due_date);
    return dueDate < new Date();
});

const daysLate = computed(() => {
    if (!bill.value || !isOverdue.value) return 0;
    const dueDate = new Date(bill.value.due_date);
    const now = new Date();
    const diffTime = Math.abs(now.getTime() - dueDate.getTime());
    return Math.ceil(diffTime / (1000 * 60 * 60 * 24)); 
});
</script>

<template>
<div class="flex-1 flex flex-col h-full relative bg-background-forest font-inter">
    <!-- Subtle Gradient -->
    <div class="absolute top-0 left-0 w-full h-96 bg-primary/5 blur-3xl pointer-events-none rounded-full -translate-y-1/2"></div>
    
    <!-- Top Header -->
    <header class="px-8 py-6 z-10 flex items-center justify-between shrink-0">
        <div class="flex items-center gap-4">
             <button @click="goBack" class="size-10 rounded-full hover:bg-white/10 flex items-center justify-center text-white transition-colors">
                 <span class="material-symbols-outlined">arrow_back</span>
             </button>
             <h1 class="text-2xl font-bold text-white">Bill Details</h1>
        </div>
        <div v-if="bill" class="flex items-center gap-3">
             <button class="flex items-center gap-2 px-4 py-2 bg-slate-800 hover:bg-slate-700 text-white text-sm font-medium rounded-lg border border-slate-700 transition-colors">
                 <span class="material-symbols-outlined text-lg">edit</span>
                 Edit
             </button>
             <button @click="handleDelete" :disabled="isDeleting" class="flex items-center gap-2 px-4 py-2 bg-red-500/10 hover:bg-red-500/20 text-red-500 text-sm font-medium rounded-lg border border-red-500/20 transition-colors disabled:opacity-50">
                 <span v-if="isDeleting" class="size-4 border-2 border-red-500 border-t-transparent rounded-full animate-spin"></span>
                 <span v-else class="material-symbols-outlined text-lg">delete</span>
                 <span v-if="!isDeleting">Delete</span>
             </button>
        </div>
    </header>

    <!-- Main Content -->
    <div v-if="isLoadingBill" class="flex-1 flex items-center justify-center z-10">
        <div class="size-12 rounded-full border-4 border-primary border-t-transparent animate-spin"></div>
    </div>
    
    <div v-else-if="!bill" class="flex-1 flex items-center justify-center flex-col gap-4 z-10 text-slate-400">
        <span class="material-symbols-outlined text-6xl">error_outline</span>
        <p>Bill not found</p>
        <button @click="goBack" class="text-primary hover:underline">Go Back</button>
    </div>

    <div v-else class="flex-1 px-8 pb-8 overflow-y-auto no-scrollbar z-10">
        <div class="max-w-6xl w-full grid grid-cols-1 lg:grid-cols-12 gap-8">
            
            <!-- Left Column: Bill Info -->
            <div class="lg:col-span-8 flex flex-col gap-6">
                
                <!-- Main Info Card -->
                <div class="bg-card-dark border border-white/5 rounded-[2rem] p-8 relative overflow-hidden">
                    <div class="flex flex-col md:flex-row md:items-start justify-between gap-6 mb-8">
                        <div>
                            <h2 class="text-3xl font-bold text-white mb-1">{{ bill.title }}</h2>
                            <p class="text-slate-400 text-sm">Invoice #{{ bill.invoice_number }}</p>
                        </div>
                        <div class="text-right">
                             <div class="text-4xl font-extrabold text-white tracking-tight">{{ formatCurrency(bill.amount) }}</div>
                             <div class="flex items-center justify-end gap-1 mt-1">
                                 <span class="size-2 rounded-full bg-primary/80"></span>
                                 <span class="text-xs text-slate-400 font-medium">USD</span>
                             </div>
                        </div>
                    </div>
                    
                    <div class="w-full h-px bg-white/5 my-8"></div>
                    
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-y-10 gap-x-8">
                        <div>
                             <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-3">VENDOR</p>
                             <div class="flex items-center gap-3">
                                 <div class="size-10 rounded-full bg-slate-700/50 flex items-center justify-center text-slate-300 overflow-hidden">
                                     <img v-if="bill.vendor?.logo_url" :src="bill.vendor.logo_url" class="w-full h-full object-cover">
                                     <span v-else class="material-symbols-outlined">domain</span>
                                 </div>
                                 <span class="text-white font-bold text-lg">{{ bill.vendor?.name || 'Unknown Vendor' }}</span>
                             </div>
                        </div>
                        <div>
                             <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-3">CATEGORY</p>
                             <div class="flex items-center gap-3">
                                 <span class="material-symbols-outlined text-primary">label</span>
                                 <span class="text-white font-bold text-lg">{{ bill.category?.name || 'Uncategorized' }}</span>
                             </div>
                        </div>
                         <div>
                             <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-3">FREQUENCY</p>
                             <div class="inline-flex items-center px-4 py-2 rounded-full bg-slate-800 text-slate-300 text-sm font-medium border border-slate-700 capitalize">
                                 {{ bill.is_recurring ? `Recurring: ${bill.recurring_frequency}` : 'One-time' }}
                             </div>
                        </div>
                         <div>
                             <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-3">STATUS</p>
                             <div class="flex items-center gap-2 text-white font-medium text-lg">
                                 <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-bold border capitalize" :class="getStatusColor(bill.status)">
                                     {{ bill.status }}
                                 </span>
                             </div>
                        </div>
                    </div>

                    <div class="mt-10" v-if="bill.notes">
                         <p class="text-[10px] font-bold text-slate-500 uppercase tracking-widest mb-3">NOTES</p>
                         <div class="bg-black/20 rounded-xl p-5 border border-white/5 italic text-slate-300">
                             "{{ bill.notes }}"
                         </div>
                    </div>
                </div>

                <!-- Attachments -->
                <div class="bg-card-dark border border-white/5 rounded-[2rem] p-8">
                     <div class="flex items-center justify-between mb-6">
                         <h3 class="text-lg font-bold text-white">Attachments</h3>
                         <button class="flex items-center gap-1.5 text-primary text-sm font-bold hover:text-white transition-colors">
                             <span class="material-symbols-outlined text-lg">add_circle</span>
                             Upload New
                         </button>
                     </div>

                     <div v-if="!bill.attachments?.length" class="text-slate-500 text-sm text-center py-4">No attachments</div>

                     <div v-for="att in bill.attachments" :key="att.id" class="bg-black/20 rounded-xl p-4 flex items-center justify-between group hover:bg-black/30 transition-colors cursor-pointer border border-transparent hover:border-white/5 mb-2">
                         <div class="flex items-center gap-4">
                             <div class="size-10 rounded-lg bg-red-500/10 flex items-center justify-center text-red-500">
                                 <span class="material-symbols-outlined text-xl">description</span>
                             </div>
                             <div>
                                 <p class="text-white font-medium group-hover:underline decoration-1 underline-offset-4">{{ att.file_name }}</p>
                                 <p class="text-xs text-slate-500">Document</p>
                             </div>
                         </div>
                         <button class="size-8 rounded-full hover:bg-white/10 flex items-center justify-center text-slate-400 hover:text-white transition-colors">
                             <span class="material-symbols-outlined">download</span>
                         </button>
                     </div>
                </div>

            </div>

            <!-- Right Column: Sidebar -->
            <div class="lg:col-span-4 space-y-6">
                
                <!-- Status Card -->
                <div class="bg-card-dark border border-white/5 rounded-[2rem] p-6 relative overflow-hidden">
                    <div class="flex items-center justify-between mb-6 relative z-10">
                        <span class="px-3 py-1 rounded-full text-[10px] font-bold border tracking-wider uppercase bg-opacity-20" :class="getStatusColor(bill.status)">● {{ bill.status }}</span>
                    </div>

                    <!-- Warning Triangle BG for Overdue -->
                    <div v-if="bill.status === 'overdue'" class="absolute -top-6 -right-6 text-red-500/5 rotate-12 pointer-events-none">
                         <span class="material-symbols-outlined text-[180px]" style="font-variation-settings: 'FILL' 1;">warning</span>
                    </div>
                    
                    <div class="relative z-10 mb-8">
                        <p class="text-slate-400 text-xs mb-1">Due Date</p>
                        <p class="text-2xl font-bold text-white">{{ formatDate(bill.due_date) }}</p>
                    </div>

                    <div v-if="bill.status === 'overdue'" class="relative z-10 bg-red-500/10 border border-red-500/10 rounded-xl flex items-center gap-3 p-4 mb-6">
                        <span class="material-symbols-outlined text-red-500">history</span>
                        <span class="text-red-400 font-bold text-sm">{{ daysLate }} days late</span>
                    </div>

                    <button v-if="bill.status !== 'paid'" @click="handlePay" :disabled="isPaying" class="relative z-10 w-full h-12 bg-primary hover:bg-primary/90 text-background-dark font-bold rounded-xl shadow-lg shadow-primary/20 transition-transform active:scale-95 flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
                        <span v-if="isPaying" class="size-5 border-2 border-background-dark border-t-transparent rounded-full animate-spin"></span>
                        <span v-else>Pay Now</span>
                        <span v-if="!isPaying" class="material-symbols-outlined text-lg">arrow_forward</span>
                    </button>
                </div>

                <!-- Activity Log -->
                <div class="bg-card-dark border border-white/5 rounded-[2rem] p-6">
                     <h3 class="text-lg font-bold text-white mb-6">Activity Log</h3>
                     
                     <div class="space-y-8 relative pl-2">
                         <!-- Vertical Line -->
                         <div class="absolute top-2 left-2 bottom-2 w-px bg-white/10"></div>
                         
                         <div v-if="isLoadingActivities" class="pl-6 text-sm text-slate-500">Loading activities...</div>
                         <div v-else-if="!activities?.length" class="pl-6 text-sm text-slate-500">No recent activity</div>

                         <div v-for="log in activities" :key="log.id" class="relative flex gap-4">
                             <div class="size-4 rounded-full bg-slate-600 border-4 border-card-dark relative z-10 shrink-0"></div>
                             <div class="flex flex-col -mt-1">
                                 <p class="text-white text-sm font-bold">{{ log.details }}</p>
                                 <p class="text-xs text-slate-500">{{ formatDateTime(log.created_at) }}</p>
                             </div>
                         </div>
                     </div>
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
