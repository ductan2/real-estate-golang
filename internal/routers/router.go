package routers

import (
	"ecommerce/internal/routers/investor"
	"ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouter
	Seller user.SellerRouter
	Admin user.AdminRouter
	Investor investor.InvestorRouter
}

var RouterGroupApp = new(RouterGroup)
