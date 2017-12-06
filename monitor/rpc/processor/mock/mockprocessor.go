// Code generated by MockGen. DO NOT EDIT.
// Source: monitor/rpc/processor/interfaces.go

// Package mockprocessor is a generated GoMock package.
package mockprocessor

import (
	reflect "reflect"

	constants "github.com/aporeto-inc/trireme-lib/constants"
	events "github.com/aporeto-inc/trireme-lib/monitor/rpc/events"
	processor "github.com/aporeto-inc/trireme-lib/monitor/rpc/processor"
	policy "github.com/aporeto-inc/trireme-lib/policy"
	gomock "github.com/golang/mock/gomock"
)

// MockProcessingUnitsHandler is a mock of ProcessingUnitsHandler interface
// nolint
type MockProcessingUnitsHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingUnitsHandlerMockRecorder
}

// MockProcessingUnitsHandlerMockRecorder is the mock recorder for MockProcessingUnitsHandler
// nolint
type MockProcessingUnitsHandlerMockRecorder struct {
	mock *MockProcessingUnitsHandler
}

// NewMockProcessingUnitsHandler creates a new mock instance
// nolint
func NewMockProcessingUnitsHandler(ctrl *gomock.Controller) *MockProcessingUnitsHandler {
	mock := &MockProcessingUnitsHandler{ctrl: ctrl}
	mock.recorder = &MockProcessingUnitsHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockProcessingUnitsHandler) EXPECT() *MockProcessingUnitsHandlerMockRecorder {
	return m.recorder
}

// CreatePURuntime mocks base method
// nolint
func (m *MockProcessingUnitsHandler) CreatePURuntime(contextID string, runtimeInfo *policy.PURuntime) error {
	ret := m.ctrl.Call(m, "CreatePURuntime", contextID, runtimeInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePURuntime indicates an expected call of CreatePURuntime
// nolint
func (mr *MockProcessingUnitsHandlerMockRecorder) CreatePURuntime(contextID, runtimeInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePURuntime", reflect.TypeOf((*MockProcessingUnitsHandler)(nil).CreatePURuntime), contextID, runtimeInfo)
}

// HandlePUEvent mocks base method
// nolint
func (m *MockProcessingUnitsHandler) HandlePUEvent(contextID string, event events.Event) error {
	ret := m.ctrl.Call(m, "HandlePUEvent", contextID, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandlePUEvent indicates an expected call of HandlePUEvent
// nolint
func (mr *MockProcessingUnitsHandlerMockRecorder) HandlePUEvent(contextID, event interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlePUEvent", reflect.TypeOf((*MockProcessingUnitsHandler)(nil).HandlePUEvent), contextID, event)
}

// MockSynchronizationHandler is a mock of SynchronizationHandler interface
// nolint
type MockSynchronizationHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSynchronizationHandlerMockRecorder
}

// MockSynchronizationHandlerMockRecorder is the mock recorder for MockSynchronizationHandler
// nolint
type MockSynchronizationHandlerMockRecorder struct {
	mock *MockSynchronizationHandler
}

// NewMockSynchronizationHandler creates a new mock instance
// nolint
func NewMockSynchronizationHandler(ctrl *gomock.Controller) *MockSynchronizationHandler {
	mock := &MockSynchronizationHandler{ctrl: ctrl}
	mock.recorder = &MockSynchronizationHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockSynchronizationHandler) EXPECT() *MockSynchronizationHandlerMockRecorder {
	return m.recorder
}

// HandleSynchronization mocks base method
// nolint
func (m *MockSynchronizationHandler) HandleSynchronization(contextID string, state events.State, RuntimeReader policy.RuntimeReader, syncType processor.SynchronizationType) error {
	ret := m.ctrl.Call(m, "HandleSynchronization", contextID, state, RuntimeReader, syncType)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleSynchronization indicates an expected call of HandleSynchronization
// nolint
func (mr *MockSynchronizationHandlerMockRecorder) HandleSynchronization(contextID, state, RuntimeReader, syncType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSynchronization", reflect.TypeOf((*MockSynchronizationHandler)(nil).HandleSynchronization), contextID, state, RuntimeReader, syncType)
}

// HandleSynchronizationComplete mocks base method
// nolint
func (m *MockSynchronizationHandler) HandleSynchronizationComplete(syncType processor.SynchronizationType) {
	m.ctrl.Call(m, "HandleSynchronizationComplete", syncType)
}

// HandleSynchronizationComplete indicates an expected call of HandleSynchronizationComplete
// nolint
func (mr *MockSynchronizationHandlerMockRecorder) HandleSynchronizationComplete(syncType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSynchronizationComplete", reflect.TypeOf((*MockSynchronizationHandler)(nil).HandleSynchronizationComplete), syncType)
}

// MockProcessor is a mock of Processor interface
// nolint
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor
// nolint
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance
// nolint
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// Start mocks base method
// nolint
func (m *MockProcessor) Start(eventInfo *events.EventInfo) error {
	ret := m.ctrl.Call(m, "Start", eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
// nolint
func (mr *MockProcessorMockRecorder) Start(eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProcessor)(nil).Start), eventInfo)
}

// Stop mocks base method
// nolint
func (m *MockProcessor) Stop(eventInfo *events.EventInfo) error {
	ret := m.ctrl.Call(m, "Stop", eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
// nolint
func (mr *MockProcessorMockRecorder) Stop(eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProcessor)(nil).Stop), eventInfo)
}

// Create mocks base method
// nolint
func (m *MockProcessor) Create(eventInfo *events.EventInfo) error {
	ret := m.ctrl.Call(m, "Create", eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
// nolint
func (mr *MockProcessorMockRecorder) Create(eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProcessor)(nil).Create), eventInfo)
}

// Destroy mocks base method
// nolint
func (m *MockProcessor) Destroy(eventInfo *events.EventInfo) error {
	ret := m.ctrl.Call(m, "Destroy", eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
// nolint
func (mr *MockProcessorMockRecorder) Destroy(eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockProcessor)(nil).Destroy), eventInfo)
}

