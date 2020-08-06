package prettytable_test

import (
	"fmt"
	"github.com/bingoohuang/prettytable"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	// Output:
	// ┌───┬────────┬───────────┬────────┬─────────────────────────────┐
	// │ # │ FIRST  │ LAST NAME │ SALARY │ WORDS                       │
	// ├───┼────────┼───────────┼────────┼─────────────────────────────┤
	// │ 1 │ Arya   │ Stark     │   3000 │                             │
	// │ 2 │ Jon    │ Snow      │   2000 │ You know nothing, Jon Snow! │
	// │ 3 │ Tyrion │ Lannister │   5000 │                             │
	// └───┴────────┴───────────┴────────┴─────────────────────────────┘
}

func TestPrint(t *testing.T) {
	a := struct {
		Name string
		Age  int    `table:"年齿"`
		Area string `table:"-"`
	}{
		Name: "foobar",
		Age:  100,
		Area: "南极",
	}
	out := prettytable.TablePrinter{}.Print(a)
	// fmt.Println(out)
	expected := `
┌───┬────────┬──────┐
│ # │ NAME   │ 年齿 │
├───┼────────┼──────┤
│ 1 │ foobar │  100 │
└───┴────────┴──────┘`

	assert.Equal(t, expected[1:], out)

	out = prettytable.TablePrinter{NoPrintRowSeq: true}.Print(a)
	// fmt.Println(out)
	expected = `
┌────────┬──────┐
│ NAME   │ 年齿 │
├────────┼──────┤
│ foobar │  100 │
└────────┴──────┘`

	assert.Equal(t, expected[1:], out)

	out = prettytable.TablePrinter{SingleRowTranspose: true}.Print(a)
	//fmt.Println(out)
	expected = `
┌───┬──────┬────────┐
│ # │ KEY  │ VALUE  │
├───┼──────┼────────┤
│ 1 │ Name │ foobar │
│ 2 │ 年齿 │ 100    │
└───┴──────┴────────┘`

	assert.Equal(t, expected[1:], out)

	out = prettytable.TablePrinter{NoPrintRowSeq: true, SingleRowTranspose: true}.Print(a)
	fmt.Println(out)
	expected = `
┌──────┬────────┐
│ KEY  │ VALUE  │
├──────┼────────┤
│ Name │ foobar │
│ 年齿 │ 100    │
└──────┴────────┘`
	assert.Equal(t, expected[1:], out)

}
