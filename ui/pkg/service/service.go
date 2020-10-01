package service

import (
	"context"
	"log"
	"scm/auth/pkg/grpc/pb"

	sdetcd "github.com/go-kit/kit/sd/etcd"
	"google.golang.org/grpc"
)

// UiService describes the service.
type UiService interface {
	// Add your methods here
	Display(ctx context.Context, s string) (rs string, err error)
}

type basicUiService struct {
	authServiceClient pb.AuthClient
}

func (b *basicUiService) Display(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Display
	reply, err := b.authServiceClient.Auth(context.Background(), &pb.AuthRequest{
		Email:   s,
		Content: "Hi! You are now in service!",
	})

	if reply != nil {
		log.Printf("Email Id : %s", reply.Id)
	}

	return rs, err
}

// NewBasicUiService returns a naive, stateless implementation of UiService.
func NewBasicUiService() UiService {
	var (
		etcdServer = "http://etcd:2379"
		prefix     = "/services/auth/"
	)

	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		log.Printf("unable to connect to etcd %s", err.Error())
		return new(basicUiService)
	}

	entries, err := client.GetEntries(prefix)
	if err != nil || len(entries) == 0 {
		log.Printf("unable to get entries %v", err)
		return new(basicUiService)
	}

	conn, err := grpc.Dial(entries[0], grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to auth %s", err.Error())
		return new(basicUiService)
	}

	return &basicUiService{
		authServiceClient: pb.NewAuthClient(conn),
	}
}

// New returns a UiService with all of the expected middleware wired in.
func New(middleware []Middleware) UiService {
	var svc UiService = NewBasicUiService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
