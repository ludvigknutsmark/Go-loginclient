var index = {

    login: function() {
          //Creates a JSON object from the HTML-form values.
          var user = new Object();
          user.operation = "0";
          user.username = document.getElementById("username").value;
          user.password = document.getElementById("password").value;
          //If a form is empty, warn and reload.
          if(!user.username || !user.password){
            alert("Fill in all fields.")
            location.reload();
          }
          //Parse the object to a JSON string.
          var jsonstring = JSON.stringify(user);

          //Send the message to window.On
          astilectron.send(jsonstring);
          /*Listen to the request. As there is a delay from the API-server
            we need to wait for 0.5 sec before receiving the message */
          astilectron.listen(function(mess) {
            setTimeout(function() {
              index.response(mess)
            }, 500);
          });
      },
      response: function(message) {
        //Main.go returns an empty message if the login was successfull
        if(!message) {
          location.replace("success.html")
        } else {
          //The div in index.html gets the error message as value.
          document.getElementById("test").innerHTML = message;
          document.getElementById("username").value = ""
          document.getElementById("password").value = ""

        }
      },

      register: function() {
            var user = new Object();
            user.operation = "1";
            user.reg_username = document.getElementById("reg_username").value;
            user.reg_password = document.getElementById("reg_password").value;
            user.reg_password_confirm = document.getElementById("reg_password_confirm").value;
            if(!user.reg_username || !user.reg_password || !user.reg_password_confirm){
              alert("Fill in all fields.")
              location.reload();
            }
            var jsonstring = JSON.stringify(user);

            astilectron.send(jsonstring);
            astilectron.listen(function(mess) {
              setTimeout(function() {
                index.reg_response(mess)
              }, 500);
            });

        },

        reg_response: function(message) {
          if(!message) {
            location.replace("index.html")
          } else {
            document.getElementById("test").innerHTML = message;
            document.getElementById("reg_password").value = ""
            document.getElementById("reg_password_confirm").value = ""
          }
        }

}
