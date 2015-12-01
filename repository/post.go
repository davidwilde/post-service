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

func (r *PostRepository) SavePostById(id string, comment Comment) {
	var buffer bytes.Buffer

	member := redis.Z{Score: float64(comment.Timecode), Member: comment.Comment}

	buffer.WriteString("video:")
	buffer.WriteString(id)
	err := r.client.ZAdd(buffer.String(), member).Err()
	if err != nil {
		panic(err)
	}
}

func (r *PostRepository) FetchPostsForId(id string) []Comment {
	var buffer bytes.Buffer

	buffer.WriteString("video:")
	buffer.WriteString(id)

	opts := redis.ZRangeByScore{Min: "0", Max: "inf"}

	result, err := r.client.ZRangeByScoreWithScores(buffer.String(), opts).Result()

	if err != nil {
		panic(err)
	}

	comments := make([]Comment, len(result))
	for i, v := range result {
		comments[i].Comment = v.Member.(string)
		comments[i].Timecode = int(v.Score)
	}

	return comments
}
