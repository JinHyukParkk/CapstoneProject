package controllers

import (
	"net/http"
	"strconv"

	"github.com/JinHyukParkk/CapstoneProject/GoogleAPI"
	"github.com/JinHyukParkk/CapstoneProject/models"
	"github.com/labstack/echo"
)

func VideoList(c echo.Context) error {
	listData, err := googleApi.ListAPI()
	check(err)
	videoList := []models.Video{}

	for _, data := range listData {
		videoList = append(videoList, models.Video{data})
	}

	u := &models.VideoListModel{
		Video_List: videoList,
		Total:      strconv.Itoa(len(listData)),
	}
	return c.JSON(http.StatusOK, u)
}