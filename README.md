# App Store Review Aggregator

This project consists of a **Go backend** that fetches and stores iOS App Store reviews, and a **React frontend** with TanStack router and query
that displays those reviews with date filtering and pagination.

---

![Recording](./recording.mp4)

## How to Run

### Clone the Repository

```bash
git clone git@github.com:rasjonell/app-reviews.git
cd app-reviews
```

### Install Dependencies, Run Tests & Run Both Services

```bash
make install
make test
make run
```

> This will start:
> - Backend API on [http://localhost:8080](http://localhost:8080)
> - Frontend on [http://localhost:3000](http://localhost:3000)

---

## Architecture

You can find server and client architecture overview in each directory's README:

- Server [README.md](./server/README.md)
- Client [README.md](./client/README.md)

---

## Features

### Backend

- Fetches iOS App Store reviews using RSS feeds.
- Stores reviews and app metadata in SQLite.
- Supports multiple apps.
- Polls app reviews every hour.
- Manual trigger to re-fetch reviews
- REST API for:
  - Adding apps
  - Fetching apps
  - Fetching reviews (filtered by date and paginated)

### Frontend

- View all tracked apps.
- View recent reviews per app.
- Filter reviews by:
  - Date (using `since` parameter)
- Paginate through reviews.
- Add new apps via UI.
