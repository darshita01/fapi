<template>
  <div class="main">
    <div class="tableName">
      <h3>&nbsp; List of Users</h3>
      <button class="btn btn-primary btn-sm" id="add" @click="addUser()">Add</button>
    </div>
    <div class="container-fluid">
      <div class="row" id="header">
        <div class="col-sm-1"><h4>Id</h4></div>
        <div class="col-sm-7"><h4>Name</h4></div>
        <div class="col-sm-4"><h4>Actions</h4></div>
      </div>
      <div class="row" v-for="(i, index) in record" :key="index">
        <div class="col-sm-1">
          {{ i.id }}
        </div>
        <div class="col-sm-7">{{ i.fname }} {{ i.mname }} {{ i.lname }}</div>
        <div class="col-sm-4" style="align-content: center">
          <button class="btn btn-success btn-sm" @click="getUser(i)">
            Show
          </button>
          <button class="btn btn-primary btn-sm" @click="editUser(i)">
            Edit
          </button>
          <button class="btn btn-danger btn-sm" @click="deleteUser(i)">
            Delete
          </button>
        </div>
      </div>
    </div>
    <div id="address" v-if="isBtnClicked">
      <div>
        {{ address }}
        <button id="close" @click="closeAddDiv()">Close</button>
      </div>
    </div>
    <form v-if="isEditClicked || isAddClicked">
      <div>
        <label>First Name:</label>
        <input v-model="fname" id="fname" name="F_name" type="text" />
      </div>
      <div>
        <label>Last Name:</label>
        <input v-model="lname" id="lname" name="L_name" type="text" />
      </div>
      <div>
        <label>Middle Name:</label>
        <input v-model="mname" id="mname" name="M_name" type="text" />
      </div>
      <div>
        <button
          v-if="isformbtn"
          type="submit"
          id="Post"
          @click="checkValidity($event)"
        >
          Edit
        </button>
        <button
          v-if="!isformbtn"
          type="submit"
          id="Post"
          @click="checkValidity($event)"
        >
          Add
        </button>
      </div>
    </form>
  </div>
</template>
<script>
export default {
  name: "App",
  data() {
    return {
      record: "",
      address: null,
      isBtnClicked: false,
      isEditClicked: false,
      isAddClicked: false,
      fname: "",
      lname: "",
      mname: "",
      id: "",
      isformbtn: false,
      msg: "",
      // "user1",
      // "user2",
      // "u
    };
  },
  mounted() {
    this.getUsers();
  },

  methods: {
    getUsers() {
      this.axios
        .get("http://localhost:8000/Users/")
        .then((res) => {
          this.record = res.data.data;
          console.log(res);
        })
        .catch((err) => {
          console.log(err);
        });
      console.log(this.fname);
    },

    getUser(i) {
      this.axios
        .get("http://localhost:8000/Users/" + i.id)
        .then((res) => {
          var address = [res.data.Address];
          console.log(address);
          this.isBtnClicked = true;
          this.address = address;
        })
        .catch((err) => {
          console.log(err);
        });
      this.isBtnClicked = false;
    },
    async editdata() {
      this.bodyFormData = new FormData();

      this.bodyFormData.append("F_name", this.fname);
      this.bodyFormData.append("M_name", this.mname);
      this.bodyFormData.append("L_name", this.lname);
      if (this.isEditClicked) {
        await this.axios({
          method: "put",
          url: "http://localhost:8000/Users/" + this.id.id,
          data: this.bodyFormData,
          headers: { "Content-Type": "multipart/form-data" },
        })
          .then(function (response) {
            //handle success
            console.log(response);
          })
          .catch(function (err) {
            //handle error
            console.log(err);
          });
        this.isEditClicked = false;
      } else if(this.isAddClicked){
        await this.axios({
          method: "post",
          url: "http://localhost:8000/Users/",
          data: this.bodyFormData,
          headers: { "Content-Type": "multipart/form-data" },
        })
          .then(function (response) {
            //handle success
            console.log(response);
          })
          .catch(function (err) {
            console.log(err);
          });
        this.isAddClicked = false;
      }
      this.getUsers();
    },
    editUser(i) {
      this.isEditClicked = true;
      this.isformbtn = true;
      this.id = i;
      this.fname = i.fname;
      this.lname = i.lname;
      this.mname = i.mname;
    },
    addUser(i) {
      this.isAddClicked = true;
      this.isformbtn = false;
      this.id = i;
      this.fname = "";
      this.lname = "";
      this.mname = "";
    },
    async deleteUser(i) {
      await this.axios
        .delete("http://localhost:8000/Users/" + i.id)
        .then((res) => {
          this.msg = res.data.message;
          console.log(res);
        })
        .catch((err) => {
          console.log(err);
        });
        alert(this.msg)

      this.getUsers();
    },
    closeAddDiv() {
      document.getElementById("address").innerHTML = "";
    },
    checkValidity(event){
      event.preventDefault()
      let x = /[0-9]/
      let y = /[`!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?~]/

      if(this.fname == "" || this.lname == "" || this.mname == ""){
        alert("Fill all the field.")
        return false
      }else if(x.test(this.fname) || x.test(this.lname) || x.test(this.mname)){
        alert("Numeric value or numbers in between are not allowed for these field.")
        return false
      }else if(y.test(this.fname) || y.test(this.lname) || y.test(this.mname)){
        alert("Special characters are not allowed.")
        return false
      } else {
        this.editdata()
      }
    }
  },
  components: {},
};
</script>

<style>
.row {
  margin: 1%;
  padding-top: 1%;
  border-top: 0.5px solid #d9dbdb;
}
.btn {
  margin-left: 1%;
  margin-right: 1%;
}
.tableName {
  margin: 1%;
  display: flex;
  gap: 2%;
}
.main {
  width: 90%;
  margin: auto;
  margin-top: 2%;
  background-color: #ffffff;
  border: 1px solid #e2e6e6;
  border-radius: 2px;
  box-shadow: 5px 5px #f1f1f1;
}
#add{
  background-color: #8d8a8a;
  border: #8d8a8a;
  padding: 0% 5%;
}
</style>
