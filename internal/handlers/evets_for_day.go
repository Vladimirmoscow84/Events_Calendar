package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/wb-go/wbf/ginext"
)

func (r *Router) EventsForDayHandler(c *ginext.Context) {

	userIDStr := c.Query("user_id")
	dateStr := c.Query("date")

	if userIDStr == "" || dateStr == "" {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "user_id and date are required",
		})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "invalid user_id",
		})
		return
	}

	day, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "invalid date format, expected YYYY-MM-DD",
		})
		return
	}

	events, err := r.Getter.EventsForDay(c.Request.Context(), userID, day)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ginext.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ginext.H{
		"result": events,
	})
}
