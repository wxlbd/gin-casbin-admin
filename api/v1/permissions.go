package v1

type PermissionsAddRequest struct {
	Name   string `json:"name" binding:"required"`
	Icon   string `json:"icon"`
	Path   string `json:"path"`
	Url    string `json:"url" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Hidden int    `json:"hidden" `
	IsMenu int    `json:"is_menu"`
	PId    int    `json:"pid"`
	Method string `json:"method" binding:"required,oneof=GET POST PUT DELETE * PATCH"`
	Status int    `json:"status"`
}

type PermissionsUpdateRequest struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Path   string `json:"path"`
	Url    string `json:"url"`
	Title  string `json:"title"`
	Hidden int    `json:"hidden"`
	IsMenu int    `json:"is_menu"`
	PId    int    `json:"p_id"`
	Method string `json:"method"`
	Status int    `json:"status"`
}

type PermissionsGetRequest struct {
	Id int `json:"id"`
}

type PermissionsDeleteRequest struct {
	Id int `json:"id"`
}

type PermissionTreeRequest struct {
}

type Permissions struct {
	Id       int            `json:"id"`
	Name     string         `json:"name"`
	Icon     string         `json:"icon"`
	Path     string         `json:"path"`
	Url      string         `json:"url"`
	Title    string         `json:"title"`
	Hidden   int            `json:"hidden"`
	IsMenu   int            `json:"is_menu"`
	Pid      int            `json:"pid"`
	Method   string         `json:"method"`
	Status   int            `json:"status"`
	Children []*Permissions `json:"children"`
}

type PermissionsTreeResponse struct {
	List []*Permissions `json:"list"`
}

type PermissionsGetResponse struct {
	Permissions
}
