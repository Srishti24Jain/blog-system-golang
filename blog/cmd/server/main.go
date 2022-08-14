package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"blog/api/delivery/httphandler"
	"blog/api/middleware"
	"blog/api/middleware/swagger"
	"blog/api/usecase"
	"blog/db"
	"blog/domain/dto"
)

func main() {

	// connect to db
	conn, err := db.Connect()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to connect to db: %+v\n", err)
	}

	//auto migrations
	conn.AutoMigrate(&dto.User{}, &dto.Post{}, &dto.Tag{}, &dto.Comment{})
	// New gin server
	r := gin.New()

	// inject middlewares
	// newLogger
	// recover
	// swagger editor

	logger, _ := zap.NewProduction()
	r.Use(middleware.JSONMiddleware())

	/*  Add a ginzap middleware, which:
	    - Logs all requests, like a combined access and error log.
	    - Logs to stdout.
		- RFC3339 with UTC time format.
	*/

	// Host Swagger middleware
	r.Use(gin.WrapH(swagger.Middleware()))

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	/* Logs all panic to error log - stack means whether output the stack info. */
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Serve UI files
	r.Use(static.Serve("/", static.LocalFile("/app/assets", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("/app/assets")
	})

	// users endpoints
	userUsecase := usecase.NewUserUsecase(conn)
	httphandler.NewUserHandler(r, userUsecase)

	//tags endpoints
	tagsUsecase := usecase.NewTagsUsecase(conn)
	httphandler.NewTagsHandler(r, tagsUsecase)

	//posts endpoints
	postUsecase := usecase.NewPostUsecase(conn)
	httphandler.NewPostHandler(r, postUsecase)

	//comments endpoints
	commentsUsecase := usecase.NewCommentsUsecase(conn)
	httphandler.NewCommentsHandler(r, commentsUsecase)

	// Start the server
	_ = r.Run(":8080")
}