// Pause mocks base method
// nolint
func (m *MockProcessor) Pause(eventInfo *events.EventInfo) error {
	ret := m.ctrl.Call(m, "Pause", eventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause
// nolint
func (mr *MockProcessorMockRecorder) Pause(eventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockProcessor)(nil).Pause), eventInfo)
}

// ReSync mocks base method
// nolint
func (m *MockProcessor) ReSync(EventInfo *events.EventInfo) error {
	ret := m.ctrl.Call(m, "ReSync", EventInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReSync indicates an expected call of ReSync
// nolint
func (mr *MockProcessorMockRecorder) ReSync(EventInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReSync", reflect.TypeOf((*MockProcessor)(nil).ReSync), EventInfo)
}

// MockRegisterer is a mock of Registerer interface
// nolint
type MockRegisterer struct {
	ctrl     *gomock.Controller
	recorder *MockRegistererMockRecorder
}

// MockRegistererMockRecorder is the mock recorder for MockRegisterer
// nolint
type MockRegistererMockRecorder struct {
	mock *MockRegisterer
}

// NewMockRegisterer creates a new mock instance
// nolint
func NewMockRegisterer(ctrl *gomock.Controller) *MockRegisterer {
	mock := &MockRegisterer{ctrl: ctrl}
	mock.recorder = &MockRegistererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockRegisterer) EXPECT() *MockRegistererMockRecorder {
	return m.recorder
}

// RegisterProcessor mocks base method
// nolint
func (m *MockRegisterer) RegisterProcessor(puType constants.PUType, p processor.Processor) error {
	ret := m.ctrl.Call(m, "RegisterProcessor", puType, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterProcessor indicates an expected call of RegisterProcessor
// nolint
func (mr *MockRegistererMockRecorder) RegisterProcessor(puType, p interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterProcessor", reflect.TypeOf((*MockRegisterer)(nil).RegisterProcessor), puType, p)
}

// GetHandler mocks base method
// nolint
func (m *MockRegisterer) GetHandler(puType constants.PUType, e events.Event) (events.EventHandler, error) {
	ret := m.ctrl.Call(m, "GetHandler", puType, e)
	ret0, _ := ret[0].(events.EventHandler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandler indicates an expected call of GetHandler
// nolint
func (mr *MockRegistererMockRecorder) GetHandler(puType, e interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandler", reflect.TypeOf((*MockRegisterer)(nil).GetHandler), puType, e)
}
