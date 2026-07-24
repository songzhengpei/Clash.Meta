package statistic

import (
	"testing"

	C "github.com/metacubex/mihomo/constant"
)

func TestManagerSeparatesDirectAndProxyTraffic(t *testing.T) {
	m := &Manager{}

	m.PushUploaded(C.Direct, 10)
	m.PushDownloaded(C.Direct, 20)
	m.PushUploaded(C.Shadowsocks, 30)
	m.PushDownloaded(C.Shadowsocks, 40)

	if up, down := m.Total(); up != 40 || down != 60 {
		t.Fatalf("Total() = (%d, %d), want (40, 60)", up, down)
	}
	if up, down := m.ProxyTotal(); up != 30 || down != 40 {
		t.Fatalf("ProxyTotal() = (%d, %d), want (30, 40)", up, down)
	}
}

func TestManagerResetClearsProxyTraffic(t *testing.T) {
	m := &Manager{}
	m.PushUploaded(C.Shadowsocks, 30)
	m.PushDownloaded(C.Shadowsocks, 40)

	m.ResetStatistic()

	if up, down := m.ProxyTotal(); up != 0 || down != 0 {
		t.Fatalf("ProxyTotal() after reset = (%d, %d), want (0, 0)", up, down)
	}
}
