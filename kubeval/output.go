package kubeval

import (
	"log"
)

// TODO (brendanryan) move these structs to `/log` once we have removed the potential
// circular dependancy between this package and `/log`

// outputManager controls how results of the `kubeval` evaluation will be recorded
// and reported to the end user.
// This interface is kept private to ensure all implementations are closed within
// this package.
type outputManager interface {
	Put(r ValidationResult) error
	Flush() error
}

const (
	outputSTD  = "stdout"
	outputJSON = "json"
	outputTAP  = "tap"
)

func validOutputs() []string { _ = "STUB: not implemented"; return nil }

func GetOutputManager(outFmt string) outputManager {
	_ = "STUB: not implemented"
	return *new(outputManager)
}

// STDOutputManager reports `kubeval` results to stdout.
type STDOutputManager struct {
}

// newSTDOutputManager instantiates a new instance of STDOutputManager.
func newSTDOutputManager() *STDOutputManager { _ = "STUB: not implemented"; return nil }

func (s *STDOutputManager) Put(result ValidationResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *STDOutputManager) Flush() error {
	_ = "STUB: not implemented"
	// no op
	return nil
}

type status string

const (
	statusInvalid = "invalid"
	statusValid   = "valid"
	statusSkipped = "skipped"
)

type dataEvalResult struct {
	Filename string   `json:"filename"`
	Kind     string   `json:"kind"`
	Status   status   `json:"status"`
	Errors   []string `json:"errors"`
}

// jsonOutputManager reports `ccheck` results to `stdout` as a json array..
type jsonOutputManager struct {
	logger *log.Logger

	data []dataEvalResult
}

func newDefaultJSONOutputManager() *jsonOutputManager { _ = "STUB: not implemented"; return nil }

func newJSONOutputManager(l *log.Logger) *jsonOutputManager { _ = "STUB: not implemented"; return nil }

func getStatus(r ValidationResult) status { _ = "STUB: not implemented"; return *new(status) }

func (j *jsonOutputManager) Put(r ValidationResult) error {
	_ = "STUB: not implemented"
	// stringify gojsonschema errors
	// use a pre-allocated slice to ensure the json will have an
	// empty array in the "zero" case
	return nil
}

func (j *jsonOutputManager) Flush() error { _ = "STUB: not implemented"; return nil }

// tapOutputManager reports `conftest` results to stdout.
type tapOutputManager struct {
	logger *log.Logger

	data []dataEvalResult
}

// newDefaultTapOutManager instantiates a new instance of tapOutputManager
// using the default logger.
func newDefaultTAPOutputManager() *tapOutputManager { _ = "STUB: not implemented"; return nil }

// newTapOutputManager constructs an instance of tapOutputManager given a
// logger instance.
func newTAPOutputManager(l *log.Logger) *tapOutputManager { _ = "STUB: not implemented"; return nil }

func (j *tapOutputManager) Put(r ValidationResult) error { _ = "STUB: not implemented"; return nil }

func (j *tapOutputManager) Flush() error { _ = "STUB: not implemented"; return nil }

// We have to skip adding 1 if it's the last error
