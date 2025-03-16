package routers

import "ecommerce/internal/routers/user"

type RouterGroup struct {
	User user.UserRouter
	Seller user.SellerRouter
	Admin user.AdminRouter
}

var RouterGroupApp = new(RouterGroup)
