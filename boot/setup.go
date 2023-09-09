package boot

import (
	"encoding/gob"
	"os"

	"go-admin/config"
	"go-admin/infrastructure/logger"
	"go-admin/infrastructure/postgres"
	"go-admin/infrastructure/session"
	controllers "go-admin/modules/controller"
	"go-admin/modules/helper"
	models "go-admin/modules/primitive/model"
	"go-admin/modules/repository"
	"go-admin/modules/service"

	log "github.com/sirupsen/logrus"
)

type HandlerSetup struct {
	UserController   controllers.UserController
	HealthController controllers.HealthController
	AuthController   controllers.AuthController
	ConfigController  controllers.ConfigController
	RoleController    controllers.RoleController
	DivisiController  controllers.DivisiController
	JabatanController controllers.JabatanController
}

func MakeHandler() HandlerSetup {
	config.Initialize()
	logger.Init(config.Conf.LogFormat, config.Conf.LogLevel)

	var err error

	//register data within the context
	gob.Register(session.UserInfo{})
	gob.Register(session.FlashMessage{})
	gob.Register(models.User{})
	gob.Register(models.Menu{})
	gob.Register(map[string]interface{}{})
	gob.Register([]helper.ValidationError{})

	//setup infrastructure postgres
	db, err := postgres.NewDatabaseClient(&config.Conf)
	if err != nil {
		log.Fatalf("failed initiate database postgres: %v", err)
		os.Exit(1)
	}

	healthRepo := repository.NewHealthRepository(db.Master)
	healthService := service.NewHealthService(healthRepo)
	healthController := controllers.NewHealthController(healthService)

	userRepo := repository.NewUserRepository(db.Master)
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	authService := service.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	configRepo := repository.NewConfigRepository(db.Master)
	configService := service.NewConfigService(configRepo)
	configController := controllers.NewConfigController(configService)

	roleRepo := repository.NewRoleRepository(db.Master)
	roleService := service.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService)

	divisiRepo := repository.NewDevisiRepository(db.Master)
	divisiService := service.NewDevisiService(divisiRepo)
	divisiController := controllers.NewDevisiController(divisiService)

	jabatanRepo := repository.NewJabatanRepository(db.Master)
	jabatanService := service.NewJabatanService(jabatanRepo)
	jabatanController := controllers.NewJabatanController(jabatanService)

	return HandlerSetup{
		UserController:    userController,
		HealthController:  healthController,
		AuthController:    authController,
		ConfigController:  configController,
		RoleController:    roleController,
		DivisiController:  divisiController,
		JabatanController: jabatanController,
	}
}
