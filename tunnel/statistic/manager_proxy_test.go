package statistic

import "testing"

func TestManagerSeparatesDirectAndProxyTraffic(t *testing.T) {
	m := &Manager{}

	m.PushUploaded("DIRECT", 10)
	m.PushDownloaded("DIRECT", 20)
	m.PushUploaded("proxy-node", 30)
	m.PushDownloaded("proxy-node", 40)

	if up, down := m.Total(); up != 40 || down != 60 {
		t.Fatalf("Total() = (%d, %d), want (40, 60)", up, down)
	}
	if up, down := m.ProxyTotal(); up != 30 || down != 40 {
		t.Fatalf("ProxyTotal() = (%d, %d), want (30, 40)", up, down)
	}
}

func TestManagerResetClearsProxyTraffic(t *testing.T) {
	m := &Manager{}
	m.PushUploaded("proxy-node", 30)
	m.PushDownloaded("proxy-node", 40)

	m.ResetStatistic()

	if up, down := m.ProxyTotal(); up != 0 || down != 0 {
		t.Fatalf("ProxyTotal() after reset = (%d, %d), want (0, 0)", up, down)
	}
}
