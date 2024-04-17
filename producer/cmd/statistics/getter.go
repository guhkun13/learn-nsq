package statistics

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func getRandomIndex(maxIndex int) int {
	return rand.Intn(maxIndex)
}

func generateULID() string {
	return ulid.Make().String()
}

func generateId() string {
	return generateULID()
}

func getRandomJobId() string {
	countItems := len(JobIds)
	randomIdx := getRandomIndex(countItems - 1)

	return JobIds[randomIdx]
}

func getRandomUserId() string {
	countItems := len(UserIds)
	randomIdx := getRandomIndex(countItems - 1)

	return UserIds[randomIdx]
}

func getRandomCompressionId() string {
	countItems := len(CompressionIds)
	randomIdx := getRandomIndex(countItems - 1)

	return CompressionIds[randomIdx]
}

func getRandomMachineId() string {
	countItems := len(MachineIds)
	randomIdx := getRandomIndex(countItems - 1)

	return MachineIds[randomIdx]
}

func getRandomCompressorId() string {
	countItems := len(CompressorIds)
	randomIdx := getRandomIndex(countItems - 1)

	return CompressorIds[randomIdx]
}

func getRandomOrizinalSize() string {
	countItems := len(OriginalSizes)
	randomIdx := getRandomIndex(countItems - 1)
	val := OriginalSizes[randomIdx]

	return fmt.Sprintf("%.2f", val)
}

func getRandomCompressedsize() string {
	countItems := len(CompressedSizes)
	randomIdx := getRandomIndex(countItems - 1)
	val := CompressedSizes[randomIdx]

	return fmt.Sprintf("%.2f", val)
}

func getRandomCompressedDurationFloat() float64 {
	countItems := len(CompressedDurations)
	randomIdx := getRandomIndex(countItems - 1)
	return CompressedDurations[randomIdx]
}

func getRandomCompressedDurationInt(val float64) int {
	return int(math.Ceil(val))
}

func getRandomCompressedDurationStr(val float64) string {
	return fmt.Sprintf("%.2f", val)
}

func getRandomCompressedRatio() string {
	countItems := len(CompressedRatios)
	randomIdx := getRandomIndex(countItems - 1)
	val := CompressedRatios[randomIdx]

	return fmt.Sprintf("%.2f", val)
}

func getTimestamp(addedSecond int) string {
	currentTime := time.Now().Add(time.Second * time.Duration(addedSecond))
	layoutFormat := time.RFC3339
	return currentTime.Format(layoutFormat)
}

func generatePayload() Payload {
	randomJobId := getRandomJobId()
	randomUserId := getRandomUserId()
	randomMachineId := getRandomMachineId()
	randomCompressionId := getRandomCompressionId()
	randomCompressorId := getRandomCompressorId()
	randomOriginalSize := getRandomOrizinalSize()
	randomCompressedSize := getRandomCompressedsize()
	randomCompressedDurationFloat := getRandomCompressedDurationFloat()
	randomCompressedDurationStr := getRandomCompressedDurationStr(randomCompressedDurationFloat)
	randomCompressedDurationInt := getRandomCompressedDurationInt(randomCompressedDurationFloat)
	randomCompressedRatio := getRandomCompressedRatio()

	return Payload{
		JobId:              randomJobId,
		TaskId:             generateId(),
		UserId:             randomUserId,
		CompressionId:      randomCompressionId,
		MachineId:          randomMachineId,
		CompressorId:       randomCompressorId,
		Filename:           "test.mp4",
		FormatFile:         "mp4",
		OriginalSize:       randomOriginalSize,
		CompressedSize:     randomCompressedSize,
		CompressedDuration: randomCompressedDurationStr,
		CompressedRatio:    randomCompressedRatio,
		StartedAt:          getTimestamp(0),
		FinishedAt:         getTimestamp(randomCompressedDurationInt),
		Timestamp:          getTimestamp(0),
	}

}
