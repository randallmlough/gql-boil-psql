package main

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/randallmloug/gqlboil/graphql_models"
)

type Resolver struct{}

func (r *mutationResolver) CreatePet(ctx context.Context, input graphql_models.PetCreateInput) (*graphql_models.PetPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreatePets(ctx context.Context, input graphql_models.PetsCreateInput) (*graphql_models.PetsPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdatePet(ctx context.Context, id string, input graphql_models.PetUpdateInput) (*graphql_models.PetPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdatePets(ctx context.Context, filter *graphql_models.PetFilter, input graphql_models.PetUpdateInput) (*graphql_models.PetsUpdatePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeletePet(ctx context.Context, id string) (*graphql_models.PetDeletePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeletePets(ctx context.Context, filter *graphql_models.PetFilter) (*graphql_models.PetsDeletePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateToy(ctx context.Context, input graphql_models.ToyCreateInput) (*graphql_models.ToyPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateToys(ctx context.Context, input graphql_models.ToysCreateInput) (*graphql_models.ToysPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateToy(ctx context.Context, id string, input graphql_models.ToyUpdateInput) (*graphql_models.ToyPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateToys(ctx context.Context, filter *graphql_models.ToyFilter, input graphql_models.ToyUpdateInput) (*graphql_models.ToysUpdatePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteToy(ctx context.Context, id string) (*graphql_models.ToyDeletePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteToys(ctx context.Context, filter *graphql_models.ToyFilter) (*graphql_models.ToysDeletePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input graphql_models.UserCreateInput) (*graphql_models.UserPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUsers(ctx context.Context, input graphql_models.UsersCreateInput) (*graphql_models.UsersPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input graphql_models.UserUpdateInput) (*graphql_models.UserPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUsers(ctx context.Context, filter *graphql_models.UserFilter, input graphql_models.UserUpdateInput) (*graphql_models.UsersUpdatePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*graphql_models.UserDeletePayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUsers(ctx context.Context, filter *graphql_models.UserFilter) (*graphql_models.UsersDeletePayload, error) {
	panic("not implemented")
}

func (r *queryResolver) Pet(ctx context.Context, id string) (*graphql_models.Pet, error) {
	panic("not implemented")
}

func (r *queryResolver) Pets(ctx context.Context, filter *graphql_models.PetFilter) ([]*graphql_models.Pet, error) {
	panic("not implemented")
}

func (r *queryResolver) Toy(ctx context.Context, id string) (*graphql_models.Toy, error) {
	panic("not implemented")
}

func (r *queryResolver) Toys(ctx context.Context, filter *graphql_models.ToyFilter) ([]*graphql_models.Toy, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id string) (*graphql_models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, filter *graphql_models.UserFilter) ([]*graphql_models.User, error) {
	panic("not implemented")
}

// Mutation returns graphql_models.MutationResolver implementation.
func (r *Resolver) Mutation() graphql_models.MutationResolver { return &mutationResolver{r} }

// Query returns graphql_models.QueryResolver implementation.
func (r *Resolver) Query() graphql_models.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
