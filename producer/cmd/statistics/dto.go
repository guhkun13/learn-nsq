package statistics

type Payload struct {
	JobId                 string `json:"job_id"`
	TaskId                string `json:"task_id"`
	UserId                string `json:"user_id"`
	CompressionId         string `json:"compression_id"`
	MachineId             string `json:"machine_id"`
	CompressorId          string `json:"compressor_id"`
	Filename              string `json:"filename"`
	FormatFile            string `json:"format_file"`
	OriginalSize          string `json:"original_size"`
	CompressedSize        string `json:"compressed_size"`
	CompressedDuration    string `json:"compressed_duration"`
	SavingSpacePercentage string `json:"saving_space_percentage"`
	StartedAt             string `json:"started_at"`
	FinishedAt            string `json:"finished_at"`
	Timestamp             string `json:"timestamp"`
}

// {
// 	"job_id": "01HVK0BN2EBX6JMKXNKNP8SM7J",
// 	"task_id": "2cf527a6-3dbe-4ac5-bdd5-19e5c7e66aab",
// 	"compression_id": "01HVK0BN2C3WKVCJ52X8TZWG76",
// 	"user_id": "",
// 	"machine_id": "b9:03:0f",
// 	"compressor_id": "compressor-1",
// 	"filename": "wijaya.jpeg",
// 	"format_file": "jpeg",
// 	"original_size": "135.65430",
// 	"compressed_size": "96.70898",
// 	"compressed_duration": "0.08386",
// 	"compressed_ratio": "28.70924",
// 	"started_at": "2024-04-16T08:28:46Z",
// 	"finished_at": "2024-04-16T08:28:46Z"
// 	"timestamp": "2024-04-16T08:28:46Z",
//   }
