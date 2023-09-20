package output

import (
	"bufio"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"os"
)

type BucketInfo struct {
	Count             int
	BucketWork        map[int]int
	FirstHopWork      map[int]int
	BucketAccountings map[int]int
	BucketPayCount    map[int]int
	BucketPayment     map[int]int
	HopWork           map[int]int
	HopPayCount       map[int]int

	File   *os.File
	Writer *bufio.Writer
}

func InitBucketInfo() *BucketInfo {
	bi := BucketInfo{}
	bi.BucketWork = make(map[int]int)
	bi.FirstHopWork = make(map[int]int)
	bi.BucketAccountings = make(map[int]int)
	bi.BucketPayCount = make(map[int]int)
	bi.BucketPayment = make(map[int]int)
	bi.HopWork = make(map[int]int)
	bi.HopPayCount = make(map[int]int)

	bi.File = MakeFile("./results/buckets.txt")
	bi.Writer = bufio.NewWriter(bi.File)
	LogExpSting(bi.Writer)
	return &bi
}

func (bi *BucketInfo) Reset() {
	bi.BucketWork = make(map[int]int)
	bi.BucketAccountings = make(map[int]int)
	bi.BucketPayCount = make(map[int]int)
	bi.BucketPayment = make(map[int]int)
	bi.HopWork = make(map[int]int)
	bi.HopPayCount = make(map[int]int)
}

func (bi *BucketInfo) Close() {
	err := bi.Writer.Flush()
	if err != nil {
		fmt.Println("Couldn't flush the remaining buffer in the writer for bucket output")
	}
	err = bi.File.Close()
	if err != nil {
		fmt.Println("Couldn't close the file with filepath: ./results/buckets.txt")
	}
}

func (bi *BucketInfo) BucketPayRatio() []float64 {
	result := make([]float64, len(bi.BucketWork))

	for b, work := range bi.BucketWork {
		result[b] = float64(bi.BucketPayCount[b]) / float64(work)
	}

	return result
}

func (bi *BucketInfo) HopPayRatio() []float64 {
	result := make([]float64, len(bi.HopWork))

	for b, work := range bi.HopWork {
		result[b] = float64(bi.HopPayCount[b]) / float64(work)
	}

	return result
}

func (bi *BucketInfo) Update(output *Route) {
	bi.Count++
	if output.failed() {
		return
	}
	route := output.RouteWithPrices
	payments := output.PaymentsWithPrices
	for h, hop := range route {
		bin := config.GetBits() - general.BitLength(hop.RequesterNode.ToInt()^hop.ProviderNode.ToInt())
		if h == 0 {
			bi.FirstHopWork[bin]++
		}
		bi.BucketWork[bin]++
		bi.BucketAccountings[bin] += hop.Price
		bi.HopWork[h]++
		for _, payment := range payments {
			if payment.Payment.FirstNodeId == hop.RequesterNode {
				bi.BucketPayCount[bin]++
				bi.BucketPayment[bin] += hop.Price
				bi.HopPayCount[h]++
			}
		}
	}
}

func (bi *BucketInfo) Log() {
	_, err := bi.Writer.WriteString(fmt.Sprintf("\n Current count: %d\n", bi.Count))
	if err != nil {
		panic(err)
	}

	_, err = bi.Writer.WriteString("BucketUse: ")
	if err != nil {
		panic(err)
	}

	for bucket := 0; bucket < 11; bucket++ {
		_, err = bi.Writer.WriteString(fmt.Sprintf(" Bucket %d used %d times \n", bucket, bi.BucketWork[bucket]))
		if err != nil {
			panic(err)
		}
	}

	_, err = bi.Writer.WriteString("BucketUse as fraction of bucket 0: ")
	if err != nil {
		panic(err)
	}

	for bucket := 0; bucket < 11; bucket++ {
		_, err = bi.Writer.WriteString(fmt.Sprintf(" Bucket %d used %.3f times on first hop and %.3f times over all \n", bucket, float64(bi.FirstHopWork[bucket])/float64(bi.FirstHopWork[0]), float64(bi.BucketWork[bucket])/float64(bi.BucketWork[0])))
		if err != nil {
			panic(err)
		}
	}

	_, err = bi.Writer.WriteString("Accounting units transfered per bucket: \n")
	if err != nil {
		panic(err)
	}

	for bucket := 0; bucket < 11; bucket++ {
		_, err = bi.Writer.WriteString(fmt.Sprintf(" In bucket %d, transfer costs %.3f \n", bucket, float64(bi.BucketAccountings[bucket])/float64(bi.BucketWork[bucket])))
		if err != nil {
			panic(err)
		}
	}

	_, err = bi.Writer.WriteString("BucketPayRatio: ")
	if err != nil {
		panic(err)
	}
	for _, ratio := range bi.BucketPayRatio() {
		_, err = bi.Writer.WriteString(fmt.Sprintf("%.3f, ", ratio))
		if err != nil {
			panic(err)
		}
	}
	_, err = bi.Writer.WriteString("\n HopPayRatio: ")
	if err != nil {
		panic(err)
	}

	for _, ratio := range bi.HopPayRatio() {
		_, err = bi.Writer.WriteString(fmt.Sprintf("%.3f, ", ratio))
		if err != nil {
			panic(err)
		}
	}
}
