package featureflag

const FlakesEnv = "FLAKES" // DEVBOX_FEATURE_FLAKES

func Flakes() bool {
	return Get(FlakesEnv).Enabled()
}

func init() {
	disabled(FlakesEnv)
}
