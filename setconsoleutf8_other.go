//go:build !windows
//+build !windows

package setconsoleutf8

func SetConsoleUTF8() error {
	return nil
}
