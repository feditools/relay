package metrics

const (
	// StatBaseDB is the base state name for database metrics
	StatBaseDB = "db"
	// StatDBQuery is the timing for a db query
	StatDBQuery = StatBaseDB + ".query"

	// StatBaseHTTP is the base state name for http metrics
	StatBaseHTTP = "http"
	// StatHTTPRequest is the timing for a http request
	StatHTTPRequest = StatBaseHTTP + ".request"

	// StatBaseSys is the base state name for system metrics
	StatBaseSys = "sys"
	// StatSysMem is the base state name for system memory metrics
	StatSysMem = StatBaseSys + ".mem"
	// StatSysMemAlloc is the gauge for memory allocation
	StatSysMemAlloc = StatSysMem + ".alloc"
	// StatSysMemAllocTotal is the gauge for memory total allocation
	StatSysMemAllocTotal = StatSysMem + ".alloc-total"
	// StatSysMemSys is the gauge for system memory usage
	StatSysMemSys = StatSysMem + ".sys"
	// StatSysMemNumGC is the counter for the number of garbage collections
	StatSysMemNumGC = StatSysMem + ".num-gc"
	// StatSysRoutines is the gauge for the number of active go routines
	StatSysRoutines = StatBaseSys + ".goroutines"
)
