phony: run
run:
	@echo "Building..."
	@go build -o ./bin/diy-docker
	@echo "Running..."
	@/usr/bin/sandbox-exec -f ./sandbox.sb ./bin/diy-docker
