package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.schibsted.io/Schibsted-fr-lab/graphql-poc/models"
)

// MainResolver is the top-level resolver for the GraphQL API
type MainResolver struct{}

// Service takes an ID as argument and returns the associated service if found
func (r *MainResolver) Service(args struct{ ID graphql.ID }) *models.ServiceResolver {
	if s, ok := models.ServicesData[args.ID]; ok {
		return &models.ServiceResolver{S: s}
	}
	return nil
}

// Status will filter the alive services if alive is given otherwise it will
// return all the services
func (r *MainResolver) Status(args struct{ Alive *bool }) *[]*models.ServiceResolver {
	out := []*models.ServiceResolver{}
	if args.Alive != nil {
		for _, s := range models.Services {
			if s.Alive == *args.Alive {
				out = append(out, &models.ServiceResolver{S: s})
			}
		}
	} else {
		for _, s := range models.Services {
			out = append(out, &models.ServiceResolver{S: s})
		}
	}
	return &out
}

// Alive returns all the alive services
func (r *MainResolver) Alive() *[]*models.ServiceResolver {
	b := true
	return r.Status(struct{ Alive *bool }{Alive: &b})
}

// Dead returns all the dead services
func (r *MainResolver) Dead() *[]*models.ServiceResolver {
	var b bool
	return r.Status(struct{ Alive *bool }{Alive: &b})
}
