package report

// Saver - the name is a convention. When an interface has only one method, on contract.
// We can use method name plus "er" word at the end of interface name
type Saver interface {
	Save() error
}
