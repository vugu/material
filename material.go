// Package material embeds the Vugu Material Design SASS files provides a way to use them in your project.

package material

// NewFileWriter returns a FileWriter.
func NewFileWriter(dir string) *FileWriter {
	 fw := &FileWriter{dir:dir}
	 return fw
}

type FileWriter struct {
	dir string
	noVerify bool
	overwrite bool
}

func (fw *FileWriter) NoVerify() *FileWriter {
	fw.noVerify = true
	return fw
}

func (fw *FileWriter) Overwrite() *FileWriter {
}

// Write will write the embedded SASS files in this packge to the specified directory.
func (fw *FileWriter) Write() error {

}

// MustWrite is like Write but panics upon error.
func (fw *FileWriter) MustWrite()  {
}
