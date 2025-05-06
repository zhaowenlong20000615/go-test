package dao

import (
	"go-test/go-blog/models"
	"log"
)

func GetCategorys() []models.Category {
	query, err := DB.Query("select * from category")
	if err != nil {
		log.Println(err)
		return nil
	}
	var categories []models.Category
	for query.Next() {
		var category models.Category
		err := query.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println(err)
		}
		categories = append(categories, category)
	}
	return categories
}
