import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query';
import { userService } from '../services/userService';
import type { ChangePasswordInput, UpdateProfileInput } from '../types';

export function useUserProfile() {
    return useQuery({
        queryKey: ['users', 'profile'],
        queryFn: userService.getProfile,
    });
}

export function useUpdateProfile() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (input: UpdateProfileInput) => userService.updateProfile(input),
        onSuccess: (data) => {
            queryClient.setQueryData(['users', 'profile'], data);
            queryClient.setQueryData(['auth', 'me'], data);
        },
    });
}

export function useChangePassword() {
    return useMutation({
        mutationFn: (input: ChangePasswordInput) => userService.changePassword(input),
    });
}
