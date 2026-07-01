package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var (
	Setting = &configs{}
	envPath = "config/config.env"
)

type configs struct {
	DBUser    dBUserConfig    `envPrefix:"DB_USER_"`
	DBProblem dBProblemConfig `envPrefix:"DB_PROBLEM_"`
	Server    serverConfig    `envPrefix:"SERVER_"`
	Log       logConfig       `envPrefix:"LOG_"`
	Redis     redisConfig     `envPrefix:"REDIS_"`
}

type dBUserConfig struct {
	Type          string `env:"TYPE,required"`
	Host          string `env:"HOST,required"`
	Port          int    `env:"PORT,required"`
	User          string `env:"USER,required"`
	Password      string `env:"PASSWORD,required"`
	DBName        string `env:"DBNAME,required"`
	MainTableName string `env:"MAIN_TABLENAME,required"`
}

type dBProblemConfig struct {
	Type              string `env:"TYPE,required"`
	Host              string `env:"HOST,required"`
	Port              int    `env:"PORT,required"`
	User              string `env:"USER,required"`
	Password          string `env:"PASSWORD,required"`
	DBName            string `env:"DBNAME,required"`
	MainTableName     string `env:"MAIN_TABLENAME,required"`
	LimitTableName    string `env:"LIMIT_TABLENAME,required"`
	SampleTableName   string `env:"SAMPLE_TABLENAME,required"`
	TestCaseTableName string `env:"TESTCASE_TABLENAME,required"`
}

type serverConfig struct {
	Port         int           `env:"PORT,required"`
	RuntimePath  string        `env:"RUNTIME_PATH,required"`
	RunMode      string        `env:"RUNMODE,required"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT,required"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT,required"`
}

type logConfig struct {
	SavePath   string `env:"SAVEPATH,required"`
	SaveName   string `env:"SAVENAME,required"`
	FileExt    string `env:"FILEEXT,required"`
	TimeFormat string `env:"TIMEFORMAT,required"`
}

type redisConfig struct {
	Host           string        `env:"HOST,required"`
	Port           int           `env:"PORT,required"`
	Password       string        `env:"PASSWORD,required"`
	DB             int           `env:"DB,required"`
	DefaultTimeout time.Duration `env:"DEFAULTTIMEOUT,required"`
	PoolSize       int           `env:"POOLSIZE,required"`
	MinIdleConns   int           `env:"MINIDLECONNS,required"`
	MaxRetries     int           `env:"MAXRETRIES,required"`
	DialTimeout    time.Duration `env:"DIALTIMEOUT,required"`
	ReadTimeout    time.Duration `env:"READTIMEOUT,required"`
	WriteTimeout   time.Duration `env:"WRITETIMEOUT,required"`
}

// Setup Initial configuration
func Setup() {
	err := godotenv.Load(envPath)
	if err != nil {
		panic(fmt.Errorf("load env file failed: %w", err))
	}

	err = env.Parse(Setting)
	if err != nil {
		panic(fmt.Errorf("parse env config failed: %w", err))
	}
}
