package routes

import (
	"errors"
	"fmt"
	"go-admin/utils"
	"html/template"
	"strconv"
	"time"

	"go-admin/boot"
	"go-admin/infrastructure/postgres"
	controllers "go-admin/modules/controller"
	"go-admin/modules/middleware"
	models "go-admin/modules/primitive/model"
	"go-admin/modules/repository"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go-admin/infrastructure/session"
	"gorm.io/gorm"
)

func BackendRoute(e *echo.Echo, handlerSetup boot.HandlerSetup) {
	db := postgres.GetConnection()
	//=========== Backend ===========//
	var userInfo session.UserInfo
	//new middleware
	mv := echoview.NewMiddleware(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials: []string{
			"partials/asside",
			"partials/brand-menu",
			"partials/chart",
			"partials/chat",
			"partials/demo-panel",
			"partials/header-mobile",
			"partials/language",
			"partials/notification",
			"partials/quick-action",
			"partials/quick-panel",
			"partials/quick-panel-toogle",
			"partials/scrolltop",
			"partials/search",
			"partials/sticky-toolbar",
			"partials/sub-header",
			"partials/user-bar",
		},
		Funcs: template.FuncMap{
			"user": func(ctx echo.Context) models.User {
				userModel, err := utils.GetUserInfoFromContext(ctx, db)
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return models.User{}
					}
					return models.User{}
				}
				return userModel
			},
			"floatToString": func(value *float64) string {
				if value == nil {
					return ""
				}
				return strconv.FormatFloat(*value, 'f', -1, 64)
			},
			"floatNotPointerToString": func(value float64) string {
				return strconv.FormatFloat(value, 'f', -1, 64)
			},
			"formatDate": func(date time.Time, layout string) string {
				return date.Format(layout)
			},
			"maritalStatus": func(data string) string {
				var result string
				if data == "S" {
					result = "Belum Pernah Menikah"
				}
				return result
			},
			"MenuParent": func(ctx echo.Context) []map[string]interface{} {
				var dataMenu map[string]interface{}
				var listOfMenu []map[string]interface{}
				result, err := session.Manager.Get(ctx, session.SessionId)
				if err != nil {
					panic(err)
				}
				userInfo = result.(session.UserInfo)

				listParentMenu, err := repository.GetListParentMenu(db)
				if err != nil {
					log.Info("ERROR GET LIST PARENT MENU", err.Error())
				}

				if len(listParentMenu) > 0 {
					for _, v := range listParentMenu {
						parentID, err := strconv.Atoi(v.ID)
						if err != nil {
							log.Info("ERROR parse string to int ", err.Error())
						}
						menus, err := repository.GetMenu(db, parentID)
						if err != nil {
							log.Info("ERROR GET MENU BY ROLE ", err.Error())
						}
						dataMenu = map[string]interface{}{
							"Name":  v.Name,
							"Icon":  v.IconClass,
							"Menus": menus,
						}
						listOfMenu = append(listOfMenu, dataMenu)
					}
				}

				return listOfMenu
			},
			"Menu": func(ctx echo.Context) []map[string]interface{} {
				var dataMenu map[string]interface{}
				var listOfMenu []map[string]interface{}
				result, err := session.Manager.Get(ctx, session.SessionId)
				if err != nil {
					panic(err)
				}
				userInfo = result.(session.UserInfo)

				menus, err := repository.GetMenuWithParentID(db, 0)
				dataMenu = map[string]interface{}{
					"Menus": menus,
				}
				listOfMenu = append(listOfMenu, dataMenu)
				return listOfMenu
			},
		},
		DisableCache: true,
	})

	fmt.Println("userInfo ", userInfo)

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	bGroup := e.Group("/check")
	backendGroup := bGroup.Group("/admin", mv, middleware.SessionMiddleware(session.Manager))
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	menus, err := repository.GetMenuIsActive(db)
	if err != nil {
		log.Info("ERROR GET MENU BY ROLE ", err.Error())
	}

	homeController := controllers.NewHomeController()
	backendGroup.GET("/home", homeController.Index)

	authController := handlerSetup.AuthController
	backendGroup.POST("/logout", authController.Logout)

	userController := handlerSetup.UserController
	userGroup := backendGroup.Group("/register", authorizationMiddleware.AuthorizationMiddleware(menus, "user"))
	userGroup.GET("", userController.Index)
	userGroup.POST("/store", userController.Store)
	userGroup.GET("/add", userController.Add)
	userGroup.GET("/profile", userController.Profile)
	userGroup.GET("/datatable", userController.List)
	userGroup.DELETE("/delete/:id", userController.Delete)
	userGroup.GET("/detail/:id", userController.View)
	userGroup.GET("/edit/:id", userController.Edit)
	userGroup.POST("/update/:id", userController.Update)

	//ConfigController
	configController := handlerSetup.ConfigController
	configGroup := backendGroup.Group("/config", authorizationMiddleware.AuthorizationMiddleware(menus, "config"))
	configGroup.GET("", configController.Index)
	configGroup.POST("/store", configController.Store)
	configGroup.GET("/datatable", configController.Datatable)
	configGroup.POST("/update/:id", configController.Update)
	configGroup.GET("/:id", configController.Detail)
	configGroup.POST("/delete/:id", configController.Delete)
	bGroup.POST("/admin/config/set-active/:id", configController.SetActive)
	bGroup.POST("/admin/config/set-inactive/:id", configController.SetInactive)

	// RoleController
	roleController := handlerSetup.RoleController
	roleGroup := backendGroup.Group("/role", authorizationMiddleware.AuthorizationMiddleware(menus, "role"))
	roleGroup.GET("", roleController.Index)
	roleGroup.GET("/datatable", roleController.List)
	roleGroup.GET("/add", roleController.Add)
	roleGroup.POST("/store", roleController.Store)
	roleGroup.GET("/edit/:id", roleController.Edit)
	roleGroup.POST("/update/:id", roleController.Update)
	roleGroup.GET("/detail/:id", roleController.View)
	roleGroup.DELETE("/delete/:id", roleController.Delete)
}
