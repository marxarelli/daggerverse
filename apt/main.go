package main

type Apt struct{}

var aptInstall = []string{"apt-get", "install", "-y"}

func (apt *Apt) Install(packages []string) WithContainerFunc {
	return WithContainerFunc(func(ctr *Container) {
		return ctr.
			WithMountedCache(
				"/var/lib/apt",
				dag.CacheVolume("/var/lib/apt"),
				ContainerWithMountedCacheOpts{Sharing: Locked},
			).
			WithMountedCache(
				"/var/cache/apt",
				dag.CacheVolume("/var/cache/apt"),
				ContainerWithMountedCacheOpts{Sharing: Locked},
			).
			WithEnvVariable("DEBIAN_FRONTEND", "noninteractive").
			WithExec(append(aptInstall, packages...))
	})
}
