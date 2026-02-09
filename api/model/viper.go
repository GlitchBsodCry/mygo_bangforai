package model



type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Logger   Logger   `mapstructure:"logger"`
	JWT      JWT      `mapstructure:"jwt"`
}

var AppConfig Config

type Server struct {
	Port      string `mapstructure:"port"`
	Name      string `mapstructure:"name"`
	Env       string `mapstructure:"env"`
	Timeout   int    `mapstructure:"timeout"`
	AllowOrigins string `mapstructure:"allow_origins"`
}


type Database struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int `mapstructure:"conn_max_lifetime"`
}

// Logger 日志配置
type Logger struct {
	Level  string `mapstructure:"level"`  // 日志级别: debug, info, warn, error
	Format string `mapstructure:"format"` // 日志格式: json, console
	Output string `mapstructure:"output"` // 输出方式: stdout, file
	File   string `mapstructure:"file"`   // 日志文件路径
	ErrorFile   string `mapstructure:"error_file"`   // 错误日志文件路径
}

// JWT 配置
type JWT struct {
	Secret     string `mapstructure:"secret"`     // JWT密钥
	ExpiresIn  int    `mapstructure:"expires_in"` // 过期时间（小时）
	Issuer     string `mapstructure:"issuer"`     // 签发者
	Subject    string `mapstructure:"subject"`    // 主题
}