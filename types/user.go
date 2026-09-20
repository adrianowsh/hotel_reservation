package types

type User struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Age       int    `bson:"age" json:"age"`
}
