// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MetricAlarmInitParameters struct {

	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to true.
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta2.Policy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// References to Policy in autoscaling to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsRefs []v1.Reference `json:"alarmActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Policy in autoscaling to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsSelector *v1.Selector `json:"alarmActionsSelector,omitempty" tf:"-"`

	// The description for the alarm.
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, LessThanOrEqualToThreshold. Additionally, the values  LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and GreaterThanUpperThreshold are used only for alarms based on anomaly detection models.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm *float64 `json:"datapointsToAlarm,omitempty" tf:"datapoints_to_alarm,omitempty"`

	// The dimensions for the alarm's associated metric.  For the list of available dimensions see the AWS documentation here.
	// +mapType=granular
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Used only for alarms based on percentiles.
	// If you specify ignore, the alarm state will not change during periods with too few data points to be statistically significant.
	// If you specify evaluate or omit this parameter, the alarm will always be evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: ignore, and evaluate.
	EvaluateLowSampleCountPercentiles *string `json:"evaluateLowSampleCountPercentiles,omitempty" tf:"evaluate_low_sample_count_percentiles,omitempty"`

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *float64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic *string `json:"extendedStatistic,omitempty" tf:"extended_statistic,omitempty"`

	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +listType=set
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// The name for the alarm's associated metric.
	// See docs for supported metrics.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQuery []MetricQueryInitParameters `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`

	// The namespace for the alarm's associated metric. See docs for the list of namespaces.
	// See docs for supported metrics.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// References to Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsRefs []v1.Reference `json:"okActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsSelector *v1.Selector `json:"okActionsSelector,omitempty" tf:"-"`

	// The period in seconds over which the specified statistic is applied.
	// Valid values are 10, 30, or any multiple of 60.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: SampleCount, Average, Sum, Minimum, Maximum
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricID *string `json:"thresholdMetricId,omitempty" tf:"threshold_metric_id,omitempty"`

	// Sets how this alarm is to handle missing data points. The following values are supported: missing, ignore, breaching and notBreaching. Defaults to missing.
	TreatMissingData *string `json:"treatMissingData,omitempty" tf:"treat_missing_data,omitempty"`

	// The unit for the alarm's associated metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricAlarmObservation struct {

	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to true.
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +listType=set
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// The description for the alarm.
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// The ARN of the CloudWatch Metric Alarm.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, LessThanOrEqualToThreshold. Additionally, the values  LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and GreaterThanUpperThreshold are used only for alarms based on anomaly detection models.
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// The number of datapoints that must be breaching to trigger the alarm.
	DatapointsToAlarm *float64 `json:"datapointsToAlarm,omitempty" tf:"datapoints_to_alarm,omitempty"`

	// The dimensions for the alarm's associated metric.  For the list of available dimensions see the AWS documentation here.
	// +mapType=granular
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Used only for alarms based on percentiles.
	// If you specify ignore, the alarm state will not change during periods with too few data points to be statistically significant.
	// If you specify evaluate or omit this parameter, the alarm will always be evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: ignore, and evaluate.
	EvaluateLowSampleCountPercentiles *string `json:"evaluateLowSampleCountPercentiles,omitempty" tf:"evaluate_low_sample_count_percentiles,omitempty"`

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *float64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	ExtendedStatistic *string `json:"extendedStatistic,omitempty" tf:"extended_statistic,omitempty"`

	// The ID of the health check.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +listType=set
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// The name for the alarm's associated metric.
	// See docs for supported metrics.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	MetricQuery []MetricQueryObservation `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`

	// The namespace for the alarm's associated metric. See docs for the list of namespaces.
	// See docs for supported metrics.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +listType=set
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// The period in seconds over which the specified statistic is applied.
	// Valid values are 10, 30, or any multiple of 60.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: SampleCount, Average, Sum, Minimum, Maximum
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	ThresholdMetricID *string `json:"thresholdMetricId,omitempty" tf:"threshold_metric_id,omitempty"`

	// Sets how this alarm is to handle missing data points. The following values are supported: missing, ignore, breaching and notBreaching. Defaults to missing.
	TreatMissingData *string `json:"treatMissingData,omitempty" tf:"treat_missing_data,omitempty"`

	// The unit for the alarm's associated metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricAlarmParameters struct {

	// Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to true.
	// +kubebuilder:validation:Optional
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta2.Policy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// References to Policy in autoscaling to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsRefs []v1.Reference `json:"alarmActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Policy in autoscaling to populate alarmActions.
	// +kubebuilder:validation:Optional
	AlarmActionsSelector *v1.Selector `json:"alarmActionsSelector,omitempty" tf:"-"`

	// The description for the alarm.
	// +kubebuilder:validation:Optional
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, LessThanOrEqualToThreshold. Additionally, the values  LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and GreaterThanUpperThreshold are used only for alarms based on anomaly detection models.
	// +kubebuilder:validation:Optional
	ComparisonOperator *string `json:"comparisonOperator,omitempty" tf:"comparison_operator,omitempty"`

	// The number of datapoints that must be breaching to trigger the alarm.
	// +kubebuilder:validation:Optional
	DatapointsToAlarm *float64 `json:"datapointsToAlarm,omitempty" tf:"datapoints_to_alarm,omitempty"`

	// The dimensions for the alarm's associated metric.  For the list of available dimensions see the AWS documentation here.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// Used only for alarms based on percentiles.
	// If you specify ignore, the alarm state will not change during periods with too few data points to be statistically significant.
	// If you specify evaluate or omit this parameter, the alarm will always be evaluated and possibly change state no matter how many data points are available.
	// The following values are supported: ignore, and evaluate.
	// +kubebuilder:validation:Optional
	EvaluateLowSampleCountPercentiles *string `json:"evaluateLowSampleCountPercentiles,omitempty" tf:"evaluate_low_sample_count_percentiles,omitempty"`

	// The number of periods over which data is compared to the specified threshold.
	// +kubebuilder:validation:Optional
	EvaluationPeriods *float64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	// +kubebuilder:validation:Optional
	ExtendedStatistic *string `json:"extendedStatistic,omitempty" tf:"extended_statistic,omitempty"`

	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +kubebuilder:validation:Optional
	// +listType=set
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// The name for the alarm's associated metric.
	// See docs for supported metrics.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// Enables you to create an alarm based on a metric math expression. You may specify at most 20.
	// +kubebuilder:validation:Optional
	MetricQuery []MetricQueryParameters `json:"metricQuery,omitempty" tf:"metric_query,omitempty"`

	// The namespace for the alarm's associated metric. See docs for the list of namespaces.
	// See docs for supported metrics.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// References to Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsRefs []v1.Reference `json:"okActionsRefs,omitempty" tf:"-"`

	// Selector for a list of Topic in sns to populate okActions.
	// +kubebuilder:validation:Optional
	OkActionsSelector *v1.Selector `json:"okActionsSelector,omitempty" tf:"-"`

	// The period in seconds over which the specified statistic is applied.
	// Valid values are 10, 30, or any multiple of 60.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The statistic to apply to the alarm's associated metric.
	// Either of the following is supported: SampleCount, Average, Sum, Minimum, Maximum
	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
	// +kubebuilder:validation:Optional
	ThresholdMetricID *string `json:"thresholdMetricId,omitempty" tf:"threshold_metric_id,omitempty"`

	// Sets how this alarm is to handle missing data points. The following values are supported: missing, ignore, breaching and notBreaching. Defaults to missing.
	// +kubebuilder:validation:Optional
	TreatMissingData *string `json:"treatMissingData,omitempty" tf:"treat_missing_data,omitempty"`

	// The unit for the alarm's associated metric.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricInitParameters struct {

	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation here.
	// +mapType=granular
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// The name for this metric.
	// See docs for supported metrics.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The namespace for this metric. See docs for the list of namespaces.
	// See docs for supported metrics.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Granularity in seconds of returned data points.
	// For metrics with regular resolution, valid values are any multiple of 60.
	// For high-resolution metrics, valid values are 1, 5, 10, 30, or any multiple of 60.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The statistic to apply to this metric.
	// See docs for supported statistics.
	Stat *string `json:"stat,omitempty" tf:"stat,omitempty"`

	// The unit for this metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricObservation struct {

	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation here.
	// +mapType=granular
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// The name for this metric.
	// See docs for supported metrics.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The namespace for this metric. See docs for the list of namespaces.
	// See docs for supported metrics.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Granularity in seconds of returned data points.
	// For metrics with regular resolution, valid values are any multiple of 60.
	// For high-resolution metrics, valid values are 1, 5, 10, 30, or any multiple of 60.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// The statistic to apply to this metric.
	// See docs for supported statistics.
	Stat *string `json:"stat,omitempty" tf:"stat,omitempty"`

	// The unit for this metric.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricParameters struct {

	// The dimensions for this metric.  For the list of available dimensions see the AWS documentation here.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Dimensions map[string]*string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// The name for this metric.
	// See docs for supported metrics.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// The namespace for this metric. See docs for the list of namespaces.
	// See docs for supported metrics.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Granularity in seconds of returned data points.
	// For metrics with regular resolution, valid values are any multiple of 60.
	// For high-resolution metrics, valid values are 1, 5, 10, 30, or any multiple of 60.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period" tf:"period,omitempty"`

	// The statistic to apply to this metric.
	// See docs for supported statistics.
	// +kubebuilder:validation:Optional
	Stat *string `json:"stat" tf:"stat,omitempty"`

	// The unit for this metric.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type MetricQueryInitParameters struct {

	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the Amazon CloudWatch User Guide.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	Metric *MetricInitParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// Granularity in seconds of returned data points.
	// For metrics with regular resolution, valid values are any multiple of 60.
	// For high-resolution metrics, valid values are 1, 5, 10, 30, or any multiple of 60.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specify exactly one metric_query to be true to use that metric_query result as the alarm.
	ReturnData *bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

