# Server

A small Go HTTP server that fetches, stores, and serves App Store reviews.
I expose a minimal JSON API and runs a background polling job to keep reviews fresh.

## Architecture Overview
- Entry Point: `cmd/server/main.go` wires dependencies, starts HTTP server on `:8080`, seeds an initial app, triggers an initial fetch, and starts the polling job with graceful shutdown.
- HTTP Layer: `internal/http/` contains request DTOs, handlers, and middlewares. Handlers validate/parse input and orchestrate services; DTOs shape request/response payloads.
- Services: `internal/service/` holds application logic.
  - `AppService` manages apps(via App Store lookup).
  - `ReviewService` fetches, transforms, and persists reviews; updates last-polled timestamps.
  - `AppStoreService` integrates with Apple iTunes APIs (app lookup + customer reviews feeds).
  - `PollingService` periodically refreshes reviews for enabled apps.
- Persistence: `internal/repo/` provides repositories over SQLite.
  - `AppRepo` stores app metadata and last poll time.
  - `ReviewRepo` stores deduplicated reviews.
- Domain Models: `internal/domain/` defines `App` and `Review` types.
- DB: `internal/db/` opens/initializes SQLite at `./data/reviews.db`.

## Data Flow
- Add App: HTTP → `AppService` resolves name via App Store → `AppRepo` upserts → handler kicks off an immediate async fetch.
- Poll Reviews: `ReviewService` calls App Store feed → maps to domain → `ReviewRepo` inserts → `AppRepo` updates `last_polled`.
- Read Reviews: HTTP parses `appId`, pagination, optional `since` → `ReviewService` queries recent reviews → JSON DTO response.

## HTTP API (JSON)
- `GET /apps` — list apps.
- `POST /apps/new` — body `{ "appId": "<itunes_id>" }` to register an app and trigger initial fetch.
- `GET /apps/{appId}/reviews` — optional query `page`, `limit`, `since` to paginate recent reviews.
- `GET /apps/{appId}/poll` — trigger an on-demand poll for a single app.

CORS: wildcard `*` with `GET, POST, OPTIONS` via middleware.

## Background Polling
- Hourly job polls enabled apps for new reviews.
- Stops gracefully on SIGINT.

## Storage
- SQLite database at `./data/reviews.db`
- Tables: `apps` (id, app_id, name, enabled, last_polled) and `reviews` (id, app_id, author, title, content, rating, timestamp).

## Running Locally
- Start: `go run ./cmd/server`
- Default port: `:8080`
- Note: Startup seeds one app (`595068606`) and fetches its reviews once.

## Dependencies
- `github.com/mattn/go-sqlite3` for SQLite access.
- External: Apple iTunes Lookup and Customer Reviews JSON feeds.
