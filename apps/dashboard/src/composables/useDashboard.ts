import { useQuery } from '@tanstack/vue-query';
import { dashboardService } from '../services/dashboardService';
import { type Ref, unref } from 'vue';

export function useDashboardStats() {
    return useQuery({
        queryKey: ['dashboard', 'stats'],
        queryFn: dashboardService.getStats,
    });
}

export function useExpensesByMonth(months: Ref<number> | number = 12) {
    return useQuery({
        queryKey: ['dashboard', 'expenses-by-month', months],
        queryFn: () => dashboardService.getExpensesByMonth(unref(months)),
    });
}

export function useExpensesByCategory() {
    return useQuery({
        queryKey: ['dashboard', 'expenses-by-category'],
        queryFn: dashboardService.getExpensesByCategory,
    });
}
