package cfgloader

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func isEnvMask(property string) bool {
	property = strings.TrimSpace(property)

	return strings.HasPrefix(property, "${") && strings.HasSuffix(property, "}")
}

func loadEnv(property string) string {
	return os.Getenv(property[2 : len(property)-1])
}

func LoadStringProp(property string) string {
	if isEnvMask(property) {
		return loadEnv(property)
	}

	return property
}

func LoadIntProp(property string) int {
	if isEnvMask(property) {
		property = loadEnv(property)
	}

	res, err := strconv.Atoi(property)
	if err != nil {
		panic(err)
	}

	return res
}

func LoadFloat64Prop(property string) float64 {
	if isEnvMask(property) {
		property = loadEnv(property)
	}

	res, err := strconv.ParseFloat(property, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func LoadBoolProp(property string) bool {
	if isEnvMask(property) {
		property = loadEnv(property)
	}

	return property == "on"
}

func LoadDurationProp(property string) time.Duration {
	return time.Duration(LoadIntProp(property))
}
