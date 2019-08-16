package expr

import (
	"testing"
)

func TestParseExpresstion(t *testing.T) {
	type args struct {
		str          string
		inputContext map[string]interface{}
	}
	var tests = []struct {
		name       string
		args       args
		wantResult float64
		wantErr    bool
	}{
		{name: "1*2+3", args: args{str: "1*2+3", inputContext: nil}, wantResult: 5, wantErr: false},
		{name: "1+2*3", args: args{str: "1+2*3", inputContext: nil}, wantResult: 7, wantErr: false},
		{name: "1+2", args: args{str: "1+2", inputContext: nil}, wantResult: 3, wantErr: false},
		{name: "1+X", args: args{str: "1+2", inputContext: map[string]interface{}{"X": 2}}, wantResult: 3, wantErr: false},
		{name: "floor(0.2)", args: args{str: "floor(0.2)", inputContext: nil}, wantResult: 0, wantErr: false},
		{name: "ceil(1.2)", args: args{str: "ceil(1.2)", inputContext: nil}, wantResult: 2, wantErr: false},
		{name: "2+(2^ceil(2))", args: args{str:"2+(2^ceil(2))" , inputContext: nil}, wantResult: 6, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ParseExpresstion(tt.args.str, tt.args.inputContext)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseExpresstion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ParseExpresstion() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}


