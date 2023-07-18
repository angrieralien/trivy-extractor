package trivy_test

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	trivy "github.com/metal-toolbox/trivy-extractor/internal"
)

func TestFakeMetrics(t *testing.T) {
	fm := trivy.FakeMetricsServicer{}
	m, _ := fm.Metrics()

	if len(m) != 7 {
		t.Fatalf("there should be %d metrics lines, there is only %d", 7, len(m))
	}
}

func TestParseMetric(t *testing.T) {
	t.Skip()
	nsTeam := trivy.NewNamespaceTeam("../data/namespaces.csv")
	vm := trivy.ParseMetrics(strings.Split(trivy.FakeMetricsData, "\n")[0], nsTeam)

	if len(vm.Labels) != len(trivy.Labels)+1 {
		t.Fatalf("should have equal label lengths. actual %d, expected %d", len(vm.Labels), len(trivy.Labels))
	}

	expected := []string{
		"activate",
		"",
		"ghcr.io",
		"app/app",
		"v1",
		"replicaset-app-1-app",
		"app-1",
		"ReplicaSet",
		"app",
		"Critical",
		"Unknown",
	}

	for i := range vm.Labels {
		if vm.Labels[i] != expected[i] {
			t.Fatalf("labels are not correct. actual %s Expected %s ", vm.Labels, expected)
		}
	}

	var expectedResult float64 = 1
	if float64(expectedResult) != vm.Value {
		t.Fatalf("Result incorrect, actual %f, expected %f", vm.Value, expectedResult)
	}
}

type FakePromServicer struct {
	Calls []trivy.VulnMetrics
}

func (p *FakePromServicer) SetTeamNamespaceVulns(vm trivy.VulnMetrics) {
	p.Calls = append(p.Calls, vm)
}

func TestReport(t *testing.T) {
	t.Skip()
	p, _ := filepath.Abs("../data/namespaces.csv")

	nsTeam := trivy.NewNamespaceTeam(p)
	ch := make(chan struct{})
	fm := &trivy.FakeMetricsServicer{}
	ps := &FakePromServicer{}

	trivy.Report(fm, ps, ch, 10*time.Millisecond, nsTeam)
	time.Sleep(11 * time.Millisecond)
	ch <- struct{}{}
	if len(ps.Calls) != 5 {
		t.Fatalf("should have processed 5 but processed only %d", len(ps.Calls))
	}
}