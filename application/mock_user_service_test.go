// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package application

import (
	"main/model"
	"sync"
)

// Ensure, that UserServiceMock does implement UserService.
// If this is not the case, regenerate this file with moq.
var _ UserService = &UserServiceMock{}

// UserServiceMock is a mock implementation of UserService.
//
//	func TestSomethingThatUsesUserService(t *testing.T) {
//
//		// make and configure a mocked UserService
//		mockedUserService := &UserServiceMock{
//			CompareAndClassifyFunc: func(users []model.User) []model.CompareResult {
//				panic("mock out the CompareAndClassify method")
//			},
//		}
//
//		// use mockedUserService in code that requires UserService
//		// and then make assertions.
//
//	}
type UserServiceMock struct {
	// CompareAndClassifyFunc mocks the CompareAndClassify method.
	CompareAndClassifyFunc func(users []model.User) []model.CompareResult

	// calls tracks calls to the methods.
	calls struct {
		// CompareAndClassify holds details about calls to the CompareAndClassify method.
		CompareAndClassify []struct {
			// Users is the users argument value.
			Users []model.User
		}
	}
	lockCompareAndClassify sync.RWMutex
}

// CompareAndClassify calls CompareAndClassifyFunc.
func (mock *UserServiceMock) CompareAndClassify(users []model.User) []model.CompareResult {
	if mock.CompareAndClassifyFunc == nil {
		panic("UserServiceMock.CompareAndClassifyFunc: method is nil but UserService.CompareAndClassify was just called")
	}
	callInfo := struct {
		Users []model.User
	}{
		Users: users,
	}
	mock.lockCompareAndClassify.Lock()
	mock.calls.CompareAndClassify = append(mock.calls.CompareAndClassify, callInfo)
	mock.lockCompareAndClassify.Unlock()
	return mock.CompareAndClassifyFunc(users)
}

// CompareAndClassifyCalls gets all the calls that were made to CompareAndClassify.
// Check the length with:
//
//	len(mockedUserService.CompareAndClassifyCalls())
func (mock *UserServiceMock) CompareAndClassifyCalls() []struct {
	Users []model.User
} {
	var calls []struct {
		Users []model.User
	}
	mock.lockCompareAndClassify.RLock()
	calls = mock.calls.CompareAndClassify
	mock.lockCompareAndClassify.RUnlock()
	return calls
}
