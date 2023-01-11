package config

type ConfigDatabase struct {
	AppName  string `env:"APP_NAME" env-default:"ELECTRONICS"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"MY_APP_PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLvl   string `env:"LOG_LEVEL" env-default:"ERROR"`
	MongoURI string `env:"MONGO_URI" env-default:"mongodb://27017:27017"`
}

var configDB ConfigDatabase

func GetConfig() *ConfigDatabase {
	return &configDB
}
