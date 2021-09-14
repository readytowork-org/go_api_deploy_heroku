module go_api_deploy_heroku

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/joho/godotenv v1.3.0
	github.com/rs/cors v1.7.0
	github.com/spf13/viper v1.7.1
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.16.0
	gorm.io/driver/postgres v1.1.1
	gorm.io/gorm v1.21.15
)
