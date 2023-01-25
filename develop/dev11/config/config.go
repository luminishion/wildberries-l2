package config

type Cfg struct {
	HttpAddr string
}

func Load() Cfg {
	// TODO FROM FILE

	return Cfg{
		HttpAddr: ":80",
	}
}
