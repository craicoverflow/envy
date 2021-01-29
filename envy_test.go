package envy

import (
	"os"
	"testing"
)

const envName = "TEST_ENV"

func TestParseBool(t *testing.T) {
	type args struct {
		value   string
		skipSet bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Throw error when environment variable is not set",
			args: args{
				value: "",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Parse 'true' as environment variable",
			args: args{
				value: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Parse 'false' as environment variable",
			args: args{
				value: "false",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Parse '0' as environment variable",
			args: args{
				value: "0",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Parse number string to bool fails",
			args: args{
				value: "22",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Parsing alphabetical character string as environment variable fails",
			args: args{
				value: "a",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := ParseBool(envName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("ParseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	type args struct {
		value   string
		skipSet bool
		base    int
		bitSize int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Should throw ErrNotFound when environment variable does not exist",
			args: args{
				value:   "1",
				base:    0,
				bitSize: 32,
				skipSet: true,
			},
			wantErr: true,
		},
		{
			name: "Integer value should parse successfully",
			args: args{
				value:   "1",
				base:    0,
				bitSize: 32,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Integer value should fail to parse",
			args: args{
				value:   "1.0",
				base:    0,
				bitSize: 64,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := ParseInt(envName, tt.args.base, tt.args.bitSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	type args struct {
		value   string
		skipSet bool
		bitSize int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Should throw ErrNotFound when environment variable does not exist",
			args: args{
				value:   "1",
				bitSize: 32,
				skipSet: true,
			},
			wantErr: true,
		},
		{
			name: "Float value should parse successfully",
			args: args{
				value:   "1.502",
				bitSize: 64,
			},
			want:    1.502,
			wantErr: false,
		},
		{
			name: "Float value should fail to parse non-float value",
			args: args{
				value:   "dsfdk",
				bitSize: 64,
			},
			want: 0.00,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := ParseFloat(envName, tt.args.bitSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		value   string
		skipSet bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get environment variable value",
			args: args{
				value: "myname",
			},
			want: "myname",
		},
		{
			name: "Should throw ErrNotFound",
			args: args{
				skipSet: true,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		// nolint
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := Get(envName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
