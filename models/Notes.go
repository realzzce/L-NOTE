package models

import (
	"time"
)

// NotesModel .
type NotesModel struct {
	NoteID int       `json:"NoteID" gorm:"column:NoteID;primary_key;auto_increment;"`
	Note   string    `json:"Note" gorm:"column:Note;"`
	Time   time.Time `json:"Time" gorm:"column:Time;"`
	UserID int       `json:"UserID" gorm:"column:UserID;"`
}

// TableName NotesModel 表
func (NotesModel) TableName() string {
	return "NotesModel"
}

// GetNote .
func GetNote(noteid string) (note NotesModel) {
	if err := GlobalDB.Where("NoteID = ?", noteid).First(&note).Error; err == nil {
		return note
	}
	return note
}

// AddNote .
func AddNote(notestr string, userid int) bool {
	var note NotesModel
	note.Note = notestr
	note.Time = time.Now().UTC()
	note.UserID = userid
	if err := GlobalDB.Create(&note).Error; err == nil {
		return true
	}
	return false
}

// UpdateNote .
func UpdateNote(notestr string, noteid int) bool {
	var note NotesModel
	if err := GlobalDB.Model(&note).Where("NoteID = ?", noteid).Update(
		map[string]interface{}{
			"Note": notestr}).Error; err == nil {
		return true
	}
	return false
}

// DeleteNote 删除话题
func DeleteNote(noteid int) bool {
	var note NotesModel
	if err := GlobalDB.
		Where("NoteID = ?", noteid).
		Delete(&note).Error; err == nil {
		return true
	}
	return false
}
