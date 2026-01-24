import api from '../lib/api';
import type {
    LoginInput,
    RegisterInput,
    AuthResponse,
    User,
    ApiResponse
} from '../types';

export const authService = {
    async register(input: RegisterInput): Promise<AuthResponse> {
        const response = await api.post<ApiResponse<AuthResponse>>('/auth/register', input);
        return response.data.data;
    },

    async login(input: LoginInput): Promise<AuthResponse> {
        const response = await api.post<ApiResponse<AuthResponse>>('/auth/login', input);
        return response.data.data;
    },

    async logout(): Promise<void> {
        await api.post('/auth/logout');
    },

    async me(): Promise<User> {
        const response = await api.get<ApiResponse<User>>('/auth/me');
        return response.data.data;
    }
};
