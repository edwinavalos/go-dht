package dht

import (
	"fmt"
	"testing"
)

func TestAllOfIt(t *testing.T) {
	err := HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := NewDHT("GPI04", Fahrenheit, "")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}

	humidity, temperature, err := dht.ReadRetry(11)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("humidity: %v\n", humidity)
	fmt.Printf("temperature: %v\n", temperature)
}
