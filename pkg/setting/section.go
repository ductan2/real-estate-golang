package setting

type Config struct {
	Database Database `mapstructure:"database"`
	Server   Server   `mapstructure:"server"`
	Redis    Redis    `mapstructure:"redis"`
	SMTP     SMTP     `mapstructure:"smtp"`
	JWT 	 JWT 	  `mapstructure:"jwt"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	MaxOpenConnections int `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int `mapstructure:"maxIdleConnections"`
	MaxLifetimeConnection int `mapstructure:"maxLifetimeConnection"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database   int `mapstructure:"database"`
}

type SMTP struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type JWT struct {
	TokenSecret string `mapstructure:"tokenSecret"`
	TokenExpirationTime string `mapstructure:"tokenExpirationTime"`
	TokenHoursToExpire int `mapstructure:"tokenHoursToExpire"`
}
