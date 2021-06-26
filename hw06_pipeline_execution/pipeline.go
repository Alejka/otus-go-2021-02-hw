package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, terminate In, stages ...Stage) Out {
	if in == nil {
		in := make(chan interface{})
		close(in)
		return in
	}

	for _, stage := range stages {
		ch := make(Bi)
		go func(ch Bi, o Out) {
			defer close(ch)
			for {
				select { // increase priority of terminate chan
				case <-terminate:
					return
				default:
				}

				select {
				case <-terminate:
					return
				case v, ok := <-o:
					if !ok {
						return
					}
					select {
					case <-terminate: // chan "ch" can be blocked at this moment
						return
					case ch <- v:
					}
				}
			}
		}(ch, in)
		in = stage(ch)
	}
	return in
}
