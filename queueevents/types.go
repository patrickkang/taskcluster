// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package queueevents

import (
	"github.com/taskcluster/taskcluster-client-go/tcclient"
)

type (
	// Message reporting a new artifact has been created for a given task.
	ArtifactCreatedMessage struct {

		// Information about the artifact that was created
		Artifact struct {

			// Mimetype for the artifact that was created.
			//
			// Max length: 255
			ContentType string `json:"contentType"`

			// Date and time after which the artifact created will be automatically
			// deleted by the queue.
			Expires tcclient.Time `json:"expires"`

			// Name of the artifact that was created, this is useful if you want to
			// attempt to fetch the artifact. But keep in mind that just because an
			// artifact is created doesn't mean that it's immediately available.
			//
			// Max length: 1024
			Name string `json:"name"`

			// This is the `storageType` for the request that was used to create the
			// artifact.
			//
			// Possible values:
			//   * "s3"
			//   * "azure"
			//   * "reference"
			//   * "error"
			StorageType string `json:"storageType"`
		} `json:"artifact"`

		// Id of the run on which artifact was created.
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int `json:"runId"`

		Status TaskStatusStructure `json:"status"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`

		// Identifier for the worker-group within which the run with the created
		// artifacted is running.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerGroup string `json:"workerGroup"`

		// Identifier for the worker within which the run with the created artifact
		// is running.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerID string `json:"workerId"`
	}

	// Message reporting that a task has complete successfully.
	TaskCompletedMessage struct {

		// Id of the run that completed the task
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int `json:"runId"`

		Status TaskStatusStructure `json:"status"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`

		// Identifier for the worker-group within which this run ran.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerGroup string `json:"workerGroup"`

		// Identifier for the worker that executed this run.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerID string `json:"workerId"`
	}

	// Message reporting that a task has been defined. The task may or may not be
	// _scheduled_ too.
	TaskDefinedMessage struct {
		Status TaskStatusStructure `json:"status"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`
	}

	// Message reporting that TaskCluster have failed to run a task.
	TaskExceptionMessage struct {

		// Id of the last run for the task, not provided if `deadline`
		// was exceeded before a run was started.
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int `json:"runId"`

		Status TaskStatusStructure `json:"status"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`

		// Identifier for the worker-group within which the last attempt of the task
		// ran. Not provided, if `deadline` was exceeded before a run was started.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerGroup string `json:"workerGroup"`

		// Identifier for the last worker that failed to report, causing the task
		// to fail. Not provided, if `deadline` was exceeded before a run
		// was started.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerID string `json:"workerId"`
	}

	// Message reporting that a task failed to complete successfully.
	TaskFailedMessage struct {

		// Id of the run that failed.
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int `json:"runId"`

		Status TaskStatusStructure `json:"status"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`

		// Identifier for the worker-group within which this run ran.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerGroup string `json:"workerGroup"`

		// Identifier for the worker that executed this run.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerID string `json:"workerId"`
	}

	// Message reporting that a task is now pending
	TaskPendingMessage struct {

		// Id of run that became pending, `run-id`s always starts from 0
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int `json:"runId"`

		Status TaskStatusStructure `json:"status"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`
	}

	// Message reporting that a given run of a task have started
	TaskRunningMessage struct {

		// Id of the run that just started, always starts from 0
		//
		// Mininum:    0
		// Maximum:    1000
		RunID int `json:"runId"`

		Status TaskStatusStructure `json:"status"`

		// Time at which the run expires and is resolved as `failed`, if the run
		// isn't reclaimed.
		TakenUntil tcclient.Time `json:"takenUntil"`

		// Message version
		//
		// Possible values:
		//   * 1
		Version int `json:"version"`

		// Identifier for the worker-group within which this run started.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerGroup string `json:"workerGroup"`

		// Identifier for the worker executing this run.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerID string `json:"workerId"`
	}

	// A representation of **task status** as known by the queue
	TaskStatusStructure struct {

		// Deadline of the task, `pending` and `running` runs are resolved as **failed** if not resolved by other means before the deadline. Note, deadline cannot be more than5 days into the future
		Deadline tcclient.Time `json:"deadline"`

		// Task expiration, time at which task definition and status is deleted. Notice that all artifacts for the must have an expiration that is no later than this.
		Expires tcclient.Time `json:"expires"`

		// Unique identifier for the provisioner that this task must be scheduled on
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		ProvisionerID string `json:"provisionerId"`

		// Number of retries left for the task in case of infrastructure issues
		//
		// Mininum:    0
		// Maximum:    999
		RetriesLeft int `json:"retriesLeft"`

		// List of runs, ordered so that index `i` has `runId == i`
		Runs []struct {

			// Reason for the creation of this run,
			// **more reasons may be added in the future**.
			//
			// Possible values:
			//   * "scheduled"
			//   * "retry"
			//   * "rerun"
			//   * "exception"
			ReasonCreated string `json:"reasonCreated"`

			// Reason that run was resolved, this is mainly
			// useful for runs resolved as `exception`.
			// Note, **more reasons may be added in the future**, also this
			// property is only available after the run is resolved.
			//
			// Possible values:
			//   * "completed"
			//   * "failed"
			//   * "deadline-exceeded"
			//   * "canceled"
			//   * "superseded"
			//   * "claim-expired"
			//   * "worker-shutdown"
			//   * "malformed-payload"
			//   * "resource-unavailable"
			//   * "internal-error"
			ReasonResolved string `json:"reasonResolved"`

			// Date-time at which this run was resolved, ie. when the run changed
			// state from `running` to either `completed`, `failed` or `exception`.
			// This property is only present after the run as been resolved.
			Resolved tcclient.Time `json:"resolved"`

			// Id of this task run, `run-id`s always starts from `0`
			//
			// Mininum:    0
			// Maximum:    1000
			RunID int `json:"runId"`

			// Date-time at which this run was scheduled, ie. when the run was
			// created in state `pending`.
			Scheduled tcclient.Time `json:"scheduled"`

			// Date-time at which this run was claimed, ie. when the run changed
			// state from `pending` to `running`. This property is only present
			// after the run has been claimed.
			Started tcclient.Time `json:"started"`

			// State of this run
			//
			// Possible values:
			//   * "pending"
			//   * "running"
			//   * "completed"
			//   * "failed"
			//   * "exception"
			State string `json:"state"`

			// Time at which the run expires and is resolved as `failed`, if the
			// run isn't reclaimed. Note, only present after the run has been
			// claimed.
			TakenUntil tcclient.Time `json:"takenUntil"`

			// Identifier for group that worker who executes this run is a part of,
			// this identifier is mainly used for efficient routing.
			// Note, this property is only present after the run is claimed.
			//
			// Syntax:     ^([a-zA-Z0-9-_]*)$
			// Min length: 1
			// Max length: 22
			WorkerGroup string `json:"workerGroup"`

			// Identifier for worker evaluating this run within given
			// `workerGroup`. Note, this property is only available after the run
			// has been claimed.
			//
			// Syntax:     ^([a-zA-Z0-9-_]*)$
			// Min length: 1
			// Max length: 22
			WorkerID string `json:"workerId"`
		} `json:"runs"`

		// Identifier for the scheduler that _defined_ this task.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		SchedulerID string `json:"schedulerId"`

		// State of this task. This is just an auxiliary property derived from state
		// of latests run, or `unscheduled` if none.
		//
		// Possible values:
		//   * "unscheduled"
		//   * "pending"
		//   * "running"
		//   * "completed"
		//   * "failed"
		//   * "exception"
		State string `json:"state"`

		// Identifier for a group of tasks scheduled together with this task, by
		// scheduler identified by `schedulerId`. For tasks scheduled by the
		// task-graph scheduler, this is the `taskGraphId`.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskGroupID string `json:"taskGroupId"`

		// Unique task identifier, this is UUID encoded as
		// [URL-safe base64](http://tools.ietf.org/html/rfc4648#section-5) and
		// stripped of `=` padding.
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		TaskID string `json:"taskId"`

		// Identifier for worker type within the specified provisioner
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 22
		WorkerType string `json:"workerType"`
	}
)
