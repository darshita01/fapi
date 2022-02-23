package business

import (
	"backend/additionalFunc"
	"backend/database"
	"database/sql"
)


type UserBusiness struct {
	db *sql.DB
}

func New (db *sql.DB) UserBusiness {
	return UserBusiness{
		db: db,
	}
}

func (h UserBusiness) EnterUserData (fName, lName, mName string) JsonResponse {
	

	var response = JsonResponse{}

	if fName == "" || lName == "" || mName == "" {
		response = JsonResponse{Type: "error", Message: "You have missed some field"}
	} else {

		additionalFunc.PrintMessage("Inserting data...")

		var lastInsertID int
		err := h.db.QueryRow(database.EnterUser, fName, lName, mName).Scan(&lastInsertID)
		additionalFunc.CheckErr(err)
		data1 := []User{}
		data1 = append(data1, User{Id: lastInsertID, User_fname: fName, User_lname: lName, User_mname: mName})
		response = JsonResponse{Type: "success", Data: data1, Message: "Data inserted in User table"}
	}
	return response
}

func (h UserBusiness) GetUserData (params map[string]string) (JsonResponse) {
	additionalFunc.PrintMessage("Getting data...")
	
	id := params["id"]
	str := "SELECT * FROM users WHERE $1 = $1 ORDER BY id;"
	var id1 int
	var f_name, l_name, m_name string
	var id2 int
	var houseNo, streetName, state, country, pincode sql.NullString

	data1 := []User{}
	address := []Addresses{}
	if id != "" {
		str = "SELECT * FROM Users WHERE id = $1"
		rows, err := h.db.Query("SELECT add_id FROM user_add WHERE user_id = $1",id)
		additionalFunc.CheckErr(err)
		for rows.Next() {
			var aId sql.NullInt64
			err := rows.Scan(&aId)
			additionalFunc.CheckErr(err)
 
			if aId.Valid{
				err1 := h.db.QueryRow("SELECT * FROM address WHERE id = $1",aId.Int64).Scan(&id2, &houseNo, &streetName, &pincode, &state, &country)
				additionalFunc.CheckErr(err1)
				address = append(address, Addresses{Id: id2, HouseNo: houseNo.String, StreetName: streetName.String, Pincode: pincode.String, State: state.String, Country: country.String})
			}
		}
		
	}
	rows3, err3 := h.db.Query(str, id)
	additionalFunc.CheckErr(err3)
	
	for rows3.Next(){
		rows3.Scan(&id1, &f_name, &l_name, &m_name)
	data1 = append(data1, User{Id: id1, User_fname: f_name, User_lname: l_name, User_mname: m_name})
	}

	response := JsonResponse{Type: "success", Data: data1, Message: "Get data successfully", Addresses: address}
	return response
}

func (h UserBusiness) EditUserData(fName, lName, mName string, params map[string]string) JsonResponse{
	additionalFunc.PrintMessage("Editing data...")
	
	id := params["id"]
	

	if fName != "" {
		err := h.db.QueryRow("UPDATE users SET \"F_name\" = $1 WHERE id = $2;", fName, id).Err()
		additionalFunc.CheckErr(err)
	}
	if lName != "" {
		err := h.db.QueryRow("UPDATE users SET \"L_name\" = $1 WHERE id = $2;", lName, id).Err()
		additionalFunc.CheckErr(err)
	}
	if mName != "" {
		err := h.db.QueryRow("UPDATE users SET \"M_name\" = $1 WHERE id = $2;", mName, id).Err()
		additionalFunc.CheckErr(err)
	}
	var id1 int
	var f_name string
	var l_name string
	var m_name string
	err := h.db.QueryRow("SELECT * FROM users WHERE id = $1;", id).Scan(&id1, &f_name, &l_name, &m_name)
	additionalFunc.CheckErr(err)

	data1 := []User{}
	data1 = append(data1, User{Id: id1, User_fname: f_name, User_lname: l_name, User_mname: m_name})

	response := JsonResponse{Type: "success", Data: data1, Message: "Data edited in User table"}
	return response
}

func (h UserBusiness) DeleteUserData (params map[string]string) (JsonResponse){
	id := params["id"]

	response := JsonResponse{}

	rows, err := h.db.Query(`SELECT add_id FROM user_add WHERE user_id = $1;`, id)
	additionalFunc.CheckErr(err)
	flag := 0
	for rows.Next() {
		var aId sql.NullInt64
		err = rows.Scan(&aId)
		additionalFunc.CheckErr(err)
		if aId.Valid {
			flag = 1
			response = JsonResponse{Type: "failure", Message: "Data can't delete because it is connected with Addresses."}
			break
		}

	}
	if flag == 0 {
		additionalFunc.PrintMessage("Deleting data...")
		_, err := h.db.Exec("DELETE FROM user_add WHERE \"id\" = $1;", id)
		additionalFunc.CheckErr(err)
		_, err1 := h.db.Exec("DELETE FROM users WHERE \"id\" = $1;", id)
		additionalFunc.CheckErr(err1)
		response = JsonResponse{Type: "success", Message: "Data deleted in User table"}
	}
	return response
}