package vo

type SellerApplyRequest struct {
	SellerID string `json:"seller_id" binding:"required"`
}

type SellerApproveRequest struct {
	SellerID string `json:"seller_id" binding:"required"`
	Approved bool `json:"approved" binding:"required"`
}

type SellerBlockRequest struct {
	SellerID string `json:"seller_id" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}

