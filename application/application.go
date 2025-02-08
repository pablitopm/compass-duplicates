package application

import (
	"main/model"

	"go.uber.org/zap"
)

type Reader interface {
	Read() ([]model.User, error)
}

type Writer interface {
	Write(result []model.CompareResult) error
}

type UserService interface {
	CompareAndClassify(users []model.User) []model.CompareResult
}

type Application struct {
	logger      *zap.Logger
	reader      Reader
	writer      Writer
	userService UserService
}

func NewApplication(logger *zap.Logger, reader Reader, writer Writer, userService UserService) *Application {
	return &Application{
		logger:      logger,
		reader:      reader,
		writer:      writer,
		userService: userService,
	}
}

func (a Application) Run() error {
	users, err := a.reader.Read()
	if err != nil {
		a.logger.Error("Failed to read input", zap.Error(err))
		return err
	}
	results := a.userService.CompareAndClassify(users)
	err = a.writer.Write(results)
	if err != nil {
		a.logger.Error("Failed to write output", zap.Error(err))
		return err
	}
	return nil
}
