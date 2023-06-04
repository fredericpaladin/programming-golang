package staff

var OverpaidLimit = 75000  // Public "exported"
var underPaidLimit = 20000 // Private "not exported"

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (o *Office) All() []Employee {
	return o.AllStaff
}
func (o *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, x := range o.AllStaff {
		if x.Salary >= OverpaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee
	for _, x := range e.AllStaff {
		if x.Salary < underPaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}
