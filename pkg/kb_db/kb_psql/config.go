package kb_psql

import (
	"fmt"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/spf13/viper"
		_ "github.com/jackc/pgx/v5/stdlib"

)


// DBConfig struct
type DBConfig struct {
	Host                  string `mapstructure:"host"`
	Username              string `mapstructure:"username"`
	Port                  string `mapstructure:"port"`
	Database              string `mapstructure:"database"`
	Password              string `mapstructure:"password"`
	SSLEnabled            string `mapstructure:"ssl_enabled"`
	MaxConnection         int32  `mapstructure:"max_connection"`
	MaxIdleConnection     int32  `mapstructure:"max_idle_connection"`
	MinConnection         int32  `mapstructure:"min_connection"`
	MaxConnectionLifetime int32  `mapstructure:"max_connection_lifetime"`
	MaxConnectionIdleTime int32  `mapstructure:"max_connection_idletime"`
	Options               string `mapstructure:"options"`
	Offset                int8   `mapstructure:"offset"`
	Limit                 int8   `mapstructure:"limit"`
}

type DBConnection *sql.DB

// NewDBConnection connect a database
func NewDBConnection(ctx context.Context, pathConfigFile, nameFileConfig, nameFileExtention string) (*sql.DB, error) {

	cfg, err := parse(pathConfigFile, nameFileConfig, nameFileExtention)
	if err != nil {
		return nil, err
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLEnabled)
	pool, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}

	// create a context for Ping
	ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	
	if err := pool.PingContext(ctxPing); err != nil {
		return nil, err
	}

	return pool, nil
}

func parse(pathConfigFile, nameFileConfig, nameFileExtention string) (*DBConfig, error) {

	viper.AddConfigPath(pathConfigFile)
	viper.SetConfigName(nameFileConfig)
	viper.SetConfigType(nameFileExtention)
	viper.AutomaticEnv()

	field, ok := viper.Get("postgresql").(map[string]interface{})
	if !ok {
		return  nil, errors.New("error, not found key postgresql in config file")
	}

	sqlConfig := DBConfig{
		Host:			 "localhost",                 
		Username:				 "postgres",             
		Port: 							 "5432",                 
		Database: 			 "service",             
		Password: 			 "postgres",             
		SSLEnabled: 		 "false",           
		MaxConnection: 	 8,        
		MaxIdleConnection: 5,   
		MinConnection: 	 1,        
		MaxConnectionLifetime: 10,
		MaxConnectionIdleTime: 5,
		Options: "",              
		Offset: 0,               
		Limit: 20,               
	}

	fInt32, ok := field["max_connection"].(int32)
	if ok {
		sqlConfig.MaxConnection = fInt32
	}

	fInt32, ok = field["min_connection"].(int32)
	if ok {
		sqlConfig.MinConnection = fInt32
	}

	fInt32, ok = field["max_connection_lifetime"].(int32)
	if ok {
		sqlConfig.MaxConnectionLifetime = fInt32
	}

	fInt32, ok = field["max_connection_idletime"].(int32)
	if ok {
		sqlConfig.MaxConnectionIdleTime = fInt32
	}

	fInt32, ok = field["max_idle_connection"].(int32)
	if ok {
		sqlConfig.MaxConnectionIdleTime = fInt32
	}

	fInt32, ok = field["max"].(int32)
	if ok {
		sqlConfig.MaxConnectionIdleTime = fInt32
	}

	fStr, ok := field["port"].(string)
	if ok {
		sqlConfig.Port = fStr
	}

	fStr, ok = field["ssl_enabled"].(string)
	if ok {
		sqlConfig.SSLEnabled = fStr
	}

	fStr, ok = field["host"].(string)
	if ok {
		sqlConfig.Host = fStr
	}

	fStr, ok = field["options"].(string)
	if ok {
		sqlConfig.Options = fStr
	}

	fInt8, ok := field["limit"].(int8)
	if ok {
		sqlConfig.Limit = fInt8
	}

	fInt8, ok = field["offset"].(int8)
	if ok {
		sqlConfig.Offset = fInt8
	}

	fStr, ok = field["database"].(string)
	if ok {
		sqlConfig.Database = fStr
	} else {
		return nil, errors.New("database can not is empty")
	}

	fStr, ok = field["username"].(string)
	if ok {
		sqlConfig.Username = fStr
	} else {
		return nil, errors.New("username can not is empty")
	}

	fStr, ok = field["password"].(string)
	if ok {
		sqlConfig.Password = fStr
	}

	// sqlConfig.SSLEnabled = field["ssl_enabled"].(bool)

	return &sqlConfig, nil
}