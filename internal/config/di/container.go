package di

import (
	"go-with-fiber/internal/config"
	dbConfig "go-with-fiber/internal/config/database"
	registerDomain "go-with-fiber/internal/domain/register"
	registerHandler "go-with-fiber/internal/handler/register"
	registerPort "go-with-fiber/internal/ports/register"

	loginDomain "go-with-fiber/internal/domain/login"
	loginHandler "go-with-fiber/internal/handler/login"
	loginPort "go-with-fiber/internal/ports/login"
)

type Container struct {
	RHandeler *registerHandler.RegisterHandler
	LHandler  *loginHandler.LoginHandler
}

func InitContainer(cfg *config.Config) *Container {

	// Initialize the database connection
	db, err := dbConfig.InitilizeDatabase(cfg)
	if err != nil {
		panic("failed to initialise database❌")
	}

	//tables created
	db.AutoMigrate(
		&registerPort.User{},
	)

	//registration module
	rPort := registerPort.NewRegisterService(db)
	rDomain := registerDomain.NewRegisterService(rPort)
	rHandelr := registerHandler.NewRegisterHandler(rDomain)

	//login module

	lPort := loginPort.NewLoginService(db)
	lDomain := loginDomain.NewLoginService(lPort)
	lHandelr := loginHandler.NewLoginHandler(lDomain)

	return &Container{
		RHandeler: rHandelr,
		LHandler:  lHandelr,
	}
}
