package commands

import (
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/internal/storage/db"
)

// Repository contains dependencies that should be used by core.
type Repository struct {
	Config  *core.Config
	Storage *db.Storage
}
