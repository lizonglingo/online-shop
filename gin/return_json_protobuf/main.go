package main

import (
	userpb "gin/return_json_protobuf/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/moreJSON", moreJSON)
	router.GET("/proto", proto)

	router.Run(":8099")
}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
	}
	msg.Name = "lzl"
	msg.Message = "more json"

	c.JSON(http.StatusOK, msg)
}

func proto(c *gin.Context) {
	teacher := &userpb.Teacher{
		Name:   "lzl",
		Course: []string{"go", "python"},
	}
	c.ProtoBuf(http.StatusOK, teacher)
}
