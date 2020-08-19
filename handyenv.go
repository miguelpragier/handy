package handy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// EnvChecker holds instructions to assert an environment variable
type EnvChecker struct {
	VarName      string
	DefaultValue string
	Mandatory    bool
	DebugPrint   bool
}

func debugLog(msg string, debugPrint bool) {
	if debugPrint {
		log.Println(msg)
	}
}

// EnvCheck Test environment variables
func EnvCheck(varName, defaultValue string, mandatory, debugPrint bool) error {
	if os.Getenv(varName) != `` {
		debugLog(fmt.Sprintf(`environment variable "%s" asserted`, varName), debugPrint)
		return nil
	}

	if defaultValue != `` {
		if err := os.Setenv(varName, defaultValue); err != nil {
			return nil
		}

		debugLog(fmt.Sprintf(`environment variable "%s" asserted with default value`, varName), debugPrint)
		return nil
	}

	if mandatory {
		return fmt.Errorf(`required environment variable "%s" isn't set`, varName)
	}

	return nil
}

// EnvStr returns the env var value as string
func EnvStr(key, defaultValue string) string {
	if os.Getenv(key) != `` {
		return os.Getenv(key)
	}

	return defaultValue
}

// EnvStrS returns the env var value as []string (string slice)
func EnvStrS(key, separator string, defaultValue []string) []string {
	if os.Getenv(key) != `` {
		return strings.Split(os.Getenv(key), separator)
	}

	if len(defaultValue) > 0 {
		return defaultValue
	}

	return []string{}
}

// EnvInt returns the env var value as int
func EnvInt(key string, defaultValue int) int {
	if os.Getenv(key) != `` {
		if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
			return i
		}
	}

	return defaultValue
}

// EnvInt64 returns the env var value as int64
func EnvInt64(key string, defaultValue int64) int64 {
	if os.Getenv(key) != `` {
		if i, err := strconv.ParseInt(os.Getenv(key), 10, 64); err == nil {
			return i
		}
	}

	return defaultValue
}

// EnvIntS returns the env var value as []int
func EnvIntS(key, separator string, defaultValue []int) []int {
	if os.Getenv(key) != `` {
		a := strings.Split(os.Getenv(key), separator)

		is := make([]int, len(a))

		for i, x := range a {
			is[i], _ = strconv.Atoi(x)
		}

		return is
	}

	if len(defaultValue) > 0 {
		return defaultValue
	}

	return []int{}
}

// EnvFloat64 returns the env var value as float64
func EnvFloat64(key string, defaultValue float64) float64 {
	if os.Getenv(key) != `` {
		if f, err := strconv.ParseFloat(os.Getenv(key), 64); err == nil {
			return f
		}
	}

	return defaultValue
}

// EnvBool returns the env var value as boolean
func EnvBool(key string, defaultValue bool) bool {
	if os.Getenv(key) != `` {
		if b, err := strconv.ParseBool(os.Getenv(key)); err == nil {
			return b
		}
	}

	return defaultValue
}

// EnvCheckerNew returns a new instance of EnvChecker to be used with EnvCheckMany()
func EnvCheckerNew(varName, defaultValue string, mandatory, debugPrint bool) EnvChecker {
	return EnvChecker{
		VarName:      varName,
		DefaultValue: defaultValue,
		Mandatory:    mandatory,
		DebugPrint:   debugPrint,
	}
}

// EnvCheckMany Test multiple environment variables at once
func EnvCheckMany(envCheckers []EnvChecker) error {
	for _, c := range envCheckers {
		if err := EnvCheck(c.VarName, c.DefaultValue, c.Mandatory, c.DebugPrint); err != nil {
			return err
		}
	}

	return nil
}

var reEnvVarRow = regexp.MustCompile(`^([A-Za-z][0-9A-Za-z_]*)=(\S+)`)

// EnvLoadFromDisk loads a file from disk, containing variables written in KEY=VALUE format
// fileName is the file name with complete path
// mustHave forces an error if the file doesn't exist
// overwriteValue when false, makes the engine skip env vars that are already definied
func EnvLoadFromDisk(fileName string, mustHave, overwriteValues bool) error {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		if mustHave {
			return err
		}

		return nil
	}

	text := string(content)

	text = strings.ReplaceAll(text, "\r\n", "\n")

	rows := strings.Split(text, "\n")

	for _, row := range rows {
		row = strings.TrimSpace(row)

		matches := reEnvVarRow.FindStringSubmatch(row)

		if len(matches) != 3 {
			continue
		}

		key := matches[1]
		value := matches[2]

		if os.Getenv(key) != "" {
			if !overwriteValues {
				continue
			}
		}

		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
