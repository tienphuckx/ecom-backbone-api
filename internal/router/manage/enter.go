package manage

type ManageRouterGroup struct {
	UserRouter  UserRouter
	AdminRouter AdminRouter
}

var RouterGroupApp = new(ManageRouterGroup)
