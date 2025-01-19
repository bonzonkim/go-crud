package networks

import (
	"fmt"
	"go-crud/services"
	"go-crud/types"

	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit		sync.Once
	userRouterInstance	*userRouter
)

type userRouter struct {
	router		 *Network
	service		 *services.UserService
}

func newUserRouter(router *Network, userService *services.UserService) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:		router,
			service:	userService,
		}
		router.registerGET("/", userRouterInstance.getUser)
		router.registerPOST("/", userRouterInstance.createUser)
		router.registerUPDATE("/", userRouterInstance.updateUser)
		router.registerDELETE("/", userRouterInstance.deleteUser)
	})
	return userRouterInstance
}

func (u *userRouter) createUser(c *gin.Context) {
	var req types.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse:  types.NewApiResponse(-1, "Binding Error while createing user", err.Error()),
		})
	} else if err := u.service.CreateUser(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse:	types.NewApiResponse(-1, "Error while creating user", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse:	types.NewApiResponse(1, "Successfully created user", nil),
		})
	}
}
func (u *userRouter) getUser(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse:	types.NewApiResponse(1, "successfully load users", nil),
		Users:			u.service.GetUsers(),
	})
}

func (u *userRouter) updateUser(c *gin.Context) {
	fmt.Println("update")
}
func (u *userRouter) deleteUser(c *gin.Context) {
	fmt.Println("delete")
}
