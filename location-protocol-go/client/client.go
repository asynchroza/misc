package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
)

type LocationInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Latitude    float32 `json:"lat"`
	Longitude   float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
}

func GetLocation(ip string) LocationInfo {

	url := "http://ip-api.com/json/" + ip

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	location := LocationInfo{}

	err := json.Unmarshal(body, &location)

	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		panic(err)
	}

	return location
}

var version = 1

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	ips := []string{"223.151.131.42", "44.197.11.96"}

	randIndex := rand.Intn(len(ips))
	loc := GetLocation(ips[randIndex])

	bytes := []byte{byte(version)}
	bytes = append(bytes, FloatToBytes(loc.Latitude)...)
	bytes = append(bytes, FloatToBytes(loc.Longitude)...)

	fmt.Println("Longitude:", loc.Longitude)
	fmt.Println("Latitude:", loc.Latitude)
	fmt.Println("Sending message:", bytes)

	for {
		_, err = conn.Write(bytes)

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func FloatToBytes(f float32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, math.Float32bits(f))
	return bytes
}
