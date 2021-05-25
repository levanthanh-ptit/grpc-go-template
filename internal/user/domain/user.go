package domain

type User struct {
	ID       ID      `json:"id" bson:"_id,omitempty"`
	Name     *string `json:"name,omitempty" bson:"name,omitempty"`
	Username string  `json:"username,omitempty" bson:"username,omitempty"`
	Password string  `json:"password,omitempty" bson:"password,omitempty"`
	Old      *int32  `json:"old,omitempty" bson:"old,omitempty"`
}
