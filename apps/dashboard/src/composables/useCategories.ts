import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query';
import { categoryService } from '../services/categoryService';
import type { CreateCategoryInput, UpdateCategoryInput } from '../types';
import { type Ref, unref } from 'vue';

export function useCategories() {
    return useQuery({
        queryKey: ['categories'],
        queryFn: categoryService.list,
    });
}

export function useCategory(id: Ref<string>) {
    return useQuery({
        queryKey: ['categories', id],
        queryFn: () => categoryService.getById(unref(id)),
        enabled: () => !!unref(id),
    });
}

export function useCreateCategory() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (input: CreateCategoryInput) => categoryService.create(input),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['categories'] });
        },
    });
}

export function useUpdateCategory() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, input }: { id: string; input: UpdateCategoryInput }) =>
            categoryService.update(id, input),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['categories'] });
        },
    });
}

export function useDeleteCategory() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: string) => categoryService.delete(id),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['categories'] });
        },
    });
}
