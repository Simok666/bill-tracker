export type UserRole = 'admin' | 'member';

export interface User {
    id: string;
    company_id: string;
    name: string;
    email: string;
    role: UserRole;
    avatar_url?: string;
    email_verified: boolean;
    created_at: string;
    updated_at: string;
}

export type BillStatus = 'draft' | 'unpaid' | 'paid' | 'overdue';
export type RecurringFrequency = 'weekly' | 'monthly' | 'yearly';
export type ActivityAction = 'created' | 'updated' | 'status_changed' | 'payment_reminder_sent' | 'attachment_added' | 'attachment_removed' | 'deleted';

export interface Vendor {
    id: string;
    company_id: string;
    name: string;
    logo_url?: string;
    contact_email?: string;
    website?: string;
    contact_info?: string;
    address?: string;
    location?: string;
    created_at: string;
    updated_at: string;
}

export interface Category {
    id: string;
    company_id: string;
    name: string;
    description?: string;
    icon?: string;
    color?: string;
    created_at: string;
    updated_at: string;
}

export interface BillAttachment {
    id: string;
    bill_id: string;
    file_name: string;
    file_url: string;
    file_type: string;
    uploaded_by: string;
    created_at: string;
    uploader?: User;
}

export interface BillActivity {
    id: string;
    bill_id: string;
    user_id?: string;
    action: ActivityAction;
    details?: string;
    created_at: string;
    user?: User;
}

export interface Bill {
    id: string;
    company_id: string;
    user_id: string;
    vendor_id?: string;
    category_id?: string;
    title: string;
    invoice_number?: string;
    amount: string; // Decimal as string from API
    currency: string;
    due_date: string;
    paid_date?: string;
    status: BillStatus;
    is_recurring: boolean;
    recurring_frequency?: RecurringFrequency;
    recurring_day?: number;
    payment_method?: string;
    notes?: string;
    created_at: string;
    updated_at: string;

    // Relations
    user?: User;
    vendor?: Vendor;
    category?: Category;
    attachments?: BillAttachment[];
    activities?: BillActivity[];
}

// Auth DTOs
export interface LoginInput {
    email: string;
    password: string;
}

export interface RegisterInput {
    name: string;
    email: string;
    password: string;
    company_name: string;
}

export interface AuthResponse {
    token: string;
    user: User;
}

// Bill DTOs
export interface CreateBillInput {
    title: string;
    vendor_id?: string;
    category_id?: string;
    invoice_number?: string;
    amount: number | string;
    currency?: string;
    due_date: string;
    is_recurring?: boolean;
    recurring_frequency?: RecurringFrequency;
    recurring_day?: number;
    payment_method?: string;
    notes?: string;
    status?: BillStatus;
}

export interface UpdateBillInput {
    title?: string;
    vendor_id?: string;
    category_id?: string;
    invoice_number?: string;
    amount?: number | string;
    currency?: string;
    due_date?: string;
    is_recurring?: boolean;
    recurring_frequency?: RecurringFrequency;
    recurring_day?: number;
    payment_method?: string;
    notes?: string;
    status?: BillStatus;
}

export interface BillFilters {
    status?: string;
    search?: string;
    page: number;
    page_size: number;
}

// Vendor DTOs
export interface CreateVendorInput {
    name: string;
    logo_url?: string;
    contact_email?: string;
    website?: string;
    contact_info?: string;
    address?: string;
    location?: string;
}

export interface UpdateVendorInput {
    name?: string;
    logo_url?: string;
    contact_email?: string;
    website?: string;
    contact_info?: string;
    address?: string;
    location?: string;
}

// Category DTOs
export interface CreateCategoryInput {
    name: string;
    description?: string;
    icon?: string;
    color?: string;
}

export interface UpdateCategoryInput {
    name?: string;
    description?: string;
    icon?: string;
    color?: string;
}

// User DTOs
export interface UpdateProfileInput {
    name?: string;
    avatar_url?: string;
}

export interface ChangePasswordInput {
    current_password: string;
    new_password: string;
}

// Dashboard DTOs
export interface DashboardStats {
    total_expense: string;
    paid_this_month: string;
    unpaid_amount: string;
    overdue_bills_count: number;
    expense_change_percent: number;
}

export interface MonthlyExpense {
    month: string;
    amount: string;
}

export interface CategoryExpense {
    category_id: string;
    category_name: string;
    amount: string;
    percentage: number;
}

// Generic API Responses
export interface PaginationMeta {
    current_page: number;
    page_size: number;
    total_items: number;
    total_pages: number;
}

export interface PaginatedResponse<T> {
    data: T[];
    meta: PaginationMeta;
}

export interface ApiResponse<T> {
    data: T;
    message?: string;
}
