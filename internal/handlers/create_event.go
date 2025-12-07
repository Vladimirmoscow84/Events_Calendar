package handlers

import (
	"net/http"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
	"github.com/wb-go/wbf/ginext"
)

func (r *Router) CreateEventHandler(c *ginext.Context) {
	var event model.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{"error": "invalid JSON body"})
		return
	}

	if event.UserID <= 0 {
		c.JSON(http.StatusBadRequest, ginext.H{"error": "nvalid user_id"})
		return
	}

	if event.Title == "" {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "title is required",
		})
		return
	}

	if event.Date.IsZero() || event.Date.After(time.Now().AddDate(100, 0, 0)) {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "invalid date",
		})
		return
	}

	id, err := r.Constructor.CreateEvent(c.Request.Context(), &event)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ginext.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ginext.H{
		"result": "event created",
		"id":     id,
	})

}
