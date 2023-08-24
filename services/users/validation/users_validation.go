package validation

// swagger:parameters CreateUsersRequest
type CreateUsersRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name"  binding:"required,min=1,max=300"`
}

// swagger:parameters UpdateUsersRequest
type UpdateUsersRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name"  binding:"required,min=1,max=300"`
}
