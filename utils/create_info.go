package utils

func CreateInfo(hosts []string, opts []string, path string) InfoData {
	var ret InfoData

	if len(hosts) == 3 {
		ret.Path = path
		ret.Env = append(ret.Env, createEnv(hosts[0], "p", "prod", "production"))
		ret.Env = append(ret.Env, createEnv(hosts[1], "s", "stag", "staging"))
		ret.Env = append(ret.Env, createEnv(hosts[2], "d", "dev", "development"))
		ret.Opts = formatOpts(opts)
	}

	return ret
}

func createEnv(host string, inst string, id string, name string) Host {
	var ret Host

	sites := GetSites(host)
	ret.Inst = inst
	ret.ID = id
	ret.Name = name
	ret.Sites = sites

	return ret
}

func formatOpts(opts []string) []Opts {
	var ret []Opts

	for _, opt := range opts {
		ret = append(ret, Opts{opt, opt})
	}

	return ret
}