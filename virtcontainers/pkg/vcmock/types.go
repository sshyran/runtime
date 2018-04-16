// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vcmock

import (
	"syscall"

	vc "github.com/kata-containers/runtime/virtcontainers"
	"github.com/sirupsen/logrus"
)

// Sandbox is a fake Sandbox type used for testing
type Sandbox struct {
	MockID          string
	MockURL         string
	MockAnnotations map[string]string
	MockContainers  []*Container
}

// Container is a fake Container type used for testing
type Container struct {
	MockID          string
	MockURL         string
	MockToken       string
	MockProcess     vc.Process
	MockPid         int
	MockSandbox     *Sandbox
	MockAnnotations map[string]string
}

// VCMock is a type that provides an implementation of the VC interface.
// It is used for testing.
type VCMock struct {
	SetLoggerFunc func(logger logrus.FieldLogger)

	CreateSandboxFunc func(sandboxConfig vc.SandboxConfig) (vc.VCSandbox, error)
	DeleteSandboxFunc func(sandboxID string) (vc.VCSandbox, error)
	ListSandboxFunc   func() ([]vc.SandboxStatus, error)
	PauseSandboxFunc  func(sandboxID string) (vc.VCSandbox, error)
	ResumeSandboxFunc func(sandboxID string) (vc.VCSandbox, error)
	RunSandboxFunc    func(sandboxConfig vc.SandboxConfig) (vc.VCSandbox, error)
	StartSandboxFunc  func(sandboxID string) (vc.VCSandbox, error)
	StatusSandboxFunc func(sandboxID string) (vc.SandboxStatus, error)
	StopSandboxFunc   func(sandboxID string) (vc.VCSandbox, error)

	CreateContainerFunc      func(sandboxID string, containerConfig vc.ContainerConfig) (vc.VCSandbox, vc.VCContainer, error)
	DeleteContainerFunc      func(sandboxID, containerID string) (vc.VCContainer, error)
	EnterContainerFunc       func(sandboxID, containerID string, cmd vc.Cmd) (vc.VCSandbox, vc.VCContainer, *vc.Process, error)
	KillContainerFunc        func(sandboxID, containerID string, signal syscall.Signal, all bool) error
	StartContainerFunc       func(sandboxID, containerID string) (vc.VCContainer, error)
	StatusContainerFunc      func(sandboxID, containerID string) (vc.ContainerStatus, error)
	StopContainerFunc        func(sandboxID, containerID string) (vc.VCContainer, error)
	ProcessListContainerFunc func(sandboxID, containerID string, options vc.ProcessListOptions) (vc.ProcessList, error)
}