import api from '../lib/api';
import type {
    Vendor,
    CreateVendorInput,
    UpdateVendorInput,
    ApiResponse
} from '../types';

export const vendorService = {
    async list(): Promise<Vendor[]> {
        const response = await api.get<ApiResponse<Vendor[]>>('/vendors');
        return response.data.data;
    },

    async getById(id: string): Promise<Vendor> {
        const response = await api.get<ApiResponse<Vendor>>(`/vendors/${id}`);
        return response.data.data;
    },

    async create(input: CreateVendorInput): Promise<Vendor> {
        const response = await api.post<ApiResponse<Vendor>>('/vendors', input);
        return response.data.data;
    },

    async update(id: string, input: UpdateVendorInput): Promise<Vendor> {
        const response = await api.put<ApiResponse<Vendor>>(`/vendors/${id}`, input);
        return response.data.data;
    },

    async delete(id: string): Promise<void> {
        await api.delete(`/vendors/${id}`);
    }
};
