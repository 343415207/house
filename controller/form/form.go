package form

type Login struct {
	User     string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type PListForm struct {
	User string `form:"phone" binding:"required"`
}

type OrderForm struct {
	CName       string `form:"c_name"`
	CPhone      string `form:"c_phone"`
	CSex        int    `form:"c_sex"`
	PID         int64  `form:"p_id"`
	IsFirst     int    `form:"is_first"`
	EscorteType int    `form:"escorte_type"`
	AccessTime  string `form:"access_time"`
	Desc        string `form:"desc"`
}
