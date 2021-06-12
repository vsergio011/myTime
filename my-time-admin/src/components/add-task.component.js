import React, { Component } from "react";
import apiService from "../services/tasks.service";
import Form from 'react-bootstrap/Form'
import { MdDeleteForever,MdModeEdit } from "react-icons/md";
import DayPickerInput from 'react-day-picker/DayPickerInput';
import 'react-day-picker/lib/style.css';
import moment from 'moment';

export default class AddTask extends Component {
  
    constructor(props) {
        super(props);
        //this.handleDayClick = this.handleDayClick.bind(this);
        this.handleDayChange = this.handleDayChange.bind(this);
        this.onChangeTitle = this.onChangeTitle.bind(this);
        this.onChangeHora = this.onChangeHora.bind(this);
        this.addTask = this.addTask.bind(this);
        
        this.state = {
            id: null,
            uid:null,
            name : null,
            Title: null,
            selectedDay: null,
            isEmpty: true,
            isDisabled: false,
            date : null,
            time: null,
            CurrentTask: null,
            currentIndex: -1
        };
      }

      handleDayChange(selectedDay, modifiers, dayPickerInput) {
        const input = dayPickerInput.getInput();
        this.setState({
          selectedDay,
          isEmpty: !input.value.trim(),
          isDisabled: modifiers.disabled === true,
        });
      }
      onChangeHora(e) {
        const time = e.target.value;
        this.setState({
          time: time
        });
        console.log(time)
      }

      onChangeTitle(e) {
        const Title = e.target.value;
        this.setState({
          Title: Title
        });
        console.log(Title)
      }

      componentDidMount() {
        console.log(this.props.match.params.id)
      }

      addTask() {

          let fecha =moment(this.state.selectedDay).format('YYYY-MM-DD')
          let fechaadd =fecha.toString()+" "+this.state.time
          let id = this.props.match.params.id
         console.log(id)
          var data = {
            Title : this.state.Title,
            Date : fechaadd,
            Uid_user : id
          };
        apiService.addTask(data)
          .then(response => {
   
            console.log(response.data);

            window.location.href = "/users/"
          })
          .catch(e => {
            console.log(e);
          });
      }

  render(){
    const { Title,date,CurrentTask,currentIndex,selectedDay,isEmpty,isDisabled,time} = this.state;
   
    return(

    <div>
    
            <div className="form-group">
                <label>Titulo</label>
                <input
                    className="form-control"
                    value={Title}
                    onChange={this.onChangeTitle}
                />
            </div>
            <br></br>
          
            <div className="form-group">
                <div>
                        <DayPickerInput
                        value={selectedDay}
                        onDayChange={this.handleDayChange}
                        dayPickerProps={{
                            selectedDays: selectedDay,
                            disabledDays: {
                            daysOfWeek: [0, 6],
                            },
                        }}
                        />
                </div>
                <br></br>
            </div>
            <div className="form-group">
                <label>Hora</label>
                <input
                    className="form-control"
                    value={time}
                    onChange={this.onChangeHora}
                />
            </div>
            <br></br>
            <Form.Group controlId="formBasicCheckbox">
               
            </Form.Group>
            <button className="btn btn-sm btn-primary"
                   onClick={this.addTask}>Nueva cita</button>
          
    </div>
    )
  }
}