<script setup lang="ts">
import { useExpensesByCategory } from '../../composables/useDashboard';
import { computed } from 'vue';

const { data: expenses, isLoading } = useExpensesByCategory();

const colors = ['#53d22d', '#6366f1', '#f59e0b', '#ef4444', '#ec4899', '#8b5cf6'];

const chartData = computed(() => {
    if (!expenses.value) return [];
    
    let currentOffset = 0;
    return expenses.value.map((item, index) => {
        const percentage = Number(item.percentage); // Ensure it's a number
        const offset = -currentOffset;
        currentOffset += percentage;
        
        return {
            ...item,
            color: colors[index % colors.length],
            dashArray: `${percentage} ${100 - percentage}`,
            dashOffset: offset
        };
    });
});

const totalAmount = computed(() => {
    if (!expenses.value) return 0;
    return expenses.value.reduce((sum, item) => sum + Number(item.amount), 0);
});

const formatCurrency = (value: number) => {
    return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD',
        notation: 'compact'
    }).format(value);
};
</script>

<template>
<div class="lg:col-span-1 rounded-xl bg-card-dark border border-border-dark p-6 flex flex-col justify-between">
<div>
<h3 class="text-white font-bold text-lg mb-1">Expense by Category</h3>
<p class="text-slate-400 text-sm">Distribution of monthly costs</p>
</div>

<div v-if="isLoading" class="flex items-center justify-center py-12">
    <div class="size-8 rounded-full border-2 border-primary border-t-transparent animate-spin"></div>
</div>

<div v-else-if="!chartData.length" class="flex items-center justify-center py-12 text-slate-500 text-sm">
    No expense data available
</div>

<div v-else class="flex items-center justify-center py-6 relative">
<!-- SVG Donut Chart -->
<svg class="transform -rotate-90" height="200" viewbox="0 0 40 40" width="200">
<circle cx="20" cy="20" fill="transparent" r="15.915" stroke="#334155" stroke-width="5"></circle>
<!-- Segments -->
<circle v-for="segment in chartData" :key="segment.category_id"
    cx="20" cy="20" fill="transparent" r="15.915" 
    :stroke="segment.color" 
    :stroke-dasharray="segment.dashArray" 
    :stroke-dashoffset="segment.dashOffset" 
    stroke-width="5">
</circle>
</svg>
<div class="absolute inset-0 flex flex-col items-center justify-center">
<span class="text-slate-400 text-xs">Total</span>
<span class="text-white font-bold text-xl">{{ formatCurrency(totalAmount) }}</span>
</div>
</div>

<div class="grid grid-cols-2 gap-3 text-sm">
<div v-for="segment in chartData" :key="segment.category_id" class="flex items-center gap-2">
<div class="size-2 rounded-full" :style="{ backgroundColor: segment.color }"></div>
<span class="text-slate-300 truncate">{{ segment.category_name }}</span>
</div>
</div>
</div>
</template>
