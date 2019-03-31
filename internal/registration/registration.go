package registration

type AuthData struct {
	Login    string `form:"login" json:"login" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LectorData struct {
	Id       int    `json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Surname  string `form:"surname" json:"surname" binding:"required"`
	Login    string `form:"login" json:"login" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type StudentData struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Surname string `form:"surname" json:"surname" binding:"required"`
	Group   string `form:"group" json:"group" binding:"required"`
}

type SubjectData struct {
	LectorID int    `form:"lector" json:"lector" binding:"required"`
	Title    string `form:"title" json:"title" binding:"required"`
}

type DeviceData struct {
	Room      string `form:"room" binding:"required"`
	MACAdress string `form:"adress" binding:"required"`
}

type LectorDataEdit struct {
	Id       int    `json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Surname  string `form:"surname" json:"surname" binding:"required"`
	Login    string `form:"login" json:"login" binding:"required"`
	Password string `form:"password" json:"password"`
}
