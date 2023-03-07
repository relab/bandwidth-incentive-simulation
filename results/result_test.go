package results

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/vmihailenco/msgpack/v4"
//	"io"
//	"os"
//	"testing"
//)
//
//func TestDecodeMessagePack(t *testing.T) {
//	filePath := "routes.mp"
//	actualFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		panic(err)
//	}
//
//	// Create a streaming decoder
//	decoder := msgpack.NewDecoder(actualFile)
//
//	// Read the MessagePack data in chunks and decode it
//	var value interface{}
//	for {
//		err = decoder.Decode(&value)
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	// Encode the Go value as JSON
//	jsonData, err := json.MarshalIndent(value, "", "  ")
//	if err != nil {
//		panic(err)
//	}
//
//	// Print the JSON data to the console
//	fmt.Println(string(jsonData))
//}
