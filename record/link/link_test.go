package link

import "testing"

func TestIsRecord(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name           string
		args           args
		wantRecordFlag bool
		wantErr        bool
	}{
		{
			name:           "t1",
			args:           args{link: "http://www.whtdld.com/"},
			wantRecordFlag: true,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordFlag, err := IsRecord(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRecordFlag != tt.wantRecordFlag {
				t.Errorf("IsRecord() = %v, want %v", gotRecordFlag, tt.wantRecordFlag)
			}
		})
	}
}
