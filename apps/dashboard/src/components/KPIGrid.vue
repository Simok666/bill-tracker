<script setup lang="ts">
import { useDashboardStats } from '../composables/useDashboard';

const { data: stats, isLoading } = useDashboardStats();

const formatCurrency = (value: string | number) => {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
    }).format(Number(value));
};

const formatPercent = (value: number) => {
    const prefix = value >= 0 ? '+' : '';
    return `${prefix}${value.toFixed(1)}%`;
};
</script>

<template>
<div v-if="isLoading" class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4 animate-pulse">
    <div v-for="i in 4" :key="i" class="h-32 rounded-xl bg-card-dark border border-border-dark"></div>
</div>
<div v-else class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4">
<!-- Total Expense -->
<div class="relative overflow-hidden rounded-xl p-6 bg-gradient-to-br from-indigo-600 to-slate-800 border border-indigo-500/30 shadow-lg shadow-indigo-900/20 group hover:shadow-indigo-900/40 transition-all">
<div class="absolute -right-6 -top-6 size-32 rounded-full bg-white/5 blur-3xl"></div>
<div class="flex flex-col gap-4 relative z-10">
<div class="flex justify-between items-start">
<div class="p-2 bg-indigo-500/20 rounded-lg text-indigo-100">
<span class="material-symbols-outlined">payments</span>
</div>
<span class="flex items-center text-xs font-bold text-indigo-200 bg-indigo-500/20 px-2 py-1 rounded-full">{{ formatPercent(stats?.expense_change_percent || 0) }}</span>
</div>
<div>
<p class="text-indigo-200 text-sm font-medium mb-1">Total Expense</p>
<h3 class="text-white text-2xl font-bold tracking-tight">{{ formatCurrency(stats?.total_expense || 0) }}</h3>
</div>
</div>
</div>
<!-- Paid This Month -->
<div class="rounded-xl p-6 bg-card-dark border border-border-dark shadow-sm hover:border-primary/50 transition-colors group">
<div class="flex flex-col gap-4">
<div class="flex justify-between items-start">
<div class="p-2 bg-primary/10 rounded-lg text-primary">
<span class="material-symbols-outlined">check_circle</span>
</div>
<!-- <span class="flex items-center text-xs font-bold text-primary bg-primary/10 px-2 py-1 rounded-full">+5%</span> -->
</div>
<div>
<p class="text-slate-400 text-sm font-medium mb-1">Paid This Month</p>
<h3 class="text-white text-2xl font-bold tracking-tight">{{ formatCurrency(stats?.paid_this_month || 0) }}</h3>
</div>
</div>
</div>
<!-- Unpaid Amount -->
<div class="rounded-xl p-6 bg-card-dark border border-border-dark shadow-sm hover:border-amber-500/50 transition-colors group">
<div class="flex flex-col gap-4">
<div class="flex justify-between items-start">
<div class="p-2 bg-amber-500/10 rounded-lg text-amber-500">
<span class="material-symbols-outlined">pending_actions</span>
</div>
<span class="flex items-center text-xs font-bold text-slate-400 bg-slate-700/50 px-2 py-1 rounded-full">--</span>
</div>
<div>
<p class="text-slate-400 text-sm font-medium mb-1">Unpaid Amount</p>
<h3 class="text-amber-500 text-2xl font-bold tracking-tight">{{ formatCurrency(stats?.unpaid_amount || 0) }}</h3>
</div>
</div>
</div>
<!-- Overdue Bills -->
<div class="rounded-xl p-6 bg-card-dark border border-red-900/50 shadow-sm hover:border-red-500/50 transition-colors group relative overflow-hidden">
<!-- Subtle red glow for urgency -->
<div class="absolute -right-10 -bottom-10 size-24 rounded-full bg-red-500/10 blur-xl"></div>
<div class="flex flex-col gap-4 relative z-10">
<div class="flex justify-between items-start">
<div class="p-2 bg-red-500/10 rounded-lg text-red-500">
<span class="material-symbols-outlined">warning</span>
</div>
<span v-if="stats?.overdue_bills_count" class="flex items-center text-xs font-bold text-red-400 bg-red-950 px-2 py-1 rounded-full border border-red-900">High Priority</span>
</div>
<div>
<p class="text-slate-400 text-sm font-medium mb-1">Overdue Bills</p>
<h3 class="text-red-500 text-2xl font-bold tracking-tight">{{ stats?.overdue_bills_count || 0 }} Bills</h3>
</div>
</div>
</div>
</div>
</template>
