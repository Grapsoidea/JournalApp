package controller

import (
	"net/http"

	"github.com/Oxynger/JournalApp/httputils"
	"github.com/Oxynger/JournalApp/model"
	"github.com/gin-gonic/gin"
)

// GetJournalSchemes Получить все схемы журналов
// @Summary Список схем журналов
// @Description Метод, который получает все списки журналов
// @Tags JournalScheme
// @Accept  json
// @Produce  json
// @Success 200 {array} model.JournalScheme
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/journal [get]
func (c *Controller) GetJournalSchemes(ctx *gin.Context) {
	schemes, err := model.JournalSchemeAll()
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, schemes)
}

// GetJournalScheme Получить схему журнала с id
// @Summary Схему журнала с id
// @Description Метод, который получает схему журнала с заданным id
// @Tags JournalScheme
// @Accept  json
// @Produce  json
// @Param journalscheme_id path string true "JournalSheme id"
// @Success 200 {object} model.JournalScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/journal/{journalscheme_id} [get]
func (c *Controller) GetJournalScheme(ctx *gin.Context) {
	id := ctx.Param("journalscheme_id")
	scheme, err := model.JournalSchemeOne(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, scheme)
}

// NewJournalScheme Создать новую схему журналов
// @Summary Новая схема журналов
// @Description Метод, который создает новую схему журналов
// @Tags JournalScheme
// @Accept  json
// @Produce  json
// @Param NewJournalScheme body model.NewJournalScheme true "New Journal Scheme"
// @Success 200 {object} model.NewJournalScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/journal [post]
func (c *Controller) NewJournalScheme(ctx *gin.Context) {
	var newJournalScheme model.NewJournalScheme
	if err := ctx.ShouldBindJSON(&newJournalScheme); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := newJournalScheme.Validation(); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := newJournalScheme.Insert()
	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, newJournalScheme)
}

// UpdateJournalScheme Изменить схему журнала с id
// @Summary Изменить схему журнала с id
// @Description Метод, который изменяет схему журнала с заданным id
// @Tags JournalScheme
// @Accept  json
// @Produce  json
// @Param journalcheme_id path string true "JournalSheme id"
// @Param UpdateJournalScheme body model.JournalScheme true "Update Journal Scheme"
// @Success 200 {object} model.JournalScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/journal/{journalscheme_id} [put]
func (c *Controller) UpdateJournalScheme(ctx *gin.Context) {
	id := ctx.Param("journalscheme_id")

	var updateJournalScheme model.UpdateJournalScheme
	if err := ctx.ShouldBindJSON(&updateJournalScheme); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := updateJournalScheme.Validation(); err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err := updateJournalScheme.Update(id)

	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, updateJournalScheme)
}

// DeleteJournalScheme Удалить схему объектов с id
// @Summary Удалить схему объектов с id
// @Description Метод, который удаляет схему объектов с заданным id
// @Tags JournalScheme
// @Accept  json
// @Produce  json
// @Param Journalscheme_id path string true "JournalSheme id"
// @Success 200 {object} model.JournalScheme
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /scheme/journal/{journalscheme_id} [delete]
func (c *Controller) DeleteJournalScheme(ctx *gin.Context) {
	id := ctx.Param("journalscheme_id")
	err := model.DeleteJournalSchemeOne(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, id)
}
