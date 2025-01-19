package types

type User struct {
	Id			int		`json:"id"`
	Name		string  `json:"name"`
	Password	string	`json:"password"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User		`json:"users"`
}

type CreateUserRequest struct {
	Id			int		`json:"id" binding:"required"`
	Name		string  `json:"name" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

func (c *CreateUserRequest) ToUser() *User {
	return &User{
		Id:		c.Id,
		Name:	c.Name,
		Password:	c.Password,
	}
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserRequest struct {
	Name				string		`json:"name" binding:"required"`
	Password			string		`json:"password" binding:"password"`
	UpdatePassword		string		`json:"updatePassword" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserRequest struct {
	Name		string		`json:"name" binding:"reuiqred"`
	Password	string		`json:"password" binding:"required"`	
}

func (d *DeleteUserRequest) ToUser() *User {
	return &User{
		Name:	d.Name,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
}
