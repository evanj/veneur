package localfile

import (
	"bytes"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stripe/veneur/samplers"
)

func TestName(t *testing.T) {
	plugin := Plugin{FilePath: "doesntexist.txt", Logger: logrus.New()}
	assert.Equal(t, plugin.Name(), "localfile")
}

func TestAppendToWriter(t *testing.T) {
	b := &bytes.Buffer{}

	metrics := []samplers.DDMetric{
		samplers.DDMetric{
			Name: "a.b.c.max",
			Value: [1][2]float64{[2]float64{1476119058,
				100}},
			Tags: []string{"foo:bar",
				"baz:quz"},
			MetricType: "gauge",
			Hostname:   "globalstats",
			DeviceName: "food",
			Interval:   0,
		},
	}

	err := appendToWriter(b, metrics, metrics[0].Hostname)
	assert.Equal(t, err, nil)
}
