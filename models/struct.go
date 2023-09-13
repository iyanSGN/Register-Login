package models

type MasterUser struct {
	Id      		int32 			`gorm:"primarykey" json:"id"`
	DepartmentId	*int32			`json:"department_id"`
	Name			string			`gorm:"type:varchar(255)" json:"name"`
	Email			string			`gorm:"type:varchar(255)" json:"email"`
	Password		string			`gorm:"type:varchar(255)" json:"password"`
}

type MasterDepartment struct {
	Id 				int32			`gorm:"primarykey" json:"id"`
	ParentId		int32			`json:"parid"`
	Name			string			`gorm:"type:varchar(255)" json:"name"`
}

func (MasterUser) TableName() string {
	return "master_user"
}

func (MasterDepartment) TableName() string {
	return "master_department"
}