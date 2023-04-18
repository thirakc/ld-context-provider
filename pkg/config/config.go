package config

type Config struct {
	ServerPort string `env:"PORT,required"`
	MongodbUri string `env:"MONGODB_URI,required"`
}
