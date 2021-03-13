package models

// UserModel .
type UserModel struct {
	UserID     int    `json:"UserID" gorm:"column:UserID;primary_key;auto_increment;"`
	Username   string `json:"Username" gorm:"column:Username;" binding:"required"`
	Password   string `json:"-" gorm:"column:Password;" binding:"required"`
	RoleID     int    `json:"RoleID" gorm:"column:RoleID;default:10;"`
	Text       string `json:"Text" gorm:"column:Text;"`
	UserClaims []ClaimsModel
}

// TableName UserModel 表
func (UserModel) TableName() string {
	return "UserModel"
}

// CheckUser .
func CheckUser(username, password string) bool {
	var user UserModel
	GlobalDB.Select([]string{"UserID", "Username", "RoleID", "Text"}).Where("username = ? AND password = ?", username, password).First(&user)
	if user.RoleID <= 10 {
		return true
	}
	return false
}
