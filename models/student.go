package models

// Project คือโครงสร้างของผลงานนักศึกษา
type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// Student คือโครงสร้างข้อมูลนักศึกษา
type Student struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Major     string    `json:"major"`
	Projects  []Project `json:"projects"` // 1 คนมีหลายผลงานได้
}