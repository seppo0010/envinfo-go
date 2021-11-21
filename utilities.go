package envinfo

func GetUtilities() []*Item {
	return getItems([]func() (*Item, error){
		GetBazelVersion,
		GetCMakeVersion,
		GetGCCVersion,
		GetClangVersion,
		GetGitVersion,
		GetMakeVersion,
		GetNinjaVersion,
		GetMercurialVersion,
		GetSubversionVersion,
		GetFFmpegVersion,
	})
}

func GetBazelVersion() (*Item, error) {
	return GetItem("bazel", "Bazel", "--version")
}

func GetCMakeVersion() (*Item, error) {
	return GetItem("cmake", "CMake", "--version")
}

func GetGCCVersion() (*Item, error) {
	return GetItem("gcc", "GCC", "--version")
}

func GetClangVersion() (*Item, error) {
	return GetItem("clang", "Clang", "--version")
}

func GetGitVersion() (*Item, error) {
	return GetItem("git", "Git", "--version")
}

func GetMakeVersion() (*Item, error) {
	return GetItem("make", "Make", "--version")
}

func GetNinjaVersion() (*Item, error) {
	return GetItem("ninja", "Ninja", "--version")
}

func GetMercurialVersion() (*Item, error) {
	return GetItem("hg", "Mercurial", "--version")
}

func GetSubversionVersion() (*Item, error) {
	return GetItem("svn", "Subversion", "--version")
}

func GetFFmpegVersion() (*Item, error) {
	return GetItem("ffmpeg", "FFmpeg", "-version")
}
