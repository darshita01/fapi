package user

type User struct {
	Id         int    `json:"id"`
	UserFname string `json:"fname"`
	UserLname string `json:"lname"`
	UserMname string `json:"mname"`
}
type JsonResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Data    []User `json:"data"`
	Addresses []Addresses `json:"Address"`
}
type Addresses struct{
	Id int `json:"id"`
	HouseNo string `json:"houseNo"`
	StreetName  string  `json:"streetName"`
	Pincode string `json:"pincode"`
	State string `json:"state"`
	Country string `json:"country"`

}
