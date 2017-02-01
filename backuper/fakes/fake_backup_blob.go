// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/pivotal-cf/pcf-backup-and-restore/backuper"
)

type FakeBackupBlob struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	IndexStub        func() string
	indexMutex       sync.RWMutex
	indexArgsForCall []struct{}
	indexReturns     struct {
		result1 string
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	IsNamedStub        func() bool
	isNamedMutex       sync.RWMutex
	isNamedArgsForCall []struct{}
	isNamedReturns     struct {
		result1 bool
	}
	BackupSizeStub        func() (string, error)
	backupSizeMutex       sync.RWMutex
	backupSizeArgsForCall []struct{}
	backupSizeReturns     struct {
		result1 string
		result2 error
	}
	BackupChecksumStub        func() (backuper.BackupChecksum, error)
	backupChecksumMutex       sync.RWMutex
	backupChecksumArgsForCall []struct{}
	backupChecksumReturns     struct {
		result1 backuper.BackupChecksum
		result2 error
	}
	StreamFromRemoteStub        func(io.Writer) error
	streamFromRemoteMutex       sync.RWMutex
	streamFromRemoteArgsForCall []struct {
		arg1 io.Writer
	}
	streamFromRemoteReturns struct {
		result1 error
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct{}
	deleteReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackupBlob) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	return fake.nameReturns.result1
}

func (fake *FakeBackupBlob) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBackupBlob) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBackupBlob) Index() string {
	fake.indexMutex.Lock()
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct{}{})
	fake.recordInvocation("Index", []interface{}{})
	fake.indexMutex.Unlock()
	if fake.IndexStub != nil {
		return fake.IndexStub()
	}
	return fake.indexReturns.result1
}

func (fake *FakeBackupBlob) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeBackupBlob) IndexReturns(result1 string) {
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBackupBlob) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	return fake.iDReturns.result1
}

func (fake *FakeBackupBlob) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeBackupBlob) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBackupBlob) IsNamed() bool {
	fake.isNamedMutex.Lock()
	fake.isNamedArgsForCall = append(fake.isNamedArgsForCall, struct{}{})
	fake.recordInvocation("IsNamed", []interface{}{})
	fake.isNamedMutex.Unlock()
	if fake.IsNamedStub != nil {
		return fake.IsNamedStub()
	}
	return fake.isNamedReturns.result1
}

func (fake *FakeBackupBlob) IsNamedCallCount() int {
	fake.isNamedMutex.RLock()
	defer fake.isNamedMutex.RUnlock()
	return len(fake.isNamedArgsForCall)
}

func (fake *FakeBackupBlob) IsNamedReturns(result1 bool) {
	fake.IsNamedStub = nil
	fake.isNamedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBackupBlob) BackupSize() (string, error) {
	fake.backupSizeMutex.Lock()
	fake.backupSizeArgsForCall = append(fake.backupSizeArgsForCall, struct{}{})
	fake.recordInvocation("BackupSize", []interface{}{})
	fake.backupSizeMutex.Unlock()
	if fake.BackupSizeStub != nil {
		return fake.BackupSizeStub()
	}
	return fake.backupSizeReturns.result1, fake.backupSizeReturns.result2
}

func (fake *FakeBackupBlob) BackupSizeCallCount() int {
	fake.backupSizeMutex.RLock()
	defer fake.backupSizeMutex.RUnlock()
	return len(fake.backupSizeArgsForCall)
}

func (fake *FakeBackupBlob) BackupSizeReturns(result1 string, result2 error) {
	fake.BackupSizeStub = nil
	fake.backupSizeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupBlob) BackupChecksum() (backuper.BackupChecksum, error) {
	fake.backupChecksumMutex.Lock()
	fake.backupChecksumArgsForCall = append(fake.backupChecksumArgsForCall, struct{}{})
	fake.recordInvocation("BackupChecksum", []interface{}{})
	fake.backupChecksumMutex.Unlock()
	if fake.BackupChecksumStub != nil {
		return fake.BackupChecksumStub()
	}
	return fake.backupChecksumReturns.result1, fake.backupChecksumReturns.result2
}

func (fake *FakeBackupBlob) BackupChecksumCallCount() int {
	fake.backupChecksumMutex.RLock()
	defer fake.backupChecksumMutex.RUnlock()
	return len(fake.backupChecksumArgsForCall)
}

func (fake *FakeBackupBlob) BackupChecksumReturns(result1 backuper.BackupChecksum, result2 error) {
	fake.BackupChecksumStub = nil
	fake.backupChecksumReturns = struct {
		result1 backuper.BackupChecksum
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupBlob) StreamFromRemote(arg1 io.Writer) error {
	fake.streamFromRemoteMutex.Lock()
	fake.streamFromRemoteArgsForCall = append(fake.streamFromRemoteArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.recordInvocation("StreamFromRemote", []interface{}{arg1})
	fake.streamFromRemoteMutex.Unlock()
	if fake.StreamFromRemoteStub != nil {
		return fake.StreamFromRemoteStub(arg1)
	}
	return fake.streamFromRemoteReturns.result1
}

func (fake *FakeBackupBlob) StreamFromRemoteCallCount() int {
	fake.streamFromRemoteMutex.RLock()
	defer fake.streamFromRemoteMutex.RUnlock()
	return len(fake.streamFromRemoteArgsForCall)
}

func (fake *FakeBackupBlob) StreamFromRemoteArgsForCall(i int) io.Writer {
	fake.streamFromRemoteMutex.RLock()
	defer fake.streamFromRemoteMutex.RUnlock()
	return fake.streamFromRemoteArgsForCall[i].arg1
}

func (fake *FakeBackupBlob) StreamFromRemoteReturns(result1 error) {
	fake.StreamFromRemoteStub = nil
	fake.streamFromRemoteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackupBlob) Delete() error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct{}{})
	fake.recordInvocation("Delete", []interface{}{})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub()
	}
	return fake.deleteReturns.result1
}

func (fake *FakeBackupBlob) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeBackupBlob) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackupBlob) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isNamedMutex.RLock()
	defer fake.isNamedMutex.RUnlock()
	fake.backupSizeMutex.RLock()
	defer fake.backupSizeMutex.RUnlock()
	fake.backupChecksumMutex.RLock()
	defer fake.backupChecksumMutex.RUnlock()
	fake.streamFromRemoteMutex.RLock()
	defer fake.streamFromRemoteMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBackupBlob) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ backuper.BackupBlob = new(FakeBackupBlob)
