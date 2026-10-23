package category

import (
	"net/http"
	"strconv"

	appErrors "bitgroup-cs/internal/errors"
	"bitgroup-cs/internal/usecase"
	"bitgroup-cs/utils"
	"github.com/gin-gonic/gin"
)

type ICategoryHandler interface {
	FetchAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type CategoryHandler struct {
	categoryUc usecase.ICategoryUseCase
}

func NewHandler(categoryUseCase usecase.ICategoryUseCase) ICategoryHandler {
	return &CategoryHandler{
		categoryUc: categoryUseCase,
	}
}

func (h *CategoryHandler) FetchAll(ctx *gin.Context) {
	categories, err := h.categoryUc.FetchAll(ctx.Request.Context())
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, CategoryResponse{}.ParseFromEntities(categories))
}

func (h *CategoryHandler) FindById(ctx *gin.Context) {
	id, err := h.getIdParam(ctx)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	category, err := h.categoryUc.FindById(ctx.Request.Context(), id)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, CategoryObject{}.ParseFromEntity(category))
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var reqBody CreateCategoryRequestDto
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		utils.SendJSONErrorResponse(ctx, appErrors.New(appErrors.ErrBadRequest, err.Error()))
		return
	}

	newCategory := reqBody.ToEntity()
	result, err := h.categoryUc.CreateCategory(ctx.Request.Context(), newCategory)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusCreated, CategoryObject{}.ParseFromEntity(result), "Category created successfully")
}

func (h *CategoryHandler) Update(ctx *gin.Context) {
	id, err := h.getIdParam(ctx)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	var reqBody UpdateCategoryRequestDto
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		utils.SendJSONErrorResponse(ctx, appErrors.New(appErrors.ErrBadRequest, err.Error()))
		return
	}

	updateData := reqBody.ToEntity(id)
	result, err := h.categoryUc.UpdateCategoryByID(ctx.Request.Context(), updateData)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, CategoryObject{}.ParseFromEntity(result), "Category updated successfully")
}

func (h *CategoryHandler) Delete(ctx *gin.Context) {
	id, err := h.getIdParam(ctx)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	err = h.categoryUc.DeleteCategoryByID(ctx.Request.Context(), id)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, gin.H{"id": id}, "Category deleted successfully")
}

func (h *CategoryHandler) getIdParam(ctx *gin.Context) (int64, error) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, appErrors.New(appErrors.ErrBadRequest, "Invalid category ID parameter")
	}
	return id, nil
}
