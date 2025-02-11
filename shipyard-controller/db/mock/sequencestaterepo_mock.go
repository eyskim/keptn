// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package db_mock

import (
	"github.com/keptn/go-utils/pkg/api/models"
	"sync"
)

// SequenceStateRepoMock is a mock implementation of db.SequenceStateRepo.
//
// 	func TestSomethingThatUsesSequenceStateRepo(t *testing.T) {
//
// 		// make and configure a mocked db.SequenceStateRepo
// 		mockedSequenceStateRepo := &SequenceStateRepoMock{
// 			CreateSequenceStateFunc: func(state models.SequenceState) error {
// 				panic("mock out the CreateSequenceState method")
// 			},
// 			DeleteSequenceStatesFunc: func(filter models.StateFilter) error {
// 				panic("mock out the DeleteSequenceStates method")
// 			},
// 			FindSequenceStatesFunc: func(filter models.StateFilter) (*models.SequenceStates, error) {
// 				panic("mock out the FindSequenceStates method")
// 			},
// 			UpdateSequenceStateFunc: func(state models.SequenceState) error {
// 				panic("mock out the UpdateSequenceState method")
// 			},
// 		}
//
// 		// use mockedSequenceStateRepo in code that requires db.SequenceStateRepo
// 		// and then make assertions.
//
// 	}
type SequenceStateRepoMock struct {
	// CreateSequenceStateFunc mocks the CreateSequenceState method.
	CreateSequenceStateFunc func(state models.SequenceState) error

	// DeleteSequenceStatesFunc mocks the DeleteSequenceStates method.
	DeleteSequenceStatesFunc func(filter models.StateFilter) error

	// FindSequenceStatesFunc mocks the FindSequenceStates method.
	FindSequenceStatesFunc func(filter models.StateFilter) (*models.SequenceStates, error)

	// UpdateSequenceStateFunc mocks the UpdateSequenceState method.
	UpdateSequenceStateFunc func(state models.SequenceState) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateSequenceState holds details about calls to the CreateSequenceState method.
		CreateSequenceState []struct {
			// State is the state argument value.
			State models.SequenceState
		}
		// DeleteSequenceStates holds details about calls to the DeleteSequenceStates method.
		DeleteSequenceStates []struct {
			// Filter is the filter argument value.
			Filter models.StateFilter
		}
		// FindSequenceStates holds details about calls to the FindSequenceStates method.
		FindSequenceStates []struct {
			// Filter is the filter argument value.
			Filter models.StateFilter
		}
		// UpdateSequenceState holds details about calls to the UpdateSequenceState method.
		UpdateSequenceState []struct {
			// State is the state argument value.
			State models.SequenceState
		}
	}
	lockCreateSequenceState  sync.RWMutex
	lockDeleteSequenceStates sync.RWMutex
	lockFindSequenceStates   sync.RWMutex
	lockUpdateSequenceState  sync.RWMutex
}

// CreateSequenceState calls CreateSequenceStateFunc.
func (mock *SequenceStateRepoMock) CreateSequenceState(state models.SequenceState) error {
	if mock.CreateSequenceStateFunc == nil {
		panic("SequenceStateRepoMock.CreateSequenceStateFunc: method is nil but SequenceStateRepo.CreateSequenceState was just called")
	}
	callInfo := struct {
		State models.SequenceState
	}{
		State: state,
	}
	mock.lockCreateSequenceState.Lock()
	mock.calls.CreateSequenceState = append(mock.calls.CreateSequenceState, callInfo)
	mock.lockCreateSequenceState.Unlock()
	return mock.CreateSequenceStateFunc(state)
}

// CreateSequenceStateCalls gets all the calls that were made to CreateSequenceState.
// Check the length with:
//     len(mockedSequenceStateRepo.CreateSequenceStateCalls())
func (mock *SequenceStateRepoMock) CreateSequenceStateCalls() []struct {
	State models.SequenceState
} {
	var calls []struct {
		State models.SequenceState
	}
	mock.lockCreateSequenceState.RLock()
	calls = mock.calls.CreateSequenceState
	mock.lockCreateSequenceState.RUnlock()
	return calls
}

