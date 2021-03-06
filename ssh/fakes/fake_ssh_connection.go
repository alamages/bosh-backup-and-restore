// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/ssh"
)

type FakeSSHConnection struct {
	StreamStub        func(cmd string, writer io.Writer) ([]byte, int, error)
	streamMutex       sync.RWMutex
	streamArgsForCall []struct {
		cmd    string
		writer io.Writer
	}
	streamReturns struct {
		result1 []byte
		result2 int
		result3 error
	}
	streamReturnsOnCall map[int]struct {
		result1 []byte
		result2 int
		result3 error
	}
	StreamStdinStub        func(cmd string, reader io.Reader) ([]byte, []byte, int, error)
	streamStdinMutex       sync.RWMutex
	streamStdinArgsForCall []struct {
		cmd    string
		reader io.Reader
	}
	streamStdinReturns struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}
	streamStdinReturnsOnCall map[int]struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}
	RunStub        func(cmd string) ([]byte, []byte, int, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		cmd string
	}
	runReturns struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}
	runReturnsOnCall map[int]struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}
	UsernameStub        func() string
	usernameMutex       sync.RWMutex
	usernameArgsForCall []struct{}
	usernameReturns     struct {
		result1 string
	}
	usernameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSSHConnection) Stream(cmd string, writer io.Writer) ([]byte, int, error) {
	fake.streamMutex.Lock()
	ret, specificReturn := fake.streamReturnsOnCall[len(fake.streamArgsForCall)]
	fake.streamArgsForCall = append(fake.streamArgsForCall, struct {
		cmd    string
		writer io.Writer
	}{cmd, writer})
	fake.recordInvocation("Stream", []interface{}{cmd, writer})
	fake.streamMutex.Unlock()
	if fake.StreamStub != nil {
		return fake.StreamStub(cmd, writer)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.streamReturns.result1, fake.streamReturns.result2, fake.streamReturns.result3
}

func (fake *FakeSSHConnection) StreamCallCount() int {
	fake.streamMutex.RLock()
	defer fake.streamMutex.RUnlock()
	return len(fake.streamArgsForCall)
}

func (fake *FakeSSHConnection) StreamArgsForCall(i int) (string, io.Writer) {
	fake.streamMutex.RLock()
	defer fake.streamMutex.RUnlock()
	return fake.streamArgsForCall[i].cmd, fake.streamArgsForCall[i].writer
}

func (fake *FakeSSHConnection) StreamReturns(result1 []byte, result2 int, result3 error) {
	fake.StreamStub = nil
	fake.streamReturns = struct {
		result1 []byte
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSSHConnection) StreamReturnsOnCall(i int, result1 []byte, result2 int, result3 error) {
	fake.StreamStub = nil
	if fake.streamReturnsOnCall == nil {
		fake.streamReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 int
			result3 error
		})
	}
	fake.streamReturnsOnCall[i] = struct {
		result1 []byte
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSSHConnection) StreamStdin(cmd string, reader io.Reader) ([]byte, []byte, int, error) {
	fake.streamStdinMutex.Lock()
	ret, specificReturn := fake.streamStdinReturnsOnCall[len(fake.streamStdinArgsForCall)]
	fake.streamStdinArgsForCall = append(fake.streamStdinArgsForCall, struct {
		cmd    string
		reader io.Reader
	}{cmd, reader})
	fake.recordInvocation("StreamStdin", []interface{}{cmd, reader})
	fake.streamStdinMutex.Unlock()
	if fake.StreamStdinStub != nil {
		return fake.StreamStdinStub(cmd, reader)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.streamStdinReturns.result1, fake.streamStdinReturns.result2, fake.streamStdinReturns.result3, fake.streamStdinReturns.result4
}

func (fake *FakeSSHConnection) StreamStdinCallCount() int {
	fake.streamStdinMutex.RLock()
	defer fake.streamStdinMutex.RUnlock()
	return len(fake.streamStdinArgsForCall)
}

func (fake *FakeSSHConnection) StreamStdinArgsForCall(i int) (string, io.Reader) {
	fake.streamStdinMutex.RLock()
	defer fake.streamStdinMutex.RUnlock()
	return fake.streamStdinArgsForCall[i].cmd, fake.streamStdinArgsForCall[i].reader
}

func (fake *FakeSSHConnection) StreamStdinReturns(result1 []byte, result2 []byte, result3 int, result4 error) {
	fake.StreamStdinStub = nil
	fake.streamStdinReturns = struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSSHConnection) StreamStdinReturnsOnCall(i int, result1 []byte, result2 []byte, result3 int, result4 error) {
	fake.StreamStdinStub = nil
	if fake.streamStdinReturnsOnCall == nil {
		fake.streamStdinReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 []byte
			result3 int
			result4 error
		})
	}
	fake.streamStdinReturnsOnCall[i] = struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSSHConnection) Run(cmd string) ([]byte, []byte, int, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		cmd string
	}{cmd})
	fake.recordInvocation("Run", []interface{}{cmd})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(cmd)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.runReturns.result1, fake.runReturns.result2, fake.runReturns.result3, fake.runReturns.result4
}

func (fake *FakeSSHConnection) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeSSHConnection) RunArgsForCall(i int) string {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].cmd
}

func (fake *FakeSSHConnection) RunReturns(result1 []byte, result2 []byte, result3 int, result4 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSSHConnection) RunReturnsOnCall(i int, result1 []byte, result2 []byte, result3 int, result4 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 []byte
			result3 int
			result4 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 []byte
		result2 []byte
		result3 int
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSSHConnection) Username() string {
	fake.usernameMutex.Lock()
	ret, specificReturn := fake.usernameReturnsOnCall[len(fake.usernameArgsForCall)]
	fake.usernameArgsForCall = append(fake.usernameArgsForCall, struct{}{})
	fake.recordInvocation("Username", []interface{}{})
	fake.usernameMutex.Unlock()
	if fake.UsernameStub != nil {
		return fake.UsernameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.usernameReturns.result1
}

func (fake *FakeSSHConnection) UsernameCallCount() int {
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	return len(fake.usernameArgsForCall)
}

func (fake *FakeSSHConnection) UsernameReturns(result1 string) {
	fake.UsernameStub = nil
	fake.usernameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSSHConnection) UsernameReturnsOnCall(i int, result1 string) {
	fake.UsernameStub = nil
	if fake.usernameReturnsOnCall == nil {
		fake.usernameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.usernameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSSHConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.streamMutex.RLock()
	defer fake.streamMutex.RUnlock()
	fake.streamStdinMutex.RLock()
	defer fake.streamStdinMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSSHConnection) recordInvocation(key string, args []interface{}) {
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

var _ ssh.SSHConnection = new(FakeSSHConnection)
