package util

import (
	"bufio"
	"fmt"
	"io"
)

func Uniq(reader io.Reader, writer io.Writer) error {
	var prev string
	row := bufio.NewScanner(reader)
	for row.Scan() {
		txt := row.Text()

		if txt == prev {
			continue
		}

		prev = txt
		fmt.Fprintln(writer, txt)

	}
	return nil
}
