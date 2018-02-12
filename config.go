package main

func Cfg(name string) string {
	cfg := make(map[string]string)

	cfg["host"] = ":3000"

	return cfg[name]
}
