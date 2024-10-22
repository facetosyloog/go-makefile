package version

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

var (
	gitTag       string = ""
	gitCommit    string = ""
	gitBranch    string = ""
	gitTreeState string = ""
	buildDate    string = ""
	version      string = ""
)

type info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitBranch    string `json:"gitBranch"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
	Version      string `json:"version"`
}

func GetInfo() {
	v := info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitBranch:    gitBranch,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Version:      version,
	}
	marshalled, err := json.MarshalIndent(&v, "", " ")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(marshalled))
}
