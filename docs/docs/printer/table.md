# TablePrinter

<!--
Replace all of the following strings with the current printer.
     table Table TablePrinter DefaultTable
-->

![TablePrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/table/animation.svg)

<p align="center"><a href="https://github.com/forvitinn/pterm/blob/master/_examples/table/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{{"Hello", "World"}}).Render()
```

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultTable.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultTable.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                        | Type                                                       |
| ------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| [CSVReader](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithCSVReader)                             | \*csv.Reader                                               |
| [Data](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithData)                                       | [][]string                                                 |
| [HasHeader](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithHasHeader)                             | ...bool                                                    |
| [HeaderStyle](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithHeaderStyle)                         | [\*Style](https://pkg.go.dev/github.com/forvitinn/pterm#Style) |
| [Separator](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithSeparator)                             | string                                                     |
| [SeparatorStyle](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithSeparatorStyle)                   | [\*Style](https://pkg.go.dev/github.com/forvitinn/pterm#Style) |
| [Style](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithStyle)                                     | [\*Style](https://pkg.go.dev/github.com/forvitinn/pterm#Style) |
| [Boxed](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithBoxed)                                     | ...bool                                                    |
| [LeftAlignment](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithLeftAlignment)                     | ...bool                                                    |
| [RightAlignment](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithRightAlignment)                   | ...bool                                                    |
| [HeaderRowSeparator](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithHeaderRowSeparator)           | string                                                     |
| [HeaderRowSeparatorStyle](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithHeaderRowSeparatorStyle) | [\*Style](https://pkg.go.dev/github.com/forvitinn/pterm#Style) |
| [RowSeparator](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithRowSeparator)                       | string                                                     |
| [RowSeparatorStyle](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithRowSeparatorStyle)             | [\*Style](https://pkg.go.dev/github.com/forvitinn/pterm#Style) |
| [Writer](https://pkg.go.dev/github.com/forvitinn/pterm#TablePrinter.WithWriter)                                   | io.Writer                                                  |

### Output functions

> This printer implements the interface [`RenderablePrinter`](https://github.com/forvitinn/pterm/blob/master/interface_renderable_printer.go)

| Function  | Description        |
| --------- | ------------------ |
| Render()  | Prints to Terminal |
| Srender() | Returns a string   |

## Related

- [Override default printers](docs/customizing/override-default-printer.md)
