package protobuff_v3

import (
	"github.com/valllabh/ocsf-tool/ocsf/schema"
	"testing"
)

func TestObjectAndEventNamespaces(t *testing.T) {
	m := NewMapper(&schema.OCSFSchema{})
	event := &Message{Name: "application", GroupKey: "Event: application", Package: m.PackageRef("events", "application")}
	object := &Message{Name: "application", GroupKey: "Object", Package: m.PackageRef("objects")}
	AddMessage(event)
	AddMessage(object)
	got, ok := GetMessage("application")
	if !ok || got != object || len(m.Messages) != 2 {
		t.Fatalf("object lookup = %v, %v; messages = %d", got, ok, len(m.Messages))
	}
}
