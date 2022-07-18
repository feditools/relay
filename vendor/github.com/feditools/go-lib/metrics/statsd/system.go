package statsd

import (
	"runtime"
	"time"

	"github.com/feditools/go-lib/metrics"
)

func (m *Module) systemCollector() {
	l := logger.WithField("func", "emitMemStats")
	l.Infof("Starting system metrics collector")
	var memStats runtime.MemStats
	var routines int
	var err error
	go func() {
		for {
			select {
			case <-time.After(m.systemCollectionRate):
				runtime.ReadMemStats(&memStats)
				m.emitMemStats(&memStats)
				routines = runtime.NumGoroutine()
				err = m.s.Gauge(metrics.StatSysRoutines, int64(routines), m.rate)
				if err != nil {
					l.Warnf("routines: %s", err.Error())
				}
			case <-m.done:
				l.Infof("Stopping system metrics collector")

				return
			}
		}
	}()
}

func (m *Module) emitMemStats(memStats *runtime.MemStats) {
	l := logger.WithField("func", "emitMemStats")
	err := m.s.Gauge(metrics.StatSysMemAlloc, int64(memStats.Alloc), m.rate)
	if err != nil {
		l.Warnf("alloc: %s", err.Error())
	}
	err = m.s.Gauge(metrics.StatSysMemAllocTotal, int64(memStats.TotalAlloc), m.rate)
	if err != nil {
		l.Warnf("alloc total: %s", err.Error())
	}
	err = m.s.Gauge(metrics.StatSysMemSys, int64(memStats.Sys), m.rate)
	if err != nil {
		l.Warnf("sys: %s", err.Error())
	}
	err = m.s.SetInt(metrics.StatSysMemNumGC, int64(memStats.NumGC), m.rate)
	if err != nil {
		l.Warnf("num gc: %s", err.Error())
	}
}
