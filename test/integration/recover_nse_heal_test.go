// +build recover_suite

package nsmd_integration_tests

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/networkservicemesh/networkservicemesh/test/kubetest"
)

func TestNSEHealLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	g := NewWithT(t)

	testNSEHeal(t, 1, map[string]int{
		"icmp-responder-nse-1": 0,
		"icmp-responder-nse-2": 0,
	}, kubetest.DefaultTestingPodFixture(g))
}

func TestNSEHealLocalToRemote(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	g := NewWithT(t)

	testNSEHeal(t, 2, map[string]int{
		"icmp-responder-nse-1": 0,
		"icmp-responder-nse-2": 1,
	}, kubetest.DefaultTestingPodFixture(g))
}

func TestNSEHealRemoteToLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	g := NewWithT(t)

	testNSEHeal(t, 2, map[string]int{
		"icmp-responder-nse-1": 1,
		"icmp-responder-nse-2": 0,
	}, kubetest.DefaultTestingPodFixture(g))
}

func TestNSEHealRemote(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	g := NewWithT(t)

	testNSEHeal(t, 2, map[string]int{
		"icmp-responder-nse-1": 1,
		"icmp-responder-nse-2": 1,
	}, kubetest.DefaultTestingPodFixture(g))
}

func TestNSEHealLocalVpp(t *testing.T) {
	g := NewWithT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSEHeal(t, 1, map[string]int{
		"vpp-agent-nse-1": 0,
		"vpp-agent-nse-2": 0,
	}, kubetest.VppAgentTestingPodFixture(g))
}

func TestNSEHealToLocalVpp(t *testing.T) {
	g := NewWithT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSEHeal(t, 2, map[string]int{
		"vpp-agent-nse-1": 1,
		"vpp-agent-nse-2": 0,
	}, kubetest.VppAgentTestingPodFixture(g))
}

func TestNSEHealToRemoteVpp(t *testing.T) {
	g := NewWithT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSEHeal(t, 2, map[string]int{
		"vpp-agent-nse-1": 0,
		"vpp-agent-nse-2": 1,
	}, kubetest.VppAgentTestingPodFixture(g))
}

func TestNSEHealRemoteVpp(t *testing.T) {
	g := NewWithT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSEHeal(t, 2, map[string]int{
		"vpp-agent-nse-1": 1,
		"vpp-agent-nse-2": 1,
	}, kubetest.VppAgentTestingPodFixture(g))
}

/**
If passed 1 both will be on same node, if not on different.
*/
func testNSEHeal(t *testing.T, nodesCount int, affinity map[string]int, fixture kubetest.TestingPodFixture) {
	g := NewWithT(t)

	k8s, err := kubetest.NewK8s(g, kubetest.ReuseNSMResources)
	defer k8s.Cleanup()
	g.Expect(err).To(BeNil())

	// Deploy open tracing to see what happening.
	nodesSetup, err := kubetest.SetupNodes(k8s, nodesCount, defaultTimeout)
	g.Expect(err).To(BeNil())
	defer k8s.ProcessArtifacts(t)

	// Run ICMP
	node := affinity["icmp-responder-nse-1"]
	nse1 := fixture.DeployNse(k8s, nodesSetup[node].Node, "icmp-responder-nse-1", defaultTimeout)

	nscPodNode := fixture.DeployNsc(k8s, nodesSetup[0].Node, "nsc-1", defaultTimeout)
	fixture.CheckNsc(k8s, nscPodNode)

	// Since all is fine now, we need to add new ICMP responder and delete previous one.
	node = affinity["icmp-responder-nse-2"]
	fixture.DeployNse(k8s, nodesSetup[node].Node, "icmp-responder-nse-2", defaultTimeout)

	logrus.Infof("Delete first NSE")
	k8s.DeletePods(nse1)

	logrus.Infof("Waiting for connection recovery...")

	k8s.WaitLogsContains(nodesSetup[0].Nsmd, "nsmd", "Heal: Connection recovered:", defaultTimeout)

	if len(nodesSetup) > 1 {
		l2, err := k8s.GetLogs(nodesSetup[1].Nsmd, "nsmd")
		g.Expect(err).To(BeNil())
		if strings.Contains(l2, "Forwarder request failed:") {
			logrus.Infof("Forwarder first attempt was failed: %v", l2)
		}
	}

	fixture.CheckNsc(k8s, nscPodNode)
}

func TestClosingNSEHealRemoteToLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	g := NewWithT(t)

	affinity := map[string]int{
		"icmp-responder-nse-1": 1,
		"icmp-responder-nse-2": 0,
	}
	fixture := kubetest.DefaultTestingPodFixture(g)

	k8s, err := kubetest.NewK8s(g, true)
	g.Expect(err).To(BeNil())
	defer k8s.Cleanup()

	// Deploy open tracing to see what happening.
	nodesSetup, err := kubetest.SetupNodes(k8s, 2, defaultTimeout)
	g.Expect(err).To(BeNil())
	defer kubetest.MakeLogsSnapshot(k8s, t)

	// Run ICMP
	node := affinity["icmp-responder-nse-1"]
	nse1 := fixture.DeployNse(k8s, nodesSetup[node].Node, "icmp-responder-nse-1", defaultTimeout)

	nscPodNode := fixture.DeployNsc(k8s, nodesSetup[0].Node, "nsc-1", defaultTimeout)
	fixture.CheckNsc(k8s, nscPodNode)

	// Delete NSE
	k8s.DeletePods(nse1)
	// Wait for DST Heal
	logrus.Infof("Waiting for connection starts recovering...")
	k8s.WaitLogsContains(nodesSetup[0].Nsmd, "nsmd", "Starting DST Heal...", defaultTimeout)
	// Delete NSC
	k8s.DeletePods(nscPodNode)

	// Run NSE and NSC
	node = affinity["icmp-responder-nse-2"]
	nse1 = fixture.DeployNse(k8s, nodesSetup[node].Node, "icmp-responder-nse-1", defaultTimeout)
	nscPodNode = fixture.DeployNsc(k8s, nodesSetup[0].Node, "nsc-1", defaultTimeout)

	fixture.CheckNsc(k8s, nscPodNode)
}
