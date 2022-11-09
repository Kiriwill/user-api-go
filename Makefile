ensure:
	@./scripts/ensure.sh

dev: ensure
	@./scripts/dev.sh <PROJECT_NAME>
