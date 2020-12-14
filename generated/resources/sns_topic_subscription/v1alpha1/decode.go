/*
	Copyright 2019 The Crossplane Authors.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package v1alpha1

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*SnsTopicSubscription)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSnsTopicSubscription(r, ctyValue)
}

func DecodeSnsTopicSubscription(prev *SnsTopicSubscription, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSnsTopicSubscription_DeliveryPolicy(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_EndpointAutoConfirms(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_Id(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_Protocol(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_ConfirmationTimeoutInMinutes(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_Endpoint(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_FilterPolicy(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_RawMessageDelivery(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_TopicArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopicSubscription_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeSnsTopicSubscription_DeliveryPolicy(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.DeliveryPolicy = ctwhy.ValueAsString(vals["delivery_policy"])
}

func DecodeSnsTopicSubscription_EndpointAutoConfirms(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.EndpointAutoConfirms = ctwhy.ValueAsBool(vals["endpoint_auto_confirms"])
}

func DecodeSnsTopicSubscription_Id(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeSnsTopicSubscription_Protocol(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.Protocol = ctwhy.ValueAsString(vals["protocol"])
}

func DecodeSnsTopicSubscription_ConfirmationTimeoutInMinutes(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.ConfirmationTimeoutInMinutes = ctwhy.ValueAsInt64(vals["confirmation_timeout_in_minutes"])
}

func DecodeSnsTopicSubscription_Endpoint(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.Endpoint = ctwhy.ValueAsString(vals["endpoint"])
}

func DecodeSnsTopicSubscription_FilterPolicy(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.FilterPolicy = ctwhy.ValueAsString(vals["filter_policy"])
}

func DecodeSnsTopicSubscription_RawMessageDelivery(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.RawMessageDelivery = ctwhy.ValueAsBool(vals["raw_message_delivery"])
}

func DecodeSnsTopicSubscription_TopicArn(p *SnsTopicSubscriptionParameters, vals map[string]cty.Value) {
	p.TopicArn = ctwhy.ValueAsString(vals["topic_arn"])
}

func DecodeSnsTopicSubscription_Arn(p *SnsTopicSubscriptionObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}