package repositories

import (
	"hexagonal/domain"
	h "hexagonal/helpers"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	connection *gorm.DB
	db         string
}

func newMysqlConnection(DBName string) *gorm.DB {
	dsn := makeDataSourceName(DBName)

	connection, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: h.ENV("DRIVER_NAME"),
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return connection
}

// data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
func makeDataSourceName(DBName string) string {
	dsn := h.ENV("DB_USERNAME")

	if h.ENV("DB_PASSWORD") != "" {
		dsn += h.ENV("DB_PASSWORD")
	}

	dsn += "@tcp(" + h.ENV("DB_HOST") + ":" + h.ENV("DB_PORT") + ")"

	dsn += "/" + DBName

	dsn += "?charset=utf8&parseTime=True&loc=Local"

	return dsn
}

func NewMysqlRepository(DBName string) mysqlRepository {
	connection := newMysqlConnection(DBName)

	repo := &mysqlRepository{
		connection: connection,
		db:         DBName,
	}

	return *repo
}

func Find(code string) (*domain.Product, error)
func Store(product *domain.Product) error
func Update(product *domain.Product) error
func FindAll() ([]*domain.Product, error)
func Delete(code string) error
