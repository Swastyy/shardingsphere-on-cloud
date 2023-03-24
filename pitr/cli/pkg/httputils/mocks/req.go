/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/httputils/req.go

// Package mock_httputils is a generated GoMock package.
package mock_httputils

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIreq is a mock of Ireq interface.
type MockIreq struct {
	ctrl     *gomock.Controller
	recorder *MockIreqMockRecorder
}

// MockIreqMockRecorder is the mock recorder for MockIreq.
type MockIreqMockRecorder struct {
	mock *MockIreq
}

// NewMockIreq creates a new mock instance.
func NewMockIreq(ctrl *gomock.Controller) *MockIreq {
	mock := &MockIreq{ctrl: ctrl}
	mock.recorder = &MockIreqMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIreq) EXPECT() *MockIreqMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockIreq) Body(b any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Body", b)
}

// Body indicates an expected call of Body.
func (mr *MockIreqMockRecorder) Body(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockIreq)(nil).Body), b)
}

// Header mocks base method.
func (m *MockIreq) Header(h map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Header", h)
}

// Header indicates an expected call of Header.
func (mr *MockIreqMockRecorder) Header(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockIreq)(nil).Header), h)
}

// Query mocks base method.
func (m_2 *MockIreq) Query(m map[string]string) {
	m_2.ctrl.T.Helper()
	m_2.ctrl.Call(m_2, "Query", m)
}

// Query indicates an expected call of Query.
func (mr *MockIreqMockRecorder) Query(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockIreq)(nil).Query), m)
}

// Send mocks base method.
func (m *MockIreq) Send(body any) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", body)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockIreqMockRecorder) Send(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockIreq)(nil).Send), body)
}