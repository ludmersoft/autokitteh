package sdktypes

import (
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	sessionv1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/sessions/v1"
)

type SessionLogRecord struct {
	object[*SessionLogRecordPB, SessionLogRecordTraits]
}

var InvalidSessionLogRecord SessionLogRecord

type SessionLogRecordPB = sessionv1.SessionLogRecord

type SessionLogRecordTraits struct{}

func (SessionLogRecordTraits) Validate(m *SessionLogRecordPB) error {
	return errors.Join(
		objectField[SessionCallAttemptStart]("call_attempt_start", m.CallAttemptStart),
		objectField[SessionCallAttemptComplete]("call_attempt_complete", m.CallAttemptComplete),
		objectField[SessionCallSpec]("call_spec", m.CallSpec),
		objectField[SessionState]("state", m.State),
	)
}

func (SessionLogRecordTraits) StrictValidate(m *SessionLogRecordPB) error {
	return errors.Join(
		mandatory("t", m.T),
		mandatory("seq", m.Seq),
		oneOfMessage(m /* ignore: */, "t", "process_id"),
	)
}

func SessionLogRecordFromProto(m *SessionLogRecordPB) (SessionLogRecord, error) {
	return FromProto[SessionLogRecord](m)
}

func StrictSessionLogRecordFromProto(m *SessionLogRecordPB) (SessionLogRecord, error) {
	return Strict(SessionLogRecordFromProto(m))
}

func (s SessionLogRecord) GetPrint() (string, bool) {
	if m := s.read(); m.Print != nil {
		return m.Print.Text, true
	}

	return "", false
}

func (s SessionLogRecord) GetState() SessionState {
	return forceFromProto[SessionState](s.read().State)
}

func (s SessionLogRecord) GetStopRequest() (string, bool) {
	if m := s.read(); m.StopRequest != nil {
		return m.StopRequest.Reason, true
	}

	return "", false
}

func NewPrintSessionLogRecord(seq int32, text string) SessionLogRecord {
	return forceFromProto[SessionLogRecord](&SessionLogRecordPB{
		T:     timestamppb.Now(),
		Print: &sessionv1.SessionLogRecord_Print{Text: text},
		Seq:   seq,
	})
}

func NewStopRequestSessionLogRecord(seq int32, reason string) SessionLogRecord {
	return forceFromProto[SessionLogRecord](&SessionLogRecordPB{
		T:           timestamppb.Now(),
		StopRequest: &sessionv1.SessionLogRecord_StopRequest{Reason: reason},
		Seq:         seq,
	})
}

func NewStateSessionLogRecord(state SessionState) SessionLogRecord {
	if !state.IsValid() {
		return InvalidSessionLogRecord
	}

	return forceFromProto[SessionLogRecord](&SessionLogRecordPB{
		T:     timestamppb.Now(),
		State: state.ToProto(),
	})
}

func NewCallAttemptStartSessionLogRecord(seq uint32, s SessionCallAttemptStart) SessionLogRecord {
	if !s.IsValid() {
		return InvalidSessionLogRecord
	}

	return forceFromProto[SessionLogRecord](&SessionLogRecordPB{
		T:                timestamppb.Now(),
		CallAttemptStart: s.ToProto(),
		Seq:              int32(seq),
	})
}

func NewCallAttemptCompleteSessionLogRecord(seq uint32, s SessionCallAttemptComplete) SessionLogRecord {
	if !s.IsValid() {
		return InvalidSessionLogRecord
	}

	return forceFromProto[SessionLogRecord](&SessionLogRecordPB{
		T:                   timestamppb.Now(),
		CallAttemptComplete: s.ToProto(),
		Seq:                 int32(seq),
	})
}

func NewCallSpecSessionLogRecord(seq uint32, s SessionCallSpec) SessionLogRecord {
	if !s.IsValid() {
		return InvalidSessionLogRecord
	}

	return forceFromProto[SessionLogRecord](&SessionLogRecordPB{
		T:        timestamppb.Now(),
		CallSpec: s.ToProto(),
		Seq:      int32(seq),
	})
}

func (r SessionLogRecord) WithoutTimestamp() SessionLogRecord {
	m := r.read()
	m.T = nil

	if m.CallAttemptStart != nil {
		m.CallAttemptStart.StartedAt = nil
	}

	if m.CallAttemptComplete != nil {
		m.CallAttemptComplete.CompletedAt = nil
	}

	return forceFromProto[SessionLogRecord](m)
}

func (r SessionLogRecord) Timestamp() time.Time {
	return r.read().T.AsTime()
}

func (r SessionLogRecord) WithProcessID(pid string) SessionLogRecord {
	m := r.read()
	m.ProcessId = pid
	return forceFromProto[SessionLogRecord](m)
}

func (r SessionLogRecord) Seq() uint32 { return uint32(r.read().Seq) }
