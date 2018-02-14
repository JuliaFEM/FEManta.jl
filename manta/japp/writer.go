package japp

// Pipe class
type Pipe struct {
	w chan<- []byte
}

// PipeFactory creates new WriterFactory
func PipeFactory(w chan<- []byte) *Pipe {
	return &Pipe{
		w: w,
	}
}

func (w *Pipe) Write(d []byte) (int, error) {
	w.w <- d
	return 2, nil
}
