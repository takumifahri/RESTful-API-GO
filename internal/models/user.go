package models 

type Roles string


type User struct {
	ID	   		uint   `json:"id" gorm:"primaryKey"`
	UniqueID 	string `json:"unique_id" gorm:"unique;not null;size:255"`
	Name 		string `json:"name" gorm:"unique;not null;size:255"`
	Hash		string `json:"-"`
	Email 		string `json:"email" gorm:"unique;not null;size:255"`
	Phone 		int64 	`json:"phone" gorm:"unique;not null;size:255"`
	Address 	string `json:"address" gorm:"not null;size:255"`
	Roles		Roles  `json:"roles" gorm:"not null;size:50;default:'user'"`
}

type RegisterRequest struct {
    Name		string `json:"name" validate:"required,min=3,max=50"`
    Password	string `json:"password" validate:"required,min=6"`  // ✅ Plaintext dari user
    Email		string `json:"email" validate:"required,email"`
    Phone		int64  `json:"phone" validate:"required"`
    Address		string `json:"address" validate:"required,min=10"`
	// Password	string 	
}