package operations_test

import (
	"encoding/json"
	"testing"
	"time"

	omise "github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	"github.com/omise/omise-go/schedule"
	r "github.com/stretchr/testify/require"
)

func TestCreateChargeScheduleMarshal(t *testing.T) {
	testdata := []struct {
		req      *CreateChargeSchedule
		expected string
	}{
		{
			req: &CreateChargeSchedule{
				Every:     3,
				Period:    schedule.PeriodDay,
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Customer:  "customer_id",
				Amount:    100000,
			},
			expected: `{"every":3,"period":"day","start_date":"2017-05-15","end_date":"2018-05-15","charge":{"customer":"customer_id","amount":100000}}`,
		},
		{
			req: &CreateChargeSchedule{
				Every:  3,
				Period: schedule.PeriodWeek,
				Weekdays: schedule.Weekdays{
					schedule.Monday,
					schedule.Saturday,
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Customer:  "customer_id",
				Amount:    100000,
			},
			expected: `{"every":3,"period":"week","start_date":"2017-05-15","end_date":"2018-05-15","on":{"weekdays":["monday","saturday"]},"charge":{"customer":"customer_id","amount":100000}}`,
		},
		{
			req: &CreateChargeSchedule{
				Every:       3,
				Period:      schedule.PeriodMonth,
				DaysOfMonth: schedule.DaysOfMonth{1, 15},
				StartDate:   "2017-05-15",
				EndDate:     "2018-05-15",
				Customer:    "customer_id",
				Amount:      100000,
			},
			expected: `{"every":3,"period":"month","start_date":"2017-05-15","end_date":"2018-05-15","on":{"days_of_month":[1,15]},"charge":{"customer":"customer_id","amount":100000}}`,
		},
		{
			req: &CreateChargeSchedule{
				Every:          3,
				Period:         schedule.PeriodMonth,
				WeekdayOfMonth: "last_thursday",
				StartDate:      "2017-05-15",
				EndDate:        "2018-05-15",
				Customer:       "customer_id",
				Amount:         100000,
			},
			expected: `{"every":3,"period":"month","start_date":"2017-05-15","end_date":"2018-05-15","on":{"weekday_of_month":"last_thursday"},"charge":{"customer":"customer_id","amount":100000}}`,
		},
	}

	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestCreateChargeSchedule_Network(t *testing.T) {
	// CustomerID must have this customer in test server
	const CustomerID = `cust_57z9e1nce0wvbbkvef1`

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd, create := &omise.Schedule{}, &CreateChargeSchedule{
		Every:  3,
		Period: schedule.PeriodWeek,
		Weekdays: schedule.Weekdays{
			schedule.Monday,
			schedule.Saturday,
		},
		EndDate:  "2018-05-15",
		Customer: CustomerID,
		Amount:   100000,
	}
	client.MustDo(schd, create)
}

func TestCreateTransferScheduleMarshal(t *testing.T) {
	testdata := []struct {
		req      *CreateTransferSchedule
		expected string
	}{
		{
			req: &CreateTransferSchedule{
				Every:     3,
				Period:    schedule.PeriodDay,
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",
				Recipient: "recipient_id",
				Amount:    100000,
			},
			expected: `{"every":3,"period":"day","start_date":"2017-05-15","end_date":"2018-05-15","transfer":{"recipient":"recipient_id","amount":100000}}`,
		},
		{
			req: &CreateTransferSchedule{
				Every:  3,
				Period: schedule.PeriodWeek,
				Weekdays: schedule.Weekdays{
					schedule.Monday,
					schedule.Saturday,
				},
				StartDate: "2017-05-15",
				EndDate:   "2018-05-15",

				Recipient:           "recipient_id",
				PercentageOfBalance: 20.35,
			},
			expected: `{"every":3,"period":"week","start_date":"2017-05-15","end_date":"2018-05-15","on":{"weekdays":["monday","saturday"]},"transfer":{"recipient":"recipient_id","percentage_of_balance":20.35}}`,
		},
		{
			req: &CreateTransferSchedule{
				Every:       3,
				Period:      schedule.PeriodMonth,
				DaysOfMonth: schedule.DaysOfMonth{1, 15},
				StartDate:   "2017-05-15",
				EndDate:     "2018-05-15",

				Recipient: "recipient_id",
				Amount:    100000,
			},
			expected: `{"every":3,"period":"month","start_date":"2017-05-15","end_date":"2018-05-15","on":{"days_of_month":[1,15]},"transfer":{"recipient":"recipient_id","amount":100000}}`,
		},
		{
			req: &CreateTransferSchedule{
				Every:          3,
				Period:         schedule.PeriodMonth,
				WeekdayOfMonth: "last_thursday",
				StartDate:      "2017-05-15",
				EndDate:        "2018-05-15",

				Recipient:           "recipient_id",
				PercentageOfBalance: 50.55,
			},
			expected: `{"every":3,"period":"month","start_date":"2017-05-15","end_date":"2018-05-15","on":{"weekday_of_month":"last_thursday"},"transfer":{"recipient":"recipient_id","percentage_of_balance":50.55}}`,
		},
	}

	for _, td := range testdata {
		b, err := json.Marshal(td.req)
		r.Nil(t, err, "err should be nothing")
		r.Equal(t, td.expected, string(b))
	}
}

func TestCreateTransferSchedule_Network(t *testing.T) {
	// RecipientID must have this recipient in test server
	const RecipientID = `recp_57z9e1nce0wvbbkvef1`

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd, create := &omise.Schedule{}, &CreateTransferSchedule{
		Every:  3,
		Period: schedule.PeriodWeek,
		Weekdays: schedule.Weekdays{
			schedule.Monday,
			schedule.Saturday,
		},
		EndDate: "2018-05-15",

		Recipient: RecipientID,
		Amount:    100000,
	}
	client.MustDo(schd, create)
}

func TestCreateSchedule(t *testing.T) {
	const (
		ScheduleID = "schd_57z9hj228pusa652nk1"
	)

	client := testutil.NewFixedClient(t)

	schd := &omise.Schedule{}
	client.MustDo(schd, &CreateChargeSchedule{})
	r.Equal(t, ScheduleID, schd.ID)
}

func TestListSchedule(t *testing.T) {
	client := testutil.NewFixedClient(t)
	schds := &omise.ScheduleList{}
	client.MustDo(schds, &ListSchedules{})

	r.Len(t, schds.Data, 2)

	r.Equal(t, "schd_57zhl296uxc7yiun6xa", schds.Data[0].ID)
	r.NotNil(t, schds.Data[0].Charge)
	r.Nil(t, schds.Data[0].Transfer)

	r.Equal(t, "schd_57zhl296uxc7yiun6xx", schds.Data[1].ID)
	r.NotNil(t, schds.Data[1].Transfer)
	r.Nil(t, schds.Data[1].Charge)
}

func TestListSchedules_Network(t *testing.T) {
	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schds, list := &omise.ScheduleList{}, &ListSchedules{
		List{
			Limit: 100,
			From:  time.Date(2017, 5, 16, 0, 0, 0, 0, time.Local),
		},
	}
	client.MustDo(schds, list)

	t.Logf("Schedules Len: %d\n", len(schds.Data))
	t.Logf("%#v\n", schds)
}

func TestRetrieveSchedule(t *testing.T) {
	ScheduleID := "schd_57z9hj228pusa652nk1"

	client := testutil.NewFixedClient(t)
	schd := &omise.Schedule{}
	client.MustDo(schd, &RetrieveSchedule{ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Transfer)
	r.Equal(t, 100000, schd.Charge.Amount)
	r.Equal(t, schedule.Active, schd.Status)
	r.Len(t, schd.NextOccurrences, 30)

	ScheduleID = "schd_57z9hj228pusa652nk2"

	schd = &omise.Schedule{}
	client.MustDo(schd, &RetrieveSchedule{ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Charge)
	r.Equal(t, 100000, *schd.Transfer.Amount)
	r.Equal(t, schedule.Active, schd.Status)
	r.Len(t, schd.NextOccurrences, 30)
}

func TestRetrieveSchedule_Network(t *testing.T) {
	// ScheduleID must have this schedule in test server
	ScheduleID := "schd_57z9hj228pusa652nk1"

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd := &omise.Schedule{}
	client.MustDo(schd, &RetrieveSchedule{ScheduleID})

	t.Logf("%#v\n", schd)
}

func TestDestroySchedule(t *testing.T) {
	ScheduleID := "schd_57z9hj228pusa652nk1"

	client := testutil.NewFixedClient(t)
	schd := &omise.Schedule{}
	client.MustDo(schd, &DestroySchedule{ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Transfer)
	r.Equal(t, 100000, schd.Charge.Amount)
	r.Equal(t, schedule.Deleted, schd.Status)
	r.Len(t, schd.NextOccurrences, 30)

	ScheduleID = "schd_57z9hj228pusa652nk2"

	schd = &omise.Schedule{}
	client.MustDo(schd, &DestroySchedule{ScheduleID})
	r.Equal(t, ScheduleID, schd.ID)
	r.Nil(t, schd.Charge)
	r.Equal(t, 100000, *schd.Transfer.Amount)
	r.Equal(t, schedule.Deleted, schd.Status)
	r.Len(t, schd.NextOccurrences, 30)
}

func TestDestroySchedule_Network(t *testing.T) {
	// ScheduleID must have this schedule in test server
	ScheduleID := "schd_57z9hj228pusa652nk1"

	testutil.Require(t, "network")
	client := testutil.NewTestClient(t)
	schd := &omise.Schedule{}
	client.MustDo(schd, &DestroySchedule{ScheduleID})

	t.Logf("%#v\n", schd)
}
