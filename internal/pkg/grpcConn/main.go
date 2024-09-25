package grpc

import (
	pbp "api_gateway/genproto/carpet_service"
	pbu "api_gateway/genproto/user_service"
	//"api_gateway/genproto/pbp"
	"api_gateway/internal/configs"
	"api_gateway/internal/pkg/logger"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	OrderService() pbp.OrderServiceClient
	AddressService() pbp.AddressesClient
	CompanyService() pbp.CompanyServiceClient
	ServiceService() pbp.ServiceServiceClient
}
type grpcClients struct {
	userService    pbu.UserServiceClient
	orderService   pbp.OrderServiceClient
	addressService pbp.AddressesClient
	companyService pbp.CompanyServiceClient
	serviceService pbp.ServiceServiceClient
}

func (g *grpcClients) UserService() pbu.UserServiceClient {
	return g.userService
}

func (g *grpcClients) OrderService() pbp.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) AddressService() pbp.AddressesClient {
	return g.addressService
}

func (g *grpcClients) CompanyService() pbp.CompanyServiceClient {
	return g.companyService
}

func (g *grpcClients) ServiceService() pbp.ServiceServiceClient {
	return g.serviceService
}

func NewGrpcClients(cnf configs.Config, log logger.ILogger) (IServiceManager, error) {
	connUserService, err :=
		grpc.NewClient(
			cnf.AuthServiceGrpcHost+cnf.AuthServiceGrpcPort,
			grpc.WithInsecure())

	if err != nil {

		log.Error("this error is  connUserService with dialing ", logger.Error(err))
		return nil, err
	}
	connCarpetService, err := grpc.NewClient(
		cnf.CarpetServiceGrpcHost+cnf.CarpetServiceGrpcPort,
		grpc.WithInsecure())
	if err != nil {
		log.Error("this error is  connTaskService with dialing ", logger.Error(err))
		return nil, err
	}
	return &grpcClients{
		userService:    pbu.NewUserServiceClient(connUserService),
		orderService:   pbp.NewOrderServiceClient(connCarpetService),
		serviceService: pbp.NewServiceServiceClient(connCarpetService),
		companyService: pbp.NewCompanyServiceClient(connCarpetService),
		addressService: pbp.NewAddressesClient(connCarpetService),
	}, nil

}
