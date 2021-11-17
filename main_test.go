package main

import (
	"reflect"
	"testing"
)

func Test_varnishstatParams_make(t *testing.T) {
	type fields struct {
		Instance string
		VSM      string
		Filter   string
	}
	tests := []struct {
		name       string
		fields     fields
		wantParams []string
	}{
		{
			name: "NoFilterFlags",
			fields: fields{
				Instance: "",
				VSM: "",
				Filter: "",
			},
			wantParams: nil,
		},
		{
			name: "FilterFlags",
			fields: fields{
				Instance: "",
				VSM: "",
				Filter: `^VBE.*\.bereq_hdrbytes,^VBE.*\.bereq_bodybytes,^VBE.*\.beresp_hdrbytes,^VBE.*\.beresp_bodybytes`,
			},
			wantParams: []string{"-f", `^VBE.*\.bereq_hdrbytes -f ^VBE.*\.bereq_bodybytes -f ^VBE.*\.beresp_hdrbytes -f ^VBE.*\.beresp_bodybytes`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &varnishstatParams{
				Instance: tt.fields.Instance,
				VSM:      tt.fields.VSM,
				Filter:   tt.fields.Filter,
			}
			if gotParams := p.make(); !reflect.DeepEqual(gotParams, tt.wantParams) {
				t.Errorf("make() = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}
