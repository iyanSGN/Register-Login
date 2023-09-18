package models

type MasterUser struct {
	Id      			uint 				`gorm:"primarykey" json:"id"`
	DepartmentId		uint				`json:"department_id"`
	Name				string				`gorm:"type:varchar(255)" json:"name"`
	Email				string				`gorm:"unique;type:varchar(255)" json:"email"`
	Password			string				`gorm:"type:varchar(255)" json:"password"`
	MasterDepartment	MasterDepartment	`gorm:"foreignKey:DepartmentId; constraint:OnDelete:CASCADE;"` 
}

type MasterDepartment struct {
	Id 				uint			`gorm:"primarykey" json:"id"`
	ParentId		uint			`json:"parent_id"`
	Name			string			`gorm:"type:varchar(255)" json:"name"`
}

func (MasterUser) TableName() string {
	return "master_user"
}

func (MasterDepartment) TableName() string {
	return "master_department"
}