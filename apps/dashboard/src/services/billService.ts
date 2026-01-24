import api from '../lib/api';
import type {
    Bill,
    CreateBillInput,
    UpdateBillInput,
    BillFilters,
    PaginatedResponse,
    ApiResponse,
    BillActivity
} from '../types';

export const billService = {
    async list(filters?: BillFilters): Promise<PaginatedResponse<Bill>> {
        const params: any = {
            page: filters?.page,
            page_size: filters?.page_size,
            search: filters?.search,
            status: filters?.status
        };

        const response = await api.get<ApiResponse<any>>('/bills', { params });
        const { items, total_items, page, page_size, total_pages } = response.data.data;

        return {
            data: items || [],
            meta: {
                current_page: page,
                page_size,
                total_items,
                total_pages
            }
        };
    },

    async getById(id: string): Promise<Bill> {
        const response = await api.get<ApiResponse<Bill>>(`/bills/${id}`);
        return response.data.data;
    },

    async create(input: CreateBillInput): Promise<Bill> {
        const response = await api.post<ApiResponse<Bill>>('/bills', input);
        return response.data.data;
    },

    async update(id: string, input: UpdateBillInput): Promise<Bill> {
        const response = await api.put<ApiResponse<Bill>>(`/bills/${id}`, input);
        return response.data.data;
    },

    async delete(id: string): Promise<void> {
        await api.delete(`/bills/${id}`);
    },

    async pay(id: string, paidDate?: Date): Promise<Bill> {
        const response = await api.post<ApiResponse<Bill>>(`/bills/${id}/pay`, {
            paid_date: paidDate || new Date()
        });
        return response.data.data;
    },

    async getActivities(id: string): Promise<BillActivity[]> {
        const response = await api.get<ApiResponse<BillActivity[]>>(`/bills/${id}/activities`);
        return response.data.data;
    }
};
