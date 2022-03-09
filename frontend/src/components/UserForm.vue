<template>
        <form class="form" v-if="isEditClicked || isAddClicked">
      <div>
        <label>First Name:</label> &nbsp;&thinsp;&thinsp;
        <input  v-model="fName" id="fName" name="F_name" type="text" />
      </div>
      <div>
        <label>Last Name:</label>&nbsp;&thinsp;&thinsp;&thinsp;&thinsp;
        <input v-model="this.$props.lname" id="lName" name="L_name" type="text" />
      </div>
      <div>
        <label>Middle Name:</label>
        <input v-model="this.$props.mname" id="mName" name="M_name" type="text" />
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
</template>
<script>
import service from "/src/services/services.js";
const services = new service();
export default{
    
data(){
        return{
            fName : this.$props.i.fname
        };
},
props: {
    isAddClicked: Boolean,
    isformbtn: Boolean,
    isEditClicked: Boolean,
    // fname: String,
    // lname: String,
    // mname: String,
    i : Object ,
},
methods: {
    checkValidity(event) {
      event.preventDefault();
    //   let x = /[0-9]/;
    //   let y = /[^a-zA-Z]+/;

    //   if (this.fName == "" || this.lName == "" || this.mName == "") {
    //     alert("Fill all the field.");
    //     return false;
    //   } else if (
    //     x.test(this.fName) ||
    //     x.test(this.lName) ||
    //     x.test(this.mName)
    //   ) {
    //     alert(
    //       "Numeric value or numbers in between are not allowed for these field."
    //     );
    //     return false;
    //   } else if (
    //     y.test(this.fName) ||
    //     y.test(this.lName) ||
    //     y.test(this.mName)
    //   ) {
    //     alert("Special characters are not allowed.");
    //     return false;
    //   } else {
    //     this.editdata();
    //   }
    },
    async editdata() {
      this.bodyFormData = new FormData();

      this.bodyFormData.append("F_name", this.fName);
      this.bodyFormData.append("M_name", this.mName);
      this.bodyFormData.append("L_name", this.lName);
      if (this.isEditClicked) {
          await services.editUser(this.bodyFormData,this.id).then(res => res)
        // await this.axios({
        //   method: "put",
        //   url: "http://localhost:8004/Users/" + this.id.id,
        //   data: this.bodyFormData,
        //   headers: { "Content-Type": "multipart/form-data" },
        // })
        //   .then(function (response) {
        //     console.log(response);
        //   })
        //   .catch(function (err) {
        //     console.log(err);
        //   });
      } else if (this.isAddClicked) {
        await this.axios({
          method: "post",
          url: "http://localhost:8004/Users/",
          data: this.bodyFormData,
          headers: { "Content-Type": "multipart/form-data" },
        })
          .then(function (response) {
            console.log(response);
          })
          .catch(function (err) {
            console.log(err);
          });
      }
      this.getUsers();
    },
}
}
</script>
<style>
.form {
  display: grid;
  width: 25%;
  font-size: medium;
  margin-bottom: 20px;
  padding: 10px;
  line-height: 25px;
}
input {
  margin: 5px;
}
</style>
