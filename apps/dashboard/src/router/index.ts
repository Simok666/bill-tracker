import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../views/Dashboard.vue';
import Bills from '../views/Bills.vue';
import MainLayout from '../layouts/MainLayout.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            component: MainLayout,
            meta: { requiresAuth: true },
            children: [
                {
                    path: '',
                    name: 'dashboard',
                    component: Dashboard
                },
                {
                    path: 'bills',
                    name: 'bills',
                    component: Bills
                },
                {
                    path: 'bills/add',
                    name: 'add-bill',
                    component: () => import('../views/AddBill.vue')
                },
                {
                    path: 'bills/:id',
                    name: 'bill-detail',
                    component: () => import('../views/BillDetail.vue')
                },
                {
                    path: 'vendors',
                    name: 'vendors',
                    component: () => import('../views/Vendors.vue')
                },
                {
                    path: 'vendors/add',
                    name: 'add-vendor',
                    component: () => import('../views/AddVendor.vue')
                },
                {
                    path: 'categories',
                    name: 'categories',
                    component: () => import('../views/Categories.vue')
                },
                {
                    path: 'categories/add',
                    name: 'add-category',
                    component: () => import('../views/AddCategory.vue')
                },
                {
                    path: 'settings',
                    name: 'settings',
                    component: () => import('../views/Settings.vue')
                }
            ]
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/Login.vue'),
            meta: { guestOnly: true }
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('../views/Register.vue'),
            meta: { guestOnly: true }
        }
    ]
});

router.beforeEach((to, _from, next) => {
    const isAuthenticated = !!localStorage.getItem('token');

    if (to.meta.requiresAuth && !isAuthenticated) {
        next('/login');
    } else if (to.meta.guestOnly && isAuthenticated) {
        next('/');
    } else {
        next();
    }
});

export default router;
