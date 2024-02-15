package entities

type Entities struct {
	Entities interface{}
}

func RegisterEntities() []Entities {
	return []Entities{
		{Entities: Role{}},
		{Entities: User{}},
	}
}
