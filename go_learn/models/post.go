package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_learn/database"
	"log"
)

type Post struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	LikeCount  int32  `gorm:"column:like_count" json:"like_count"`
}
/*对应结构体去db中拿取信息*/
type Comment struct {

}

func  GetPost()  []*Post{
	var results []*Post
	err := database.DBInstance.Find(&results).Error
	if err != nil {
		log.Println("ERROR:", err)
		fmt.Println(err)
		return nil
	}
	return results
}

func  GetPostById(postid int64)  *Post{
	var results *Post
	err := database.DBInstance.
		Where("id = ?",postid).
		Find(&results).Error
	if err != nil {
		log.Println("ERROR:", err)
		fmt.Println(err)
		return nil
	}
	return results
}

func LikeById(postid int64) error{
	err := database.DBInstance.Model(&Post{}).
		Where("id = ?",postid).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error
	return  err
}