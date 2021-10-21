package model

// User of service
type User struct {
	ID       int
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IDRole   int    `json:"id_role"`
	Role     string `json:"role"`
}

// Road in service
type Road struct {
	NameRoad    string `json:"name_road"`
	NameCourort string `json:"name_courort"`
}

// Courorts in service
type Courorts struct {
	NameCourort string `json:"name_courort"`
	City        string `json:"city"`
}

// Courort in courorrts
type Courort struct {
	TypeRoad   string `json:"type_road"`
	NameRoad   string `json:"name_road"`
	WorkStatus int    `json:"work_status"`
}
