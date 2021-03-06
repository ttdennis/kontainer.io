package iptables_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/firewall/iptables"
	"github.com/kontainerooo/kontainer.ooo/pkg/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var iptablesIsPresent = 1
var isRestore = 0

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", fmt.Sprintf("GO_IPT_IS_PRESENT=%d", iptablesIsPresent), fmt.Sprintf("IS_RESTORE=%d", isRestore)}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	if os.Getenv("IS_RESTORE") == "1" {
		f, _ := os.Create("test")
		b, _ := ioutil.ReadAll(os.Stdin)
		f.WriteString(string(b))
	}

	if os.Getenv("GO_IPT_IS_PRESENT") == "1" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func simpleNewInet(s string) abstraction.Inet {
	v, _ := abstraction.NewInet(s)
	return v
}

var _ = Describe("iptables", func() {
	iptables.ExecCommand = fakeExecCommand
	Describe("Create new service", func() {
		It("Should create a new service", func() {
			_, err := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error with missing iptables", func() {
			iptablesIsPresent = 0
			_, err := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())
			Ω(err).Should(HaveOccurred())
			iptablesIsPresent = 1
		})

		It("Should error with dberror", func() {
			db := testutils.NewMockDB()
			db.SetError(1)
			_, err := iptables.NewService("iptables", "iptables-restore", db)
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Create a new rule", func() {
		It("Should create a new chain", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
				Name:  "KROO-TEST",
				Table: "nat",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new chain in default table", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
				Name: "KROO-TEST",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new JumpToChainRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.JumpToChainRuleType, iptables.JumpToChainRule{
				From: "INPUT",
				To:   "OUTPUT",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new IsolationRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.IsolationRuleType, iptables.IsolationRule{
				SrcNetwork: "br-0815",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new OutgoingInRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.OutgoingInRuleType, iptables.OutgoingInRule{
				SrcNetwork: "br-0815",
				SrcIP:      simpleNewInet("127.0.0.1/0"),
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new OutgoingOutRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.OutgoingOutRuleType, iptables.OutgoingOutRule{
				SrcNetwork: "br-0815",
				SrcIP:      simpleNewInet("127.0.0.1/0"),
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new LinkContainerPortToRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.LinkContainerPortToRuleType, iptables.LinkContainerPortToRule{
				SrcIP:      simpleNewInet("127.0.0.1/0"),
				DstIP:      simpleNewInet("127.0.0.1/0"),
				SrcNetwork: "br-0815",
				DstNetwork: "br-0815",
				Protocol:   "tcp",
				DstPort:    uint16(8080),
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new LinkContainerPortFromRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.LinkContainerPortFromRuleType, iptables.LinkContainerPortFromRule{
				SrcIP:      simpleNewInet("127.0.0.1/0"),
				DstIP:      simpleNewInet("127.0.0.1/0"),
				SrcNetwork: "br-0815",
				DstNetwork: "br-0815",
				Protocol:   "tcp",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new LinkContainerToRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.LinkContainerToRuleType, iptables.LinkContainerToRule{
				SrcIP:      simpleNewInet("127.0.0.1/0"),
				DstIP:      simpleNewInet("127.0.0.1/0"),
				SrcNetwork: "br-0815",
				DstNetwork: "br-0815",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new LinkContainerFromRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.LinkContainerFromRuleType, iptables.LinkContainerFromRule{
				SrcIP:      simpleNewInet("127.0.0.1/0"),
				DstIP:      simpleNewInet("127.0.0.1/0"),
				SrcNetwork: "br-0815",
				DstNetwork: "br-0815",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new ConnectContainerFromRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.ConnectContainerFromRuleType, iptables.ConnectContainerFromRule{
				SrcIP:      simpleNewInet("127.0.0.1/0"),
				DstIP:      simpleNewInet("127.0.0.1/0"),
				SrcNetwork: "br-0815",
				DstNetwork: "br-0815",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new ConnectContainerToRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.ConnectContainerToRuleType, iptables.ConnectContainerToRule{
				SrcIP:      simpleNewInet("127.0.0.1/0"),
				DstIP:      simpleNewInet("127.0.0.1/0"),
				SrcNetwork: "br-0815",
				DstNetwork: "br-0815",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new AllowPortInRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.AllowPortInRuleType, iptables.AllowPortInRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new AllowPortOutRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new NatOutRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.NatOutRuleType, iptables.NatOutRule{})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should create a new NatMaskRule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.NatMaskRuleType, iptables.NatMaskRule{
				SrcIP:      simpleNewInet("127.0.0.1"),
				SrcNetwork: "br-0815",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error on invalid rule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.CreateRule(iptables.AllowPortInRuleType, iptables.CreateChainRule{
				Name:  "KROO-TEST",
				Table: "nat",
			})

			Ω(err).Should(HaveOccurred())
		})

		It("Should error on existing rule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			ipts.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
				Name:  "KROO-TEST",
				Table: "nat",
			})

			err := ipts.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
				Name:  "KROO-TEST",
				Table: "nat",
			})

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Remove a rule", func() {
		It("Should remove rule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			ipts.CreateRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			err := ipts.RemoveRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Should error on invalid rule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			ipts.CreateRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			err := ipts.RemoveRule(iptables.CreateChainRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			Ω(err).Should(HaveOccurred())
		})

		It("Should error on non-removable rule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			ipts.CreateRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
				Name:  "KROO-TEST",
				Table: "nat",
			})

			err := ipts.RemoveRule(iptables.CreateChainRuleType, iptables.CreateChainRule{
				Name:  "KROO-TEST",
				Table: "nat",
			})

			Ω(err).Should(HaveOccurred())
		})

		It("Should error on non-existing rule", func() {
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			err := ipts.RemoveRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "-m lalala",
			})

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Restore rules", func() {
		It("Should print all rules", func() {
			isRestore = 1
			ipts, _ := iptables.NewService("iptables", "iptables-restore", testutils.NewMockDB())

			ipts.CreateRule(iptables.AllowPortOutRuleType, iptables.AllowPortOutRule{
				Protocol: "tcp",
				Port:     uint16(53),
				Chain:    "INPUT",
			})

			err := ipts.RestoreRules()

			file, _ := ioutil.ReadFile("test")
			Ω(string(file)).Should(Equal("-A INPUT -p tcp -m tcp --dport 53 -m state --state NEW,ESTABLISHED -j ACCEPT\n"))

			os.Remove("test")

			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
