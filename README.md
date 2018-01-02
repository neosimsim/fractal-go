# fractal-go
This is my *play around* project to use and learn `go`.

Here I want to play around with `go`'s concurrency feature
while calculating the Mandelbrot set.

In the same turn I try to get into `Tcl/Tk` hence I decided to
write the GUI part with `Tcl/Tk`.

To run the drawing call

	go build && ./fractal-go | ./colorIter | ./draw

Yes, the `./colorIter` step could be done in `./draw` but I couldn't resist to
write some `awk` plus I have no idea (yet) how to do the same in `Tcl/Tk`.
