package main

import (
	"io"
	"log/slog"
	"math/rand/v2"
	"os"
)


func main(){
	handlerOpt := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		//AddSource: true,
	}				

	
	logFile,err := os.OpenFile("app.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666);
	if err!=nil{
		slog.Error("Error opening file : "+err.Error())
	}

	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stderr,logFile)

	logger := slog.New(slog.NewJSONHandler(multiWriter,handlerOpt))
	slog.SetDefault(logger)

	user := slog.Group("user","id",rand.IntN(70002003),)

	//creating a child logger
	IdLogger := logger.With(user)

	slog.Debug("Debug code")
	slog.Info("hello",user)
	IdLogger.Error("error occured")

}