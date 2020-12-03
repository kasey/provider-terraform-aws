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

package v1alpha1func EncodeElastictranscoderPreset(r ElastictranscoderPreset) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeElastictranscoderPreset_Container(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPreset_Description(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPreset_Id(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPreset_Name(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPreset_Type(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPreset_VideoCodecOptions(r.Spec.ForProvider, ctyVal)
	EncodeElastictranscoderPreset_AudioCodecOptions(r.Spec.ForProvider.AudioCodecOptions, ctyVal)
	EncodeElastictranscoderPreset_Thumbnails(r.Spec.ForProvider.Thumbnails, ctyVal)
	EncodeElastictranscoderPreset_Video(r.Spec.ForProvider.Video, ctyVal)
	EncodeElastictranscoderPreset_VideoWatermarks(r.Spec.ForProvider.VideoWatermarks, ctyVal)
	EncodeElastictranscoderPreset_Audio(r.Spec.ForProvider.Audio, ctyVal)
	EncodeElastictranscoderPreset_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeElastictranscoderPreset_Container(p *ElastictranscoderPresetParameters, vals map[string]cty.Value) {
	vals["container"] = cty.StringVal(p.Container)
}

func EncodeElastictranscoderPreset_Description(p *ElastictranscoderPresetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElastictranscoderPreset_Id(p *ElastictranscoderPresetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElastictranscoderPreset_Name(p *ElastictranscoderPresetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElastictranscoderPreset_Type(p *ElastictranscoderPresetParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeElastictranscoderPreset_VideoCodecOptions(p *ElastictranscoderPresetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.VideoCodecOptions {
		mVals[key] = cty.StringVal(value)
	}
	vals["video_codec_options"] = cty.MapVal(mVals)
}

func EncodeElastictranscoderPreset_AudioCodecOptions(p *AudioCodecOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AudioCodecOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPreset_AudioCodecOptions_BitDepth(v, ctyVal)
		EncodeElastictranscoderPreset_AudioCodecOptions_BitOrder(v, ctyVal)
		EncodeElastictranscoderPreset_AudioCodecOptions_Profile(v, ctyVal)
		EncodeElastictranscoderPreset_AudioCodecOptions_Signed(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["audio_codec_options"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPreset_AudioCodecOptions_BitDepth(p *AudioCodecOptions, vals map[string]cty.Value) {
	vals["bit_depth"] = cty.StringVal(p.BitDepth)
}

func EncodeElastictranscoderPreset_AudioCodecOptions_BitOrder(p *AudioCodecOptions, vals map[string]cty.Value) {
	vals["bit_order"] = cty.StringVal(p.BitOrder)
}

func EncodeElastictranscoderPreset_AudioCodecOptions_Profile(p *AudioCodecOptions, vals map[string]cty.Value) {
	vals["profile"] = cty.StringVal(p.Profile)
}

func EncodeElastictranscoderPreset_AudioCodecOptions_Signed(p *AudioCodecOptions, vals map[string]cty.Value) {
	vals["signed"] = cty.StringVal(p.Signed)
}

func EncodeElastictranscoderPreset_Thumbnails(p *Thumbnails, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Thumbnails {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPreset_Thumbnails_Interval(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_MaxHeight(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_MaxWidth(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_PaddingPolicy(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_Resolution(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_SizingPolicy(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_AspectRatio(v, ctyVal)
		EncodeElastictranscoderPreset_Thumbnails_Format(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["thumbnails"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPreset_Thumbnails_Interval(p *Thumbnails, vals map[string]cty.Value) {
	vals["interval"] = cty.StringVal(p.Interval)
}

func EncodeElastictranscoderPreset_Thumbnails_MaxHeight(p *Thumbnails, vals map[string]cty.Value) {
	vals["max_height"] = cty.StringVal(p.MaxHeight)
}

func EncodeElastictranscoderPreset_Thumbnails_MaxWidth(p *Thumbnails, vals map[string]cty.Value) {
	vals["max_width"] = cty.StringVal(p.MaxWidth)
}

func EncodeElastictranscoderPreset_Thumbnails_PaddingPolicy(p *Thumbnails, vals map[string]cty.Value) {
	vals["padding_policy"] = cty.StringVal(p.PaddingPolicy)
}

func EncodeElastictranscoderPreset_Thumbnails_Resolution(p *Thumbnails, vals map[string]cty.Value) {
	vals["resolution"] = cty.StringVal(p.Resolution)
}

func EncodeElastictranscoderPreset_Thumbnails_SizingPolicy(p *Thumbnails, vals map[string]cty.Value) {
	vals["sizing_policy"] = cty.StringVal(p.SizingPolicy)
}

func EncodeElastictranscoderPreset_Thumbnails_AspectRatio(p *Thumbnails, vals map[string]cty.Value) {
	vals["aspect_ratio"] = cty.StringVal(p.AspectRatio)
}

func EncodeElastictranscoderPreset_Thumbnails_Format(p *Thumbnails, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeElastictranscoderPreset_Video(p *Video, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Video {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPreset_Video_DisplayAspectRatio(v, ctyVal)
		EncodeElastictranscoderPreset_Video_FixedGop(v, ctyVal)
		EncodeElastictranscoderPreset_Video_KeyframesMaxDist(v, ctyVal)
		EncodeElastictranscoderPreset_Video_MaxHeight(v, ctyVal)
		EncodeElastictranscoderPreset_Video_MaxWidth(v, ctyVal)
		EncodeElastictranscoderPreset_Video_PaddingPolicy(v, ctyVal)
		EncodeElastictranscoderPreset_Video_Resolution(v, ctyVal)
		EncodeElastictranscoderPreset_Video_AspectRatio(v, ctyVal)
		EncodeElastictranscoderPreset_Video_SizingPolicy(v, ctyVal)
		EncodeElastictranscoderPreset_Video_Codec(v, ctyVal)
		EncodeElastictranscoderPreset_Video_FrameRate(v, ctyVal)
		EncodeElastictranscoderPreset_Video_MaxFrameRate(v, ctyVal)
		EncodeElastictranscoderPreset_Video_BitRate(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["video"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPreset_Video_DisplayAspectRatio(p *Video, vals map[string]cty.Value) {
	vals["display_aspect_ratio"] = cty.StringVal(p.DisplayAspectRatio)
}

func EncodeElastictranscoderPreset_Video_FixedGop(p *Video, vals map[string]cty.Value) {
	vals["fixed_gop"] = cty.StringVal(p.FixedGop)
}

func EncodeElastictranscoderPreset_Video_KeyframesMaxDist(p *Video, vals map[string]cty.Value) {
	vals["keyframes_max_dist"] = cty.StringVal(p.KeyframesMaxDist)
}

func EncodeElastictranscoderPreset_Video_MaxHeight(p *Video, vals map[string]cty.Value) {
	vals["max_height"] = cty.StringVal(p.MaxHeight)
}

func EncodeElastictranscoderPreset_Video_MaxWidth(p *Video, vals map[string]cty.Value) {
	vals["max_width"] = cty.StringVal(p.MaxWidth)
}

func EncodeElastictranscoderPreset_Video_PaddingPolicy(p *Video, vals map[string]cty.Value) {
	vals["padding_policy"] = cty.StringVal(p.PaddingPolicy)
}

func EncodeElastictranscoderPreset_Video_Resolution(p *Video, vals map[string]cty.Value) {
	vals["resolution"] = cty.StringVal(p.Resolution)
}

func EncodeElastictranscoderPreset_Video_AspectRatio(p *Video, vals map[string]cty.Value) {
	vals["aspect_ratio"] = cty.StringVal(p.AspectRatio)
}

func EncodeElastictranscoderPreset_Video_SizingPolicy(p *Video, vals map[string]cty.Value) {
	vals["sizing_policy"] = cty.StringVal(p.SizingPolicy)
}

func EncodeElastictranscoderPreset_Video_Codec(p *Video, vals map[string]cty.Value) {
	vals["codec"] = cty.StringVal(p.Codec)
}

func EncodeElastictranscoderPreset_Video_FrameRate(p *Video, vals map[string]cty.Value) {
	vals["frame_rate"] = cty.StringVal(p.FrameRate)
}

func EncodeElastictranscoderPreset_Video_MaxFrameRate(p *Video, vals map[string]cty.Value) {
	vals["max_frame_rate"] = cty.StringVal(p.MaxFrameRate)
}

func EncodeElastictranscoderPreset_Video_BitRate(p *Video, vals map[string]cty.Value) {
	vals["bit_rate"] = cty.StringVal(p.BitRate)
}

func EncodeElastictranscoderPreset_VideoWatermarks(p *VideoWatermarks, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.VideoWatermarks {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPreset_VideoWatermarks_MaxWidth(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_VerticalAlign(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_VerticalOffset(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_Opacity(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_SizingPolicy(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_Target(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_HorizontalAlign(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_HorizontalOffset(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_Id(v, ctyVal)
		EncodeElastictranscoderPreset_VideoWatermarks_MaxHeight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["video_watermarks"] = cty.SetVal(valsForCollection)
}

func EncodeElastictranscoderPreset_VideoWatermarks_MaxWidth(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["max_width"] = cty.StringVal(p.MaxWidth)
}

func EncodeElastictranscoderPreset_VideoWatermarks_VerticalAlign(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["vertical_align"] = cty.StringVal(p.VerticalAlign)
}

func EncodeElastictranscoderPreset_VideoWatermarks_VerticalOffset(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["vertical_offset"] = cty.StringVal(p.VerticalOffset)
}

func EncodeElastictranscoderPreset_VideoWatermarks_Opacity(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["opacity"] = cty.StringVal(p.Opacity)
}

func EncodeElastictranscoderPreset_VideoWatermarks_SizingPolicy(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["sizing_policy"] = cty.StringVal(p.SizingPolicy)
}

func EncodeElastictranscoderPreset_VideoWatermarks_Target(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["target"] = cty.StringVal(p.Target)
}

func EncodeElastictranscoderPreset_VideoWatermarks_HorizontalAlign(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["horizontal_align"] = cty.StringVal(p.HorizontalAlign)
}

func EncodeElastictranscoderPreset_VideoWatermarks_HorizontalOffset(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["horizontal_offset"] = cty.StringVal(p.HorizontalOffset)
}

func EncodeElastictranscoderPreset_VideoWatermarks_Id(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElastictranscoderPreset_VideoWatermarks_MaxHeight(p *VideoWatermarks, vals map[string]cty.Value) {
	vals["max_height"] = cty.StringVal(p.MaxHeight)
}

func EncodeElastictranscoderPreset_Audio(p *Audio, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Audio {
		ctyVal = make(map[string]cty.Value)
		EncodeElastictranscoderPreset_Audio_Channels(v, ctyVal)
		EncodeElastictranscoderPreset_Audio_Codec(v, ctyVal)
		EncodeElastictranscoderPreset_Audio_SampleRate(v, ctyVal)
		EncodeElastictranscoderPreset_Audio_AudioPackingMode(v, ctyVal)
		EncodeElastictranscoderPreset_Audio_BitRate(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["audio"] = cty.ListVal(valsForCollection)
}

func EncodeElastictranscoderPreset_Audio_Channels(p *Audio, vals map[string]cty.Value) {
	vals["channels"] = cty.StringVal(p.Channels)
}

func EncodeElastictranscoderPreset_Audio_Codec(p *Audio, vals map[string]cty.Value) {
	vals["codec"] = cty.StringVal(p.Codec)
}

func EncodeElastictranscoderPreset_Audio_SampleRate(p *Audio, vals map[string]cty.Value) {
	vals["sample_rate"] = cty.StringVal(p.SampleRate)
}

func EncodeElastictranscoderPreset_Audio_AudioPackingMode(p *Audio, vals map[string]cty.Value) {
	vals["audio_packing_mode"] = cty.StringVal(p.AudioPackingMode)
}

func EncodeElastictranscoderPreset_Audio_BitRate(p *Audio, vals map[string]cty.Value) {
	vals["bit_rate"] = cty.StringVal(p.BitRate)
}

func EncodeElastictranscoderPreset_Arn(p *ElastictranscoderPresetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}