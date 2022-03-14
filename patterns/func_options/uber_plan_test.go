package func_options

type options struct {
	cache  bool
	logger logger
}

type UberOption interface {
	apply(*options)
}

type cacheOption bool

func (c cacheOption) apply(opts *options) {
	opts.cache = bool(c)
}

func WithUberCache(c bool) UberOption {
	return cacheOption(c)
}

type loggerOption struct {
	Log logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.Log
}

func WithUberLogger(log logger) UberOption {
	return loggerOption{Log: log}
}

// Open creates a connection.
func Open(driverId int, orderIds []int, opts ...UberOption) MatchEngine {
	matchEngine := MatchEngine{DriverId: driverId, OrderIds: orderIds}

	options := options{
		cache:  false,
		logger: &FileLogger{},
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return matchEngine
}
