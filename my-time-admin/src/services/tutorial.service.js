import http from "../http-common";
// Firebase App (the core Firebase SDK) is always required and must be listed first
import firebase from "firebase/app";
import "firebase/auth";
class TutorialDataService {

  
  constructor(token){

    firebase.initializeApp(firebaseConfig);
    this.token = token
  }
  getAll() {
    const token = 'Bearer '+localStorage.token;
    console.log("to"+token.length)
    return http.get("/users",{ mode:'no-cors',headers: {"Authorization" : `${token}`} });
  }

  get(id) {
    console.log("estoy aqui en el get")
    if(localStorage.token){
     
    }
    console.log()
    const token = 'eyJhbGciOiJSUzI1NiIsImtpZCI6ImNjM2Y0ZThiMmYxZDAyZjBlYTRiMWJkZGU1NWFkZDhiMDhiYzUzODYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibWluYW1lIiwicGljdHVyZSI6Imh0dHA6Ly93d3cuZXhhbXBsZS5jb20vMTIzNDU2NzgvcGhvdG8ucG5nIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL215dGltZS04MjI5MSIsImF1ZCI6Im15dGltZS04MjI5MSIsImF1dGhfdGltZSI6MTYyMDEyMTM0NCwidXNlcl9pZCI6Im5uYUdsa093OXVQSlZycWYxcHdwcHhPU1hPYzIiLCJzdWIiOiJubmFHbGtPdzl1UEpWcnFmMXB3cHB4T1NYT2MyIiwiaWF0IjoxNjIwMTIxMzQ0LCJleHAiOjE2MjAxMjQ5NDQsImVtYWlsIjoibWllbWFpbEBnbWFsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJwaG9uZV9udW1iZXIiOiIrMzQ2NTg0NTIzNjUiLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIiszNDY1ODQ1MjM2NSJdLCJlbWFpbCI6WyJtaWVtYWlsQGdtYWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.j9R_lJO5zh1WCgPXzpwhpl9kcYtHa81qD4n8s2JxNfTzucLmYPG2fSYUrVa14XxLGqaJE7HvsrwR1hOTnuhUj26PuqbVd3fQQIKi4O6_W1NEFd98Skpe2SS5OXNmbbVuNSJiMv-k7cKI8NtkYxeU68NSX3HIxrOyzbiQ5BkaoWGs1ZkHifT1JEL7dvLjcDMrYk2lGjtPNF_9ea79XwzbpQ3_buwZLv8bR525OmtPDrOqD6odKExYbjOyWUv491fL0J2lujrSK-DgQtWdUZ7vYX-9A_tg9kqLcp5E3EI7jZdH-b0zvb6B_IoGbiUpi2vwTVua0-9qtj1fe0DUmOBm4g'
    const token2 = 'eyJhbGciOiJSUzI1NiIsImtpZCI6ImNjM2Y0ZThiMmYxZDAyZjBlYTRiMWJkZGU1NWFkZDhiMDhiYzUzODYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibWluYW1lIiwicGljdHVyZSI6Imh0dHA6Ly93d3cuZXhhbXBsZS5jb20vMTIzNDU2NzgvcGhvdG8ucG5nIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL215dGltZS04MjI5MSIsImF1ZCI6Im15dGltZS04MjI5MSIsImF1dGhfdGltZSI6MTYyMDIwNTkzNCwidXNlcl9pZCI6Im5uYUdsa093OXVQSlZycWYxcHdwcHhPU1hPYzIiLCJzdWIiOiJubmFHbGtPdzl1UEpWcnFmMXB3cHB4T1NYT2MyIiwiaWF0IjoxNjIwMjA1OTM0LCJleHAiOjE2MjAyMDk1MzQsImVtYWlsIjoibWllbWFpbEBnbWFsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJwaG9uZV9udW1iZXIiOiIrMzQ2NTg0NTIzNjUiLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIiszNDY1ODQ1MjM2NSJdLCJlbWFpbCI6WyJtaWVtYWlsQGdtYWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.QifIVe6iLUmTNLaJU9TzWAdw-pOGiDkhVl8scezNVj5coHxbOndM85BIB66DlyvYMWNT9Xbx6I9tE0J2f0-TJi7tjMTRoNGXXaxNrE1Ei0XNNxoYeulY0_4vvpc4yRt3i7yV3NFi9F7GHKQT2RVmD_hVeMHuTTcylOT0OKWgS1LiO2mGXHi_A9F7q4rhk874vkwtoIkhjq2ZMPU6Gttk8QFFNxRY22murqk5zCzNWl2Rvk0W_9vkF6dwuDqQ1WvPKFozm-lbSeqsvZr6IyTvOXN-bFLIf9OcnBziajOVqF2X4jX000lAbRTx_X1JnmUqhpBrKgd0lfM4XExN8VB4Qw'
    console.log(token.length+"  "+token2.length)
    return http.get(`/user/${id}`,{ mode:'no-cors',headers: {"Authorization" : `Bearer ${token}`} });
  }

  create(data) {
    return http.post("/tutorials", data);
  }
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

  update(id, data) {
    return http.put(`/tutorials/${id}`, data);
  }

  delete(id) {
    return http.delete(`/tutorials/${id}`);
  }

  deleteAll() {
    return http.delete(`/tutorials`);
  }

  findByTitle(title) {
    return http.get(`/tutorials?title=${title}`);
  }
}

export default new TutorialDataService();