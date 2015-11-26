package repository

import (
	"bytes"

	"gopkg.in/redis.v3"
)

func NewPostRepository(redis *redis.Client) *PostRepository {
	return &PostRepository{redis}
}

type PostRepository struct {
	client *redis.Client
}

func (r *PostRepository) SavePostById(comment Comment) {
	var buffer bytes.Buffer

	member := redis.Z{Score: float64(comment.Timecode), Member: comment.Comment}

	buffer.WriteString("video:")
	buffer.WriteString(string(comment.Id))
	r.client.ZAdd(buffer.String(), member)
}
