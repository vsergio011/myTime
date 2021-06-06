import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import NavigationRoutes from './helpers/NavigationRoutes'
import 'react-native-gesture-handler';

class App extends React.Component {
  render() {
    return (
      <NavigationRoutes/>
    );
  }
}

export default App;