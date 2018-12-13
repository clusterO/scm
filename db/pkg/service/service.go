package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

// DbService describes the service.
type DbService interface {
	// Add your methods here
	Connect(ctx context.Context, s string) (rs string, err error)
}

type basicDbService struct{}

func (b *basicDbService) Connect(ctx context.Context, s string) (rs string, err error) {

	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
			{
				hello
			}
		`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	return "SOS", err
}

// NewBasicDbService returns a naive, stateless implementation of DbService.
func NewBasicDbService() DbService {
	return &basicDbService{}
}

// New returns a DbService with all of the expected middleware wired in.
func New(middleware []Middleware) DbService {
	var svc DbService = NewBasicDbService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
