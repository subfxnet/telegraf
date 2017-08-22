// +build linux

package unbound

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/influxdata/telegraf/testutil"
)

func TestGather(t *testing.T) {
	u := Unbound{
		path: "unbound-control",
	}
	// overwriting exec commands with mock commands
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	var acc testutil.Accumulator

	err := u.Gather(&acc)
	if err != nil {
		t.Fatal(err)
	}
	tags := map[string]string{"host": "subfx.net"}
	fields := map[string]interface{}{
		"total.num.queries":			int64(2),
		"total.num.queries_ip_ratelimited":	int64(0),
		"total.num.cachehits":			int64(0),
		"total.num.cachemiss":			int64(2),
		"total.num.prefetch":			int64(0),
		"total.num.zero_ttl":			int64(0),
		"total.num.recursivereplies":		int64(2),
		"total.requestlist.avg":		int64(0),
		"total.requestlist.max":		int64(0),
		"total.requestlist.overwritten":	int64(0),
		"total.requestlist.exceeded":		int64(0),
		"total.requestlist.current.all":	int64(0),
		"total.requestlist.current.user":	int64(0),
		"total.recursion.time.avg":		float64(0.100296),
		"total.recursion.time.median":		float64(0),  //not sure if this is float or int
		"total.tcpusage":			float64(0),  //not sure if this is float or int
		"time.now":				float64(1503366298.735759),
		"time.up":				float64(3289817.217414),
		"time.elapsed":				float64(116.558025),
		"mem.cache.rrset":			int64(87907),
		"mem.cache.message":			int64(69885),
		"mem.mod.iterator":			int64(16548),
		"mem.mod.validator":			int64(67815),
		"mem.mod.respip":			int64(0),
		"num.query.type.A":			int64(1),
		"num.query.type.MX":			int64(1),
		"num.query.class.IN":			int64(2),
		"num.query.opcode.QUERY":		int64(2),
		"num.query.tcp":			int64(0),
		"num.query.tcpout":			int64(0),
		"num.query.ipv6":			int64(0),
		"num.query.flags.QR":			int64(0),
		"num.query.flags.AA":			int64(0),
		"num.query.flags.TC":			int64(0),
		"num.query.flags.RD":			int64(2),
		"num.query.flags.RA":			int64(0),
		"num.query.flags.Z":			int64(0),
		"num.query.flags.AD":			int64(2),
		"num.query.flags.CD":			int64(0),
		"num.query.edns.present":		int64(2),
		"num.query.edns.DO":			int64(0),
		"num.answer.rcode.NOERROR":		int64(2),
		"num.answer.rcode.FORMERR":		int64(0),
		"num.answer.rcode.SERVFAIL":		int64(0),
		"num.answer.rcode.NXDOMAIN":		int64(0),
		"num.answer.rcode.NOTIMPL":		int64(0),
		"num.answer.rcode.REFUSED":		int64(0),
		"num.answer.secure":			int64(0),
		"num.answer.bogus":			int64(0),
		"num.rrset.bogus":			int64(0),
		"unwanted.queries":			int64(0),
		"unwanted.replies":			int64(0),
		"msg.cache.count":			int64(10),
		"rrset.cache.count":			int64(72),
		"infra.cache.count":			int64(19),
		"key.cache.count":			int64(0)
	}

	acc.AssertContainsTaggedFields(t, "unbound", fields, tags)
}

// fackeExecCommand is a helper function that mock
// the exec.Command call (and call the test binary)
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

