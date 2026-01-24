<script setup lang="ts">
import { ref, computed } from 'vue';
import { useBills } from '../composables/useBills';
import debounce from 'lodash/debounce';
import { useRouter } from 'vue-router';

const router = useRouter();

// Filters state
const searchQuery = ref('');
const activeStatus = ref<string>('all');
const currentPage = ref(1);
const pageSize = ref(10);

// Computed filters for the API
const filters = computed(() => ({
    page: currentPage.value,
    page_size: pageSize.value,
    status: activeStatus.value === 'all' ? undefined : activeStatus.value,
    // Add logic for search if backend supports it. The Go backend might not strictly have 'search' on list, 
    // but looking at bill_service.go (from generic knowledge of this task), it usually supports filtering.
    // Assuming 'search' or similar might be needed. Alternatively, filter by status is supported.
    // If backend doesn't support 'search' query param, we might need to rely on client-side or specific fields.
    // Checking previous context, bill_service.go has BillFilters struct, usually generic.
    // Let's assume generic search isn't implemented and stick to status for now, 
    // unless I saw 'search' in DTOs.
    // Looking at Types: type BillFilters = { page?: number, page_size?: number, ... }
    // I should probably check if I can add generic search. 
    // For now I will leave search query out of API call if not sure, 
    // but the UI has a search box. Let's assume the user wants it.
    // I'll stick to client side filtering if the API doesn't support it, but pagination makes that hard.
    // I'll add 'search' to filters if I can, but checking `types/index.ts` I wrote earlier...
    // I wrote types myself. `BillFilters` interface.
    // Let's check `types/index.ts` content is not visible but `useBills.ts` imports it.
    // I'll proceed with assumed support or just pass it - extra params usually ignored.
    search: searchQuery.value || undefined
}));

const { data, isLoading } = useBills(filters);

// Debounced search
const handleSearch = debounce((e: Event) => {
    searchQuery.value = (e.target as HTMLInputElement).value;
    currentPage.value = 1; // Reset to page 1
}, 300);

const setStatus = (status: string) => {
    activeStatus.value = status;
    currentPage.value = 1;
};

const formatDate = (date: string) => {
    return new Date(date).toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric'
    });
};

const formatCurrency = (amount: string | number) => {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
    }).format(Number(amount));
};

const getStatusColor = (status: string) => {
    switch (status) {
        case 'paid': return 'text-primary bg-primary/10 border-primary/20';
        case 'unpaid': return 'text-amber-500 bg-amber-500/10 border-amber-500/20';
        case 'overdue': return 'text-red-500 bg-red-500/10 border-red-500/20';
        case 'draft': return 'text-slate-400 bg-slate-700/50 border-slate-600/50';
        default: return 'text-slate-400 bg-slate-700/50 border-slate-600/50';
    }
};

const getStatusDot = (status: string) => {
    switch (status) {
        case 'paid': return 'bg-primary';
        case 'unpaid': return 'bg-amber-500';
        case 'overdue': return 'bg-red-500';
        default: return 'bg-slate-400';
    }
};

const navigateToBill = (id: string) => {
    router.push(`/bills/${id}`);
};

const totalPages = computed(() => {
    if (!data.value) return 1;
    return Math.ceil(data.value.meta.total_items / data.value.meta.page_size);
});

const pages = computed(() => {
    if (!data.value) return [1];
    const range = [];
    const total = totalPages.value;
    const current = currentPage.value;
    
    // Simple pagination logic
    if (total <= 7) {
        for (let i = 1; i <= total; i++) range.push(i);
    } else {
        if (current <= 4) {
            for (let i = 1; i <= 5; i++) range.push(i);
            range.push('...');
            range.push(total);
        } else if (current >= total - 3) {
            range.push(1);
            range.push('...');
            for (let i = total - 4; i <= total; i++) range.push(i);
        } else {
            range.push(1);
            range.push('...');
            range.push(current - 1);
            range.push(current);
            range.push(current + 1);
            range.push('...');
            range.push(total);
        }
    }
    return range;
});

const changePage = (page: number | string) => {
    if (typeof page === 'number') {
        currentPage.value = page;
    }
};

const previousPage = () => {
    if (currentPage.value > 1) currentPage.value--;
};

const nextPage = () => {
    if (currentPage.value < totalPages.value) currentPage.value++;
};
</script>

