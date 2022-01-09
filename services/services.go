package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/muhriddinsalohiddin/online_store_api/config"
	pbCatalog "github.com/muhriddinsalohiddin/online_store_api/genproto/catalog_service"
	pbOrder "github.com/muhriddinsalohiddin/online_store_api/genproto/order_service"
)

type IServiceManager interface {
	CatalogService() pbCatalog.CatalogServiceClient
	OrderService() pbOrder.OrderServiceClient
}

type serviceManager struct {
	catalogService pbCatalog.CatalogServiceClient
	orderService   pbOrder.OrderServiceClient
}

func (s *serviceManager) CatalogService() pbCatalog.CatalogServiceClient {
	return s.catalogService
}

func (s *serviceManager) OrderService() pbOrder.OrderServiceClient {
	return s.orderService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CatalogServiceHost, conf.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.OrderServiceHost, conf.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		catalogService: pbCatalog.NewCatalogServiceClient(connCatalog),
		orderService:   pbOrder.NewOrderServiceClient(connOrder),
	}

	return serviceManager, nil
}
