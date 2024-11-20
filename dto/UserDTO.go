package dto

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDTO struct {
	Id        int    `json:"id"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      Role   `json:"role"`
	Password  string `json:"password"`
}
type Role struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Active      bool         `json:"active"`
	RoleModules []RoleModule `json:"role_module"`
}
type RoleModule struct {
	Id     int    `json:"id"`
	Active bool   `json:"active"`
	Module Module `json:"module"`
}
type Module struct {
	Id         uint        `json:"id"`
	Name       string      `json:"name"`
	Icon       string      `json:"icon"`
	Order      int         `json:"order"`
	Active     bool        `json:"active"`
	ModuleRole interface{} `json:"module_role"`
}
type ModuleRole struct {
	Id     uint   `json:"id"`
	Module Module `json:"module"`
	Active bool   `json:"active"`
}