// TestHelperProcess isn't a real test. It's used to mock exec.Command
// For example, if you run:
// GO_WANT_HELPER_PROCESS=1 go test -test.run=TestHelperProcess -- chrony tracking
// it returns below mockData.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	mockData := `thread0.num.queries=2
thread0.num.queries_ip_ratelimited=0
thread0.num.cachehits=0
thread0.num.cachemiss=2
thread0.num.prefetch=0
thread0.num.zero_ttl=0
thread0.num.recursivereplies=2
thread0.requestlist.avg=0
thread0.requestlist.max=0
thread0.requestlist.overwritten=0
thread0.requestlist.exceeded=0
thread0.requestlist.current.all=0
thread0.requestlist.current.user=0
thread0.recursion.time.avg=0.100296
thread0.recursion.time.median=0
thread0.tcpusage=0
total.num.queries=2
total.num.queries_ip_ratelimited=0
total.num.cachehits=0
total.num.cachemiss=2
total.num.prefetch=0
total.num.zero_ttl=0
total.num.recursivereplies=2
total.requestlist.avg=0
total.requestlist.max=0
total.requestlist.overwritten=0
total.requestlist.exceeded=0
total.requestlist.current.all=0
total.requestlist.current.user=0
total.recursion.time.avg=0.100296
total.recursion.time.median=0
total.tcpusage=0
time.now=1503366298.735759
time.up=3289817.217414
time.elapsed=116.558025
mem.cache.rrset=87907
mem.cache.message=69885
mem.mod.iterator=16548
mem.mod.validator=67815
mem.mod.respip=0
histogram.000000.000000.to.000000.000001=0
histogram.000000.000001.to.000000.000002=0
histogram.000000.000002.to.000000.000004=0
histogram.000000.000004.to.000000.000008=0
histogram.000000.000008.to.000000.000016=0
histogram.000000.000016.to.000000.000032=0
histogram.000000.000032.to.000000.000064=0
histogram.000000.000064.to.000000.000128=0
histogram.000000.000128.to.000000.000256=0
histogram.000000.000256.to.000000.000512=0
histogram.000000.000512.to.000000.001024=0
histogram.000000.001024.to.000000.002048=0
histogram.000000.002048.to.000000.004096=0
histogram.000000.004096.to.000000.008192=0
histogram.000000.008192.to.000000.016384=0
histogram.000000.016384.to.000000.032768=1
histogram.000000.032768.to.000000.065536=0
histogram.000000.065536.to.000000.131072=0
histogram.000000.131072.to.000000.262144=1
histogram.000000.262144.to.000000.524288=0
histogram.000000.524288.to.000001.000000=0
histogram.000001.000000.to.000002.000000=0
histogram.000002.000000.to.000004.000000=0
histogram.000004.000000.to.000008.000000=0
histogram.000008.000000.to.000016.000000=0
histogram.000016.000000.to.000032.000000=0
histogram.000032.000000.to.000064.000000=0
histogram.000064.000000.to.000128.000000=0
histogram.000128.000000.to.000256.000000=0
histogram.000256.000000.to.000512.000000=0
histogram.000512.000000.to.001024.000000=0
histogram.001024.000000.to.002048.000000=0
histogram.002048.000000.to.004096.000000=0
histogram.004096.000000.to.008192.000000=0
histogram.008192.000000.to.016384.000000=0
histogram.016384.000000.to.032768.000000=0
histogram.032768.000000.to.065536.000000=0
histogram.065536.000000.to.131072.000000=0
histogram.131072.000000.to.262144.000000=0
histogram.262144.000000.to.524288.000000=0
num.query.type.A=1
num.query.type.MX=1
num.query.class.IN=2
num.query.opcode.QUERY=2
num.query.tcp=0
num.query.tcpout=0
num.query.ipv6=0
num.query.flags.QR=0
num.query.flags.AA=0
num.query.flags.TC=0
num.query.flags.RD=2
num.query.flags.RA=0
num.query.flags.Z=0
num.query.flags.AD=2
num.query.flags.CD=0
num.query.edns.present=2
num.query.edns.DO=0
num.answer.rcode.NOERROR=2
num.answer.rcode.FORMERR=0
num.answer.rcode.SERVFAIL=0
num.answer.rcode.NXDOMAIN=0
num.answer.rcode.NOTIMPL=0
num.answer.rcode.REFUSED=0
num.answer.secure=0
num.answer.bogus=0
num.rrset.bogus=0
unwanted.queries=0
unwanted.replies=0
msg.cache.count=10
rrset.cache.count=72
infra.cache.count=19
key.cache.count=0
`

	args := os.Args
	cmd, args := args[3], args[4:]

	if cmd == "unbound-control" {
		fmt.Fprint(os.Stdout, mockData)
	} else {
		fmt.Fprint(os.Stdout, "command not found")
		os.Exit(1)
	}
	os.Exit(0)
}
