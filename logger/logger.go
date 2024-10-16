package logger

import (
	"cyberghostvpn-gui/tools"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _loggerFilename string

type zerologger struct {
	*zerolog.Logger
}

type LoggerOptions struct {
	Console       bool   // Enable/disable console output (default: false)
	JSONFormatted bool   // Enable/disable JSON formatting log (default: true)
	FileDirectory string // File directory (default: "./logs/")
	FileName      string // File name (default: "_lastlog.log")
	FilePath      string // File path (default: "./logs/_lastlog.log")
	Level         string // Level of log (default: "info") - Options: debug, info, warn, error, fatal, panic
	MaxFileSize   int    // size max in MB
	MaxBackups    int    // Max number of files to keep
	MaxInDays     int    // Max number of days to write into a file
	Compressed    bool   // Enable/disable compression
}

var currentLoggerConfig *LoggerOptions
var currentLogger *zerologger
var loggerWriters []io.Writer

// AddLoggerUIWriter adds a writer to the list of writers that will be used to
// print the logs. The writer should be a UI component that can display text.
// The timeFormat parameter is optional and defaults to the format
// "2006-02-01 15:04:05.000" if an empty string is passed.
func AddLoggerUIWriter(writer io.Writer, timeFormat string) {
	if writer == nil {
		return
	}
	if len(timeFormat) < 1 {
		timeFormat = "2006-02-01 15:04:05.000"
	}
	loggerWriters = append(loggerWriters, zerolog.ConsoleWriter{Out: writer, TimeFormat: timeFormat})
}

// LoggerInit sets up the logging configuration based on the provided options.
// If the options parameter is not nil, it assigns the options to currentLoggerConfig.
// Otherwise, it calls checkLoggerConfig to set default configuration settings.
func LoggerInit(options *LoggerOptions) {
	if options != nil {
		currentLoggerConfig = options
	} else {
		checkLoggerConfig()
	}
}

// GetNewLoggerFileName returns the filename for the new logger based on the current timestamp.
// If the _loggerFilename is not empty, it returns _loggerFilename.
// Otherwise, it generates a filename using the current timestamp in the format "2006-01-02_15-04-05_output.logs".
func GetNewLoggerFileName() string {
	if len(_loggerFilename) > 0 {
		return _loggerFilename
	}
	return fmt.Sprintf("%s_output.logs", time.Now().Format("2006-01-02_15-04-05"))
}

// GetNewLoggerOptions returns the default LoggerOptions.
func GetNewLoggerOptions() LoggerOptions {
	return LoggerOptions{}
}

// GetLoggerWriters returns a pointer to the loggerWriters slice, which is a slice of writer that log will be written to.
// The returned slice is a pointer and can be modified by the caller.
func GetLoggerWriters() *[]io.Writer {
	return &loggerWriters
}

// SetLogLevel sets the global log level according to the given parameter.
//
// It checks the log level setting in the configuration and sets the global log
// level accordingly. If the log level is not recognized, it defaults to WarnLevel.
func SetLogLevel(level string) {
	// Check config
	cfg := checkLoggerConfig()

	switch strings.ToLower(cfg.Level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
}

// If no settings has been set, these default settings will be applied
func checkLoggerConfig() LoggerOptions {
	if currentLoggerConfig == nil {
		if tools.IsFileExists("./logs/last.log") {
			os.Rename("./logs/last.log", "./logs/previous.log")
		}
		fileDir := "./logs/"
		fileName := "last.log" //GetNewLoggerFileName()
		currentLoggerConfig = &LoggerOptions{
			Console:       false,
			JSONFormatted: false,
			Compressed:    true,
			FileDirectory: fileDir,
			FileName:      fileName,
			FilePath:      fmt.Sprintf("%s%s", fileDir, fileName),
			Level:         "info",
			MaxFileSize:   50, // size max in MB
			MaxBackups:    3,  // Max number of files to keep
			MaxInDays:     1,  // Max number of days to write into a file
		}
	}

	return *currentLoggerConfig
}

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func configureLogger() *zerologger {
	// Check config
	cfg := checkLoggerConfig()

	// Set time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	// Set log outputs
	if currentLoggerConfig.Console {
		loggerWriters = append(loggerWriters, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
	} else {
		loggerWriters = append(loggerWriters, newLoggerRollingFile())
	}
	mw := io.MultiWriter(loggerWriters...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = "15:04:05.000"
	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", !cfg.Console).
		Bool("jsonLogOutput", cfg.JSONFormatted).
		Str("logDirectory", cfg.FileDirectory).
		Str("fileName", cfg.FileName).
		Int("maxSizeMB", cfg.MaxFileSize).
		Int("maxBackups", cfg.MaxBackups).
		Int("maxAgeInDays", cfg.MaxInDays).
		Bool("compression", cfg.Compressed).
		Msg("logging configured")

	return &zerologger{
		Logger: &logger,
	}
}

// newLoggerRollingFile sets up and returns a new rolling file writer for logging.
// It ensures the log directory exists and creates a rolling logger with specified
// configuration options such as file path, max size, max backups, max age, and compression.
// If the log directory cannot be created, it logs the error and returns nil.
func newLoggerRollingFile() io.Writer {
	// Check config
	cfg := checkLoggerConfig()

	if err := os.MkdirAll(cfg.FileDirectory, 0744); err != nil {
		log.Error().Err(err).Str("path", cfg.FileDirectory).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxFileSize, // MB
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxInDays, // Days
		Compress:   cfg.Compressed,
	}
}

// initLogger initializes the logger engine according to the current configuration.
// It checks the configuration options, sets up the logger engine, and configures
// the log level according to the given configuration.
func initLogger() {
	// Check config
	checkLoggerConfig()

	// Configure the logger engine
	currentLogger = configureLogger()
	setLogLevel()
	currentLogger.Info().Msg("Logger initialized and ready")
}

// setLogLevel sets the global log level according to the logger configuration.
// It checks the log level setting in the configuration and sets the global log
// level accordingly. If the log level is not recognized, it defaults to WarnLevel.
func setLogLevel() {
	// Check config
	cfg := checkLoggerConfig()

	switch strings.ToLower(cfg.Level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
}

// GetCurrentLogger returns the current logger instance. If the logger is not
// initialized, it initializes the logger before returning it. This ensures
// that the logger is always ready for logging operations.
func GetCurrentLogger() *zerologger {
	if currentLogger == nil {
		initLogger()
	}

	return currentLogger
}

// Debug logs a message at DebugLevel. The message includes the text passed as
// argument.
func Debug(text string) {
	GetCurrentLogger().Debug().Msg(text)
}

// Debugf logs a message at DebugLevel. The message includes the text passed as
// argument and is formatted according to the format string and any additional
// values passed as arguments.
func Debugf(format string, values ...interface{}) {
	GetCurrentLogger().Debug().Msgf(format, values...)
}

// Info logs a message at InfoLevel. The message includes the text passed as
// argument.
func Info(text string) {
	GetCurrentLogger().Info().Msg(text)
}

// Infof logs a message at InfoLevel. The message includes the text passed as
// argument and is formatted according to the format string and any additional
// values passed as arguments.
func Infof(format string, values ...interface{}) {
	GetCurrentLogger().Info().Msgf(format, values...)
}

// Warn logs a message at WarnLevel. The message includes the text passed as
// argument.
func Warn(text string) {
	GetCurrentLogger().Warn().Msg(text)
}

// Warnf logs a message at WarnLevel. The message includes the text passed as
// argument and is formatted according to the format string and any additional
// values passed as arguments.
func Warnf(format string, values ...interface{}) {
	GetCurrentLogger().Warn().Msgf(format, values...)
}

// Error logs a message at ErrorLevel. The message includes the text passed as
// argument.
func Error(text string) {
	GetCurrentLogger().Error().Msg(text)
}

// Errorf logs a message at ErrorLevel. The message includes the text passed as
// argument and is formatted according to the format string and any additional
// values passed as arguments.
func Errorf(format string, values ...interface{}) {
	GetCurrentLogger().Error().Msgf(format, values...)
}

// Panic logs a message at PanicLevel. The message includes the text passed as argument.
func Panic(text string) {
	GetCurrentLogger().Panic().Msg(text)
}

// Panicf logs a message at PanicLevel. The message includes the text passed as
// argument and is formatted according to the format string and any additional
// values passed as arguments.
func Panicf(format string, values ...interface{}) {
	GetCurrentLogger().Panic().Msgf(format, values...)
}

// Fatal logs a message at FatalLevel. The message includes the text passed as
// argument.
func Fatal(text string) {
	GetCurrentLogger().Fatal().Msg(text)
}

// Fatalf logs a message at FatalLevel. The message includes the text passed as
// argument and is formatted according to the format string and any additional
// values passed as arguments.
func Fatalf(format string, values ...interface{}) {
	GetCurrentLogger().Fatal().Msgf(format, values...)
}
