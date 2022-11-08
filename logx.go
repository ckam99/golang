package logx

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Error(stout interface{})
	Errorf(f string, stout ...interface{})
	Info(stout interface{})
	Infof(f string, stout ...interface{})
	Warn(stout interface{})
	Warnf(f string, stout ...interface{})
	Fatal(stout interface{})
	Fatalf(f string, stout ...interface{})
	Close()
	GetLogFile() *os.File
}

type logger struct {
	logger *log.Logger
	file   *os.File
}

func New() Logger {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	return &logger{
		logger: log.New(file, "", log.LstdFlags),
		file:   file,
	}
}

// Error implements Logger
func (l *logger) Error(stout interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" [ERROR]", stout)
	l.logger.Println("[ERROR]", stout)
}

// Infof implements Logger
func (l *logger) Errorf(f string, stout ...interface{}) {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+" [ERROR] "+f, stout...)
	l.logger.Printf("[ERROR] "+f, stout...)

}

// Info implements Logger
func (l *logger) Info(stout interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" [INFO]", stout)
	l.logger.Println("[INFO]", stout)
}

// Infof implements Logger
func (l *logger) Infof(f string, stout ...interface{}) {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+" [INFO] "+f, stout...)
	l.logger.Printf("[INFO] "+f, stout...)
}

// Warning implements Logger
func (l *logger) Warn(stout interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" [WARNING]", stout)
	l.logger.Println("[WARNING]", stout)
}

func (l *logger) Warnf(f string, stout ...interface{}) {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+" [WARNING] "+f, stout...)
	l.logger.Printf("[WARNING] "+f, stout...)
}

// Warning implements Logger
func (l *logger) Fatal(stout interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+" [FATAL]", stout)
	l.logger.Fatalln("[FATAL]", stout)
}

// Warning implements Logger
func (l *logger) Fatalf(f string, stout ...interface{}) {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+" [FATAL] "+f, stout...)
	l.logger.Fatalf("[FATAL] "+f, stout...)
}

func (l *logger) Close() {
	l.file.Close()
}

func (l *logger) GetLogFile() *os.File {
	return l.file
}
