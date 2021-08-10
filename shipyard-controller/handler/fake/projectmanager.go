// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/keptn/shipyard-controller/common"
	"github.com/keptn/keptn/shipyard-controller/models"
	"github.com/keptn/keptn/shipyard-controller/operations"
	"sync"
)

// IProjectManagerMock is a mock implementation of handler.IProjectManager.
//
// 	func TestSomethingThatUsesIProjectManager(t *testing.T) {
//
// 		// make and configure a mocked handler.IProjectManager
// 		mockedIProjectManager := &IProjectManagerMock{
// 			CreateFunc: func(params *operations.CreateProjectParams) (error, common.RollbackFunc) {
// 				panic("mock out the Create method")
// 			},
// 			DeleteFunc: func(projectName string) (string, error) {
// 				panic("mock out the Delete method")
// 			},
// 			GetFunc: func() ([]*models.ExpandedProject, error) {
// 				panic("mock out the Get method")
// 			},
// 			GetByNameFunc: func(projectName string) (*models.ExpandedProject, error) {
// 				panic("mock out the GetByName method")
// 			},
// 			UpdateFunc: func(params *operations.UpdateProjectParams) (error, common.RollbackFunc) {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedIProjectManager in code that requires handler.IProjectManager
// 		// and then make assertions.
//
// 	}
type IProjectManagerMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(params *operations.CreateProjectParams) (error, common.RollbackFunc)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(projectName string) (string, error)

	// GetFunc mocks the Get method.
	GetFunc func() ([]*models.ExpandedProject, error)

	// GetByNameFunc mocks the GetByName method.
	GetByNameFunc func(projectName string) (*models.ExpandedProject, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(params *operations.UpdateProjectParams) (error, common.RollbackFunc)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Params is the params argument value.
			Params *operations.CreateProjectParams
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
		}
		// Get holds details about calls to the Get method.
		Get []struct {
		}
		// GetByName holds details about calls to the GetByName method.
		GetByName []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Params is the params argument value.
			Params *operations.UpdateProjectParams
		}
	}
	lockCreate    sync.RWMutex
	lockDelete    sync.RWMutex
	lockGet       sync.RWMutex
	lockGetByName sync.RWMutex
	lockUpdate    sync.RWMutex
}

// Create calls CreateFunc.
func (mock *IProjectManagerMock) Create(params *operations.CreateProjectParams) (error, common.RollbackFunc) {
	if mock.CreateFunc == nil {
		panic("IProjectManagerMock.CreateFunc: method is nil but IProjectManager.Create was just called")
	}
	callInfo := struct {
		Params *operations.CreateProjectParams
	}{
		Params: params,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(params)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedIProjectManager.CreateCalls())
func (mock *IProjectManagerMock) CreateCalls() []struct {
	Params *operations.CreateProjectParams
} {
	var calls []struct {
		Params *operations.CreateProjectParams
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *IProjectManagerMock) Delete(projectName string) (string, error) {
	if mock.DeleteFunc == nil {
		panic("IProjectManagerMock.DeleteFunc: method is nil but IProjectManager.Delete was just called")
	}
	callInfo := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(projectName)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedIProjectManager.DeleteCalls())
func (mock *IProjectManagerMock) DeleteCalls() []struct {
	ProjectName string
} {
	var calls []struct {
		ProjectName string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *IProjectManagerMock) Get() ([]*models.ExpandedProject, error) {
	if mock.GetFunc == nil {
		panic("IProjectManagerMock.GetFunc: method is nil but IProjectManager.Get was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc()
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedIProjectManager.GetCalls())
func (mock *IProjectManagerMock) GetCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetByName calls GetByNameFunc.
func (mock *IProjectManagerMock) GetByName(projectName string) (*models.ExpandedProject, error) {
	if mock.GetByNameFunc == nil {
		panic("IProjectManagerMock.GetByNameFunc: method is nil but IProjectManager.GetByName was just called")
	}
	callInfo := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	mock.lockGetByName.Lock()
	mock.calls.GetByName = append(mock.calls.GetByName, callInfo)
	mock.lockGetByName.Unlock()
	return mock.GetByNameFunc(projectName)
}

// GetByNameCalls gets all the calls that were made to GetByName.
// Check the length with:
//     len(mockedIProjectManager.GetByNameCalls())
func (mock *IProjectManagerMock) GetByNameCalls() []struct {
	ProjectName string
} {
	var calls []struct {
		ProjectName string
	}
	mock.lockGetByName.RLock()
	calls = mock.calls.GetByName
	mock.lockGetByName.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *IProjectManagerMock) Update(params *operations.UpdateProjectParams) (error, common.RollbackFunc) {
	if mock.UpdateFunc == nil {
		panic("IProjectManagerMock.UpdateFunc: method is nil but IProjectManager.Update was just called")
	}
	callInfo := struct {
		Params *operations.UpdateProjectParams
	}{
		Params: params,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(params)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedIProjectManager.UpdateCalls())
func (mock *IProjectManagerMock) UpdateCalls() []struct {
	Params *operations.UpdateProjectParams
} {
	var calls []struct {
		Params *operations.UpdateProjectParams
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