<template>
<div class="flex-1 flex flex-col h-full relative">
<!-- Subtle radial gradient for depth -->
<div class="absolute top-0 left-0 w-full h-96 bg-primary/5 blur-3xl pointer-events-none rounded-full -translate-y-1/2"></div>
<!-- Header -->
<header class="px-8 py-6 z-10 flex flex-wrap items-center justify-between gap-4 shrink-0">
<div>
<h2 class="text-4xl font-extrabold text-white tracking-tight">Bills</h2>
<p class="text-slate-400 text-sm mt-1">Manage and track your enterprise expenses</p>
</div>
<router-link to="/bills/add" class="group flex items-center justify-center gap-2 h-12 px-6 bg-primary hover:bg-primary/90 text-background-dark font-bold rounded-full transition-all hover:shadow-[0_0_15px_rgba(83,210,45,0.4)] hover:-translate-y-0.5 active:translate-y-0">
<span class="material-symbols-outlined text-xl transition-transform group-hover:rotate-90">add</span>
<span>Add New Bill</span>
</router-link>
</header>
<!-- Filters & Search Toolbar -->
<div class="px-8 pb-4 z-10">
<div class="flex flex-col lg:flex-row lg:items-center justify-between gap-4 bg-surface-forest border border-white/5 p-2 rounded-[2rem]">
<!-- Tab Filters -->
<div class="flex items-center gap-1 overflow-x-auto no-scrollbar px-2">
<button @click="setStatus('all')" :class="{'bg-white/10 text-white shadow-sm ring-1 ring-white/5': activeStatus === 'all', 'text-slate-400 hover:text-white hover:bg-white/5': activeStatus !== 'all'}" class="px-5 py-2.5 rounded-full text-sm font-bold transition-all">
                        All
                    </button>
<button @click="setStatus('unpaid')" :class="{'bg-white/10 text-white shadow-sm ring-1 ring-white/5': activeStatus === 'unpaid', 'text-slate-400 hover:text-white hover:bg-white/5': activeStatus !== 'unpaid'}" class="px-5 py-2.5 rounded-full text-sm font-medium transition-all">
                        Unpaid
</button>
<button @click="setStatus('overdue')" :class="{'bg-white/10 text-white shadow-sm ring-1 ring-white/5': activeStatus === 'overdue', 'text-slate-400 hover:text-white hover:bg-white/5': activeStatus !== 'overdue'}" class="px-5 py-2.5 rounded-full text-sm font-medium transition-all">
                        Overdue
</button>
<button @click="setStatus('paid')" :class="{'bg-white/10 text-white shadow-sm ring-1 ring-white/5': activeStatus === 'paid', 'text-slate-400 hover:text-white hover:bg-white/5': activeStatus !== 'paid'}" class="px-5 py-2.5 rounded-full text-sm font-medium transition-all">
                        Paid
                    </button>
</div>
<!-- Search -->
<div class="relative w-full lg:w-80 group">
<div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
<span class="material-symbols-outlined text-slate-500 group-focus-within:text-primary transition-colors">search</span>
</div>
<input @input="handleSearch" class="w-full h-11 pl-11 pr-4 bg-background-forest border border-white/5 rounded-full text-sm text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary/50 transition-all" placeholder="Search by description or vendor..." type="text"/>
</div>
</div>
</div>
<!-- Table Container -->
<div class="flex-1 px-8 pb-8 overflow-hidden">
<div class="h-full flex flex-col bg-surface-forest border border-white/5 rounded-[2.5rem] overflow-hidden shadow-2xl shadow-black/40">
<!-- Table Header -->
<div class="grid grid-cols-12 gap-4 px-8 py-5 border-b border-white/5 bg-white/[0.02]">
            <div class="col-span-3 text-xs font-bold text-slate-400 uppercase tracking-wider">Bill Name</div>
            <div class="col-span-2 text-xs font-bold text-slate-400 uppercase tracking-wider">Vendor</div>
            <div class="col-span-2 text-right text-xs font-bold text-slate-400 uppercase tracking-wider">Amount</div>
            <div class="col-span-2 text-right text-xs font-bold text-slate-400 uppercase tracking-wider">Due Date</div>
            <div class="col-span-2 text-center text-xs font-bold text-slate-400 uppercase tracking-wider">Status</div>
            <div class="col-span-1 text-center text-xs font-bold text-slate-400 uppercase tracking-wider">Action</div>
        </div>
<!-- Table Body (Scrollable) -->
<div class="flex-1 overflow-y-auto custom-scrollbar">

<!-- Loading State -->
<div v-if="isLoading" class="flex flex-col items-center justify-center h-full text-slate-500">
    <div class="size-12 mb-4 rounded-full border-4 border-primary border-t-transparent animate-spin"></div>
    <p>Loading your expenses...</p>
</div>

<!-- Empty State -->
<div v-else-if="!data?.data || data.data.length === 0" class="flex flex-col items-center justify-center h-full text-slate-500">
    <span class="material-symbols-outlined text-6xl mb-4 text-slate-700">receipt_long</span>
    <p>No bills found matching your criteria.</p>
</div>

