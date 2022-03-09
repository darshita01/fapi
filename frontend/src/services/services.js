import axios from 'axios'

export default class service {
     async getUsers() {
         try {
            const res = await axios.get("http://localhost:8004/Users/")
            return res.data.data;
         }catch(error) {
            console.log(error);
         }
      }
    async getUser(i) {
      try{
    const response =  await axios.get("http://localhost:8004/Users/" + i.id)
    return response.data.Address      
   }catch(error){
      console.log(error);
   }
}
   async deleteUser(i){
      try{
         const res = await axios.delete("http://localhost:8004/Users/" + i.id)
         let msg = res.data.message;
         alert(msg)
      }catch(error){
         console.log(error);
      }
   }
   async editUser(bodyFormData,i){
      await this.axios({
           method: "put",
           url: "http://localhost:8004/Users/" + i.id,
           data: bodyFormData,
           headers: { "Content-Type": "multipart/form-data" },
         })
           .then(function (response) {
             console.log(response);
           })
           .catch(function (err) {
             console.log(err);
           });
}
}

