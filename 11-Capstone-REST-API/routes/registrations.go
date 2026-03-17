package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.com/rest-api/models"
	"go.com/rest-api/util"
)

func registerToAnEvent(ctx *gin.Context) {
	userID := ctx.GetInt64(util.USER_ID)
	
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}	
	
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event for the given event ID."})
		return
		}	
		
	err = event.Register(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user to the event with event ID."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func unregisterToAnEvent(ctx *gin.Context) {
	userID := ctx.GetInt64(util.USER_ID)
	
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}	
	
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event for the given event ID."})
		return
	}	
		
	err = event.Unregister(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister user to the event with event ID."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Unregistered"})
}