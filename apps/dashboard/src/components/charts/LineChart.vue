<script setup lang="ts">
import { useExpensesByMonth } from '../../composables/useDashboard';
import { computed } from 'vue';

const { data: expenses, isLoading } = useExpensesByMonth(12);

const chartData = computed(() => {
    if (!expenses.value || expenses.value.length === 0) return { path: '', area: '', labels: [] };

    const data = expenses.value;
    const values = data.map(d => Number(d.amount));
    const max = Math.max(...values, 100); // Ensure non-zero max
    
    // Scale to viewBox 100x50, keeping 5px padding top/bottom
    // Y range: 5 to 45 (height 40)
    // X range: 0 to 100
    
    const points = values.map((val, index) => {
        const x = (index / (values.length - 1)) * 100;
        const normalizedVal = val / max;
        const y = 45 - (normalizedVal * 40); // Invert Y
        return `${x},${y}`;
    });

    const pathD = `M${points.join(' L')}`;
    const areaD = `${pathD} V50 H0 Z`;
    
    const labels = data.map(d => {
        const date = new Date(d.month + '-01'); // Append day to make parsing easier
        return date.toLocaleDateString('en-US', { month: 'short' });
    });

    // Pick ~6 labels to show evenly
    const displayLabels = [];
    const step = Math.ceil(labels.length / 6);
    for (let i = 0; i < labels.length; i += step) {
        displayLabels.push(labels[i]);
    }

    return { path: pathD, area: areaD, labels: displayLabels };
});
</script>

<template>
<div class="lg:col-span-2 rounded-xl bg-card-dark border border-border-dark p-6 flex flex-col">
<div class="flex justify-between items-start mb-6">
<div>
<h3 class="text-white font-bold text-lg mb-1">Cash Flow Trend</h3>
<p class="text-slate-400 text-sm">Actual Spending over Time</p>
</div>
<select class="bg-slate-900 border border-slate-700 text-slate-300 text-sm rounded-lg focus:ring-primary focus:border-primary block p-2">
<option>Last 12 Months</option>
</select>
</div>

<div v-if="isLoading" class="flex-1 flex items-center justify-center">
    <div class="size-8 rounded-full border-2 border-primary border-t-transparent animate-spin"></div>
</div>

<div v-else-if="!chartData.path" class="flex-1 flex items-center justify-center text-slate-500 text-sm">
    No expense history available
</div>

<div v-else class="flex-1 min-h-[200px] w-full relative">
<!-- Grid Lines -->
<div class="absolute inset-0 flex flex-col justify-between text-slate-600 text-xs pointer-events-none">
<div class="border-b border-slate-700/50 w-full h-0"></div>
<div class="border-b border-slate-700/50 w-full h-0"></div>
<div class="border-b border-slate-700/50 w-full h-0"></div>
<div class="border-b border-slate-700/50 w-full h-0"></div>
</div>
<!-- Chart SVG -->
<svg class="absolute inset-0 w-full h-full" preserveAspectRatio="none" viewBox="0 0 100 50">
<!-- Gradient Definition -->
<defs>
<linearGradient id="chartGradient" x1="0" x2="0" y1="0" y2="1">
<stop offset="0%" stop-color="#53d22d" stop-opacity="0.2"></stop>
<stop offset="100%" stop-color="#53d22d" stop-opacity="0"></stop>
</linearGradient>
</defs>
<!-- The Line -->
<path :d="chartData.path" fill="none" stroke="#53d22d" stroke-width="0.8" vector-effect="non-scaling-stroke"></path>
<path :d="chartData.area" fill="url(#chartGradient)" stroke="none"></path>
</svg>
</div>
<div class="flex justify-between mt-4 px-2 text-slate-500 text-xs font-medium uppercase tracking-wider">
<span v-for="label in chartData.labels" :key="label">{{ label }}</span>
</div>
</div>
</template>
