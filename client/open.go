package client

// OpenFile open file by default $EDITOR env
func OpenFile(defaultEditor string, file string) {
	if defaultEditor != "" {
		RawOpenCommand(defaultEditor, file)
	} else {
		RawOpenEditorCommand(file)
	}
}
