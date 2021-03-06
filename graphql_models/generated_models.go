// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_models

type BooleanFilter struct {
	EqualTo    *bool `json:"equalTo"`
	NotEqualTo *bool `json:"notEqualTo"`
}

type FloatFilter struct {
	EqualTo           *float64  `json:"equalTo"`
	NotEqualTo        *float64  `json:"notEqualTo"`
	LessThan          *float64  `json:"lessThan"`
	LessThanOrEqualTo *float64  `json:"lessThanOrEqualTo"`
	MoreThan          *float64  `json:"moreThan"`
	MoreThanOrEqualTo *float64  `json:"moreThanOrEqualTo"`
	In                []float64 `json:"in"`
	NotIn             []float64 `json:"notIn"`
}

type IDFilter struct {
	EqualTo    *string  `json:"equalTo"`
	NotEqualTo *string  `json:"notEqualTo"`
	In         []string `json:"in"`
	NotIn      []string `json:"notIn"`
}

type IntFilter struct {
	EqualTo           *int  `json:"equalTo"`
	NotEqualTo        *int  `json:"notEqualTo"`
	LessThan          *int  `json:"lessThan"`
	LessThanOrEqualTo *int  `json:"lessThanOrEqualTo"`
	MoreThan          *int  `json:"moreThan"`
	MoreThanOrEqualTo *int  `json:"moreThanOrEqualTo"`
	In                []int `json:"in"`
	NotIn             []int `json:"notIn"`
}

type Pet struct {
	ID        string  `json:"id"`
	Owner     *User   `json:"owner"`
	Name      *string `json:"name"`
	Breed     *string `json:"breed"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
	Toys      []*Toy  `json:"toys"`
}

type PetCreateInput struct {
	OwnerID   *string `json:"ownerId"`
	Name      *string `json:"name"`
	Breed     *string `json:"breed"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
}

type PetDeletePayload struct {
	ID string `json:"id"`
}

type PetFilter struct {
	Search *string   `json:"search"`
	Where  *PetWhere `json:"where"`
}

type PetPayload struct {
	Pet *Pet `json:"pet"`
}

type PetUpdateInput struct {
	OwnerID   *string `json:"ownerId"`
	Name      *string `json:"name"`
	Breed     *string `json:"breed"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
}

type PetWhere struct {
	ID        *IDFilter     `json:"id"`
	Owner     *UserWhere    `json:"owner"`
	Name      *StringFilter `json:"name"`
	Breed     *StringFilter `json:"breed"`
	CreatedAt *IntFilter    `json:"createdAt"`
	UpdatedAt *IntFilter    `json:"updatedAt"`
	DeletedAt *IntFilter    `json:"deletedAt"`
	Toys      *ToyWhere     `json:"toys"`
	Or        *PetWhere     `json:"or"`
	And       *PetWhere     `json:"and"`
}

type PetsCreateInput struct {
	Pets []*PetCreateInput `json:"pets"`
}

type PetsDeletePayload struct {
	Ids []string `json:"ids"`
}

type PetsPayload struct {
	Pets []*Pet `json:"pets"`
}

type PetsUpdatePayload struct {
	Ok bool `json:"ok"`
}

type StringFilter struct {
	EqualTo            *string  `json:"equalTo"`
	NotEqualTo         *string  `json:"notEqualTo"`
	In                 []string `json:"in"`
	NotIn              []string `json:"notIn"`
	StartWith          *string  `json:"startWith"`
	NotStartWith       *string  `json:"notStartWith"`
	EndWith            *string  `json:"endWith"`
	NotEndWith         *string  `json:"notEndWith"`
	Contain            *string  `json:"contain"`
	NotContain         *string  `json:"notContain"`
	StartWithStrict    *string  `json:"startWithStrict"`
	NotStartWithStrict *string  `json:"notStartWithStrict"`
	EndWithStrict      *string  `json:"endWithStrict"`
	NotEndWithStrict   *string  `json:"notEndWithStrict"`
	ContainStrict      *string  `json:"containStrict"`
	NotContainStrict   *string  `json:"notContainStrict"`
}

type Toy struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Color       *string `json:"color"`
	Pet         *Pet    `json:"pet"`
	CreatedAt   *int    `json:"createdAt"`
	DeletedAt   *int    `json:"deletedAt"`
	UpdatedAt   *int    `json:"updatedAt"`
}

type ToyCreateInput struct {
	Description string  `json:"description"`
	Color       *string `json:"color"`
	PetID       *string `json:"petId"`
	CreatedAt   *int    `json:"createdAt"`
	DeletedAt   *int    `json:"deletedAt"`
	UpdatedAt   *int    `json:"updatedAt"`
}

type ToyDeletePayload struct {
	ID string `json:"id"`
}

type ToyFilter struct {
	Search *string   `json:"search"`
	Where  *ToyWhere `json:"where"`
}

type ToyPayload struct {
	Toy *Toy `json:"toy"`
}

type ToyUpdateInput struct {
	Description *string `json:"description"`
	Color       *string `json:"color"`
	PetID       *string `json:"petId"`
	CreatedAt   *int    `json:"createdAt"`
	DeletedAt   *int    `json:"deletedAt"`
	UpdatedAt   *int    `json:"updatedAt"`
}

type ToyWhere struct {
	ID          *IDFilter     `json:"id"`
	Description *StringFilter `json:"description"`
	Color       *StringFilter `json:"color"`
	Pet         *PetWhere     `json:"pet"`
	CreatedAt   *IntFilter    `json:"createdAt"`
	DeletedAt   *IntFilter    `json:"deletedAt"`
	UpdatedAt   *IntFilter    `json:"updatedAt"`
	Or          *ToyWhere     `json:"or"`
	And         *ToyWhere     `json:"and"`
}

type ToysCreateInput struct {
	Toys []*ToyCreateInput `json:"toys"`
}

type ToysDeletePayload struct {
	Ids []string `json:"ids"`
}

type ToysPayload struct {
	Toys []*Toy `json:"toys"`
}

type ToysUpdatePayload struct {
	Ok bool `json:"ok"`
}

type User struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	Name      *string `json:"name"`
	Password  string  `json:"password"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	OwnerPets []*Pet  `json:"ownerPets"`
}

type UserCreateInput struct {
	Email     string  `json:"email"`
	Name      *string `json:"name"`
	Password  string  `json:"password"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
}

type UserDeletePayload struct {
	ID string `json:"id"`
}

type UserFilter struct {
	Search *string    `json:"search"`
	Where  *UserWhere `json:"where"`
}

type UserPayload struct {
	User *User `json:"user"`
}

type UserUpdateInput struct {
	Email     *string `json:"email"`
	Name      *string `json:"name"`
	Password  *string `json:"password"`
	DeletedAt *int    `json:"deletedAt"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
}

type UserWhere struct {
	ID        *IDFilter     `json:"id"`
	Email     *StringFilter `json:"email"`
	Name      *StringFilter `json:"name"`
	Password  *StringFilter `json:"password"`
	DeletedAt *IntFilter    `json:"deletedAt"`
	CreatedAt *IntFilter    `json:"createdAt"`
	UpdatedAt *IntFilter    `json:"updatedAt"`
	OwnerPets *PetWhere     `json:"ownerPets"`
	Or        *UserWhere    `json:"or"`
	And       *UserWhere    `json:"and"`
}

type UsersCreateInput struct {
	Users []*UserCreateInput `json:"users"`
}

type UsersDeletePayload struct {
	Ids []string `json:"ids"`
}

type UsersPayload struct {
	Users []*User `json:"users"`
}

type UsersUpdatePayload struct {
	Ok bool `json:"ok"`
}
