package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/jwt"
	"github.com/go-dev-frame/sponge/pkg/logger"

	"test-user-server/internal/config"
	"test-user-server/internal/handler"
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
		g.Use(middleware.Auth(
			middleware.WithSignKey([]byte(jwtCfg.SigningKey)),
			middleware.WithExtraVerify(func(claims *jwt.Claims, c *gin.Context) error {
				logger.Info("middleware.Auth", logger.Any("claims", claims))
				return nil
			}),
		))
	}
	railsCfg := config.Get().Rails
	if railsCfg.SecretKeyBase != "change-me" {
		g.Use(middleware.RailsCookieAuthMiddleware(railsCfg.SecretKeyBase, railsCfg.CookieName))
		g.Use(VerifyRailsSessionUserIdIs(railsCfg.UserID))
	}

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
