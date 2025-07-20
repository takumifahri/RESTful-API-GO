package models 

type Roles string


type User struct {
	ID	   		uint   `json:"id" gorm:"primaryKey"`
	UniqueID 	string `json:"unique_id" gorm:"unique;not null;size:255"`
	Name 		string `json:"name" gorm:"unique;not null;size:255"`
	Hash		string `json:"-"`
	Password 	string `json:"password" gorm:"not null;size:255"`
	Email 		string `json:"email" gorm:"unique;not null;size:255"`
	Phone 		int64 	`json:"phone" gorm:"unique;not null;size:255"`
	Address 	string `json:"address" gorm:"not null;size:255"`
	Roles		Roles  `json:"roles" gorm:"not null;size:50;default:'user'"`
}

type RegisterRequest struct {
	Name		string `json:"name"`
	Password	string `json:"password"`
	// Password	string 	
}