package envinfo

func GetUtilities() []*Item {
	return getItems([]func() *Item{
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

func GetBazelVersion() *Item {
	return GetItem("bazel", "Bazel")
}

func GetCMakeVersion() *Item {
	return GetItem("cmake", "CMake")
}

func GetGCCVersion() *Item {
	return GetItem("gcc", "GCC")
}

func GetClangVersion() *Item {
	return GetItem("clang", "Clang")
}

func GetGitVersion() *Item {
	return GetItem("git", "Git")
}

func GetMakeVersion() *Item {
	return GetItem("make", "Make")
}

func GetNinjaVersion() *Item {
	return GetItem("ninja", "Ninja")
}

func GetMercurialVersion() *Item {
	return GetItem("hg", "Mercurial")
}

func GetSubversionVersion() *Item {
	return GetItem("svn", "Subversion")
}

func GetFFmpegVersion() *Item {
	return GetItem("ffmpeg", "FFmpeg")
}
