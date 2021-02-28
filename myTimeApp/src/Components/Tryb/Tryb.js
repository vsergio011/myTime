import React from 'react';
import { View,Image, Alert, FlatList, TouchableOpacity, Text, ActivityIndicator } from 'react-native';
import { Avatar,Card,Button,Rating, AirbnbRating} from "react-native-elements";
const axios = require('axios');

class Tryb extends React.Component {

    
 //Define your state for your component. 
 state = {
    //Assing a array to your pokeList state
    pokeList: [],
    //Have a loading state where when data retrieve returns data. 
    loading: true
}
//Define your componentDidMount lifecycle hook that will retrieve data.
//Also have the async keyword to indicate that it is asynchronous. 
async componentDidMount() {
    //Have a try and catch block for catching errors.
    try {
        //Assign the promise unresolved first then get the data using the json method. 
        const pokemonApiCall = await fetch('https://pokeapi.co/api/v2/pokemon/');
        const pokemon = await pokemonApiCall.json();
        this.setState({pokeList: pokemon.results, loading: false});
    } catch(err) {
        console.log("Error fetching data-----------", err);
    }
}
onPressButton() {  
    Alert.alert('You clicked the button!')  
} 
 async getUser(){
    try {
        const response = await axios.get('https://pokeapi.co/api/v2/pokemon/');
        console.log(response.json);
        
      } catch (error) {
        console.error(error);
      }
 }
//Define your renderItem method the callback for the FlatList for rendering each item, and pass data as a argument. 
renderItem(data) {
    return <TouchableOpacity>
                <View >
                    <Text >{data.item.name}</Text>
                    <Image source={{uri: 'https://res.cloudinary.com/aa1997/image/upload/v1535930682/pokeball-image.jpg'}} 
                            />
                </View>
            </TouchableOpacity>
} 
render() {
   return(     <View><Button  
    onPress={this.onPressButton}  
    title="Press Me"  
/><Button  
    onPress={this.getUser}  
    title="Press Me"  
/></View>   );

}

}



export default Tryb;