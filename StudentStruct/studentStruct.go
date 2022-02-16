package studentstruct

type Person struct {
ID        int `json:"id"`
FirstName string `json:"firstname"`
LastName  string `json:"lastname"`
Age       int `json:"age"`
}

type Total struct{
	Total int `json:"total"`
}