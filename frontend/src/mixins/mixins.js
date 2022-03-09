export default{

    methods: {
        checkValidity: function(){
            let x = /[0-9]/;
            let y = /[^a-zA-Z]+/;

            if (this.fname == "" || this.lname == "" || this.mname == "") {
                alert("Fill all the field.");
                return false;
            } else if (
            x.test(this.fname) ||
            x.test(this.lname) ||
            x.test(this.mname)
            ) {
                alert(
                    "Numeric value or numbers in between are not allowed for these field."
                );
                return false;
            } else if (
            y.test(this.fname) ||
            y.test(this.lname) ||
            y.test(this.mname)
            ) {
                alert("Special characters are not allowed.");
                return false;
             }
             else return true;
        }
    }
}