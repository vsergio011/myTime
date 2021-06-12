import React, { Component } from "react";
import apiService from "../services/tasks.service";
import Form from 'react-bootstrap/Form'
import DayPickerInput from 'react-day-picker/DayPickerInput';
import 'react-day-picker/lib/style.css';
import moment from 'moment';

export default class updateTask extends Component {
  
    constructor(props) {
        super(props);
        //this.handleDayClick = this.handleDayClick.bind(this);
        this.handleDayChange = this.handleDayChange.bind(this);
        this.onChangeTitle = this.onChangeTitle.bind(this);
        this.onChangeHora = this.onChangeHora.bind(this);
        this.updateTask = this.updateTask.bind(this);
        
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
            Hora : null,
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

      onChangeTitle(e) {
        const Title = e.target.value;
        this.setState({
          Title: Title
        });
        console.log(Title)
      }
      onChangeHora(e) {
        const time = e.target.value;
        this.setState({
          time: time
        });
        console.log(time)
      }

      componentDidMount() {
        this.retrieveTask();
      }
      // fetch task by id
      retrieveTask() {
        console.log("get task"+this.props.match.params.id)
        apiService.getTask(this.props.match.params.id)
          .then(response => {
            this.setState({
                Title: response.data.Title,
                date : response.data.Date,
                time : response.data.Date.substring(response.data.Date.indexOf("T")+1,response.data.Date.indexOf("Z")),
                selectedDay : new Date(response.data.Date),
                uid : response.data.Uid_user,
                id : this.props.match.params.id
                
            });
            console.log(response.data);
            console.log(response.data.Date.substring(response.data.Date.indexOf("T")+1,response.data.Date.indexOf("Z")));
          })
          .catch(e => {
            console.log(e);
          });
      }

      // update the task
      updateTask() {

        let fecha =moment(this.state.selectedDay).format('YYYY-MM-DD')
        let fechaadd =fecha.toString()+" "+this.state.time
        console.log(fechaadd+"   "+this.props.match.params.id)
        var data = {
          Id : parseInt(this.props.match.params.id),
          Title : this.state.Title,
          Date : fechaadd,
          Uid_user : this.state.uid
        };

        apiService.updateTask(data)
          .then(response => {
   
            console.log(response.data);
          })
          .catch(e => {
            console.log(e);
          });

      }

  render(){
    const { Title,date,CurrentTask,currentIndex,selectedDay,isEmpty,isDisabled,Hora,time} = this.state;
   
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
            <Form.Group controlId="formBasicCheckbox">  </Form.Group>
            <button className="btn btn-sm btn-primary" onClick={this.updateTask}>Update</button>
          
    </div>
    )
  }
}