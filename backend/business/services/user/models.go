package user

type User struct {
	Id        int    `json:"id"`
	F_name string `json:"F_name"`
	L_name string `json:"L_name"`
	M_name string `json:"M_name"`
}
type JsonResponse struct {
	Type      string      //`json:"type"`
	Message   string      //`json:"message"`
	Data      []User      //`json:"data"`
	Addresses []Addresses //`json:"Address"`
}
type Addresses struct {
	Id         int    `json:"id"`
	HouseNo    string `json:"houseNo"`
	StreetName string `json:"streetName"`
	Pincode    int    `json:"pincode"`
	State      string `json:"state"`
	Country    string `json:"country"`
}