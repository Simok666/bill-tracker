import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query';
import { computed, ref } from 'vue';
import { authService } from '../services/authService';
import type { LoginInput, RegisterInput, User } from '../types';
import { useRouter } from 'vue-router';

const token = ref(localStorage.getItem('token'));
const currentUser = ref<User | null>(null);
const isAuthenticated = computed(() => !!token.value);

export function useAuth() {
    const queryClient = useQueryClient();
    const router = useRouter();

    const { data: user, isLoading: isLoadingUser } = useQuery({
        queryKey: ['auth', 'me'],
        queryFn: authService.me,
        enabled: isAuthenticated,
        retry: false,
        staleTime: 5 * 60 * 1000, // 5 minutes
    });

    const loginMutation = useMutation({
        mutationFn: (input: LoginInput) => authService.login(input),
        onSuccess: (data) => {
            token.value = data.token;
            localStorage.setItem('token', data.token);
            queryClient.setQueryData(['auth', 'me'], data.user);
            router.push('/');
        },
    });

    const registerMutation = useMutation({
        mutationFn: (input: RegisterInput) => authService.register(input),
        onSuccess: (data) => {
            token.value = data.token;
            localStorage.setItem('token', data.token);
            queryClient.setQueryData(['auth', 'me'], data.user);
            router.push('/');
        },
    });

    const logoutMutation = useMutation({
        mutationFn: () => authService.logout(),
        onSettled: () => {
            token.value = null;
            currentUser.value = null;
            localStorage.removeItem('token');
            queryClient.clear();
            router.push('/login');
        },
    });

    return {
        token,
        currentUser: user,
        isAuthenticated,
        isLoadingUser,
        login: loginMutation.mutate,
        loginAsync: loginMutation.mutateAsync,
        isLoggingIn: loginMutation.isPending,
        loginError: loginMutation.error,
        register: registerMutation.mutate,
        registerAsync: registerMutation.mutateAsync,
        isRegistering: registerMutation.isPending,
        registerError: registerMutation.error,
        logout: logoutMutation.mutate,
        isLoggingOut: logoutMutation.isPending,
    };
}
