package model

//City 城市struct
type City struct {
	Name string
	URL  string
}

//Person 个人信息
type Person struct {
	Name      string
	City      string
	Salary    string
	Education string
	Age       int
	Height    int
	Introduce string
	Cover     string
}