// DeleteSequenceStates calls DeleteSequenceStatesFunc.
func (mock *SequenceStateRepoMock) DeleteSequenceStates(filter models.StateFilter) error {
	if mock.DeleteSequenceStatesFunc == nil {
		panic("SequenceStateRepoMock.DeleteSequenceStatesFunc: method is nil but SequenceStateRepo.DeleteSequenceStates was just called")
	}
	callInfo := struct {
		Filter models.StateFilter
	}{
		Filter: filter,
	}
	mock.lockDeleteSequenceStates.Lock()
	mock.calls.DeleteSequenceStates = append(mock.calls.DeleteSequenceStates, callInfo)
	mock.lockDeleteSequenceStates.Unlock()
	return mock.DeleteSequenceStatesFunc(filter)
}

// DeleteSequenceStatesCalls gets all the calls that were made to DeleteSequenceStates.
// Check the length with:
//     len(mockedSequenceStateRepo.DeleteSequenceStatesCalls())
func (mock *SequenceStateRepoMock) DeleteSequenceStatesCalls() []struct {
	Filter models.StateFilter
} {
	var calls []struct {
		Filter models.StateFilter
	}
	mock.lockDeleteSequenceStates.RLock()
	calls = mock.calls.DeleteSequenceStates
	mock.lockDeleteSequenceStates.RUnlock()
	return calls
}

// FindSequenceStates calls FindSequenceStatesFunc.
func (mock *SequenceStateRepoMock) FindSequenceStates(filter models.StateFilter) (*models.SequenceStates, error) {
	if mock.FindSequenceStatesFunc == nil {
		panic("SequenceStateRepoMock.FindSequenceStatesFunc: method is nil but SequenceStateRepo.FindSequenceStates was just called")
	}
	callInfo := struct {
		Filter models.StateFilter
	}{
		Filter: filter,
	}
	mock.lockFindSequenceStates.Lock()
	mock.calls.FindSequenceStates = append(mock.calls.FindSequenceStates, callInfo)
	mock.lockFindSequenceStates.Unlock()
	return mock.FindSequenceStatesFunc(filter)
}

// FindSequenceStatesCalls gets all the calls that were made to FindSequenceStates.
// Check the length with:
//     len(mockedSequenceStateRepo.FindSequenceStatesCalls())
func (mock *SequenceStateRepoMock) FindSequenceStatesCalls() []struct {
	Filter models.StateFilter
} {
	var calls []struct {
		Filter models.StateFilter
	}
	mock.lockFindSequenceStates.RLock()
	calls = mock.calls.FindSequenceStates
	mock.lockFindSequenceStates.RUnlock()
	return calls
}

// UpdateSequenceState calls UpdateSequenceStateFunc.
func (mock *SequenceStateRepoMock) UpdateSequenceState(state models.SequenceState) error {
	if mock.UpdateSequenceStateFunc == nil {
		panic("SequenceStateRepoMock.UpdateSequenceStateFunc: method is nil but SequenceStateRepo.UpdateSequenceState was just called")
	}
	callInfo := struct {
		State models.SequenceState
	}{
		State: state,
	}
	mock.lockUpdateSequenceState.Lock()
	mock.calls.UpdateSequenceState = append(mock.calls.UpdateSequenceState, callInfo)
	mock.lockUpdateSequenceState.Unlock()
	return mock.UpdateSequenceStateFunc(state)
}

// UpdateSequenceStateCalls gets all the calls that were made to UpdateSequenceState.
// Check the length with:
//     len(mockedSequenceStateRepo.UpdateSequenceStateCalls())
func (mock *SequenceStateRepoMock) UpdateSequenceStateCalls() []struct {
	State models.SequenceState
} {
	var calls []struct {
		State models.SequenceState
	}
	mock.lockUpdateSequenceState.RLock()
	calls = mock.calls.UpdateSequenceState
	mock.lockUpdateSequenceState.RUnlock()
	return calls
}
