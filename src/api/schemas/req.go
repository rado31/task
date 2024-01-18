package schemas

import (
	"encoding/json"
	"strconv"
	"task/src/utils"

	"github.com/gin-gonic/gin"
)

type User_id struct {
	ID int `json:"id" validate:"gt=0"`
}

type QueryParam struct {
	Page     int    `json:"page" validate:"omitempty,gt=0"`
	Count    int    `json:"count" validate:"omitempty,gt=0"`
	Order_by string `json:"order_by" validate:"oneof=id name surname age"`
	Desc     string `json:"desc" validate:"oneof=desc asc"`
}

type Create struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic"`
}

type Update struct {
	ID          int    `json:"id" validate:"gt=0"`
	Name        string `json:"name" validate:"required"`
	Surname     string `json:"surname" validate:"required"`
	Patronymic  string `json:"patronymic" validate:"omitnil"`
	Age         int    `json:"age" validate:"gt=0,lt=100"`
	Gender      string `json:"gender" validate:"oneof=male female"`
	Nationality string `json:"nationality" validate:"required"`
}

func Validate_id(ctx *gin.Context) {
	var schema User_id

	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)
	schema.ID = int_id

	errors := utils.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Next()
}

func Validate_query(ctx *gin.Context) {
	var schema QueryParam

	page_str := ctx.Query("page")
	page, _ := strconv.Atoi(page_str)
	count_str := ctx.Query("count")
	count, _ := strconv.Atoi(count_str)
	order_by := ctx.Query("order_by")
	desc := ctx.Query("desc")

	schema.Page = page
	schema.Count = count
	schema.Order_by = order_by
	schema.Desc = desc

	errors := utils.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Next()
}

func Validate_create(ctx *gin.Context) {
	var schema Create

	data, _ := ctx.GetRawData()
	json.Unmarshal(data, &schema)

	errors := utils.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Set("data", schema)
	ctx.Next()
}

func Validate_update(ctx *gin.Context) {
	var schema Update

	data, _ := ctx.GetRawData()
	json.Unmarshal(data, &schema)

	errors := utils.Validation_errors(&schema)

	if errors != nil {
		ctx.JSON(400, errors)
		ctx.Abort()
	}

	ctx.Set("data", schema)
	ctx.Next()
}
