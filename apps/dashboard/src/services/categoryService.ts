import api from '../lib/api';
import type {
    Category,
    CreateCategoryInput,
    UpdateCategoryInput,
    ApiResponse
} from '../types';

export const categoryService = {
    async list(): Promise<Category[]> {
        const response = await api.get<ApiResponse<Category[]>>('/categories');
        return response.data.data;
    },

    async getById(id: string): Promise<Category> {
        const response = await api.get<ApiResponse<Category>>(`/categories/${id}`);
        return response.data.data;
    },

    async create(input: CreateCategoryInput): Promise<Category> {
        const response = await api.post<ApiResponse<Category>>('/categories', input);
        return response.data.data;
    },

    async update(id: string, input: UpdateCategoryInput): Promise<Category> {
        const response = await api.put<ApiResponse<Category>>(`/categories/${id}`, input);
        return response.data.data;
    },

    async delete(id: string): Promise<void> {
        await api.delete(`/categories/${id}`);
    }
};
