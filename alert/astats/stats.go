package astats

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "n9e"
	subsystem = "alert"
)

type Stats struct {
	AlertNotifyTotal            *prometheus.CounterVec
	AlertNotifyErrorTotal       *prometheus.CounterVec
	CounterAlertsTotal          *prometheus.CounterVec
	GaugeAlertQueueSize         prometheus.Gauge
	CounterRuleEval             *prometheus.CounterVec
	CounterQueryDataErrorTotal  *prometheus.CounterVec
	CounterQueryDataTotal       *prometheus.CounterVec
	CounterVarFillingQuery      *prometheus.CounterVec
	CounterRecordEval           *prometheus.CounterVec
	CounterRecordEvalErrorTotal *prometheus.CounterVec
	CounterMuteTotal            *prometheus.CounterVec
	CounterRuleEvalErrorTotal   *prometheus.CounterVec
	CounterHeartbeatErrorTotal  *prometheus.CounterVec
	CounterSubEventTotal        *prometheus.CounterVec
	GaugeQuerySeriesCount       *prometheus.GaugeVec
	GaugeRuleEvalDuration       *prometheus.GaugeVec
	GaugeNotifyRecordQueueSize  prometheus.Gauge
}

func NewSyncStats() *Stats {
	CounterRuleEval := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "rule_eval_total",
		Help:      "Number of rule eval.",
	}, []string{})

	CounterRuleEvalErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "rule_eval_error_total",
		Help:      "Number of rule eval error.",
	}, []string{"datasource", "stage", "busi_group", "rule_id"})

	CounterQueryDataErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "query_data_error_total",
		Help:      "Number of rule eval query data error.",
	}, []string{"datasource"})

	CounterQueryDataTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "query_data_total",
		Help:      "Number of rule eval query data.",
	}, []string{"datasource", "rule_id"})

	CounterRecordEval := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "record_eval_total",
		Help:      "Number of record eval.",
	}, []string{"datasource"})

	CounterRecordEvalErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "record_eval_error_total",
		Help:      "Number of record eval error.",
	}, []string{"datasource"})

	AlertNotifyTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alert_notify_total",
		Help:      "Number of send msg.",
	}, []string{"channel"})

	AlertNotifyErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alert_notify_error_total",
		Help:      "Number of send msg.",
	}, []string{"channel"})

	// 产生的告警总量
	CounterAlertsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alerts_total",
		Help:      "Total number alert events.",
	}, []string{"cluster", "type", "busi_group"})

	// 内存中的告警事件队列的长度
	GaugeAlertQueueSize := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "alert_queue_size",
		Help:      "The size of alert queue.",
	})

	CounterMuteTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "mute_total",
		Help:      "Number of mute.",
	}, []string{"group", "rule_id", "mute_rule_id", "datasource_id"})

	CounterSubEventTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "sub_event_total",
		Help:      "Number of sub event.",
	}, []string{"group"})

	CounterHeartbeatErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "heartbeat_error_count",
		Help:      "Number of heartbeat error.",
	}, []string{})

	GaugeQuerySeriesCount := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "eval_query_series_count",
		Help:      "Number of series retrieved from data source after query.",
	}, []string{"rule_id", "datasource_id", "ref"})
	// 通知记录队列的长度
	GaugeNotifyRecordQueueSize := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "notify_record_queue_size",
		Help:      "The size of notify record queue.",
	})

	GaugeRuleEvalDuration := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "rule_eval_duration_ms",
		Help:      "Duration of rule eval in milliseconds.",
	}, []string{"rule_id", "datasource_id"})

	CounterVarFillingQuery := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "var_filling_query_total",
		Help:      "Number of var filling query.",
	}, []string{"rule_id", "datasource_id", "ref", "typ"})

	prometheus.MustRegister(
		CounterAlertsTotal,
		GaugeAlertQueueSize,
		AlertNotifyTotal,
		AlertNotifyErrorTotal,
		CounterRuleEval,
		CounterQueryDataTotal,
		CounterQueryDataErrorTotal,
		CounterRecordEval,
		CounterRecordEvalErrorTotal,
		CounterMuteTotal,
		CounterRuleEvalErrorTotal,
		CounterHeartbeatErrorTotal,
		CounterSubEventTotal,
		GaugeQuerySeriesCount,
		GaugeRuleEvalDuration,
		GaugeNotifyRecordQueueSize,
		CounterVarFillingQuery,
	)

	return &Stats{
		CounterAlertsTotal:          CounterAlertsTotal,
		GaugeAlertQueueSize:         GaugeAlertQueueSize,
		AlertNotifyTotal:            AlertNotifyTotal,
		AlertNotifyErrorTotal:       AlertNotifyErrorTotal,
		CounterRuleEval:             CounterRuleEval,
		CounterQueryDataTotal:       CounterQueryDataTotal,
		CounterQueryDataErrorTotal:  CounterQueryDataErrorTotal,
		CounterRecordEval:           CounterRecordEval,
		CounterRecordEvalErrorTotal: CounterRecordEvalErrorTotal,
		CounterMuteTotal:            CounterMuteTotal,
		CounterRuleEvalErrorTotal:   CounterRuleEvalErrorTotal,
		CounterHeartbeatErrorTotal:  CounterHeartbeatErrorTotal,
		CounterSubEventTotal:        CounterSubEventTotal,
		GaugeQuerySeriesCount:       GaugeQuerySeriesCount,
		GaugeRuleEvalDuration:       GaugeRuleEvalDuration,
		GaugeNotifyRecordQueueSize:  GaugeNotifyRecordQueueSize,
		CounterVarFillingQuery:      CounterVarFillingQuery,
	}
}
