package models


type User struct {
	Id        		int    	  `json:"_id" bson:"_id"`
	Email     		string    `json:"email" bson:"email" binding:"required"`
	Password  		string    `json:"password" bson:"password" binding:"required"`
	Phone     		string    `json:"phone" bson:"phone" binding:"required"`
	Login     		string    `json:"login" bson:"login" binding:"required"`
	Balance   		int	   	  `json:"balance" bson:"balance"`
	CryptoBalance 	int		  `json:"crypto_balance" bson:"crypto_balance"`
	UsingPromo		string    `json:"using_promo" bson:"using_promo"`
	Documents 		UserDocs  `json:"documents" bson:"documents"`
	PromoCode 		UserPromo `json:"promo_code" bson:"promo_code"`
} 

type UserDocs struct {
	DriverLicense 		string `json:"driver_license"`
	IdCard        		string `json:"id_card"`
	Passport      		string `json:"passport"`
	Sex           		int8   `json:"sex"`
	BirthDate     		string `json:"birth_date"`
	DocsNumber	  		string `json:"docs_number"`
	IdentificationCode	string `json:"identification_code"`
	PassportGiver		string `json:"passport_giver"`
}

type UserPromo struct {
	Name 	string `json:"name" bson:"name"`
	Count 	int32  `json:"count" bson:"count"`
}