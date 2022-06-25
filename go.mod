module wallet-engine

// +heroku goVersion go1.16

go 1.16

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.8.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.5
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	gorm.io/driver/postgres v1.3.7
	gorm.io/gorm v1.23.6
)

require golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
