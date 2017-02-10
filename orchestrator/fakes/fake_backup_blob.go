// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/pivotal-cf/pcf-backup-and-restore/orchestrator"
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
	SizeStub        func() (string, error)
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct{}
	sizeReturns     struct {
		result1 string
		result2 error
	}
	ChecksumStub        func() (orchestrator.BackupChecksum, error)
	checksumMutex       sync.RWMutex
	checksumArgsForCall []struct{}
	checksumReturns     struct {
		result1 orchestrator.BackupChecksum
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
	StreamToRemoteStub        func(io.Reader) error
	streamToRemoteMutex       sync.RWMutex
	streamToRemoteArgsForCall []struct {
		arg1 io.Reader
	}
	streamToRemoteReturns struct {
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

func (fake *FakeBackupBlob) Size() (string, error) {
	fake.sizeMutex.Lock()
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct{}{})
	fake.recordInvocation("Size", []interface{}{})
	fake.sizeMutex.Unlock()
	if fake.SizeStub != nil {
		return fake.SizeStub()
	}
	return fake.sizeReturns.result1, fake.sizeReturns.result2
}

func (fake *FakeBackupBlob) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakeBackupBlob) SizeReturns(result1 string, result2 error) {
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupBlob) Checksum() (orchestrator.BackupChecksum, error) {
	fake.checksumMutex.Lock()
	fake.checksumArgsForCall = append(fake.checksumArgsForCall, struct{}{})
	fake.recordInvocation("Checksum", []interface{}{})
	fake.checksumMutex.Unlock()
	if fake.ChecksumStub != nil {
		return fake.ChecksumStub()
	}
	return fake.checksumReturns.result1, fake.checksumReturns.result2
}

func (fake *FakeBackupBlob) ChecksumCallCount() int {
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	return len(fake.checksumArgsForCall)
}

func (fake *FakeBackupBlob) ChecksumReturns(result1 orchestrator.BackupChecksum, result2 error) {
	fake.ChecksumStub = nil
	fake.checksumReturns = struct {
		result1 orchestrator.BackupChecksum
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

func (fake *FakeBackupBlob) StreamToRemote(arg1 io.Reader) error {
	fake.streamToRemoteMutex.Lock()
	fake.streamToRemoteArgsForCall = append(fake.streamToRemoteArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("StreamToRemote", []interface{}{arg1})
	fake.streamToRemoteMutex.Unlock()
	if fake.StreamToRemoteStub != nil {
		return fake.StreamToRemoteStub(arg1)
	}
	return fake.streamToRemoteReturns.result1
}

func (fake *FakeBackupBlob) StreamToRemoteCallCount() int {
	fake.streamToRemoteMutex.RLock()
	defer fake.streamToRemoteMutex.RUnlock()
	return len(fake.streamToRemoteArgsForCall)
}

func (fake *FakeBackupBlob) StreamToRemoteArgsForCall(i int) io.Reader {
	fake.streamToRemoteMutex.RLock()
	defer fake.streamToRemoteMutex.RUnlock()
	return fake.streamToRemoteArgsForCall[i].arg1
}

func (fake *FakeBackupBlob) StreamToRemoteReturns(result1 error) {
	fake.StreamToRemoteStub = nil
	fake.streamToRemoteReturns = struct {
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
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	fake.streamFromRemoteMutex.RLock()
	defer fake.streamFromRemoteMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.streamToRemoteMutex.RLock()
	defer fake.streamToRemoteMutex.RUnlock()
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

var _ orchestrator.BackupBlob = new(FakeBackupBlob)
