package handlers

import (
	"net/http"

	"github.com/wb-go/wbf/ginext"
)

func (r *Router) DeleteEventHandler(c *ginext.Context) {
	var req struct {
		EventID int `json:"event_id"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "invalid JSON body",
		})
		return
	}

	if req.EventID <= 0 {
		c.JSON(http.StatusBadRequest, ginext.H{
			"error": "invalid event_id",
		})
		return
	}

	err = r.Constructor.Delete(c.Request.Context(), req.EventID)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, ginext.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ginext.H{
		"result": "event deleted",
	})
}