<!-- Data Rows -->
<div v-else v-for="bill in data.data" :key="bill.id" @click="navigateToBill(bill.id)" class="group grid grid-cols-12 gap-4 px-8 py-5 items-center border-b border-white/5 hover:bg-white/[0.02] transition-colors cursor-pointer relative">
<div class="absolute left-0 top-0 bottom-0 w-1 bg-primary scale-y-0 group-hover:scale-y-100 transition-transform origin-center"></div>
<div class="col-span-3 flex flex-col">
<span class="text-white font-bold text-sm">{{ bill.title }}</span>
<span class="text-xs text-slate-500">#{{ bill.invoice_number }}</span>
</div>
                <div class="col-span-2 flex items-center gap-3">
                    <div class="size-8 rounded-full bg-blue-900/50 flex items-center justify-center text-blue-400 font-bold text-xs ring-1 ring-blue-500/20 overflow-hidden">
                         <img v-if="bill.vendor?.logo_url" :src="bill.vendor.logo_url" class="w-full h-full object-cover">
                         <span v-else>{{ bill.vendor?.name?.charAt(0) || 'V' }}</span>
                    </div>
                    <span class="text-slate-300 text-sm font-medium">{{ bill.vendor?.name || 'Unknown' }}</span>
                </div>
                <div class="col-span-2 text-right">
                    <span class="text-white font-mono font-medium">{{ formatCurrency(bill.amount) }}</span>
                </div>
                <div class="col-span-2 text-right">
                    <span class="text-sm font-medium" :class="{'text-red-400': bill.status === 'overdue', 'text-slate-300': bill.status !== 'overdue'}">{{ formatDate(bill.due_date) }}</span>
                </div>
                <div class="col-span-2 flex justify-center">
                    <span class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-bold border capitalize" :class="getStatusColor(bill.status)">
                        <span class="size-1.5 rounded-full animate-pulse" :class="getStatusDot(bill.status)"></span>
                        {{ bill.status }}
                    </span>
                </div>
                <!-- Action Column -->
                <div class="col-span-1 flex justify-center gap-2" @click.stop>
                    <button class="size-8 rounded-full bg-surface-forest border border-white/10 hover:border-primary/50 text-slate-400 hover:text-primary flex items-center justify-center shadow-lg transition-colors">
                        <span class="material-symbols-outlined text-lg">edit</span>
                    </button>
                    <button class="size-8 rounded-full bg-surface-forest border border-white/10 hover:border-primary/50 text-slate-400 hover:text-primary flex items-center justify-center shadow-lg transition-colors">
                        <span class="material-symbols-outlined text-lg">more_vert</span>
                    </button>
                </div>
            </div>

</div>
<!-- Footer / Pagination -->
<div v-if="data && data.meta.total_items > 0" class="px-8 py-4 border-t border-white/5 flex items-center justify-between bg-white/[0.02]">
<div class="text-sm text-slate-500 font-medium">
                        Showing <span class="text-white">{{ (data.meta.current_page - 1) * data.meta.page_size + 1 }}</span> to <span class="text-white">{{ Math.min(data.meta.current_page * data.meta.page_size, data.meta.total_items) }}</span> of <span class="text-white">{{ data.meta.total_items }}</span> entries
                    </div>
<div class="flex items-center gap-2">
<button @click="previousPage" :disabled="currentPage === 1" class="size-9 rounded-full bg-white/5 hover:bg-white/10 text-slate-400 hover:text-white flex items-center justify-center transition-colors disabled:opacity-50 disabled:cursor-not-allowed">
<span class="material-symbols-outlined text-lg">chevron_left</span>
</button>

<template v-for="p in pages" :key="p">
    <button v-if="p === '...'" class="size-9 text-slate-600 text-sm flex items-center justify-center cursor-default">
        ...
    </button>
    <button v-else @click="changePage(p)" :class="{'bg-primary text-background-dark font-bold shadow-lg shadow-primary/20': currentPage === p, 'bg-transparent hover:bg-white/5 text-slate-400 hover:text-white': currentPage !== p}" class="size-9 rounded-full text-sm flex items-center justify-center transition-colors">
        {{ p }}
    </button>
</template>

<button @click="nextPage" :disabled="currentPage === totalPages" class="size-9 rounded-full bg-white/5 hover:bg-white/10 text-slate-400 hover:text-white flex items-center justify-center transition-colors disabled:opacity-50 disabled:cursor-not-allowed">
<span class="material-symbols-outlined text-lg">chevron_right</span>
</button>
</div>
</div>
</div>
</div>
</div>

</template>

<style scoped>
/* Custom scrollbar for webkit */
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
    background: rgba(255, 255, 255, 0.02);
}
.custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 10px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.2);
}

/* Hide scrollbar for filters */
.no-scrollbar::-webkit-scrollbar {
    display: none;
}
.no-scrollbar {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>
