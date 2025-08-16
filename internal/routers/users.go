package routers

import (
	"github.com/gin-gonic/gin"
	sponge_middleware "github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/jwt"
	"github.com/go-dev-frame/sponge/pkg/logger"

	"test-user-server/internal/config"
	"test-user-server/internal/handler"
	"test-user-server/internal/middleware"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		usersRouter(group, handler.NewUsersHandler())
	})
}

func usersRouter(group *gin.RouterGroup, h handler.UsersHandler) {
	g := group.Group("/users")

	// not change-me signing key will make routes use jwt authentication
	jwtCfg := config.Get().JWT
	if jwtCfg.SigningKey != "change-me" {
		g.Use(sponge_middleware.Auth(
			sponge_middleware.WithSignKey([]byte(jwtCfg.SigningKey)),
			sponge_middleware.WithExtraVerify(func(claims *jwt.Claims, c *gin.Context) error {
				logger.Info("middleware.Auth", logger.Any("claims", claims))
				return nil
			}),
		))
	}
	g.Use(middleware.RailsCookieAuthMiddleware("b1870c9c2d472d577b91a25f3ae9daa626725afffa70876d2fd9e004720e9a4f822bdcf0ddc07f3c54ae110d9ff852d5b5f648be56a275338f028287f90e8a85", "_coreui_pro_rails_starter_session"))

	// If jwt authentication is not required for all routes, authentication middleware can be added
	// separately for only certain routes. In this case, g.Use(middleware.Auth()) above should not be used.

	g.POST("/", h.Create)          // [post] /api/v1/users
	g.DELETE("/:id", h.DeleteByID) // [delete] /api/v1/users/:id
	g.PUT("/:id", h.UpdateByID)    // [put] /api/v1/users/:id
	g.GET("/:id", h.GetByID)       // [get] /api/v1/users/:id
	g.POST("/list", h.List)        // [post] /api/v1/users/list

	g.POST("/delete/ids", h.DeleteByIDs)   // [post] /api/v1/users/delete/ids
	g.POST("/condition", h.GetByCondition) // [post] /api/v1/users/condition
	g.POST("/list/ids", h.ListByIDs)       // [post] /api/v1/users/list/ids
	g.GET("/list", h.ListByLastID)         // [get] /api/v1/users/list
}
