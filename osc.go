package main

import (
	"encoding/json"
	"fmt"
	"github.com/loffa/gosc"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"time"
)

var (
	cli *gosc.Client
)

func setupOSCClient() {
	address := fmt.Sprintf("%s:%d", viper.GetString("osc.host"), viper.GetInt("osc.port"))
	client, err := gosc.NewClient(address)
	if err != nil {
		log.Fatalln(err)
	}
	cli = client
	log.Println("Fetching console info...")
	info, err := cli.CallMessage("/info")
	if err != nil {
		log.Fatalln("Could not get mixer info:", err)
	}
	log.Println("Got info from the OSC server:", info.Arguments)
	err = cli.HandleMessageFunc("/*", oscMessageHandler)
	if err != nil {
		log.Fatalln("Could not register handler:", err)
	}

	go xremoteSender()
}

func xremoteSender() {
	ticker := time.NewTicker(time.Second * 9)
	err := cli.EmitMessage("/xremote")
	if err != nil {
		log.Println(err)
	}
	for {
		select {
		case _ = <-ticker.C:
			err := cli.EmitMessage("/xremote")
			if err != nil {
				log.Println(err)
			}
		}
	}
}

type dataField struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}
type mqttPayload []*dataField

func oscMessageHandler(msg *gosc.Message) {
	prefix := viper.GetString("mqtt.topic_prefix")
	topic := fmt.Sprintf("%s%s", prefix, msg.Address)
	res := make(mqttPayload, 0, len(msg.Arguments))
	for _, v := range msg.Arguments {
		res = append(res, &dataField{
			Type:  reflect.TypeOf(v).Name(),
			Value: v,
		})
	}
	data, _ := json.Marshal(res)
	err := mq.Publish(topic, 0, false, data).Error()
	if err != nil {
		log.Println("MQTT: Could not publish:", err)
	}
}
