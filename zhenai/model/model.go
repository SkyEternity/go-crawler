package model

//City 城市struct
type City struct {
	Name string
	URL  string
}

//Person 个人信息
type Person struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Salary    string `json:"salary"`
	Education string `json:"education"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Introduce string `json:"introduce"`
	Cover     string `json:"cover"`
}
