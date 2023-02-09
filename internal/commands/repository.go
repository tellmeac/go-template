package commands

import (
	"github.com/tellmeac/go-template/internal/core"
	"github.com/tellmeac/go-template/internal/storage/db"
)

type Repository struct {
	Config  *core.Config
	Storage *db.Storage
}
