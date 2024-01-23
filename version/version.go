package version

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

var (
	Version   = "development"
	BuildTime = ""
)

type VersionInfo struct {
	V string `json:"version"`
	T string `json:"build_time"`
}

func init() {
	checkVersionCommand()
}

func checkVersionCommand() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Println(Version) //nolint
			os.Exit(0)
		case "version-json":
			jsonBytes, err := json.Marshal(VersionInfo{Version, BuildTime})
			if err != nil {
				panic("Error encoding version to JSON: " + err.Error())
			}
			fmt.Print(string(jsonBytes)) //nolint
			os.Exit(0)
		default:
		}
	}
}

// Get returns the version information for the given Go executable
func Get(executable string) (*VersionInfo, error) {
	cmd := exec.Command(executable, "version-json")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var versionInfo VersionInfo
	err = json.Unmarshal(out, &versionInfo)
	return &versionInfo, err
}
