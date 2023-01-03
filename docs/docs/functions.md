# Global Functions

> PTerm exposes some global functions. They are listet here.

## Print Methods

> The global exposted print functions are listed here. Use them as an alternative for the print functions in`fmt`, because they support colors and various other checks.

> [!TIP]
> Any of the above functions can be used in a string context by putting an `S` in front of the function name. Example: `Sprint(...) `.

|Function Name|Description|
|-------------|-----------|
|[Print](https://pkg.go.dev/github.com/gemini/pterm#Print)|Print formats using the default formats for its operands and writes to standard output. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.|
|[Printf](https://pkg.go.dev/github.com/gemini/pterm#Printf)|Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.|
|[Printfln](https://pkg.go.dev/github.com/gemini/pterm#Printfln)|Printfln formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered with a new line at the end.|
|[Println](https://pkg.go.dev/github.com/gemini/pterm#Println)|Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.|
|[Printo](https://pkg.go.dev/github.com/gemini/pterm#Printo)|Printo overrides the current line in a terminal. If the current line is empty, the text will be printed like with `pterm.Print`. Example: pterm.Printo("Hello, World") time.Sleep(time.Second) pterm.Printo("Hello, Earth!")|

## Disable/Enable features

|Function Name|Description|
|-------------|-----------|
|[DisableStyling](https://pkg.go.dev/github.com/gemini/pterm#DisableStyling)|DisableStyling enables all PTerm styling.|
|[EnableStyling](https://pkg.go.dev/github.com/gemini/pterm#EnableStyling)|EnableStyling disables all PTerm styling.|
|[DisableColor](https://pkg.go.dev/github.com/gemini/pterm#DisableColor)|DisableColor disables colors.|
|[EnableColor](https://pkg.go.dev/github.com/gemini/pterm#EnableColor)|EnableColor enables colors.|
|[DisableDebugMessages](https://pkg.go.dev/github.com/gemini/pterm#DisableDebugMessages)|DisableDebugMessages disables the output of debug printers.|
|[EnableDebugMessages](https://pkg.go.dev/github.com/gemini/pterm#EnableDebugMessages)|EnableDebugMessages enables the output of debug printers.|
|[DisableOutput](https://pkg.go.dev/github.com/gemini/pterm#DisableOutput)|DisableOutput disables the output of PTerm.|
|[EnableOutput](https://pkg.go.dev/github.com/gemini/pterm#EnableOutput)|EnableOutput enables the output of PTerm.|

## Helper functions

|Function Name|Description|
|-------------|-----------|
|[GetTerminalHeight](https://pkg.go.dev/github.com/gemini/pterm#GetTerminalHeight)|GetTerminalHeight returns the terminal height of the active terminal.|
|[GetTerminalSize](https://pkg.go.dev/github.com/gemini/pterm#GetTerminalSize)|GetTerminalSize returns the width and the height of the active terminal.|
|[GetTerminalWidth](https://pkg.go.dev/github.com/gemini/pterm#GetTerminalWidth)|GetTerminalWidth returns the terminal width of the active terminal.|
|[RemoveColorFromString](https://pkg.go.dev/github.com/gemini/pterm#RemoveColorFromString)|RemoveColorFromString removes color codes from a string.|
|[SetDefaultOutput](https://pkg.go.dev/github.com/gemini/pterm#SetDefaultOutput)|SetDefaultOutput sets the default output of pterm.|