type MetricQueryObservation struct {

	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the Amazon CloudWatch User Guide.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	Metric *MetricObservation `json:"metric,omitempty" tf:"metric,omitempty"`

	// Granularity in seconds of returned data points.
	// For metrics with regular resolution, valid values are any multiple of 60.
	// For high-resolution metrics, valid values are 1, 5, 10, 30, or any multiple of 60.
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specify exactly one metric_query to be true to use that metric_query result as the alarm.
	ReturnData *bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

type MetricQueryParameters struct {

	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the Amazon CloudWatch User Guide.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	// +kubebuilder:validation:Optional
	Metric *MetricParameters `json:"metric,omitempty" tf:"metric,omitempty"`

	// Granularity in seconds of returned data points.
	// For metrics with regular resolution, valid values are any multiple of 60.
	// For high-resolution metrics, valid values are 1, 5, 10, 30, or any multiple of 60.
	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// Specify exactly one metric_query to be true to use that metric_query result as the alarm.
	// +kubebuilder:validation:Optional
	ReturnData *bool `json:"returnData,omitempty" tf:"return_data,omitempty"`
}

// MetricAlarmSpec defines the desired state of MetricAlarm
type MetricAlarmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricAlarmParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MetricAlarmInitParameters `json:"initProvider,omitempty"`
}

// MetricAlarmStatus defines the observed state of MetricAlarm.
type MetricAlarmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricAlarmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MetricAlarm is the Schema for the MetricAlarms API. Provides a CloudWatch Metric Alarm resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MetricAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.comparisonOperator) || (has(self.initProvider) && has(self.initProvider.comparisonOperator))",message="spec.forProvider.comparisonOperator is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.evaluationPeriods) || (has(self.initProvider) && has(self.initProvider.evaluationPeriods))",message="spec.forProvider.evaluationPeriods is a required parameter"
	Spec   MetricAlarmSpec   `json:"spec"`
	Status MetricAlarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricAlarmList contains a list of MetricAlarms
type MetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricAlarm `json:"items"`
}

// Repository type metadata.
var (
	MetricAlarm_Kind             = "MetricAlarm"
	MetricAlarm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricAlarm_Kind}.String()
	MetricAlarm_KindAPIVersion   = MetricAlarm_Kind + "." + CRDGroupVersion.String()
	MetricAlarm_GroupVersionKind = CRDGroupVersion.WithKind(MetricAlarm_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricAlarm{}, &MetricAlarmList{})
}
