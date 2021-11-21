package envinfo

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

type SDKManagerPackage struct {
	APILevels    []string
	BuildTools   []string
	SystemImages []string
	// TODO: detect NDK
}

func GetAndroidVersions() (*SDKManagerPackage, error) {
	sdkManagerList := getAndroidSdkManagerList()
	if sdkManagerList == "" {
		return nil, nil
	}
	return parseSDKManagerOutput(sdkManagerList), nil
}

func getAndroidSdkManagerList() string {
	sdkManagerList := ""
	for _, path := range []string{
		"sdkmanager",
		path.Join(os.Getenv("ANDROID_HOME"), "tools/bin/sdkmanager"),
		path.Join(os.Getenv("ANDROID_HOME"), "cmdline-tools/latest/bin/sdkmanager"),
		path.Join(os.Getenv("HOME"), "Library/Android/sdk/tools/bin/sdkmanager"),
	} {
		sdkManagerListBytes, _ := exec.Command(path, "--list").Output()
		sdkManagerList = string(sdkManagerListBytes)
		if sdkManagerList != "" {
			return sdkManagerList
		}
	}
	return ""
}

func parseSDKManagerOutput(output string) *SDKManagerPackage {
	availableStrings := strings.Split(output, "Available")
	systemImagesRegex := regexp.MustCompile(`system-images;([\S \t]+)`)
	apiLevelsRegex := regexp.MustCompile(`platforms;android-(\d+)[\S\s]`)
	buildToolsRegex := regexp.MustCompile(`build-tools;([\d|.]+)[\S\s]`)
	installed := availableStrings[0]
	apiLevels := apiLevelsRegex.FindAllString(installed, -1)
	if apiLevels == nil {
		apiLevels = []string{}
	}
	for i, apiLevel := range apiLevels {
		apiLevels[i] = strings.TrimSpace(apiLevel)[18:]
	}

	buildTools := buildToolsRegex.FindAllString(installed, -1)
	if buildTools == nil {
		buildTools = []string{}
	}
	for i, buildTool := range buildTools {
		buildTools[i] = strings.TrimSpace(buildTool)[12:]
	}
	rawSystemImages := systemImagesRegex.FindAllString(installed, -1)
	systemImages := make([]string, 0, len(rawSystemImages))
	for _, line := range rawSystemImages {
		split := strings.Split(line[14:], "|")
		if len(split) < 3 {
			continue
		}
		path := strings.Split(strings.TrimSpace(split[0]), ";")[0]
		description := strings.Split(strings.TrimSpace(split[2]), " SystemImage")[0]
		systemImages = append(systemImages, fmt.Sprintf("%s | %s", path, description))
	}
	return &SDKManagerPackage{
		APILevels:    apiLevels,
		BuildTools:   buildTools,
		SystemImages: systemImages,
	}
}
