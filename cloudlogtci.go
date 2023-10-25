package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v2"
)

var config Config
var rx1, rx2 Radio

// Config from YAML
type Config struct {
	CloudLog struct {
		Server string `yaml:"server"`
		API    string `yaml:"api"`
	} `yaml:"cloudlog"`
	TCI struct {
		Host string `yaml:"host"`
	} `yaml:"tci"`
}

// Radio Struct
type Radio struct {
	Name  string
	VfoA  string
	VfoB  string
	Mode  string
	Split bool
}

func loadConfig(cfg Config) {
	ex, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	configFile := filepath.Dir(ex) + "/config.yaml"
	yamlFile, err := ioutil.ReadFile(configFile)
	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	} else {
		cfg.CloudLog.Server = strings.TrimRight(cfg.CloudLog.Server, "/")
	}
}

func updateVFO(rx string, vfo string, freq string) {
	//vfo:0,0,375500
	if rx == "0" { //RX1 (0)
		if vfo == "0" { //VFO A (0)
			rx1.VfoA = freq
			if !rx1.Split {

				updateCloudLog(rx1)
			}
		} else { //VFO B (1)
			rx1.VfoB = freq
			if rx1.Split {

				updateCloudLog(rx1)
			}
		}
	} else { //RX2 (1)
		if vfo == "0" { //VFO A (0)
			rx2.VfoA = freq
			if !rx2.Split {

				updateCloudLog(rx2)
			}
		} else { //VFO B (1)
			rx2.VfoB = freq
			if rx2.Split {

				updateCloudLog(rx2)
			}
		}
	}
}

func updateSplit(rx string, rxSplit string) {
	if rx == "0" { //RX1
		b, err := strconv.ParseBool(rxSplit)
		if err != nil {
			log.Fatalf("error: %v", err)
		} else {
			rx1.Split = b
			updateCloudLog(rx1)
		}
	} else { //RX2
		b, err := strconv.ParseBool(rxSplit)
		if err != nil {
			log.Fatalf("error: %v", err)
		} else {
			rx2.Split = b
			updateCloudLog(rx2)
		}
	}
}

func updateMode(rx string, rxMode string) {
	if rx == "0" { //RX1
		rx1.Mode = fixMode(rxMode)
		updateCloudLog(rx1)
	} else { //RX2
		rx2.Mode = fixMode(rxMode)
		updateCloudLog(rx2)
	}
}

func fixMode(rxMode string) string {
	rxMode = strings.ToUpper(rxMode)

	// Thetis Modes: AM,SAM,DSB,LSB,USB,NFM,FM,DIGL,DIGU,CWL, CWU, CW;
	// No valid modes are exist for generic digital
	switch rxMode {
	case "CWL": //Thetis
		rxMode = "CW"
	case "CWU": //Thetis
		rxMode = "CW"
	//case "LSB":
	//	rxMode = "SSB"
	//case "USB":
	//	rxMode = "SSB"
	case "DIGIL": //ExpertSDR
		rxMode = "LSB"
	case "DIGL": //Thetis
		rxMode = "LSB"
	case "DIGIU": //ExpertSDR
		rxMode = "USB"
	case "DIGU": //Thetis
		rxMode = "USB"
	case "DSB":
		rxMode = "AM"
	case "SAM":
		rxMode = "AM"
	case "DRM":
		rxMode = "AM"
	case "NFM":
		rxMode = "FM"
	case "WFM":
		rxMode = "FM"
	}
	return rxMode
}

func sendTCI(c *websocket.Conn, message string) {
	log.Println("Sent:", message)
	err := c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write close:", err)
		return
	}
}

func updateCloudLog(rx Radio) {

	var frequency string
	//Only send TX VFO
	if rx.Split {
		frequency = rx.VfoB
	} else {
		frequency = rx.VfoA
	}

	timestamp := time.Now().UTC().Format("2006/01/02 15:04:05")

	//{"key":"YOUR_API_KEY", "radio":"FT-950","frequency":14075,"mode":"SSB","timestamp":"2006/01/02 15:04:05"}
	requestBody, err := json.Marshal(map[string]string{
		"key":       config.CloudLog.API,
		"radio":     rx.Name,
		"frequency": frequency,
		"mode":      rx.Mode,
		"timestamp": timestamp,
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", config.CloudLog.Server+"/index.php/api/radio", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	debug := false

	if string(body) != "{\"status\":\"success\"}" || debug {
		fmt.Println(string(body))
	}
}

func connectTCI(u url.URL) *websocket.Conn {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		//log.Fatal("dial:", err)
		for { // Redial
			log.Println("Reconnecting...")
			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Printf("dial err:" + err.Error())
				log.Println("Wait 5 seconds...")
				time.Sleep(time.Second * 5)
				continue
			}
			c = conn
			break
		}
	}
	return c
}

func main() {

	loadConfig(config)

	fmt.Println("CloudLogTCI 2021.11.0")
	fmt.Println("CloudLog Server:", config.CloudLog.Server)
	fmt.Println("TCI Server:", config.TCI.Host)

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: config.TCI.Host}
	log.Printf("connecting to %s", u.String())

	c := connectTCI(u)
	defer c.Close()

	sendTCI(c, "CloudLogTCI Connected")
	log.Println("Connected to TCI:", time.Now().UTC().Format("2006/01/02 15:04:05"))

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				log.Println("Disconnected from TCI")
				c = connectTCI(u)
				sendTCI(c, "CloudLogTCI Reconnected")
				log.Println("Reconnected to TCI", time.Now().UTC().Format("2006/01/02 15:04:05"))
			}
			// Print all messages
			//log.Printf("recv: %s", message)

			tciRaw := strings.TrimRight(string(message), ";")
			tciArray := strings.Split(tciRaw, ":")

			switch tciArray[0] {
			case "device":
				rx1.Name = tciArray[1] + " RX1"
				rx2.Name = tciArray[1] + " RX2"
			case "vfo":
				//vfo:0,0,375500
				tciValue := strings.Split(tciArray[1], ",")
				//RX, VFO, Freq
				updateVFO(tciValue[0], tciValue[1], tciValue[2])
			case "split_enable":
				//split_enable:0,true/false;
				tciValue := strings.Split(tciArray[1], ",")
				updateSplit(tciValue[0], tciValue[1])
			case "modulation":
				//modulation:0,lsb;
				tciValue := strings.Split(tciArray[1], ",")
				updateMode(tciValue[0], tciValue[1])
			}
		}
	}()

	for {
		select {
		case <-done:
			log.Println("Bye!")
			return
		case <-interrupt:
			log.Println("Bye!")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}
