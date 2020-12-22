package service

import "github.com/JUkhan/goapp/entity"

type VideoService interface {
	Add(entity.Video) entity.Video
	FindAll() []entity.Video
}

func NewVideoService() VideoService {
	return &videoStorage{
		[]entity.Video{},
	}
}

type videoStorage struct {
	videos []entity.Video
}

func (s *videoStorage) Add(item entity.Video) entity.Video {
	s.videos = append(s.videos, item)
	return item
}

func (s *videoStorage) FindAll() []entity.Video {
	return s.videos
}
