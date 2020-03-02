package model

type User struct {
	ID   uint   `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `json:"name" gorm:"type:varchar(8);NOT NULL"`
	Job string `json:"job" gorm:"varchar(50);NOT NULL"`
	Role string `json:"role" gomr:"varchar(50);NOT NULL'"`
	Mobile string `json:"mobile" gorm:"type:char(11);NOT NULL"`
	Password string `json:"password" grom:"type:varchar(128)"`
	Sex  bool   `json:"sex"`
}
