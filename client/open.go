package client

// OpenFile open file by default $EDITOR env
func OpenFile(defaultEditor string, file string) {
	if defaultEditor != "" {
		rawOpenCommand(defaultEditor, file)
	} else {
		rawOpenEditorCommand(file)
	}
}
