package controllers

import (
	"go_learn/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

func GetPosts(c *gin.Context){
	posts := models.GetPost()
	c.JSON(http.StatusOK, gin.H{
		"status"	:	0,
		"data"	:	posts,
	})
}

func GetUserId(c *gin.Context) {
	fmt.Println("in here")

	username:= c.Query("username")
	passwd := c.Query("passwd")
	fmt.Println(username)
	fmt.Println(passwd)
	if username == "abc" && passwd == "abc"{
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg": "sucess!",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg": "Not Found!",
		})
	}
}

func LikeById(c *gin.Context) {
	postid := c.Param("postid")
	if postid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "postid 不能为空",
		})
		return
	}

	postId, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "postid 格式错误",
		})
		return
	}

	err = models.LikeById(postId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "系统错误, err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
	})
	return
}