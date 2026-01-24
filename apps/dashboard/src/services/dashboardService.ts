import api from '../lib/api';
import type {
    DashboardStats,
    MonthlyExpense,
    CategoryExpense,
    ApiResponse
} from '../types';

export const dashboardService = {
    async getStats(): Promise<DashboardStats> {
        const response = await api.get<ApiResponse<DashboardStats>>('/dashboard/stats');
        return response.data.data;
    },

    async getExpensesByMonth(months: number = 12): Promise<MonthlyExpense[]> {
        const response = await api.get<ApiResponse<MonthlyExpense[]>>('/dashboard/expenses-by-month', {
            params: { months }
        });
        return response.data.data;
    },

    async getExpensesByCategory(): Promise<CategoryExpense[]> {
        const response = await api.get<ApiResponse<CategoryExpense[]>>('/dashboard/expenses-by-category');
        return response.data.data;
    }
};
