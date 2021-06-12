import React, { Component } from "react";
import apiService from "../services/tasks.service";
import { Link } from "react-router-dom";
import ModalForm from './ModalForm';
export default class UsersList extends Component {
  constructor(props) {
    super(props);
    this.onChangeSearchTitle = this.onChangeSearchTitle.bind(this);
    this.retrieveTutorials = this.retrieveTutorials.bind(this);
    this.refreshList = this.refreshList.bind(this);
    this.setActiveTutorial = this.setActiveTutorial.bind(this);
    this.removeAllTutorials = this.removeAllTutorials.bind(this);
    this.searchTitle = this.searchTitle.bind(this);

    this.state = {
      ListUsers: [],
      currentUser: null,
      currentIndex: -1,
      searchTitle: "",
      isOpen: false
    };
  }
  openModal = () => this.setState({ isOpen: true });
  closeModal = () => this.setState({ isOpen: false });
  componentDidMount() {
    this.retrieveTutorials();
  }

  onChangeSearchTitle(e) {
    const searchTitle = e.target.value;

    this.setState({
      searchTitle: searchTitle
    });
  }

  retrieveTutorials() {
    console.log("get all")
    apiService.getAll()
      .then(response => {
        this.setState({
          ListUsers: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  refreshList() {
    this.retrieveTutorials();
    this.setState({
      currentUser: null,
      currentIndex: -1
    });
  }

  setActiveTutorial(user, index) {
    this.setState({
      currentUser: user,
      currentIndex: index
    });
  }

  removeAllTutorials() {
    apiService.deleteAll()
      .then(response => {
        console.log(response.data);
        this.refreshList();
      })
      .catch(e => {
        console.log(e);
      });
  }

  citasUser(usuario) {
    this.openModal()
    console.log(usuario)
  }

  searchTitle() {
    apiService.get(this.state.searchTitle)
      .then(response => {
        this.setState({
          ListUsers: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  goToAddTask(userID){
    window.location.href = "/add/"+userID
  }

  render() {
    const { searchTitle, ListUsers, currentUser, currentIndex } = this.state;

    return (
      <div className="list row">
        <div className="col-md-8">
          <div className="input-group mb-3">
            <input
              type="text"
              className="form-control"
              placeholder="Search by title"
              value={searchTitle}
              onChange={this.onChangeSearchTitle}
            />
            <div className="input-group-append">
              <button
                className="btn btn-outline-secondary"
                type="button"
                onClick={this.searchTitle}
              >
                Search
              </button>
            </div>
          </div>
        </div>
        <div className="col-md-6">
          <h4>Users</h4>

          <ul className="list-group">
            {ListUsers &&
              ListUsers.map((user, index) => (
                <li
                  className={
                    "list-group-item " +
                    (index === currentIndex ? "active" : "")
                  }
                  onClick={() => this.setActiveTutorial(user, index)}
                  key={index}
                >
                  {user.Email}
                </li>
              ))}
          </ul>

                
        </div>
        <div className="col-md-6">
          {currentUser ? (
            <div>
              <h4>User</h4>
              <div>
                <label>
                  <strong>Nombre:</strong>
                </label>{" "}
                {currentUser.Name}
              </div>

              <div>
                <label>
                  <strong>Email:</strong>
                </label>{" "}
                {currentUser.Email}
              </div>
              <div>
                <label>
                  <strong>Telefono:</strong>
                </label>{" "}
                {currentUser.PhoneNumber}
                
              </div>

              <div>
                <label>
                  <strong>Activo:</strong>
                </label>{" "}
                {currentUser.activo ? "activo" : "activo"}
              </div>
              <div>
                <button className="m-3 btn btn-sm btn-primary" onClick={this.citasUser.bind(this,currentUser.Email)}>Citas</button>
              </div>

              <div>
                <button type="button" onClick={this.goToAddTask.bind(this,currentUser.Id)} class="btn btn-warning buttons1" >Añadir task</button>
              </div>

              <Link
                to={"/tutorials/" + currentUser.id}
                className="badge badge-warning"
              >
                Edit
              </Link>
            </div>
          ) : (
            <div>
              <br />
              <p>Please click on a User...</p>
            </div>
          )}
        </div>
        { this.state.isOpen ? 
          <ModalForm 
            closeModal={this.closeModal} 
            isOpen={this.state.isOpen} 
            handleSubmit={this.handleSubmit}
            ID = {currentUser.Id}
          /> 
          : 
          null 
        }
 
      </div>
    );
  }
}