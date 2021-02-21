import React from 'react';
import { StyleSheet, View,ScrollView  } from 'react-native';
import { Avatar,Card,Text,Button,Rating, AirbnbRating} from "react-native-elements";
class HomeList extends React.Component {
  render() {
    return (
        <ScrollView>
         <Card containerStyle={{backgroundColor: 'powderblue'}}>
            <Card.Title>Company</Card.Title>
            <View style={{ flexDirection: "row", height: 100, padding: 0 }}>
                <View style={{flex: 0.3}}>
                    <Avatar
                        rounded
                        source={{
                        uri: 'https://happytravel.viajes/wp-content/uploads/2020/04/146-1468479_my-profile-icon-blank-profile-picture-circle-hd.png',
                        }}
                        size="large"
                    />
                </View>
                <View style={{flexDirection: "column",flex: 0.7,}}>
                    <View style={{flex: 0.6}}>
                        <Text>description</Text>
                    </View>   
                    <View style={{flex: 0.4,width: "80%"}}>
                        <Button 
                        title="Outline button" 
                        type="solid"
                        color="red"
                         />
                    </View>

                </View>
            </View>
        </Card>
        </ScrollView>
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

export default HomeList;