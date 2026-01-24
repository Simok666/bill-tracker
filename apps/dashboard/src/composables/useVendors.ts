import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query';
import { vendorService } from '../services/vendorService';
import type { CreateVendorInput, UpdateVendorInput } from '../types';
import { type Ref, unref } from 'vue';

export function useVendors() {
    return useQuery({
        queryKey: ['vendors'],
        queryFn: vendorService.list,
    });
}

export function useVendor(id: Ref<string>) {
    return useQuery({
        queryKey: ['vendors', id],
        queryFn: () => vendorService.getById(unref(id)),
        enabled: () => !!unref(id),
    });
}

export function useCreateVendor() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (input: CreateVendorInput) => vendorService.create(input),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['vendors'] });
        },
    });
}

export function useUpdateVendor() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, input }: { id: string; input: UpdateVendorInput }) =>
            vendorService.update(id, input),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['vendors'] });
        },
    });
}

export function useDeleteVendor() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => vendorService.delete(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['vendors'] });
        },
    });
}
