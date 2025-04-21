
package model

func AllModels() []interface{} {
	// Init models here
	return []interface{}{
		&User{},
		&UserInfo{},
		&Seller{},
		&Project{},
		&SubProject{},
		&Listing{},
		&LoanSupport{},
		&ProjectManager{},
		&Investor{},
	}
}
