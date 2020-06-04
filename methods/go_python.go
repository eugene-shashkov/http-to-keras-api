package methods

import (
	"bytes"
	"fmt"
	"os/exec"
)

func (m *Method) GoPython(file string, j string) string {
	cmd := exec.Command("python3", file, j)
	// cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	var outStr string
	outStr = out.String()
	return outStr
}
