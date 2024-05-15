package mediatools

// MediaToolsAvailable checks availability of ffmpeg and ffprobe
func MediaToolsAvailable() (bool, bool) {
	return HasFFMPEG(), HasFFProbe()
}
