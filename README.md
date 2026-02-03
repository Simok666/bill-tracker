# Bill Reminder (Bill Tracker)

A comprehensive monorepo application for tracking and managing corporate bills, recurrence, and reminders.

## 🚀 Features

- **Dashboard**: High-level overview of bill statuses and summaries.
- **Bill Management**: Add, update, and track bills with specific categories and vendors.
- **Recurring Bills**: Support for recurring billing cycles.
- **Category & Vendor Management**: Organize bills by custom categories and service providers.
- **Authentication**: Secure user login and profile management.
- **Responsive Design**: Modern UI built with Vue 3 and Tailwind CSS.

## 🛠️ Tech Stack

### Backend
- **Language**: Go (Golang)
- **Framework**: [Gin](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Middleware**: CORS handling, JWT (planned/integrated)

### Frontend
- **Framework**: [Vue.js 3](https://vuejs.org/) (Composition API)
- **Build Tool**: [Vite](https://vitejs.dev/)
- **Styling**: [Tailwind CSS](https://tailwindcss.com/)
- **State Management/Data Fetching**: [TanStack Vue Query](https://tanstack.com/query/latest/docs/framework/vue/overview)
- **Routing**: Vue Router

## 📂 Project Structure

```text
bill-reminder/
├── apps/
│   ├── backend/      # Go API implementation
│   │   ├── cmd/      # Entry points (main.go)
│   │   ├── internal/ # Core logic (models, services, routes)
│   │   └── pkg/      # Shared utilities
│   └── dashboard/    # Vue.js frontend
│       ├── src/
│       │   ├── views/    # Page components
│       │   ├── services/ # API client logic
│       │   └── components/
├── package.json      # Monorepo configuration
└── README.md         # You are here
```

## 🏁 Getting Started

### Prerequisites
- **Go**: v1.25.6 or higher
- **Node.js**: v18.x or higher
- **npm**: v9.x or higher
- **PostgreSQL**: Running instance for the backend

### Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd bill-reminder
   ```

2. **Backend Setup**:
   ```bash
   cd apps/backend
   # Copy env example and configure your database
   cp .env.example .env
   # Install dependencies
   go mod download
   ```

3. **Frontend Setup**:
   ```bash
   cd apps/dashboard
   # Install dependencies
   npm install
   ```

### Running Locally

- **Run Backend**:
  ```bash
  cd apps/backend
  go run cmd/server/main.go
  ```

- **Run Frontend**:
  ```bash
  # From the root directory (using monorepo scripts)
  npm run dev:dashboard
  
  # OR from the dashboard directory
  cd apps/dashboard
  npm run dev
  ```

## 📄 License
[MIT](LICENSE)
puki
