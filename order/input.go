package order

type OrderInput struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}

type OrderSearchInput struct {
	Column string `json:"column" binding:"required"`
	Value  string `json:"value" binding:"required"`
}
