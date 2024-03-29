package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description"`
}

type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=2"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
