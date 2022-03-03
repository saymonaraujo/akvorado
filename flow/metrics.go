package flow

import (
	"flowexporter/reporter"
)

type metrics struct {
	trafficBytes         *reporter.CounterVec
	trafficPackets       *reporter.CounterVec
	trafficPacketSizeSum *reporter.SummaryVec
	trafficErrors        *reporter.CounterVec

	decoderStats  *reporter.CounterVec
	decoderErrors *reporter.CounterVec
	decoderTime   *reporter.SummaryVec

	netflowErrors             *reporter.CounterVec
	netflowStats              *reporter.CounterVec
	netflowSetRecordsStatsSum *reporter.CounterVec
	netflowSetStatsSum        *reporter.CounterVec
	netflowTimeStatsSum       *reporter.SummaryVec
	netflowTemplatesStats     *reporter.CounterVec
}

func (c *Component) initMetrics() {
	c.metrics.trafficBytes = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "traffic_bytes",
			Help: "Bytes received by the application.",
		},
		[]string{"remote_ip", "type"},
	)
	c.metrics.trafficPackets = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "traffic_packets",
			Help: "Packets received by the application.",
		},
		[]string{"remote_ip", "type"},
	)
	c.metrics.trafficPacketSizeSum = c.r.SummaryVec(
		reporter.SummaryOpts{
			Name:       "traffic_summary_size_bytes",
			Help:       "Summary of packet size.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"remote_ip", "type"},
	)
	c.metrics.trafficErrors = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "traffic_errors",
			Help: "Errors while receiving packets by the application.",
		},
		[]string{"type"},
	)

	c.metrics.decoderStats = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "decoder_count",
			Help: "Decoder processed count.",
		},
		[]string{"name"},
	)
	c.metrics.decoderErrors = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "decoder_error_count",
			Help: "Decoder processed error count.",
		},
		[]string{"name"},
	)
	c.metrics.decoderTime = c.r.SummaryVec(
		reporter.SummaryOpts{
			Name:       "summary_decoding_time_us",
			Help:       "Decoding time summary.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"name"},
	)

	c.metrics.netflowErrors = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "nf_errors_count",
			Help: "Netflows processed errors.",
		},
		[]string{"router", "error"},
	)
	c.metrics.netflowStats = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "nf_count",
			Help: "Netflows processed.",
		},
		[]string{"router", "version"},
	)
	c.metrics.netflowSetRecordsStatsSum = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "nf_flowset_records_sum",
			Help: "Netflows FlowSets sum of records.",
		},
		[]string{"router", "version", "type"},
	)
	c.metrics.netflowSetStatsSum = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "nf_flowset_sum",
			Help: "Netflows FlowSets sum.",
		},
		[]string{"router", "version", "type"},
	)
	c.metrics.netflowTimeStatsSum = c.r.SummaryVec(
		reporter.SummaryOpts{
			Name:       "nf_delay_summary_seconds",
			Help:       "Netflows time difference between time of flow and processing.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"router", "version"},
	)
	c.metrics.netflowTemplatesStats = c.r.CounterVec(
		reporter.CounterOpts{
			Name: "nf_templates_count",
			Help: "Netflows Template count.",
		},
		[]string{"router", "version", "obs_domain_id", "template_id", "type"},
	)
}
