import api from '../lib/api';
import type {
    User,
    UpdateProfileInput,
    ChangePasswordInput,
    ApiResponse
} from '../types';

export const userService = {
    async getProfile(): Promise<User> {
        const response = await api.get<ApiResponse<User>>('/users/profile');
        return response.data.data;
    },

    async updateProfile(input: UpdateProfileInput): Promise<User> {
        const response = await api.put<ApiResponse<User>>('/users/profile', input);
        return response.data.data;
    },

    async changePassword(input: ChangePasswordInput): Promise<void> {
        await api.put('/users/password', input);
    }
};
