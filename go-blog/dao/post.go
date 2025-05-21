package dao

import (
	"go-test/go-blog/models"
	"time"
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

func GetPost(query models.PostQuery) ([]models.PostMore, error) {

	rows, err := DB.Query("SELECT * FROM post LIMIT ? OFFSET ?", query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.PostMore
	for rows.Next() {
		post := models.PostMore{}
		ts := struct {
			CreateAt int64 `json:"createdAt"`
			UpdateAt int64 `json:"updatedAt"`
		}{}
		err = rows.Scan(&post.Pid, &post.Title, &post.Slug, &post.Content,
			&post.CategoryName, &post.CategoryId, &post.UserId, &post.ViewCount,
			&post.Type, &ts.CreateAt, &ts.UpdateAt)

		if err != nil {
			return nil, err
		}
		post.CreateAt = time.Unix(ts.CreateAt, 0).Format("2006-01-02 15:04:05")
		post.UpdateAt = time.Unix(ts.UpdateAt, 0).Format("2006-01-02 15:04:05")
		posts = append(posts, post)
	}
	return posts, nil
}
