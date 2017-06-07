package rocks

import (
	"github.com/uber-go/dosa"
	"time"
)

type Meeting struct {
	dosa.Entity `dosa:"primaryKey=Handle"`
	Handle       string
        MeetingUUID  dosa.UUID
}

type Question struct {
	dosa.Entity               `dosa:"primaryKey=(MeetingUUID, QuestionUUID)"`
	MeetingUUID, QuestionUUID dosa.UUID
	Question                  string
	Votes                     int32
	Created                   time.Time
	Token			  string
}
