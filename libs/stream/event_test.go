package stream

import (
	"testing"
)

type PartyCreated struct{}
type PartyUpdated struct{}
type PartyCreatedNow struct{}

type testProtoTable struct {
	event         any
	expStreamName string
	expSubject    string
}

func TestProtoToEvent(t *testing.T) {
	tables := []testProtoTable{
		{
			event:         PartyCreated{},
			expStreamName: "PARTY",
			expSubject:    "PARTY.created",
		},
		{
			event:         PartyUpdated{},
			expStreamName: "PARTY",
			expSubject:    "PARTY.updated",
		},
		{
			event:         PartyCreatedNow{},
			expStreamName: "PARTY",
			expSubject:    "PARTY.created.now",
		},
	}

	for _, table := range tables {
		e := eventFromProtobufMessage(table.event)
		if e.streamName != table.expStreamName {
			t.Errorf("Expected StreamName %s to equal %s", e.streamName, table.expStreamName)
		}
		if e.subject != table.expSubject {
			t.Errorf("Expected Subject %s to equal %s", e.subject, table.expSubject)
		}
	}
}
