// +build linux

package unbound

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"strconv"
)

var (
	execCommand = exec.Command // execCommand is used to mock commands in tests.
)

type Unbound struct {
	path    string
	UseSudo bool
}

// clean this up
var sampleConfig = `
  use_sudo = false
  path = ""
`

func (u *Unbound) Description() string {
	return "Get metrics from the Unbound DNS caching server."
}

func (u *Unbound) SampleConfig() string {
	return sampleConfig
}

func (u *Unbound) Gather(acc telegraf.Accumulator) error {
	if len(u.path) == 0 {
		return errors.New("unbound-control not found: verify that unbound is installed and that unbound-control is in your PATH")
	}

	name := f.path
	var arg []string

	if f.UseSudo {
		name = "sudo"
		arg = append(arg, f.path)
	}

	args := append(arg, "stats")

	cmd := execCommand(name, args...)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to run command %s: %s - %s", strings.Join(cmd.Args, " "), err, string(out))
	}

	fields, err := processUnboundStats(string(out))
	if err != nil {
		return err
	}
	//check this function interface
	acc.AddFields("unbound", fields, nil)
	return nil
}

func processUnboundStats(out string) (map[string]interface{}, error) {
	// Main stats processing here
}

func init() {
	u := Unbound{}
	path, _ := exec.LookPath("unbound-control")
	if len(path) > 0 {
		u.path = path
	}
	inputs.Add("unbound", func() telegraf.Input {
		return &u
	})
}
