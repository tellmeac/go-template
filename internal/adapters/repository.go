package adapters

import (
	"context"
	"fmt"
	"github.com/tellmeac/go-template/internal/adapters/ent"
	"github.com/tellmeac/go-template/internal/adapters/ent/templates"
	"github.com/tellmeac/go-template/internal/pkg/errors"
)

func NewDAO(client *ent.Client) DAO {
	return DAO{client: client}
}

type DAO struct {
	client *ent.Client
}

func (dao DAO) Get(ctx context.Context, target string) (string, error) {
	template, err := dao.client.Templates.Query().Where(templates.Name(target)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return "", fmt.Errorf("user template not exists: %w", errors.ErrNotFound)
	case err != nil:
		return "", err
	}

	return template.Template, nil
}

func (dao DAO) Upsert(ctx context.Context, target string, desired string) error {
	return dao.client.Templates.Create().
		SetName(target).
		SetTemplate(desired).
		OnConflict().UpdateNewValues().Exec(ctx)
}
