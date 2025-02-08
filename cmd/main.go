package main

import (
	"main/application"
	"main/cache"
	"main/reader"
	"main/service"
	"main/writer"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	inputFilePath := "./input.csv"
	outputFilePath := "./output.csv"

	csvReader := reader.NewCSVReader(inputFilePath, logger)
	write := writer.NewCSVWriter(outputFilePath, logger)

	// write := writer.NewStdOut()

	cache := cache.NewMemoryCache()
	userService := service.NewUserService(cache)

	err := application.NewApplication(logger, csvReader, write, userService).Run()
	if err != nil {
		logger.Error("Failed to run the application", zap.Error(err))
		os.Exit(1)
	}
}
