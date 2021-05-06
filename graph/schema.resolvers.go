package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"suburiQL/graph/generated"
	"suburiQL/graph/model"
)

var todos []*model.Todo

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := model.Todo{ID: "1", Text: "hoge", Done: false}
	todos = append(todos, &todo)
	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todos, nil
}

func (r *queryResolver) Template(ctx context.Context) (*model.Template, error) {
	temp := model.Template{Name: "hoge", Query: nil}
	temp2 := model.Template{Name: "fuga", Query: &temp}
	return &temp2, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
