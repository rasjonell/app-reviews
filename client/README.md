# Client

A minimal React + TypeScript client for browsing apps and viewing their reviews.

## Architecture Overview

- Frontend Stack: React, TypeScript, Vite, Tailwind
- Routing: TanStack Router with file‑based routes in `src/routes`
  - Routes: `/` (apps list) and `/reviews/:appId` (paginated reviews with filtering)
- Data Layer: TanStack React Query for fetching/caching, pagination, and cache invalidation
  - Provider: `src/integrations/tanstack-query/root-provider.tsx`
  - Hooks: `src/hooks/useApps.ts`, `src/hooks/useReviews.ts`, `src/hooks/useRefresh.ts`
- API Access: Lightweight `fetch` wrapper in `src/lib/api.ts`
  - Endpoints used: `/apps`, `/apps/new`, `/apps/:appId/reviews`, `/apps/:appId/poll`
- UI Components: Reusable pieces in `src/components` (`AppCard`, `AppList`, `AddAppModal`, `ReviewCard`, plus loading skeletons)
- Types: Shared domain types in `src/types.ts` (`App`, `Review`, etc.)
- Devtools: TanStack Devtools with Router + React Query panels

## Project Structure

- `src/routes` — File‑based routes
- `src/hooks` — Query hooks and mutations for apps/reviews
- `src/lib/api.ts` — API calls and response handling
- `src/components` — UI components and skeletons
- `src/integrations/tanstack-query` — QueryClient provider and devtools integration
- `src/styles.css` — Tailwind entry and base styles
- `vite.config.ts` — Vite + TanStack Router plugin

## Development

- Start dev server: `npm run dev` (default on `http://localhost:3000`)
- Build for production: `npm run build`

## Data Flow

UI components → hooks (`react-query`) → `src/lib/api.ts` → backend → cached state → UI updates.

