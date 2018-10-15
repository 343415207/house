package form

type Login struct {
	User     string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type PListForm struct {
	User string `form:"phone" binding:"required"`
}

type OrderForm struct {
}
