package envy

import (
	"errors"
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
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := ParseBool(envName)
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
		skipSet bool
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
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := ParseInt(envName, tt.args.base, tt.args.bitSize)
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

func TestGet(t *testing.T) {
	type args struct {
		value   string
		skipSet bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Get environmet variable value",
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
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv(envName)
			if !tt.args.skipSet {
				os.Setenv(envName, tt.args.value)
			}
			got, err := Get(envName)
			if baseErr := errors.Unwrap(err); (baseErr != nil && tt.wantErr != nil) &&
				baseErr.Error() != tt.wantErr.Error() {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
