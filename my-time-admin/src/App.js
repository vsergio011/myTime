import React, { Component } from "react";
import { BrowserRouter as Router,Switch, Route, Link,Redirect } from "react-router-dom";
import "./css/bootstrap.min.css";
import "./css/style.css";
import "./App.css";

import AddTask from "./components/add-task.component";
import usersList from "./components/users-list.component";
import login from "./components/login";
import editarTarea from "./components/updateTask";
const loginok = false
function PrivateRoute({ children, ...rest }) {
  return (
    <Route {...rest} render={() => {
      return localStorage.token.length > 1 
        ? children
        : <Redirect to='/login' />
    }} />
  )
}
class App extends Component {
  render() {
    return (
      <div>
        <nav className="navbar navbar-expand navbar-dark bg-dark">
          <a href="/" className="navbar-brand">
            MyTime
          </a>
          <div className="navbar-nav mr-auto">
            <li className="nav-item">
              <Link to={"/users"} className="nav-link">
                Users
              </Link>
            </li>
          </div>
        </nav>

        <div className="container mt-3">
        <PrivateRoute path='/users'/>  
        <PrivateRoute path='/add'/>         
          <Switch>
            <Route exact path={["/", "/login"]} component={login} />
            <Route exact path={["/", "/users"]} component={usersList} />
            <Route exact path="/add/:id" component={AddTask} />
            
            <Route path="/task/:id" component={editarTarea} />       
          </Switch>
        
        </div>
      </div>
    );
  }
}

export default App;