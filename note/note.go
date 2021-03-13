package note

import (
	"net/http"
	"other/L-NOTE/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetNote .
func GetNote(c *gin.Context) {
	NoteID := c.Query("noteid")
	note := models.GetNote(NoteID)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    note,
	})
}

// AddNote .
func AddNote(c *gin.Context) {
	Notestr := c.Query("notestr")
	userid := c.Query("userid")
	UserID, errA := strconv.Atoi(userid)
	if errA != nil || UserID <= 0 {
		UserID = 1
	}
	status := models.AddNote(Notestr, UserID)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    status,
	})
}

// UpdateNote .
func UpdateNote(c *gin.Context) {
	Notestr := c.Query("notestr")
	noteid := c.Query("noteid")
	NoteID, errA := strconv.Atoi(noteid)
	if errA != nil || NoteID <= 0 {
		NoteID = 1
	}
	status := models.UpdateNote(Notestr, NoteID)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    status,
	})
}

// DeleteNote .
func DeleteNote(c *gin.Context) {
	noteid := c.Query("noteid")
	NoteID, errA := strconv.Atoi(noteid)
	if errA != nil || NoteID <= 0 {
		NoteID = 1
	}
	status := models.DeleteNote(NoteID)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    status,
	})
}
