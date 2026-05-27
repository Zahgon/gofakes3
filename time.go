package gofakes3

import "time"

type TimeSource interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type TimeSourceAdvancer interface {
	TimeSource
	Advance(by time.Duration)
}

// FixedTimeSource provides a source of time that always returns the
// specified time.
func FixedTimeSource(at time.Time) TimeSourceAdvancer {
	_ = "STUB: not implemented"
	return *new(TimeSourceAdvancer)
}

func DefaultTimeSource() TimeSource {
	_ = "STUB: not implemented"
	return *

	// XXX: uses time.FixedZone to 'fake' the GMT timezone that S3 uses
	// (which is basically just UTC with a different name) to avoid
	// time.LoadLocation, which requires zoneinfo.zip to be available and
	// can break spectacularly on Windows (https://github.com/golang/go/issues/21881)
	// or Docker.
	new(TimeSource)
}

type locatedTimeSource struct {
	timeLocation *time.Location
}

func (l *locatedTimeSource) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (l *locatedTimeSource) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type fixedTimeSource struct {
	time time.Time
}

func (l *fixedTimeSource) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (l *fixedTimeSource) Since(t time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (l *fixedTimeSource) Advance(by time.Duration) { _ = "STUB: not implemented"; return }
