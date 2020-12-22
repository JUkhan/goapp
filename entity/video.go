package entity

type Person struct {
	Name string `json:"name" binding:"required,min=5"`
	Age  int8   `json:"age" binding:"required,gte=20,lte=130"`
}
type Video struct {
	Title       string `json:"title" binding:"required,min=10,max=120" validate:"is-cool"`
	Description string `json:"description" binding:"required,min=10,max=120"`
	URL         string `json:"url" binding:"required,url`
	Authore     Person `json:"authore" binding:"required"`
}
