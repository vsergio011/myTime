import React, { Component } from 'react';
import apiService from "../services/tutorial.service";
import {
  BrowserRouter as Router,
  Route,
  Link,
  Redirect
} from 'react-router-dom'
import AddTutorial from "./add-tutorial.component";

import PropTypes from 'prop-types'
const Protected = () => <h3>Protected</h3>
const loginok = false
function PrivateRoute({ children, ...rest }) {
  return (
    <Route {...rest} render={() => {
      return localStorage.isLogged === 'true'
        ? children
        : <Redirect to='/login' />
    }} />
  )
}

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
      redirectToTarget = () => {
        this.context.push('/add')
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
              localStorage.isLogged = 'true'
            }else{
              localStorage.isLogged = 'false'
            } 
              console.log("el ddf"+localStorage.isLogged)
            
           // window.location.href = "/protected"
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
        <PrivateRoute path='/protected'>
          <Protected />
        </PrivateRoute>
    </div>
    
    )
  }
}
   