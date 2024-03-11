package main

var aptInstall = []string{"apt-get", "install", "-y"}

func New(ctr *Container) *Apt {
	return &Apt{ctr}
}

type Apt struct {
	*Container
}

func (apt *Apt) withApt(f WithContainerFunc) *Container {
	return apt.
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
		With(f).
		WithoutEnvVariable("DEBIAN_FRONTEND").
		WithoutMount("/var/cache/apt").
		WithoutMount("/var/lib/apt")
}

func (apt *Apt) Install(packages []string) *Container {
	return apt.withApt(func(ctr *Container) {
		return ctr.WithExec(append(aptInstall, packages...))
	})
}
