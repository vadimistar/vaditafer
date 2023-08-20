package taf

import (
	"reflect"
	"testing"
)

func Test_createdAt(t *testing.T) {
	tests := []struct {
		code       string
		wantDay    int
		wantHour   int
		wantMinute int
		wantErr    bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "ABBFBDFBFB",
			wantErr: true,
		},
		{
			code:    "2013514",
			wantErr: true,
		},
		{
			code:       "201351Z",
			wantDay:    20,
			wantHour:   13,
			wantMinute: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotT, err := createdAt(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("createdAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(gotT.Day(), tt.wantDay) {
				t.Errorf("day = %v, want %v", gotT.Day(), tt.wantDay)
			}
			if !reflect.DeepEqual(gotT.Hour(), tt.wantHour) {
				t.Errorf("hour = %v, want %v", gotT.Hour(), tt.wantHour)
			}
			if !reflect.DeepEqual(gotT.Minute(), tt.wantMinute) {
				t.Errorf("minute = %v, want %v", gotT.Minute(), tt.wantMinute)
			}
		})
	}
}

func Test_period(t *testing.T) {
	tests := []struct {
		code         string
		wantFromDay  int
		wantFromHour int
		wantToDay    int
		wantToHour   int
		wantErr      bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "3234424242",
			wantErr: true,
		},
		{
			code:    "24FS/SF33",
			wantErr: true,
		},
		{
			code:         "2002/2004",
			wantFromDay:  20,
			wantFromHour: 2,
			wantToDay:    20,
			wantToHour:   4,
		},
		{
			code:         "3102/0104",
			wantFromDay:  31,
			wantFromHour: 2,
			wantToDay:    1,
			wantToHour:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotFrom, gotTo, err := period(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("period() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(gotFrom.Day(), tt.wantFromDay) {
				t.Errorf("from date = %v, want %v", gotFrom.Day(), tt.wantFromDay)
			}
			if !reflect.DeepEqual(gotFrom.Hour(), tt.wantFromHour) {
				t.Errorf("from hour = %v, want %v", gotFrom.Hour(), tt.wantFromHour)
			}
			if !reflect.DeepEqual(gotTo.Day(), tt.wantToDay) {
				t.Errorf("to date = %v, want %v", gotTo.Day(), tt.wantToDay)
			}
			if !reflect.DeepEqual(gotTo.Hour(), tt.wantToHour) {
				t.Errorf("to hour = %v, want %v", gotTo.Hour(), tt.wantToHour)
			}
		})
	}
}

func Test_from(t *testing.T) {
	tests := []struct {
		code       string
		wantDay    int
		wantHour   int
		wantMinute int
		wantErr    bool
	}{
		{
			code:    "",
			wantErr: true,
		},
		{
			code:    "32242333",
			wantErr: true,
		},
		{
			code:       "FM300204",
			wantDay:    30,
			wantHour:   2,
			wantMinute: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			gotT, err := from(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("from() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(gotT.Day(), tt.wantDay) {
				t.Errorf("day = %v, want %v", gotT.Day(), tt.wantDay)
			}
			if !reflect.DeepEqual(gotT.Hour(), tt.wantHour) {
				t.Errorf("hour = %v, want %v", gotT.Hour(), tt.wantHour)
			}
			if !reflect.DeepEqual(gotT.Minute(), tt.wantMinute) {
				t.Errorf("minute = %v, want %v", gotT.Minute(), tt.wantMinute)
			}
		})
	}
}
