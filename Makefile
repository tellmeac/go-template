.PHONY: migration
migration:
	go run -mod=mod ./internal/store/sqlstore/ent/migrate/main.go $(NAME)