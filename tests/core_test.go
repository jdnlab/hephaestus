package htest

import "testing"

func TestParseStdIn(t *testing.T) {

  t.Log("Testing SSH: Expected Result - 0")
  coret := NewCoreTest()

  acceptedVerbs := map[string]bool {
    "permit": true,
    "deny": true,
  }

  acceptedFlow :=  map[string]bool {
    "inbound": true,
    "outbound": true,
  }

  acceptedProto := map[string]bool {
    "ssh": true,
    "icmp": true,
    "http": true,
    "https": true,
  }

  for verb := range acceptedVerbs {
    for flow := range acceptedFlow {
      for proto := range acceptedProto {
        ta := verb + " " + flow + " " + proto
        if res := coret.validateHuInt(&ta); res != (0) {
          t.Errorf("FAIL : Non-zero exit code returned (expected 0)")
        }
      }
    }
  }
}