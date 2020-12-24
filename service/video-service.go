package service

import (
	"github.com/JUkhan/goapp/entity"
	"github.com/JUkhan/goapp/repository"
)

type (
	VideoService interface {
		Add(*entity.Video)
		FindAll() []entity.Video
		Update(*entity.Video)
		Delete(*entity.Video)
	}
	videoService struct {
		videoRepo repository.VideoRepository
	}
)

func NewVideoService(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepo: repo,
	}
}

func (s *videoService) Add(item *entity.Video) {
	s.videoRepo.Save(item)
}

func (s *videoService) FindAll() []entity.Video {
	return s.videoRepo.FindAll()
}
func (s *videoService) Update(item *entity.Video) {
	s.videoRepo.Update(item)
}
func (s *videoService) Delete(item *entity.Video) {
	s.videoRepo.Delete(item)
}
