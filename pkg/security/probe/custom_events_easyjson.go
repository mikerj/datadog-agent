//go:build linux
// +build linux

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package probe

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe(in *jlexer.Lexer, out *RulesetLoadedEvent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "policies":
			if in.IsNull() {
				in.Skip()
				out.PoliciesLoaded = nil
			} else {
				in.Delim('[')
				if out.PoliciesLoaded == nil {
					if !in.IsDelim(']') {
						out.PoliciesLoaded = make([]*PolicyLoaded, 0, 8)
					} else {
						out.PoliciesLoaded = []*PolicyLoaded{}
					}
				} else {
					out.PoliciesLoaded = (out.PoliciesLoaded)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PolicyLoaded
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PolicyLoaded)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.PoliciesLoaded = append(out.PoliciesLoaded, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "policies_ignored":
			if in.IsNull() {
				in.Skip()
				out.PoliciesIgnored = nil
			} else {
				if out.PoliciesIgnored == nil {
					out.PoliciesIgnored = new(PoliciesIgnored)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.PoliciesIgnored).UnmarshalJSON(data))
				}
			}
		case "macros_loaded":
			if in.IsNull() {
				in.Skip()
				out.MacrosLoaded = nil
			} else {
				in.Delim('[')
				if out.MacrosLoaded == nil {
					if !in.IsDelim(']') {
						out.MacrosLoaded = make([]string, 0, 4)
					} else {
						out.MacrosLoaded = []string{}
					}
				} else {
					out.MacrosLoaded = (out.MacrosLoaded)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.MacrosLoaded = append(out.MacrosLoaded, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe(out *jwriter.Writer, in RulesetLoadedEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix[1:])
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"policies\":"
		out.RawString(prefix)
		if in.PoliciesLoaded == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.PoliciesLoaded {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.PoliciesIgnored != nil {
		const prefix string = ",\"policies_ignored\":"
		out.RawString(prefix)
		out.Raw((*in.PoliciesIgnored).MarshalJSON())
	}
	{
		const prefix string = ",\"macros_loaded\":"
		out.RawString(prefix)
		if in.MacrosLoaded == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.MacrosLoaded {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RulesetLoadedEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RulesetLoadedEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe1(in *jlexer.Lexer, out *RuleLoaded) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "expression":
			out.Expression = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe1(out *jwriter.Writer, in RuleLoaded) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"expression\":"
		out.RawString(prefix)
		out.String(string(in.Expression))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleLoaded) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleLoaded) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe1(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe2(in *jlexer.Lexer, out *RuleIgnored) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "version":
			out.Version = string(in.String())
		case "expression":
			out.Expression = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe2(out *jwriter.Writer, in RuleIgnored) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"expression\":"
		out.RawString(prefix)
		out.String(string(in.Expression))
	}
	{
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RuleIgnored) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RuleIgnored) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe2(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe3(in *jlexer.Lexer, out *PolicyLoaded) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Version":
			out.Version = string(in.String())
		case "rules_loaded":
			if in.IsNull() {
				in.Skip()
				out.RulesLoaded = nil
			} else {
				in.Delim('[')
				if out.RulesLoaded == nil {
					if !in.IsDelim(']') {
						out.RulesLoaded = make([]*RuleLoaded, 0, 8)
					} else {
						out.RulesLoaded = []*RuleLoaded{}
					}
				} else {
					out.RulesLoaded = (out.RulesLoaded)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *RuleLoaded
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(RuleLoaded)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.RulesLoaded = append(out.RulesLoaded, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rules_ignored":
			if in.IsNull() {
				in.Skip()
				out.RulesIgnored = nil
			} else {
				in.Delim('[')
				if out.RulesIgnored == nil {
					if !in.IsDelim(']') {
						out.RulesIgnored = make([]*RuleIgnored, 0, 8)
					} else {
						out.RulesIgnored = []*RuleIgnored{}
					}
				} else {
					out.RulesIgnored = (out.RulesIgnored)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *RuleIgnored
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(RuleIgnored)
						}
						(*v8).UnmarshalEasyJSON(in)
					}
					out.RulesIgnored = append(out.RulesIgnored, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe3(out *jwriter.Writer, in PolicyLoaded) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Version\":"
		out.RawString(prefix[1:])
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"rules_loaded\":"
		out.RawString(prefix)
		if in.RulesLoaded == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.RulesLoaded {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.RulesIgnored) != 0 {
		const prefix string = ",\"rules_ignored\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.RulesIgnored {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PolicyLoaded) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PolicyLoaded) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe3(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe4(in *jlexer.Lexer, out *NoisyProcessEvent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "pid_count":
			out.Count = uint64(in.Uint64())
		case "threshold":
			out.Threshold = int64(in.Int64())
		case "control_period":
			out.ControlPeriod = time.Duration(in.Int64())
		case "discarded_until":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DiscardedUntil).UnmarshalJSON(data))
			}
		case "process":
			(out.Process).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe4(out *jwriter.Writer, in NoisyProcessEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix[1:])
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"pid_count\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Count))
	}
	{
		const prefix string = ",\"threshold\":"
		out.RawString(prefix)
		out.Int64(int64(in.Threshold))
	}
	{
		const prefix string = ",\"control_period\":"
		out.RawString(prefix)
		out.Int64(int64(in.ControlPeriod))
	}
	{
		const prefix string = ",\"discarded_until\":"
		out.RawString(prefix)
		out.Raw((in.DiscardedUntil).MarshalJSON())
	}
	{
		const prefix string = ",\"process\":"
		out.RawString(prefix)
		(in.Process).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NoisyProcessEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NoisyProcessEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe4(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe5(in *jlexer.Lexer, out *EventLostWrite) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "map":
			out.Name = string(in.String())
		case "per_event":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Lost = make(map[string]uint64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v13 uint64
					v13 = uint64(in.Uint64())
					(out.Lost)[key] = v13
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe5(out *jwriter.Writer, in EventLostWrite) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix[1:])
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"map\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"per_event\":"
		out.RawString(prefix)
		if in.Lost == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v14First := true
			for v14Name, v14Value := range in.Lost {
				if v14First {
					v14First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v14Name))
				out.RawByte(':')
				out.Uint64(uint64(v14Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventLostWrite) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventLostWrite) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe5(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe6(in *jlexer.Lexer, out *EventLostRead) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "map":
			out.Name = string(in.String())
		case "lost":
			out.Lost = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe6(out *jwriter.Writer, in EventLostRead) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix[1:])
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"map\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"lost\":"
		out.RawString(prefix)
		out.Float64(float64(in.Lost))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventLostRead) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe6(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventLostRead) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe6(l, v)
}
func easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe7(in *jlexer.Lexer, out *AbnormalPathEvent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "triggering_event":
			if in.IsNull() {
				in.Skip()
				out.Event = nil
			} else {
				if out.Event == nil {
					out.Event = new(EventSerializer)
				}
				(*out.Event).UnmarshalEasyJSON(in)
			}
		case "path_resolution_error":
			out.PathResolutionError = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe7(out *jwriter.Writer, in AbnormalPathEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix[1:])
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"triggering_event\":"
		out.RawString(prefix)
		if in.Event == nil {
			out.RawString("null")
		} else {
			(*in.Event).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"path_resolution_error\":"
		out.RawString(prefix)
		out.String(string(in.PathResolutionError))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AbnormalPathEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF8f9ddd1EncodeGithubComDataDogDatadogAgentPkgSecurityProbe7(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AbnormalPathEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF8f9ddd1DecodeGithubComDataDogDatadogAgentPkgSecurityProbe7(l, v)
}
