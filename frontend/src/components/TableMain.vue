<template>
  <div>
    <div class="tableName">
      <h5>&nbsp; List of Users</h5>
      <button class="btn btn-primary btn-sm" id="add" @click="addUser()">
        Add
      </button>
    </div>
    <div class="container-fluid">
      <div class="row" id="header">
        <div class="col-sm-1"><h6>Id</h6></div>
        <div class="col-sm-7"><h6>Name</h6></div>
        <div class="col-sm-4"><h6>Actions</h6></div>
      </div>
      <div class="row" v-for="(i, index) in record" :key="index">
        <div class="col-sm-1">
          {{ i.id }}
        </div>
        <div class="col-sm-7">{{ i.fname }} {{ i.mname }} {{ i.lname }}</div>
        <div class="col-sm-4" style="align-content: center">
          <button
            class="btn btn-success btn-sm"
            @click="onShowSection(i), (j = i)"
          >
            Show
          </button>
          <button class="btn btn-primary btn-sm" @click="editUser(i)">
            Edit
          </button>
          <button class="btn btn-danger btn-sm" @click="deleteUser(i)">
            Delete
          </button>
        </div>
        <div v-if="openSections.indexOf(i.id) !== -1">
          <div class="cardgroup">
            <div v-for="(i, index) in i.address" :key="index" class="card">
              <div class="addressHeader">Address id : {{ i.id }}</div>
              <div class="addressText">
                {{ i.houseNo }}, {{ i.streetName }}, {{ i.state }},
                {{ i.country }}, {{ i.pincode }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <user-form :isAddClicked="isAddClicked" :isformbtn="isformbtn" :isEditClicked="isEditClicked" :i="i ? i : "></user-form>
  </div>
</template>

<script>
import UserForm from "./UserForm.vue";
import service from "/src/services/services.js";
const services = new service();
export default {
  components: { UserForm },
  data() {
    return {
      address: null,
      record: "",
      fname: "",
      lname: "",
      mname: "",
      i: null,
      isAddClicked: false,
      isEditClicked: false,
      isformbtn: false,
      openSections: [],
    };
  },
  mounted() {
    services.getUsers().then((res) => (this.record = res));
  },
  methods: {
    addUser() {
      this.isAddClicked = true;
      this.isformbtn = false;
    },
    editUser(i) {
      this.isEditClicked = true;
      this.isformbtn = true;
      this.i = i.i;
      // this.fname = i.fname;
      // this.lname = i.lname;
      // this.mname = i.mname;
    },
    async onShowSection(userInfo) {
      if (!userInfo.address) {
        try {
          await services
            .getUser(userInfo)
            .then((res) => (userInfo.address = res));
        } catch (error) {
          console.log(error);
        }
      }
      if (this.openSections.indexOf(userInfo.id) !== -1) {
        this.openSections = this.openSections.filter(
          (id) => id !== userInfo.id
        );
      } else if (userInfo.address.length === 0) {
        alert("address is not available");
      } else {
        this.openSections = [...this.openSections, userInfo.id];
      }
    },
    async deleteUser(i) {
      await services.deleteUser(i);
      services.getUsers().then((res) => (this.record = res));
    },
  },
};
</script>

<style>
button {
  background-color: #031155;
  border: #8d8a8a;
  color: aliceblue;
}
.row {
  margin: 1%;
  padding-top: 1%;
  border-top: 0.5px solid #d9dbdb;
}
.tableName {
  margin: 1%;
  display: flex;
  gap: 2%;
}
.addressHeader {
  border: 1px solid rgb(224, 224, 224);
  padding: 5px;
  background-color: rgb(243, 243, 243);
}
.card {
  display: flex;
  margin: 10px;
  border: 1px solid rgb(172, 168, 168);
  width: 10%;
}
.btn {
  margin-left: 1%;
  margin-right: 1%;
}
.addressText {
  padding: 5px;
}
.cardgroup {
  display: flex;
}
#add {
  background-color: #031155;
  border: #8d8a8a;
  padding: 0% 5%;
  margin-left: 0%;
}
#header {
  color: #8d8a8a;
}
</style>
