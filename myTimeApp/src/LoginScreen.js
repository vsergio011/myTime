import React from 'react';
import { StyleSheet, Text, View,Button } from 'react-native';
import Login from './Components/Login/Login'
class LoginScreen extends React.Component {
  render() {
    return (
      <Login/>
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