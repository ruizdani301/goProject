package models


type Program struct {
	UID	string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Num1 int	`json:"num1,omitempty"`
	Num2 int	`json:"num2,omitempty"`
	Operation  []Operation `json:"operation,omitempty"` // operation apunta a type Operation
}

type Operation struct {
	UID	string `json:"uid,omitempty"`
	Add int	`json:"add,omitempty"`
	Sub int	`json:"sub,omitempty"`	
	Mult int `json:"mult,omitempty"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	StatusCode int `json:"status_code,omitempty"`
	Result	interface{} `json:"results,omitempty"`

}

type QueryResults struct {
	Response []Program
}
