package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/wuttachaichumyen/student-portfolio-api/models" // แก้ YOUR_USERNAME ให้ตรงกับ go.mod
)

// จำลอง Database ด้วยตัวแปร Global (ในงานจริงเราจะต่อ SQL ตรงนี้)
var students = []models.Student{
	{
		ID:        "66001",
		FirstName: "Somsri",
		LastName:  "Jaidee",
		Email:     "somsri@spu.ac.th",
		Major:     "Computer Science",
		Projects: []models.Project{
			{Title: "AI Chatbot", Description: "Chatbot using OpenAI API", Link: "github.com/somsri/ai"},
		},
	},
}

// GetStudents: ดึงรายชื่อนักศึกษาทั้งหมด
func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

// CreateStudent: เพิ่มนักศึกษาใหม่
func CreateStudent(c *gin.Context) {
	var newStudent models.Student

	// ผูก JSON ที่ส่งมา เข้ากับ Struct
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เพิ่มลงใน Slice (จำลองการ save ลง DB)
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

// GetStudentByID: ค้นหานักศึกษาตาม ID
func GetStudentByID(c *gin.Context) {
	id := c.Param("id")

	for _, s := range students {
		if s.ID == id {
			c.JSON(http.StatusOK, s)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}