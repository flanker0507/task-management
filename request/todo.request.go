package request

type TodoUpdateeRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Note       string `json:"note" form:"note"`
	IsComplete bool   `json:"is_complete" form:"is_complete"`
	UserId     uint   `json:"user_id" form:"user_id"`
}

type TodoUpdateRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Note       string `json:"note" form:"note"`
	IsComplete bool   `json:"is_complete" form:"is_complete" validate:"required"`
}

func (TodoUpdateeRequest) TableName() string {
	return "todos"
}
