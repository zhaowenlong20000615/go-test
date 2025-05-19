package dao

import (
	"go-test/go-blog/models"
)

func SavePost(post models.PostReq) (models.Post, error) {
	ret, err := DB.Exec("INSERT INTO post (title, slug, content, markdown, category_id, user_id, view_count, type) VALUES (?,?,?,?,?,?,?,?)",
		post.Title,
		post.Slug,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
	)
	m := models.Post{}
	if err != nil {
		return m, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return m, err
	}
	m.Pid = int(id)
	return m, nil
}
