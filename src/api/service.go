package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"task/src/api/schemas"
	"task/src/utils"

	"github.com/gin-gonic/gin"
)

func get_one(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	user := get_one_(int_id)

	if user.ID == 0 {
		ctx.JSON(404, gin.H{"message": "There is no user with this ID"})
		return
	}

	ctx.JSON(200, user)
}

func get_all(ctx *gin.Context) {
	page_str := ctx.Query("page")
	page_int, _ := strconv.Atoi(page_str)
	count_str := ctx.Query("count")
	count, _ := strconv.Atoi(count_str)
	order_by := ctx.Query("order_by")
	desc := ctx.Query("desc")

	// set default values if not presented
	if page_int == 0 {
		page_int = 1
	}

	// set default values if not presented
	if count == 0 {
		count = 10
	}

	page := page_int*count - count

	users, err := get_all_(page, count, order_by, desc)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	if len(users) == 0 {
		ctx.JSON(200, gin.H{"message": "No users in database"})
		return
	}

	ctx.JSON(200, users)
}

func create(ctx *gin.Context) {
	var user schemas.Create

	schema := ctx.MustGet("data")
	buf, _ := json.Marshal(schema)
	json.Unmarshal(buf, &user)

	result, err := utils.Get_data(user.Name, user.Surname, user.Patronymic)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	error := create_(result)

	if error != nil {
		fmt.Println(error.Error())
		ctx.JSON(500, error.Error())
		return
	}

	ctx.JSON(201, gin.H{"message": "Has been successfully created"})
}

func update(ctx *gin.Context) {
	var user schemas.Update

	schema := ctx.MustGet("data")
	buf, _ := json.Marshal(schema)
	json.Unmarshal(buf, &user)

	err := update_(user)

	if err != nil {
		fmt.Println(500, err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, gin.H{"message": "Has been updated successfully"})
}

func remove(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	affected, err := remove_(int_id)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	if affected == 0 {
		ctx.JSON(404, gin.H{"message": "There is no user with this ID"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Has been successfully deleted"})
}
