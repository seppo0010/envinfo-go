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
	return GetItem("bazel", "Bazel")
}

func GetCMakeVersion() (*Item, error) {
	return GetItem("cmake", "CMake")
}

func GetGCCVersion() (*Item, error) {
	return GetItem("gcc", "GCC")
}

func GetClangVersion() (*Item, error) {
	return GetItem("clang", "Clang")
}

func GetGitVersion() (*Item, error) {
	return GetItem("git", "Git")
}

func GetMakeVersion() (*Item, error) {
	return GetItem("make", "Make")
}

func GetNinjaVersion() (*Item, error) {
	return GetItem("ninja", "Ninja")
}

func GetMercurialVersion() (*Item, error) {
	return GetItem("hg", "Mercurial")
}

func GetSubversionVersion() (*Item, error) {
	return GetItem("svn", "Subversion")
}

func GetFFmpegVersion() (*Item, error) {
	return GetItem("ffmpeg", "FFmpeg")
}
