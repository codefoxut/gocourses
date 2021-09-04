package solid

import (
	"fmt"
)

type Document struct {}

type Machine interface {
	Print(d Document)
	Scan(d Document)
	Fax(d Document)
}

type MultiFunctionPrinter struct {}

func (m MultiFunctionPrinter) Print(d Document){}

func (m MultiFunctionPrinter) Scan(d Document) {}

func  (m MultiFunctionPrinter) Fax(d Document){}

type OldFashionedPrinter struct {}

func (m OldFashionedPrinter) Print(d Document){}

func (m OldFashionedPrinter) Scan(d Document) {fmt.Println("Cant execute this operation")}

func (m OldFashionedPrinter) Fax(d Document){ fmt.Println("Cant execute this operation")}


type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type FaxMachine  interface {
	Fax(d Document)
}

type MyPrinter struct {}

func (m MyPrinter) Print(d Document){}

func (m MyPrinter) Scan(d Document) {}

type PhotoCopier struct {}

func (p PhotoCopier) Print(d Document){}

func (p PhotoCopier) Scan(d Document) {}


type MultiFunctionDevice interface {
	Printer
	Scanner
	FaxMachine
}

// interface combination + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document){
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}








func main(){

}