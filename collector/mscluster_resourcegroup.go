package collector

import (
	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("mscluster_resourcegroup", newMSCluster_ResourceGroupCollector)
}

// A MSCluster_ResourceGroupCollector is a Prometheus collector for WMI MSCluster_ResourceGroup metrics
type MSCluster_ResourceGroupCollector struct {
	AutoFailbackType    *prometheus.Desc
	Characteristics     *prometheus.Desc
	ColdStartSetting    *prometheus.Desc
	DefaultOwner        *prometheus.Desc
	FailbackWindowEnd   *prometheus.Desc
	FailbackWindowStart *prometheus.Desc
	FailoverPeriod      *prometheus.Desc
	FailoverThreshold   *prometheus.Desc
	FaultDomain         *prometheus.Desc
	Flags               *prometheus.Desc
	GroupType           *prometheus.Desc
	PlacementOptions    *prometheus.Desc
	Priority            *prometheus.Desc
	ResiliencyPeriod    *prometheus.Desc
	State               *prometheus.Desc
	UpdateDomain        *prometheus.Desc
}

func newMSCluster_ResourceGroupCollector() (Collector, error) {
	const subsystem = "mscluster_resourcegroup"
	return &MSCluster_ResourceGroupCollector{
		AutoFailbackType: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "auto_failback_type"),
			"(AutoFailbackType)",
			[]string{"name"},
			nil,
		),
		Characteristics: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "characteristics"),
			"(Characteristics)",
			[]string{"name"},
			nil,
		),
		ColdStartSetting: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "cold_start_setting"),
			"(ColdStartSetting)",
			[]string{"name"},
			nil,
		),
		DefaultOwner: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "default_owner"),
			"(DefaultOwner)",
			[]string{"name"},
			nil,
		),
		FailbackWindowEnd: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "failback_window_end"),
			"(FailbackWindowEnd)",
			[]string{"name"},
			nil,
		),
		FailbackWindowStart: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "failback_window_start"),
			"(FailbackWindowStart)",
			[]string{"name"},
			nil,
		),
		FailoverPeriod: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "failover_period"),
			"(FailoverPeriod)",
			[]string{"name"},
			nil,
		),
		FailoverThreshold: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "failover_threshold"),
			"(FailoverThreshold)",
			[]string{"name"},
			nil,
		),
		FaultDomain: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "fault_domain"),
			"(FaultDomain)",
			[]string{"name"},
			nil,
		),
		Flags: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "flags"),
			"(Flags)",
			[]string{"name"},
			nil,
		),
		GroupType: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "group_type"),
			"(GroupType)",
			[]string{"name"},
			nil,
		),
		PlacementOptions: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "placement_options"),
			"(PlacementOptions)",
			[]string{"name"},
			nil,
		),
		Priority: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "priority"),
			"(Priority)",
			[]string{"name"},
			nil,
		),
		ResiliencyPeriod: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "resiliency_period"),
			"(ResiliencyPeriod)",
			[]string{"name"},
			nil,
		),
		State: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "state"),
			"(State)",
			[]string{"name"},
			nil,
		),
		UpdateDomain: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "update_domain"),
			"(UpdateDomain)",
			[]string{"name"},
			nil,
		),
	}, nil
}

// MSCluster_ResourceGroup docs:
// - <add link to documentation here>
type MSCluster_ResourceGroup struct {
	Name string

	AutoFailbackType    uint
	Characteristics     uint
	ColdStartSetting    uint
	DefaultOwner        uint
	FailbackWindowEnd   int
	FailbackWindowStart int
	FailoverPeriod      uint
	FailoverThreshold   uint
	FaultDomain         uint
	Flags               uint
	GroupType           uint
	PlacementOptions    uint
	Priority            uint
	ResiliencyPeriod    uint
	State               uint
	UpdateDomain        uint
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *MSCluster_ResourceGroupCollector) Collect(ctx *ScrapeContext, ch chan<- prometheus.Metric) error {
	var dst []MSCluster_ResourceGroup
	q := queryAll(&dst)
	if err := wmi.QueryNamespace(q, &dst, "root/MSCluster"); err != nil {
		return err
	}

	for _, v := range dst {

		ch <- prometheus.MustNewConstMetric(
			c.AutoFailbackType,
			prometheus.GaugeValue,
			float64(v.AutoFailbackType),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Characteristics,
			prometheus.GaugeValue,
			float64(v.Characteristics),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.ColdStartSetting,
			prometheus.GaugeValue,
			float64(v.ColdStartSetting),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.DefaultOwner,
			prometheus.GaugeValue,
			float64(v.DefaultOwner),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.FailbackWindowEnd,
			prometheus.GaugeValue,
			float64(v.FailbackWindowEnd),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.FailbackWindowStart,
			prometheus.GaugeValue,
			float64(v.FailbackWindowStart),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.FailoverPeriod,
			prometheus.GaugeValue,
			float64(v.FailoverPeriod),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.FailoverThreshold,
			prometheus.GaugeValue,
			float64(v.FailoverThreshold),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.FaultDomain,
			prometheus.GaugeValue,
			float64(v.FaultDomain),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Flags,
			prometheus.GaugeValue,
			float64(v.Flags),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.GroupType,
			prometheus.GaugeValue,
			float64(v.GroupType),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.PlacementOptions,
			prometheus.GaugeValue,
			float64(v.PlacementOptions),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.Priority,
			prometheus.GaugeValue,
			float64(v.Priority),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.ResiliencyPeriod,
			prometheus.GaugeValue,
			float64(v.ResiliencyPeriod),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.State,
			prometheus.GaugeValue,
			float64(v.State),
			v.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.UpdateDomain,
			prometheus.GaugeValue,
			float64(v.UpdateDomain),
			v.Name,
		)

	}

	return nil
}
