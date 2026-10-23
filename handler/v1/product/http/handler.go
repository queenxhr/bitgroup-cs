package product

import (
	"net/http"
	"strconv"

	appErrors "bitgroup-cs/internal/errors"
	"bitgroup-cs/internal/usecase"
	"bitgroup-cs/utils"
	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	FetchAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ProductHandler struct {
	productUc usecase.IProductUseCase
}

func NewHandler(productUseCase usecase.IProductUseCase) IProductHandler {
	return &ProductHandler{
		productUc: productUseCase,
	}
}

func (h *ProductHandler) FetchAll(ctx *gin.Context) {
	products, err := h.productUc.FetchAll(ctx.Request.Context())
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, ProductResponse{}.ParseFromEntities(products))
}

func (h *ProductHandler) FindById(ctx *gin.Context) {
	id, err := h.getIdParam(ctx)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	product, err := h.productUc.FindById(ctx.Request.Context(), id)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, ProductObject{}.ParseFromEntity(product))
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var reqBody CreateProductRequestDto
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		utils.SendJSONErrorResponse(ctx, appErrors.New(appErrors.ErrBadRequest, err.Error()))
		return
	}

	newProduct := reqBody.ToEntity()
	result, err := h.productUc.CreateProduct(ctx.Request.Context(), newProduct)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusCreated, ProductObject{}.ParseFromEntity(result), "Product created successfully")
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	id, err := h.getIdParam(ctx)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	var reqBody UpdateProductRequestDto
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		utils.SendJSONErrorResponse(ctx, appErrors.New(appErrors.ErrBadRequest, err.Error()))
		return
	}

	updateData := reqBody.ToEntity(id)
	result, err := h.productUc.UpdateProductByID(ctx.Request.Context(), updateData)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, ProductObject{}.ParseFromEntity(result), "Product updated successfully")
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id, err := h.getIdParam(ctx)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	err = h.productUc.DeleteProductByID(ctx.Request.Context(), id)
	if err != nil {
		utils.SendJSONErrorResponse(ctx, err)
		return
	}

	utils.JSON(ctx, http.StatusOK, gin.H{"id": id}, "Product deleted successfully")
}

func (h *ProductHandler) getIdParam(ctx *gin.Context) (int64, error) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, appErrors.New(appErrors.ErrBadRequest, "Invalid product ID parameter")
	}
	return id, nil
}
