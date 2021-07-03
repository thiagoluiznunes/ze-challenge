package entity

type Partner struct {
	ID   string `json:"id,omitempty" bson:"_id"`
	Name string `json:"name,omitempty" bson:"name"`
}

func (p *Partner) validate() (err error) {

	return nil
}
