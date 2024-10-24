package user

type UserRouterGroup struct {
	UserRouter    UserRouter
	ProductRouter ProductRouter
}

var RouterGroupApp = new(UserRouterGroup)
