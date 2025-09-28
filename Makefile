.PHONY: install test backend frontend run clean

install:
	@echo "Installing Server dependencies..."
	cd server && go mod tidy
	@echo "Installing Client dependencies..."
	cd client && npm install

test:
	@echo "Running Tests..."
	cd server && go test ./internal/service && \
		go test ./internal/repo && \
		go test ./internal/http/handlers && \
		go test ./internal/http/middleware

backend:
	@echo "Starting Go backend server..."
	cd server && go run cmd/server/main.go

frontend:
	@echo "Starting frontend dev server..."
	cd client && npm run dev

run:
	@echo "Starting backend and frontend..."
	make -j2 backend frontend

clean:
	@echo "🛑 Stopping all background jobs..."
	killall -q go node || true
