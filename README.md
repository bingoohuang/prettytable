# prettytable

golang render renders the Table in a human-readable "pretty" format.

Render renders the Table in a human-readable "pretty" format. Example:

```go

import "github.com/bingoohuang/prettytable"

type Person struct {
	FirstName string `table:"FIRST"`
	LastName  string
	Salary    int
	Words     string
	Address   string `table:"-"`
}

func ExampleTablePrinter() {
	persons := []Person{
		{FirstName: "Arya", LastName: "Stark", Salary: 3000, Words: ""},
		{FirstName: "Jon", LastName: "Snow", Salary: 2000, Words: "You know nothing, Jon Snow!"},
		{FirstName: "Tyrion", LastName: "Lannister", Salary: 5000, Words: ""},
	}

	out := prettytable.TablePrinter{}.Print(&persons)
	fmt.Println(out)
}
```

Output:

```
┌───┬────────┬───────────┬────────┬─────────────────────────────┐
│ # │ FIRST  │ LAST NAME │ SALARY │ WORDS                       │
├───┼────────┼───────────┼────────┼─────────────────────────────┤
│ 1 │ Arya   │ Stark     │   3000 │                             │
│ 2 │ Jon    │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 3 │ Tyrion │ Lannister │   5000 │                             │
└───┴────────┴───────────┴────────┴─────────────────────────────┘
```

```go
func ExampleTablePrinter_SingleRowTranspose() {
	p := &Person{FirstName: "Jon", LastName: "Snow", Salary: 2000, Words: "You know nothing, Jon Snow!"}
	out := prettytable.TablePrinter{SingleRowTranspose: true}.Print(p)
	fmt.Println(out)
}
```

Output:

```
┌───┬───────────┬─────────────────────────────┐
│ # │ KEY       │ VALUE                       │
├───┼───────────┼─────────────────────────────┤
│ 1 │ FIRST     │ Jon                         │
│ 2 │ Last Name │ Snow                        │
│ 3 │ Salary    │ 2000                        │
│ 4 │ Words     │ You know nothing, Jon Snow! │
└───┴───────────┴─────────────────────────────┘
```

## Thanks

1. [jedib0t/go-pretty](https://github.com/jedib0t/go-pretty)
