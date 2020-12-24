package resource

import (
	"fmt"
	"log"
	"time"

	go_logger "github.com/phachon/go-logger"
	"github.com/spf13/viper"
)

var (
	Logger *go_logger.Logger
)

// InitLogger 初始化 logger
func InitLogger() {
	log.Println("init logger start")
	Logger = go_logger.NewLogger()
	logFormat := "[%timestamp_format%] [%level_string%] %function%:%line% %body%"
	logDir := viper.GetString("logDir")
	current := time.Now()
	// file adapter config
	fileConfig := &go_logger.FileConfig{
		//Filename: "./reviewer-test.log", // The file name of the logger output, does not exist automatically
		// If you want to separate separate logs into files, configure LevelFileName parameters.
		LevelFileName: map[int]string{
			Logger.LoggerLevel("error"): fmt.Sprintf("./spiderErr%d.log", current.Unix()),   // The error level log is written to the error.log file.
			Logger.LoggerLevel("info"):  fmt.Sprintf("./spiderInfo%d.log", current.Unix()),  // The info level log is written to the info.log file.
			Logger.LoggerLevel("debug"): fmt.Sprintf("./spiderDebug%d.log", current.Unix()), // The debug level log is written to the debug.log file.
		},
		MaxSize:    0,         // File maximum (KB), default 0 is not limited
		MaxLine:    0,         // The maximum number of lines in the file, the default 0 is not limited
		DateSlice:  "y",       // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
		JsonFormat: false,     // Whether the file data is written to JSON formatting
		Format:     logFormat, // JsonFormat is false, logger message written to file format string
	}
	// add output to the file
	_ = Logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	log.Println("log to dir : " + logDir)
}
