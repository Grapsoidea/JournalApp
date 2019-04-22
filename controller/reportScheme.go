package controller

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model"
	"github.com/gin-gonic/gin"
)

// GetReportSchemes Получить все схемы отчетов
// @Summary Список схем отчетов
// @Description Метод, который получает все списки отчетов
// @Tags ReportScheme
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ReportScheme
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/report [get]
func (c *Controller) GetReportSchemes(ctx *gin.Context) {
	schemes, err := model.ReportSchemeAll()
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, schemes)
}

// GetReportScheme Получить схему отчетов с id
// @Summary Схему отчетов с id
// @Description Метод, который получает схему отчетов с заданным id
// @Tags ReportScheme
// @Accept  json
// @Produce  json
// @Param reportscheme_id path string true "ReportScheme id"
// @Success 200 {object} model.ReportScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/report/{reportscheme_id} [get]
func (c *Controller) GetReportScheme(ctx *gin.Context) {
	id := ctx.Param("reportscheme_id")
	scheme, err := model.ReportSchemeOne(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, scheme)
}

// NewReportScheme Создать новую схему отчетов
// @Summary Новая схема отчетов
// @Description Метод, который создает новую схему отчетов
// @Tags ReportScheme
// @Accept  json
// @Produce  json
// @Param NewReportScheme body model.NewReportScheme true "New Report Scheme"
// @Success 200 {object} model.NewReportScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/report [post]
func (c *Controller) NewReportScheme(ctx *gin.Context) {
	var newReportScheme model.NewReportScheme
	if err := ctx.ShouldBindJSON(&newReportScheme); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := newReportScheme.Validation(); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := newReportScheme.Insert()
	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, newReportScheme)
}

// UpdateReportScheme Изменить схему отчетов с id
// @Summary Изменить схему отчетов с id
// @Description Метод, который изменяет схему отчетов с заданным id
// @Tags ReportScheme
// @Accept  json
// @Produce  json
// @Param reportscheme_id path string true "ReportScheme id"
// @Param UpdateReportScheme body model.ReportScheme true "Update Report Scheme"
// @Success 200 {object} model.ReportScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/report/{reportscheme_id} [put]
func (c *Controller) UpdateReportScheme(ctx *gin.Context) {
	id := ctx.Param("reportscheme_id")

	var updateReportScheme model.UpdateReportScheme
	if err := ctx.ShouldBindJSON(&updateReportScheme); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := updateReportScheme.Validation(); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := updateReportScheme.Update(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, updateReportScheme)
}

// DeleteReportScheme Удалить схему объектов с id
// @Summary Удалить схему объектов с id
// @Description Метод, который удаляет схему объектов с заданным id
// @Tags ReportScheme
// @Accept  json
// @Produce  json
// @Param Reportscheme_id path string true "ReportSheme id"
// @Success 200 {string} string    "5ca10d9d015c736a72b7b3ba"
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/report/{reportscheme_id} [delete]
func (c *Controller) DeleteReportScheme(ctx *gin.Context) {
	id := ctx.Param("reportscheme_id")
	err := model.DeleteReportSchemeOne(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, id)
}
