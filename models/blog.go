package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Blog struct
{
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Body string `json:"body" bson:"body"`
	CreateDate time.Time `json:"CreateDate" bson:"CreateDate"`
}

type Blogs []Blog

func (blog *Blog) Validate() (bool) {

	if blog.Title == "" {
		return false
	}
	if blog.Body == "" {
		return false
	}

	return true
}