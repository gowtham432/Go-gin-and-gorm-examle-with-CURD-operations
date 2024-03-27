package models

type UserProfileEntity struct {
	UserId    int    `gorm:"column:user_id;primary_key;autoIncrement;"`
	FirstName string `gorm:"column:firstname" json:"firstName"`
	LastName  string `gorm:"column:lastname" json:"lastName"`
	UserName  string `gorm:"column:username" gorm:"unique;not null" json:"userName"`
	Email     string `gorm:"column:email" gorm:"unique;not null" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	Mobile    string `gorm:"column:mobile" gorm:"unique;not null" json:"mobile"`
}

func (UserProfileEntity) TableName() string {
	return "user_profile"
}
