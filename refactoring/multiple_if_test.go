package refactoring

import "testing"

func TestOptimize(t *testing.T) {
	type args struct {
		o OrderInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "日本拼车请求", args: args{o: OrderInfo{
			ProductId:   jpProductId,
			ExtraType:   carpoolExtraType,
			OrderStatus: 0,
		}}, wantErr: false},
		{name: "日本open_ride请求", args: args{o: OrderInfo{
			ProductId:   jpProductId,
			ExtraType:   openRideExtraType,
			OrderStatus: 0,
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Optimize(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Optimize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
