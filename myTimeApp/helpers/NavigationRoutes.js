import React, { Component } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import 'react-native-gesture-handler';
import HomeScreen from '../src/HomeScreen';
import FriendsScreen from '../src/FriendScreen';
import LoginScreen from '../src/LoginScreen';

const Stack = createStackNavigator();
export class NavigationRoutes extends Component {
  render() {
    return(
      <NavigationContainer>
        <Stack.Navigator initialRouteName="Login">
        <Stack.Screen
            name="Home"
            component={HomeScreen}
            options={{ title: 'Overview' }}
        />
        <Stack.Screen
            name="Login"
            component={LoginScreen}
          />
          <Stack.Screen
            name="Friends"
            component={FriendsScreen}
          />
        </Stack.Navigator>
      </NavigationContainer>
    )
  }

}

export default NavigationRoutes