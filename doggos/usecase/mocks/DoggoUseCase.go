// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "DoggosPkg/doggos/usecase/model"

	mock "github.com/stretchr/testify/mock"
)

// DoggoUseCase is an autogenerated mock type for the DoggoUseCase type
type DoggoUseCase struct {
	mock.Mock
}

// GetDoggos provides a mock function with given fields: page, limit, breedID
func (_m *DoggoUseCase) GetDoggos(page int, limit int, breedID string) ([]model.Doggo, error) {
	ret := _m.Called(page, limit, breedID)

	var r0 []model.Doggo
	if rf, ok := ret.Get(0).(func(int, int, string) []model.Doggo); ok {
		r0 = rf(page, limit, breedID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Doggo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(page, limit, breedID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
