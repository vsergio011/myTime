import React, { Component } from 'react';
import apiService from "../services/tasks.service";

import AddTutorial from "./add-task.component";
import PropTypes from 'prop-types'


export default class login extends Component {
    constructor(props) {
        super(props);
        this.onChangeEmail = this.onChangeEmail.bind(this);
        this.onChangePassword = this.onChangePassword.bind(this);
        this.checkUser = this.checkUser.bind(this);
    
    
        this.state = {
            password : "",
            email: ""
        };
      }

      componentDidMount() {
        localStorage.token = "";
      }
      
      onChangeEmail(e) {
        const email = e.target.value;
        this.setState({
          email: email
        });
        console.log(this.state.email)
      }
      onChangePassword(e) {
        const password = e.target.value;
        this.setState({
          password: password
        });
      }
      static contextTypes = {
        router: PropTypes.object
      }
     
      checkUser() {
      
        var data = {
          email : this.state.email,
          password : this.state.password
        };
    
        console.log(data)
          apiService.getkey(data)
          .then(response => {
            this.setState({
              //tutorials: response.data
            });
            
            console.log("el token "+localStorage.token.length)
            if(localStorage.token.length > 1 ){
              //localStorage.isLogged = 'true'
              window.location.href = "/users"
            }else{
              localStorage.isLogged = 'false'
            } 
              console.log("el ddf"+localStorage.isLogged)
            
            
           // console.log("tokle"+response.data.idToken);
            
           
          })
          .catch(e => {
            console.log(e);
          });
         
      }
  render() {
    const { email, password } = this.state;
    return(
    <div>

    
        <h3>Sign In</h3>

        <div className="form-group">
            <label>Email address</label>
            <input
                type="email" 
                className="form-control"
                placeholder="Enter email"
                value={email}
                onChange={this.onChangeEmail}
            />
        </div>

        <div className="form-group">
            <label>Password</label>
            <input 
                type="password" 
                className="form-control" 
                placeholder="Enter password"
                value={password}
                onChange={this.onChangePassword} 
            />
        </div>

        <div className="form-group">
            <div className="custom-control custom-checkbox">
                <input type="checkbox" className="custom-control-input" id="customCheck1" />
                <label className="custom-control-label" htmlFor="customCheck1">Remember me</label>
            </div>
        </div>
        
        <button 
        
        className="btn btn-primary btn-block"
        onClick={this.checkUser}>
            Submit
        </button>
        <p className="forgot-password text-right">
            
        </p>
    
    </div>
    
    )
  }
}
   