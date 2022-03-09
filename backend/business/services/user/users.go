package user

import (
	"backend/additionalFunc"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Userstruct struct{
	db *sqlx.DB
}

func New(db *sqlx.DB) Userstruct{
	return Userstruct{
		db: db,
	}
}

func (h Userstruct) EnterUserData (fName, lName, mName string) JsonResponse {
	

	var response = JsonResponse{}

	if fName == "" || lName == "" || mName == "" {
		response = JsonResponse{Type: "error", Message: "You have missed some field"}
	} else {

		additionalFunc.PrintMessage("Inserting data...")

		var lastInsertID int
		err := h.db.QueryRow("INSERT INTO users(F_name, L_name, M_name) VALUES($1,$2,$3) returning id;", fName, lName, mName).Scan(&lastInsertID)
		additionalFunc.CheckErr(err)
		data1 := []User{}
		data1 = append(data1, User{Id: lastInsertID, F_name: fName, L_name: lName, M_name: mName})
		response = JsonResponse{Type: "success", Data: data1, Message: "Data inserted in User table"}
	}
	return response
}

func (h Userstruct) EditUserData(fName, lName, mName string, id string) JsonResponse{
	additionalFunc.PrintMessage("Editing data...")
	
	if fName != "" {
		err := h.db.QueryRow("UPDATE users SET F_name = $1 WHERE id = $2;", fName, id).Err()
		additionalFunc.CheckErr(err)
	}
	if lName != "" {
		err := h.db.QueryRow("UPDATE users SET L_name = $1 WHERE id = $2;", lName, id).Err()
		additionalFunc.CheckErr(err)
	}
	if mName != "" {
		err := h.db.QueryRow("UPDATE users SET M_name = $1 WHERE id = $2;", mName, id).Err()
		additionalFunc.CheckErr(err)
	}
	var id1 int
	var f_name string
	var l_name string
	var m_name string
	err := h.db.QueryRow("SELECT * FROM users WHERE id = $1;", id).Scan(&id1, &f_name, &l_name, &m_name)
	additionalFunc.CheckErr(err)
	iD,_ := strconv.Atoi(id)
	if iD != id1 {
		response := JsonResponse{Type: "failure", Message: "id is not in table"}
		return response
	}
	data1 := []User{}
	data1 = append(data1, User{Id: id1, F_name: f_name, L_name: l_name, M_name: m_name})

	response := JsonResponse{Type: "success", Data: data1, Message: "Data edited in User table"}
	return response
}

func (h Userstruct) DeleteUserData (id string) (JsonResponse){

	response := JsonResponse{}
	var id1 int
    err := h.db.QueryRow("SELECT id FROM users WHERE id = $1",id).Scan(&id1)
	error:= additionalFunc.CheckErr(err)
	if error != nil{
		response = JsonResponse{Type: "failure", Message: error.Error()}
		return response
	}
	if id1 == 0 {
		response = JsonResponse{Type: "failure", Message: "id is not in the table"}
		return response
	}

	rows, err1 := h.db.Query(`SELECT add_id FROM user_add WHERE user_id = $1;`, id)
	additionalFunc.CheckErr(err1)
	flag := 0
	for rows.Next() {
		var aId sql.NullInt64
		err := rows.Scan(&aId)
		additionalFunc.CheckErr(err)
		if aId.Valid {
			flag = 1
			response = JsonResponse{Type: "failure", Message: "Data can't delete because it is connected with Addresses."}
			break
		}

	}
	if flag == 0 {
		additionalFunc.PrintMessage("Deleting data...")
		_, err := h.db.Exec("DELETE FROM user_add WHERE id = $1;", id)
		additionalFunc.CheckErr(err)
		_, err1 := h.db.Exec("DELETE FROM users WHERE id = $1;", id)
		additionalFunc.CheckErr(err1)
		response = JsonResponse{Type: "success", Message: "Data deleted in User table"}
	}
	return response
}

func (h Userstruct) GetUserData (id string) (JsonResponse) {

	var response JsonResponse
	fmt.Println("Serving Get Request")

	str := "SELECT * FROM users WHERE $1 = $1 ORDER BY id;"

	userSlice := make([]User, 0)
	
	addressSlice := make([]Addresses, 0)
	
	if id != "" {
		
		// temp id print
		fmt.Println(id)
		
		str = "SELECT * FROM users WHERE id = $1"
		
		rows, err := h.db.Queryx("select * from address where id in (select add_id from user_add where user_id  = $1)", id)
		additionalFunc.CheckErr(err)

		defer rows.Close()
		
		for rows.Next() {
			var tempAddressStruct Addresses

			err := rows.StructScan(&tempAddressStruct)

			additionalFunc.CheckErr(err)

			addressSlice = append(addressSlice, tempAddressStruct)
		}
	}

	rows, err := h.db.Queryx(str, id)
	additionalFunc.CheckErr(err)
	
	
	for rows.Next() {
		var tempUserStruct User
		err := rows.StructScan(&tempUserStruct)
		additionalFunc.CheckErr(err)

		userSlice = append(userSlice, tempUserStruct)
	}

	defer rows.Close()

	response = JsonResponse{Type: "success", Data: userSlice, Message: "Get data successfully", Addresses: addressSlice}
	return response

}


























































