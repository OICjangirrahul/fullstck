package userstruct



type User struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=5,max=20"`
	Age  int    `json:"age" validate:"gte=18,lte=120"`
}

// Response represents the response structure
type Response struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

var users = make(map[string]User)