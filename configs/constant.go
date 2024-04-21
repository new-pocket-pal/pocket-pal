package configs

type Config struct {
	Port           string         `json:"port" mapstructure:"port"`
	ProjectRootDir string         `json:"project_root_dir" mapstructure:"project_root_dir"`
	Redis          RedisConfig    `json:"redis" mapstructure:"redis"`
	Postgres       PostgresConfig `json:"postgres" mapstructure:"postgres"`
	ENV            string         `json:"env" mapstructure:"env"`
	GpayPath       string         `json:"gpayPath" mapstructure:"gpayPath"`
	JWT            JWTConfig      `json:"jwt" mapstructure:"jwt"`
}

type RedisConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}

type JWTConfig struct {
	SecretKey  string `json:"secret_key" mapstructure:"secret_key"`
	ExpireTime int64  `json:"exp_time" mapstructure:"exp_time"`
}

type PostgresConfig struct {
	Host      string `json:"host" mapstructure:"host"`
	Port      int    `json:"port" mapstructure:"port"`
	DbName    string `json:"dbname" mapstructure:"dbname"`
	User      string `json:"user" mapstructure:"user"`
	Pass      string `json:"password" mapstructure:"password"`
	SSLMode   string `json:"sslmode" mapstructure:"sslmode"`
	Prefix    string `json:"prefix" mapstructure:"prefix"`
	DebugMode bool   `json:"debug_mode" mapstructure:"debug_mode"`
}
