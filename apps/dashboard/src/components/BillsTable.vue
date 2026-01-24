<script setup lang="ts">
import { useBills } from '../composables/useBills';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const filters = ref({
    page: 1,
    page_size: 5
    // Default sorting is by due_date from backend logic if not specified, 
    // or we might need to add sort param if supported.
});

const { data, isLoading } = useBills(filters);

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
</script>

<template>
<div class="rounded-xl bg-card-dark border border-border-dark overflow-hidden flex flex-col">
<div class="p-6 border-b border-border-dark flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
<div>
<h3 class="text-white font-bold text-lg">Upcoming Bills</h3>
<p class="text-slate-400 text-sm">Manage your payable accounts</p>
</div>
<router-link to="/bills/add" class="flex items-center gap-2 bg-primary hover:bg-green-400 text-slate-900 px-4 py-2 rounded-full text-sm font-bold transition-colors">
<span class="material-symbols-outlined text-[18px]">add</span>
                            Add New Bill
                        </router-link>
</div>
<div class="overflow-x-auto">
<table class="w-full text-left text-sm text-slate-400">
<thead class="bg-slate-900/50 text-xs uppercase font-bold text-slate-500">
<tr>
<th class="px-6 py-4" scope="col">Vendor</th>
<th class="px-6 py-4" scope="col">Due Date</th>
<th class="px-6 py-4" scope="col">Amount</th>
<th class="px-6 py-4" scope="col">Status</th>
<th class="px-6 py-4 text-right" scope="col">Action</th>
</tr>
</thead>
<tbody class="divide-y divide-border-dark">

<tr v-if="isLoading">
    <td colspan="5" class="px-6 py-8 text-center">
        <div class="flex justify-center">
            <div class="size-6 rounded-full border-2 border-primary border-t-transparent animate-spin"></div>
        </div>
    </td>
</tr>

<tr v-else-if="!data?.data.length">
    <td colspan="5" class="px-6 py-8 text-center text-slate-500">
        No bills found.
    </td>
</tr>

<tr v-for="bill in data?.data" :key="bill.id" @click="navigateToBill(bill.id)" class="hover:bg-slate-800/50 transition-colors cursor-pointer">
<td class="px-6 py-4 font-medium text-white flex items-center gap-3">
<div class="size-8 rounded-full bg-slate-700 flex items-center justify-center text-slate-300 overflow-hidden">
    <img v-if="bill.vendor?.logo_url" :src="bill.vendor.logo_url" class="w-full h-full object-cover">
    <span v-else class="material-symbols-outlined text-[18px]">receipt_long</span>
</div>
                                        {{ bill.title }}
                                    </td>
<td class="px-6 py-4 font-medium" :class="{'text-red-400': bill.status === 'overdue'}">{{ formatDate(bill.due_date) }}</td>
<td class="px-6 py-4 text-white font-bold">{{ formatCurrency(bill.amount) }}</td>
<td class="px-6 py-4">
<span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-bold border capitalize" :class="getStatusColor(bill.status)">
<span class="size-1.5 rounded-full" :class="getStatusDot(bill.status)"></span>
                                            {{ bill.status }}
                                        </span>
</td>
<td class="px-6 py-4 text-right">
<button v-if="bill.status !== 'paid'" class="text-primary hover:text-green-300 font-medium text-sm" @click.stop="navigateToBill(bill.id)">Pay Now</button>
<button v-else class="text-slate-400 hover:text-white transition-colors" @click.stop="navigateToBill(bill.id)">
<span class="material-symbols-outlined">visibility</span>
</button>
</td>
</tr>

</tbody>
</table>
</div>
<div class="p-4 border-t border-border-dark flex justify-center">
<router-link to="/bills" class="text-sm font-medium text-slate-400 hover:text-primary transition-colors flex items-center gap-1">
                            View All Bills <span class="material-symbols-outlined text-[16px]">arrow_forward</span>
</router-link>
</div>
</div>
</template>
