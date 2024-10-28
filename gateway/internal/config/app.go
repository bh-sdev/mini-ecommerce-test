package config

type App struct {
	AppPort int16
	// JWTSecret string
	ProductServiceConfig
	OrderServiceConfig
}

type ProductServiceConfig struct {
	ProductServiceUrl string `env:"PRODUCT_SERVICE_ADDRESS" envDefault:"localhost:50051"`
}

type OrderServiceConfig struct {
	OrderServiceUrl string `env:"ORDER_SERVICE_ADDRESS" envDefault:"localhost:50052"`
}
