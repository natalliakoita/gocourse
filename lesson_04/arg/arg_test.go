package arg

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// test fot area circle
func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value equal 0",
			fields: fields{
				radius: 0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal -1",
			fields: fields{
				radius: -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				radius: 1,
			},
			want:    float64(math.Pi),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.radius,
			}
			got, err := c.Area()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// test fot string when outpur radius circle
func TestCircle_String(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "whether there is a value radius in string",
			fields: fields{
				radius: 0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.radius,
			}
			got := c.String()
			s := fmt.Sprintf("%.2f", c.Radius)
			require.Contains(t, got, s)
		})
	}
}

// test fot perimetr circle
func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value equal 0",
			fields: fields{
				radius: 0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal -1",
			fields: fields{
				radius: -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				radius: 1,
			},
			want:    float64(6.283185307179586),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.radius,
			}
			got, err := c.Perimeter()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// test fot area rectangle
func TestRectangle_Area(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value of height equal 0",
			fields: fields{
				height: 0,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal 0",
			fields: fields{
				height: 1,
				width:  0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of height equal -1",
			fields: fields{
				height: -1,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal -1",
			fields: fields{
				height: 1,
				width:  -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				height: 1,
				width:  1,
			},
			want:    float64(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.height,
				Width:  tt.fields.width,
			}
			got, err := r.Area()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

// test fot string when outpur height and width rectangle
func TestRectangle_String(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "whether there is a value height in string",
			fields: fields{
				height: 1,
			},
			want: "",
		},

		{
			name: "whether there is a value width in string",
			fields: fields{
				width: 1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.height,
				Width:  tt.fields.width,
			}
			got := r.String()
			h := fmt.Sprintf("%.2f", r.Height)
			require.Contains(t, got, h)
			w := fmt.Sprintf("%.2f", r.Width)
			require.Contains(t, got, w)
		})
	}
}

// test fot perimeter rectangle
func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "when value of height equal 0",
			fields: fields{
				height: 0,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal 0",
			fields: fields{
				height: 1,
				width:  0,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of height equal -1",
			fields: fields{
				height: -1,
				width:  1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value of width equal -1",
			fields: fields{
				height: 1,
				width:  -1,
			},
			want:    float64(0),
			wantErr: true,
		},

		{
			name: "when value equal 1",
			fields: fields{
				height: 1,
				width:  1,
			},
			want:    float64(4),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Height: tt.fields.height,
				Width:  tt.fields.width,
			}
			got, err := r.Perimeter()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDescribeShape(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		wantErr bool
	}{
		{
			name: "success for rectangle",
			shape: Rectangle{
				Height: 1,
				Width:  1,
			},
			wantErr: false,
		},

		{
			name: "success for circle",
			shape: Circle{
				Radius: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DescribeShape(tt.shape)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
