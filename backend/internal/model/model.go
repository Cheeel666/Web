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

// Response to user
type Response struct {
	Status string `json:"status"`
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

// Comment of courort
type Comment struct {
	ID        int    `json:"comment_id"`
	Email     string `json:"email"`
	Text      string `json:"text"`
	IDCourort int    `json:"id_courort"`
}
