package output

import (
	"bufio"
	"fmt"
	"os"
)

type RerouteInfo struct {
	SuccessfullChunks int
	Count             int
	File              *os.File
	Writer            *bufio.Writer
}

func InitRerouteInfo() *RerouteInfo {
	rri := RerouteInfo{}
	rri.File = MakeFile("./results/reroute.txt")
	rri.Writer = bufio.NewWriter(rri.File)
	LogExpSting(rri.Writer)
	return &rri
}

func (rri *RerouteInfo) Close() {
	err := rri.Writer.Flush()
	if err != nil {
		fmt.Println("Couldn't flush the remaining buffer in the writer for reroute output")
	}
	err = rri.File.Close()
	if err != nil {
		fmt.Println("Couldn't close the file with filepath: ./results/reroute.txt")
	}
}

func (rri *RerouteInfo) Reset() {
	rri.SuccessfullChunks = 0
	rri.Count = 0
}

func (rri *RerouteInfo) Update(output *Route) {
	if output.Found {
		rri.SuccessfullChunks++
		rri.Count += output.NumRetries
	}
}

func (rri *RerouteInfo) Log() {
	avgRetryPerSuccess := float64(rri.Count) / float64(rri.SuccessfullChunks)
	_, err := rri.Writer.WriteString(fmt.Sprintf("Average retry per successfull downloads: %.2f%%  \n", avgRetryPerSuccess))
	if err != nil {
		panic(err)
	}
}
