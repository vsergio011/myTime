import React from 'react';
import { View,Image,  FlatList, TouchableOpacity, Text, ActivityIndicator } from 'react-native';
import { Avatar,Card,Button,Rating, AirbnbRating} from "react-native-elements";

fetchUsersWithFetchAPI = () => {
    this.setState({...this.state, isFetching: true});
    fetch(USER_SERVICE_URL)
        .then(response => response.json())
        .then(result => {
            this.setState({users: result, isFetching: false})
        })
        .catch(e => {
            console.log(e);
            this.setState({...this.state, isFetching: false});
        });
};
fetchUsers = this.fetchUsersWithFetchAPI