package models

// ClaimsModel .
type ClaimsModel struct {
	ClaimsID int    `json:"ClaimsID" gorm:"column:ClaimsID;primary_key;auto_increment;"`
	UserID   int    `json:"UserID" gorm:"column:UserID;" binding:"required"`
	Type     int    `json:"Type" gorm:"column:Type;"`
	Value    string `json:"Value" gorm:"column:Value;"`
}

// TableName ClaimsModel 表
func (ClaimsModel) TableName() string {
	return "ClaimsModel"
}

// GetUserClaims .
func GetUserClaims(username string) (claims []ClaimsModel) {
	var user UserModel
	GlobalDB.Select([]string{"UserID", "Username"}).Where("Username = ?", username).First(&user)
	GlobalDB.Select([]string{"ClaimsID", "UserID", "Type", "Value"}).Where("UserID = ?", user.UserID).First(&claims)

	return
}
