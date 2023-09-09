package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"go-admin/config"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	defaultConnMaxLifeTime int = 5  // default max 5 minutes lifetime
	defaultMaxOpenConns    int = 10 // default max 10 open connections
	defaultMaxIdleConns    int = 10 // default max 10 idle connections
)

type HandlerDatabase struct {
	Master *gorm.DB
}

func NewDatabaseClient(conf *config.Config) (HandlerDatabase, error) {
	dbHost := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s fallback_application_name=go-soskomlap-api-master TimeZone=Asia/Jakarta",
		conf.Postgres.Master.Host,
		conf.Postgres.Master.Port,
		conf.Postgres.Master.User,
		conf.Postgres.Master.Password,
		conf.Postgres.Master.DBName,
		"disable")

	dbConn, err := loadPsqlDb(conf.LogMode, dbHost, conf.Postgres.MaxOpenConnections, conf.Postgres.MaxIdleConnections, conf.Postgres.ConnMaxLifetime)
	if err != nil {
		log.Printf("failed to connect database master instance: %v", err)
		return HandlerDatabase{}, err
	}

	//set connection
	SetConnection(dbConn)

	return HandlerDatabase{
		Master: dbConn,
	}, nil
}

func loadPsqlDb(logMode bool, psqlInfo string, maxOpenConn, maxIdleConn, maxLifetime int) (*gorm.DB, error) {
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	// checking if connection to db has been established
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	if maxLifetime == 0 {
		maxLifetime = defaultConnMaxLifeTime
	}

	if maxIdleConn == 0 {
		maxIdleConn = defaultMaxIdleConns
	}

	if maxOpenConn == 0 {
		maxOpenConn = defaultMaxOpenConns
	}

	conn.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)
	conn.SetMaxOpenConns(maxOpenConn)
	conn.SetMaxIdleConns(maxIdleConn)

	if err != nil {
		return nil, err
	}

	var loggerGorm *gorm.Config
	if logMode {
		loggerGorm = &gorm.Config{
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		loggerGorm = &gorm.Config{
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	}

	return gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), loggerGorm)
}

var connDB *gorm.DB

// GetConnection : Get Available Connection
func GetConnection() *gorm.DB {
	return connDB
}

// SetConnection : Set Available Connection
func SetConnection(connect *gorm.DB) {
	connDB = connect
}
