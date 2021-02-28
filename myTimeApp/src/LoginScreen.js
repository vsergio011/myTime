import React from 'react';
import { StyleSheet, Text, View,Button } from 'react-native';
//import Login from './Components/Login/Login'
import Tryb from './Components/Tryb/Tryb'


class LoginScreen extends React.Component {
  render() {
    return (
      <Tryb/>
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

export default LoginScreen;