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

// ServicesSlice is a slice of pointer to service
type ServicesSlice []*Service

// Services is a slice of service
var Services = ServicesSlice{
	{ID: "0", Name: "Gomonit", Host: "Scaleway", Alive: true},
	{ID: "1", Name: "Goploader", Host: "Scaleway", Alive: true},
	{ID: "2", Name: "Gopaste", Host: "Scaleway", Alive: false},
	{ID: "3", Name: "Markdownblog", Host: "OVH", Alive: false},
}

// AddService adds a service
func AddService(s *Service) {
	Services = append(Services, s)
	ServicesData[s.ID] = s
}

// DeleteService will delete a service
func DeleteService(s *Service) {
	delete(ServicesData, s.ID)
	for k, v := range Services {
		if v.ID == s.ID {
			Services = append(Services[:k], Services[k+1:]...)
			break
		}
	}
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
