setup-migrations:
	@if ! command -v migrate &> /dev/null || [ ! -f "./migrate" ]; then \
		echo "Installing golang-migrate..."; \
		go get github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
		go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
		CGO_ENABLED=1 go build -tags "sqlite3" -o migrate github.com/golang-migrate/migrate/v4/cmd/migrate; \
		go mod tidy; \
	else \
		echo "Migrate is already installed and built."; \
	fi

migrate-up:
	./migrate -path ./migrations -database "sqlite3://./data/taskmanager.db" up

migrate-down:
	./migrate -path ./migrations -database "sqlite3://./data/taskmanager.db" down

build:
	@mkdir -p bin
	@if [ ! -f "./data/taskmanager.db" ]; then \
		echo "Database file missing. Running migrations..."; \
		make migrate-up; \
	else \
		echo "Database exists. Skipping migrations."; \
	fi
	go build -o bin/task-manager main.go

run: build
	./bin/task-manager

clean:
	rm -f bin/task-manager

mrproper: clean
	@echo "Warning: This will delete all database files!"
	@read -p "Are you sure? (y/N) " confirm && if [ "$$confirm" = "y" ]; then rm -f data/taskmanager.db; fi
