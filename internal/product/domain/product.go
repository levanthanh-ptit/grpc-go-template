package domain

type Product struct {
	ID   ID      `json:"id" bson:"_id,omitempty"`
	Name *string `json:"name,omitempty" bson:"name,omitempty"`
}
