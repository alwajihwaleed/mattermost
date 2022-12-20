// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	filestore "github.com/mattermost/mattermost-server/v6/shared/filestore"
	mock "github.com/stretchr/testify/mock"

	zip "archive/zip"
)

// ComplianceExporter is an autogenerated mock type for the ComplianceExporter type
type ComplianceExporter struct {
	mock.Mock
}

// CsvExport provides a mock function with given fields: zipFile, fileBackend
func (_m *ComplianceExporter) CsvExport(zipFile *zip.Writer, fileBackend filestore.FileBackend) error {
	ret := _m.Called(zipFile, fileBackend)

	var r0 error
	if rf, ok := ret.Get(0).(func(*zip.Writer, filestore.FileBackend) error); ok {
		r0 = rf(zipFile, fileBackend)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExportedEntitiesCount provides a mock function with given fields:
func (_m *ComplianceExporter) ExportedEntitiesCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}