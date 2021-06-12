import React, { Component } from "react";
import apiService from "../services/tasks.service";
import Modal from 'react-bootstrap/Modal'
import { MdDeleteForever,MdModeEdit } from "react-icons/md";
import moment from 'moment';
export default class ModalForm extends Component {

    constructor(props) {
        super(props);
       
    
        this.state = {
            name : null,
            ListTasks: [],
            CurrentTask: null,
            currentIndex: -1
        };
      }
  handleChange = (e) => this.setState({name: e.target.value})
  componentDidMount() {
    this.retrieveTasks();
  }
  retrieveTasks() {
    console.log("get all")
    apiService.getTasks(this.props.ID)
      .then(response => {
        this.setState({
          ListTasks: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }
  setActiveTask(user, index) {
    this.setState({
      currentTask: user,
      currentIndex: index
    });
  }
  redirectToTarget = () => {
    this.context.push('/task/4')
  }

  goToUpdateTask(task){
    window.location.href = "/task/"+task
  }
  deleteTask(task) {
    
    let fecha =moment(this.state.selectedDay).format('YYYY-MM-DD')
    let fechaadd =fecha.toString()+" 10:00:00"
    var data = {
      Id : task,
      Title : this.state.Title,
      Date : fechaadd,
      Uid_user : this.state.uid
    };
  apiService.deleteTask(data)
    .then(response => {

      console.log(response.data);
    })
    .catch(e => {
      console.log(e);
    });
    this.retrieveTasks();
   
}

  render(){
    const { ListTasks,CurrentTask,currentIndex} = this.state;
    return(
        
      <Modal 
        show={this.props.isOpen} 
        onHide={this.props.closeModal}
        dialogClassName="modal-90w"
      >
      <Modal.Header closeButton>
        <Modal.Title>Citas del usuario</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <ul className="list-group">
                {ListTasks &&
                ListTasks.map((task, index) => (
                    <li
                    className={
                        "list-group-item " +
                        (index === currentIndex ? "active" : "")
                    }
                    onClick={() => this.setActiveTask(task, index)}
                    key={index}
                    >
                    {task.Title} - {task.Date} -- {task.Id}
                    <div class="div-buttons1"> 
                        <button type="button" onClick={this.goToUpdateTask.bind(this,task.Id)} class="btn btn-warning buttons1" ><MdModeEdit/></button>
                        <button type="button" onClick={this.deleteTask.bind(this,task.Id)} class="btn btn-danger buttons1"><MdDeleteForever/></button> 
                    </div>
                   
                    </li>
                ))}
            </ul>

            <div className="col-md-6">
        </div>
      </Modal.Body>
      <Modal.Footer>
         
      </Modal.Footer>
    </Modal>
    )
  }
}