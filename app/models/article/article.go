package article

import (
	"goblogCalmk/app/models"
	"goblogCalmk/pkg/model"
	"goblogCalmk/pkg/route"
	"goblogCalmk/pkg/types"
	"strconv"
)

type Article struct {
	models.BaseModel

	Title string
	Body  string
}

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.BaseModel.ID), 10))
}
