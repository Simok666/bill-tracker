import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query';
import { billService } from '../services/billService';
import type { BillFilters, CreateBillInput, UpdateBillInput } from '../types';
import { type Ref, unref } from 'vue';

export function useBills(filters: Ref<BillFilters>) {
    return useQuery({
        queryKey: ['bills', filters],
        queryFn: () => billService.list(unref(filters)),
    });
}

export function useBill(id: Ref<string | undefined>) {
    return useQuery({
        queryKey: ['bills', id],
        queryFn: () => billService.getById(unref(id)!),
        enabled: () => !!unref(id),
    });
}

export function useBillActivities(id: Ref<string | undefined>) {
    return useQuery({
        queryKey: ['bills', id, 'activities'],
        queryFn: () => billService.getActivities(unref(id)!),
        enabled: () => !!unref(id),
    });
}

export function useCreateBill() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (input: CreateBillInput) => billService.create(input),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['bills'] });
            queryClient.invalidateQueries({ queryKey: ['dashboard'] });
        },
    });
}

export function useUpdateBill() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, input }: { id: string; input: UpdateBillInput }) =>
            billService.update(id, input),
        onSuccess: (data) => {
            queryClient.invalidateQueries({ queryKey: ['bills'] });
            queryClient.invalidateQueries({ queryKey: ['bills', data.id] });
            queryClient.invalidateQueries({ queryKey: ['dashboard'] });
        },
    });
}

export function useDeleteBill() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => billService.delete(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['bills'] });
            queryClient.invalidateQueries({ queryKey: ['dashboard'] });
        },
    });
}

export function usePayBill() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, date }: { id: string; date?: Date }) =>
            billService.pay(id, date),
        onSuccess: (data) => {
            queryClient.invalidateQueries({ queryKey: ['bills'] });
            queryClient.invalidateQueries({ queryKey: ['bills', data.id] });
            queryClient.invalidateQueries({ queryKey: ['dashboard'] });
        },
    });
}
