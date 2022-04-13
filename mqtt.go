package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	mq mqtt.Client
)

func setupMQTTClient() {
	mqttOptions := mqtt.NewClientOptions()
	mqttOptions.SetClientID(viper.GetString("mqtt.client_id"))
	mqttOptions.SetUsername(viper.GetString("mqtt.username"))
	mqttOptions.SetPassword(viper.GetString("mqtt.password"))
	mqttOptions.SetMaxReconnectInterval(time.Second * 5)
	mqttOptions.SetConnectTimeout(time.Second)
	mqttOptions.SetCleanSession(viper.GetBool("mqtt.clean_session"))
	mqttOptions.SetAutoReconnect(true)
	mqttOptions.SetOnConnectHandler(connectHandler)
	mqttOptions.SetConnectionLostHandler(connectionLostHandler)
	mqttOptions.SetOrderMatters(true)
	mqttOptions.SetKeepAlive(viper.GetDuration("mqtt.keep_alive"))
	mqttOptions.AddBroker(viper.GetString("mqtt.broker"))

	mq = mqtt.NewClient(mqttOptions)
	for token := mq.Connect(); token.Wait() && token.Error() != nil; token = mq.Connect() {
		log.Println("MQTT: error connecting:", token.Error())
		time.Sleep(time.Second * 5)
	}
}

func connectionLostHandler(_ mqtt.Client, err error) {
	log.Println("MQTT: Connection lost:", err)
}

func connectHandler(_ mqtt.Client) {
	log.Println("MQTT: Connected!")
}
