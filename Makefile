.PHONY: migration
migration:
	go run -mod=mod ./cmd/migration/main.go $(NAME)