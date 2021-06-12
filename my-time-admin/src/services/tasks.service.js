import http from "../http-common";
// Firebase App (the core Firebase SDK) is always required and must be listed first
import firebase from "firebase/app";
import "firebase/auth";

// Requests to database
class TasksDataService {

  
  constructor(token){
    var firebaseConfig = {
      apiKey: "AIzaSyDZD3Dyzt5ahIfnIWWkhqYF14UJn5EoP-g",
      authDomain: "mytime-82291.firebaseapp.com",
      databaseURL: "https://mytime-82291.firebaseio.com",
      projectId: "mytime-82291",
      storageBucket: "mytime-82291.appspot.com",
      messagingSenderId: "4866724698",
      appId: "1:4866724698:web:5b35f96a9542d9b97fd963",
      measurementId: "G-10WMJVPNRD",
    };
    //init the firebase database
    firebase.initializeApp(firebaseConfig);
    this.token = token
  }
  //GET - retrieve all users
  getAll() {
    const token = 'Bearer '+localStorage.token;
    console.log("to"+token.length)
    return http.get("/users",{ mode:'no-cors',headers: {"Authorization" : `${token}`} });
  }
  // Get - retrive all tasks
  getTasks(user) {
    const token = 'Bearer '+localStorage.token;
    console.log("to"+token.length)
    return http.get(`/taskscreatedby/${user}`,{ mode:'no-cors',headers: {"Authorization" : `${token}`} });
  }
  // GET - retrieve task by ID
  getTask(id) {
    const token = 'Bearer '+localStorage.token;
    console.log("to"+token.length)
    return http.get(`/task/${id}`,{ mode:'no-cors',headers: {"Authorization" : `${token}`} });
  }

  //GET - retroeve user by ID
  get(id) {
    console.log("estoy aqui en el get")
    if(localStorage.token){
     
    }
    console.log()
   
    return http.get(`/user/${id}`,{ mode:'no-cors',headers: {"Authorization" : `Bearer token`} });
  }

  create(data) {
    return http.post("/tutorials", data);
  }

  //Retrieve the token from firebase
  getkey(data) {
    var email = data.email
    var password = data.password

    firebase.auth().signInWithEmailAndPassword(email, password)
      .then((userCredential) => {
        // Signed in
        var user = userCredential.user;
        return user.getIdToken(true).then(idToken => {
          console.log('fire'+idToken.length)
          localStorage.token = idToken;
          this.getLoginUserInfo(data)
        });
        
      })
      .catch(function(error) {
        // Handle Errors here.
        console.log("estoy en el catch")
        localStorage.token = "";
        var errorCode = error.code;
        var errorMessage = error.message;
        if (errorCode === 'auth/wrong-password') {
          alert('Wrong password.');
        } else {
          alert(errorMessage);
        }
        console.log(error);
      });

    return this.getLoginUserInfo(data)
  }
  
  getLoginUserInfo(data){
    //encontrado el usuario se realiza una consulta al servicio go de los datos del usuario
    this.token = "prueba"
    console.log(this.token)
    return http.post("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyDZD3Dyzt5ahIfnIWWkhqYF14UJn5EoP-g", {
        headers:{
          'Content-Type' : 'application/json; charset=UTF-8'
        },
        email: 'miemail@gmal.com',
        password: 'mipass'
      });
  }
  // POST - Update task
  updateTask(data){
    console.log("los datos"+data.Id)
    //encontrado el usuario se realiza una consulta al servicio go de los datos del usuario
    const token = 'Bearer '+localStorage.token;
  
    let axiosConfig = {
      headers: {"Authorization" : `${token}`},
    };
    
    return http.post("/updatetask",data, axiosConfig);
  }

  addTask(data){
    //encontrado el usuario se realiza una consulta al servicio go de los datos del usuario
    const token = 'Bearer '+localStorage.token;
  
    let axiosConfig = {
      headers: {"Authorization" : `${token}`},
    };
    
    return http.post("/addtask",data, axiosConfig);
  }

  deleteTask(data){
    //encontrado el usuario se realiza una consulta al servicio go de los datos del usuario
    const token = 'Bearer '+localStorage.token;
  
    let axiosConfig = {
      headers: {"Authorization" : `${token}`},
    };
    
    return http.post("/deletetask",data, axiosConfig);
  }

}

export default new TasksDataService();