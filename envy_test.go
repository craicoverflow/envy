package envy

import (
	"errors"
	"os"
	"testing"
)

func TestParseBool(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "Throw error when environment variable is not set",
			args: args{
				value: "",
			},
			want:    false,
			wantErr: ErrNotFound,
		},
		{
			name: "Parse 'true' as environment variable",
			args: args{
				value: "true",
			},
			want:    true,
			wantErr: nil,
		},
		{
			name: "Parse 'false' as environment variable",
			args: args{
				value: "false",
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "Parse '0' as environment variable",
			args: args{
				value: "0",
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "Parse number string to bool fails",
			args: args{
				value: "22",
			},
			want:    false,
			wantErr: ErrSyntax,
		},
		{
			name: "Parsing alphabetical character string as environment variable fails",
			args: args{
				value: "a",
			},
			want:    false,
			wantErr: ErrSyntax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envKey := "ENV_KEY"
			os.Setenv(envKey, tt.args.value)
			got, err := ParseBool(envKey)
			if baseErr := errors.Unwrap(err); baseErr != tt.wantErr {
				t.Errorf("ParseBool() error = %v, wantErr %v", baseErr, tt.wantErr)
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
		base    int
		bitSize int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{
			name: "Integer value should parse successfully",
			args: args{
				value:   "1",
				base:    0,
				bitSize: 32,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Float value should fail to parse",
			args: args{
				value:   "1.0",
				base:    0,
				bitSize: 64,
			},
			want:    0,
			wantErr: ErrSyntax,
		},
		{
			name: "Should be invalid bit size error",
			args: args{
				value:   "18446744073709551616",
				base:    0,
				bitSize: 1<<32 - 1,
			},
			want:    0,
			wantErr: errors.New("invalid bit size 4294967295"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			envKey := "ENV_KEY"
			os.Setenv(envKey, tt.args.value)
			got, err := ParseInt(envKey, tt.args.base, tt.args.bitSize)
			if baseErr := errors.Unwrap(err); (baseErr != nil && tt.wantErr != nil) &&
				baseErr.Error() != tt.wantErr.Error() {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
