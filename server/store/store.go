package store

import (
	"context"

	"github.com/MalukiMuthusi/edufi/server/graph/model"
)

type Store interface {
	CreateModule(ctx context.Context, input model.NewModule) (*string, error)
	// UpdateModule(ctx context.Context, input model.NewModule) (*model.Module, error)
	// DeleteModule(ctx context.Context, input string) (*model.Module, error)
	ListModules(ctx context.Context) ([]*model.Module, error)
	GetModule(ctx context.Context, id string) (*model.Module, error)
}

func init() {}
