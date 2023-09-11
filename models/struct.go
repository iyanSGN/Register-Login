package models

type MasterUser struct {
	Id      		int32 			`gorm:"primarykey" json:"id"`
	Name			string			`gorm:"type:varchar(255)" json:"name"`
	Email			string			`gorm:"type:varchar(255)" json:"email"`
	Password		string			`gorm:"type:varchar(255)" json:"password"`
}

func (MasterUser) TableName() string {
	return "master_user"
}