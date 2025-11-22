package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Config 全局配置
type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Log      LogConfig      `yaml:"log"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Port    int    `yaml:"port"`
	Mode    string `yaml:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	Database        string        `yaml:"database"`
	Charset         string        `yaml:"charset"`
	ParseTime       bool          `yaml:"parseTime"`
	Loc             string        `yaml:"loc"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string        `yaml:"secret"`
	Issuer string        `yaml:"issuer"`
	Expire time.Duration `yaml:"expire"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level    string `yaml:"level"`
	Format   string `yaml:"format"`
	Output   string `yaml:"output"`
	FilePath string `yaml:"filePath"`
}

var GlobalConfig *Config

// LoadConfig 加载配置
func LoadConfig(configPath string) error {
	// 从YAML文件加载配置
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	// 从环境变量覆盖配置
	loadFromEnv(&config)

	GlobalConfig = &config
	return nil
}

// loadFromEnv 从环境变量加载配置
func loadFromEnv(config *Config) {
	// 应用配置
	if port := os.Getenv("APP_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.App.Port = p
		}
	}
	if mode := os.Getenv("APP_MODE"); mode != "" {
		config.App.Mode = mode
	}

	// 数据库配置
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Database.Host = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Database.Port = p
		}
	}
	if username := os.Getenv("DB_USERNAME"); username != "" {
		config.Database.Username = username
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		config.Database.Password = password
	}
	if database := os.Getenv("DB_DATABASE"); database != "" {
		config.Database.Database = database
	}
	if charset := os.Getenv("DB_CHARSET"); charset != "" {
		config.Database.Charset = charset
	}

	// JWT配置
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		config.JWT.Secret = secret
	}
	if issuer := os.Getenv("JWT_ISSUER"); issuer != "" {
		config.JWT.Issuer = issuer
	}
	if expire := os.Getenv("JWT_EXPIRE"); expire != "" {
		if d, err := time.ParseDuration(expire); err == nil {
			config.JWT.Expire = d
		}
	}

	// 日志配置
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Log.Level = level
	}
	if format := os.Getenv("LOG_FORMAT"); format != "" {
		config.Log.Format = format
	}
	if output := os.Getenv("LOG_OUTPUT"); output != "" {
		config.Log.Output = output
	}
	if filePath := os.Getenv("LOG_FILE_PATH"); filePath != "" {
		config.Log.FilePath = filePath
	}
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + strconv.Itoa(c.Port) + ")/" + c.Database + "?charset=" + c.Charset + "&parseTime=" + strconv.FormatBool(c.ParseTime) + "&loc=" + c.Loc
}

// Init 初始化配置
func Init() {
	if err := LoadConfig("config/config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Configuration loaded successfully")
}