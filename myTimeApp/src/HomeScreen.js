import React from 'react';
import { StyleSheet, Text, View,Button } from 'react-native';
import HomeList from './Components/HomeList/HomeList'
class HomeScreen extends React.Component {
  render() {
    return (
      <HomeList/>
    );
  }
}

const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff',
      alignItems: 'center',
      justifyContent: 'center',
    },
  });

export default HomeScreen;