package portfolio

import (
	"math"
	"testing"
	"time"
)

func TestPortfolio_Profit(t *testing.T) {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	invalidEndDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		portfolio Portfolio
		start     time.Time
		end       time.Time
		want      int
		wantErr   error
	}{
		{
			name: "valid dates with profit",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 100)
					s.AddPrice(endDate, 150)
					return s
				}(),
			}),
			start:   startDate,
			end:     endDate,
			want:    50,
			wantErr: nil,
		},
		{
			name: "invalid date range",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 100)
					s.AddPrice(invalidEndDate, 150)
					return s
				}(),
			}),
			start:   startDate,
			end:     invalidEndDate,
			want:    0,
			wantErr: invalidDateRangeErr,
		},
		{
			name: "missing price data",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 100)
					return s
				}(),
			}),
			start:   startDate,
			end:     endDate,
			want:    0,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.portfolio.Profit(tt.start, tt.end)
			if err != tt.wantErr {
				t.Errorf("Portfolio.Profit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Portfolio.Profit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortfolio_AnnualizedReturn(t *testing.T) {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		portfolio Portfolio
		start     time.Time
		end       time.Time
		want      float64
		wantErr   error
	}{
		{
			name: "one year return",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 100)
					s.AddPrice(endDate, 110)
					return s
				}(),
			}),
			start:   startDate,
			end:     endDate,
			want:    0.10,
			wantErr: nil,
		},
		{
			name: "zero value portfolio",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 0)
					s.AddPrice(endDate, 100)
					return s
				}(),
			}),
			start:   startDate,
			end:     endDate,
			want:    0,
			wantErr: portfolioZeroValueErr,
		},
		{
			name: "negative return",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 100)
					s.AddPrice(endDate, 90)
					return s
				}(),
			}),
			start:   startDate,
			end:     endDate,
			want:    -0.10,
			wantErr: nil,
		},
		{
			name: "same day for start and end dates",
			portfolio: NewPortfolio([]Stock{
				func() Stock {
					s := NewStock()
					s.AddPrice(startDate, 100)
					return s
				}(),
			}),
			start:   startDate,
			end:     startDate,
			want:    0,
			wantErr: invalidDateRangeErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.portfolio.AnnualizedReturn(tt.start, tt.end)
			if err != tt.wantErr {
				t.Errorf("Portfolio.AnnualizedReturn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			if math.Abs(got-tt.want) < 1e-05 {
				t.Errorf("Portfolio.AnnualizedReturn() = %v, want %v, abs error %v", got, tt.want, math.Abs(got-tt.want))
			}
		})
	}
}
