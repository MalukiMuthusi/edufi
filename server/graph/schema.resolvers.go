package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MalukiMuthusi/edufi/server/graph/generated"
	"github.com/MalukiMuthusi/edufi/server/graph/model"
)

func (r *learningObjectiveResolver) Module(ctx context.Context, obj *model.LearningObjective) (*model.Module, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateModule(ctx context.Context, input model.NewModule) (string, error) {
	id, err := r.S.CreateModule(ctx, input)
	return *id, err
}

func (r *mutationResolver) UpdateModule(ctx context.Context, input model.NewModule) (*model.Module, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteModule(ctx context.Context, input string) (*model.Module, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Modules(ctx context.Context) ([]*model.Module, error) {
	return r.S.ListModules(ctx)
}

func (r *queryResolver) Module(ctx context.Context, id string) (*model.Module, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchModule(ctx context.Context, text string) ([]*model.Module, error) {
	panic(fmt.Errorf("not implemented"))
}

// LearningObjective returns generated.LearningObjectiveResolver implementation.
func (r *Resolver) LearningObjective() generated.LearningObjectiveResolver {
	return &learningObjectiveResolver{r}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type learningObjectiveResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
