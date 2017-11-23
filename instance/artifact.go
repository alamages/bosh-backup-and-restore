package instance

import (
	"fmt"
	"io"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	"github.com/pkg/errors"
)

//go:generate counterfeiter -o fakes/fake_logger.go . Logger
type Logger interface {
	Debug(tag, msg string, args ...interface{})
	Info(tag, msg string, args ...interface{})
	Error(tag, msg string, args ...interface{})
}

func NewBackupArtifact(job orchestrator.Job, instance orchestrator.InstanceIdentifer, remoteRunner RemoteRunner, logger Logger) *Artifact {
	var name string
	if job.HasNamedBackupArtifact() {
		name = job.BackupArtifactName()
	} else {
		name = job.Name()
	}
	return &Artifact{
		isNamed:           job.HasNamedBackupArtifact(),
		artifactDirectory: job.BackupArtifactDirectory(),
		name:              name,
		instance:          instance,
		remoteRunner:      remoteRunner,
		Logger:            logger,
	}
}

func NewRestoreArtifact(job orchestrator.Job, instance orchestrator.InstanceIdentifer, remoteRunner RemoteRunner, logger Logger) *Artifact {
	var name string
	if job.HasNamedRestoreArtifact() {
		name = job.RestoreArtifactName()
	} else {
		name = job.Name()
	}
	return &Artifact{
		isNamed:           job.HasNamedRestoreArtifact(),
		artifactDirectory: job.RestoreArtifactDirectory(),
		name:              name,
		instance:          instance,
		remoteRunner:      remoteRunner,
		Logger:            logger,
	}
}

type Artifact struct {
	isNamed           bool
	index             string
	artifactDirectory string
	name              string
	instance          orchestrator.InstanceIdentifer
	Logger
	remoteRunner      RemoteRunner
}

func (b *Artifact) StreamFromRemote(writer io.Writer) error {
	b.Logger.Debug("bbr", "Streaming backup from instance %s/%s", b.instance.Name(), b.instance.ID())
	err := b.remoteRunner.compressDirectory(b.artifactDirectory, writer)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error streaming backup from remote instance. Error: %s", err.Error()))
	}

	return nil
}

func (b *Artifact) StreamToRemote(reader io.Reader) error {
	err := b.remoteRunner.createDirectory(b.artifactDirectory)
	if err != nil {
		return errors.Wrap(err, "Creating backup directory on the remote failed")
	}

	return b.remoteRunner.extractArchive(reader, b.artifactDirectory)
}

func (b *Artifact) Size() (string, error) {
	b.Logger.Debug("bbr", "Calculating size of backup on %s/%s", b.instance.Name(), b.instance.ID())

	size, err := b.remoteRunner.sizeOf(b.artifactDirectory)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Unable to check size of %s", b.artifactDirectory))
	}

	return size, nil
}

func (b *Artifact) Checksum() (orchestrator.BackupChecksum, error) {
	b.Logger.Debug("bbr", "Calculating shasum for remote files on %s/%s", b.instance.Name(), b.instance.ID())

	backupChecksum, err := b.remoteRunner.checksumDirectory(b.artifactDirectory)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to calculate backup checksum")
	}

	return backupChecksum, nil
}

func (b *Artifact) Delete() error {
	b.Logger.Debug("bbr", "Deleting artifact directory on %s/%s", b.instance.Name(), b.instance.ID())

	err := b.remoteRunner.removeDirectory(b.artifactDirectory)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to delete artifact directory on instance %s/%s", b.instance.Name(), b.instance.ID()))
	}

	return nil
}

func (b *Artifact) HasCustomName() bool {
	return b.isNamed
}

func (b *Artifact) Name() string {
	return b.name
}

func (b *Artifact) InstanceIndex() string {
	return b.instance.Index()
}

func (b *Artifact) InstanceName() string {
	return b.instance.Name()
}

func (b *Artifact) logAndRun(cmd, label string) ([]byte, []byte, int, error) {
	b.Logger.Debug("bbr", "Running %s on %s/%s", label, b.instance.Name(), b.instance.ID())

	stdout, stderr, exitCode, err := b.remoteRunner.connection.Run(cmd)
	b.Logger.Debug("bbr", "Stdout: %s", string(stdout))
	b.Logger.Debug("bbr", "Stderr: %s", string(stderr))

	if err != nil {
		b.Logger.Debug("bbr", "Error running %s on instance %s/%s. Exit code %d, error: %s", label, b.instance.Name(), b.instance.ID(), exitCode, err.Error())
	}

	return stdout, stderr, exitCode, err
}
