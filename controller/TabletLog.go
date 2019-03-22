package controller

import "github.com/gin-gonic/gin"

// AddTablelog Сохранение логов
// @Summary Сохранение логов
// @Description Сохранение логов на сервер
// @Tags Logs
// @Accept  json
// @Produce  json
// @Param errorLog query string true "Log with error"
// @Success 200 {string} string "answer"
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /journals/{journal_id}/items/{item_id} [put]
func (c *Controller) AddTablelog(ctx *gin.Context) {
}
