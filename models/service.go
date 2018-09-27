package models

import graphql "github.com/graph-gophers/graphql-go"

// Service is a simple service
type Service struct {
	ID    graphql.ID
	Name  string
	Host  string
	Alive bool
}

// ServicesData is a map of services with a graphql.ID as key
var ServicesData = make(map[graphql.ID]*Service)

// Services is a slice of service
var Services = []*Service{
	{ID: "0", Name: "Gomonit", Host: "Scaleway", Alive: true},
	{ID: "1", Name: "Goploader", Host: "Scaleway", Alive: true},
	{ID: "2", Name: "Gopaste", Host: "Scaleway", Alive: false},
	{ID: "3", Name: "Markdownblog", Host: "OVH", Alive: false},
}

// ServiceResolver is the resolver associated to a Service
type ServiceResolver struct {
	S *Service
}

// ID returns the service ID
func (s *ServiceResolver) ID() graphql.ID {
	return s.S.ID
}

// Name returns the service name
func (s *ServiceResolver) Name() string {
	return s.S.Name
}

// Host returns the service host
func (s *ServiceResolver) Host() string {
	return s.S.Host
}

// Alive returns if the service is alive or not
func (s *ServiceResolver) Alive() bool {
	return s.S.Alive
}

func init() {
	for _, s := range Services {
		ServicesData[s.ID] = s
	}
}
