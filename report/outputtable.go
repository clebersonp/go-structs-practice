package report

// Outputtable embedded Saver and Displayer interfaces
// When a struct implements an interface, it needs to implement all interface methods
// to guaranty that it will follow interface contracts
type Outputtable interface {
	Saver
	Displayer
}
