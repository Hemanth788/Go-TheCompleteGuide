package routes

import (
	"github.com/gin-gonic/gin"
	"go.com/rest-api/middleware"
)

func RegisterEventRoutes(router *gin.Engine) {

	// event routes
	router.GET("/events/:id", getEventByID)
	router.GET("/events", getEvents)

	authedRouter := router.Group("/")
	authedRouter.Use(middleware.Authenticate)

	// authneticated event routes
	authedRouter.POST("/events", createEvent)
	authedRouter.PUT("/events/:id", updateEvent)
	authedRouter.DELETE("/events/:id", deleteEvent)

	// auth routes
	router.POST("/signup", signup)
	router.POST("/login", login)

	// registration routes
	authedRouter.POST("/events/:id/register", registerToAnEvent)
	authedRouter.DELETE("/events/:id/register", unregisterToAnEvent)
}
