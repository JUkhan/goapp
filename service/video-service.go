package service

type VideoServices interface {
	Add(entity.Video) entity.Video
	FindAll() []entity.Video
}
