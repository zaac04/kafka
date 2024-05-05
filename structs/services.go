package structs

type RegisterService struct {
	Name     string `json:"name"`
	EndPoint string `json:"endpoint,omitempty"`
}

type UpdateService struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Service struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	EndPoint string `json:"endpoint,omitempty"`
}

type SharedSVConfig struct {
	Ip        string `json:"ip"`
	KafkaPort string `json:"kafka_port"`
	RedisPort string `json:"redis_port"`
}
