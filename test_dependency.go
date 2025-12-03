package main

import (
	"fmt"
	"github.com/wujunwei928/edge-tts-go/edge_tts"
)

func main() {
	// Test basic functionality
	fmt.Println("Testing edge-tts-go dependency...")

	// Try to list voices - this should work if the dependency is functional
	voices, err := edge_tts.ListVoices("")
	if err != nil {
		fmt.Printf("Error listing voices: %v\n", err)
		return
	}

	fmt.Printf("Successfully found %d voices\n", len(voices))
	if len(voices) > 0 {
		fmt.Printf("First voice: %s (%s)\n", voices[0].FriendlyName, voices[0].ShortName)
	}

	// Test creating a communicate connection (basic text-to-speech test)
	conn, err := edge_tts.NewCommunicate("Hello world", edge_tts.SetVoice("en-US-AriaNeural"))
	if err != nil {
		fmt.Printf("Error creating communicate connection: %v\n", err)
		return
	}

	// Test streaming (this will tell us if the core functionality works)
	audioData, err := conn.Stream()
	if err != nil {
		fmt.Printf("Error streaming audio: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated audio data: %d bytes\n", len(audioData))
	fmt.Println("edge-tts-go dependency appears to be working correctly!")
}