package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	. "gopkg.in/check.v1"
	"testing"
)

var _ = Suite(&MessageTests{})

type MessageTests struct{}

var msgResult *Message

func BenchmarkParseMessage(b *testing.B) {
	rawMsg := []byte("8=FIX.4.29=10435=D34=249=TW52=20140515-19:49:56.65956=ISLD11=10021=140=154=155=TSLA60=00010101-00:00:00.00010=039")

	var msg *Message
	for i := 0; i < b.N; i++ {
		msg, _ = ParseMessage(rawMsg)
	}

	msgResult = msg
}

func (s *FieldMapTests) TestReverseRoute(c *C) {
	msg, err := ParseMessage([]byte("8=FIX.4.29=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"))
	c.Check(err, IsNil)

	builder := msg.ReverseRoute()

	targetCompID := new(fix.StringValue)
	err = builder.Header().GetField(tag.TargetCompID, targetCompID)
	c.Check(err, IsNil)
	c.Check(targetCompID.Value, Equals, "TW")

	targetSubID := new(fix.StringValue)
	err = builder.Header().GetField(tag.TargetSubID, targetSubID)
	c.Check(err, IsNil)
	c.Check(targetSubID.Value, Equals, "KK")

	targetLocationID := new(fix.StringValue)
	err = builder.Header().GetField(tag.TargetLocationID, targetLocationID)
	c.Check(err, IsNil)
	c.Check(targetLocationID.Value, Equals, "JV")

	senderCompID := new(fix.StringValue)
	err = builder.Header().GetField(tag.SenderCompID, senderCompID)
	c.Check(err, IsNil)
	c.Check(senderCompID.Value, Equals, "ISLD")

	senderSubID := new(fix.StringValue)
	err = builder.Header().GetField(tag.SenderSubID, senderSubID)
	c.Check(err, IsNil)
	c.Check(senderSubID.Value, Equals, "AP")

	senderLocationID := new(fix.StringValue)
	err = builder.Header().GetField(tag.SenderLocationID, senderLocationID)
	c.Check(err, IsNil)
	c.Check(senderLocationID.Value, Equals, "RY")

	deliverToCompID := new(fix.StringValue)
	err = builder.Header().GetField(tag.DeliverToCompID, deliverToCompID)
	c.Check(err, IsNil)
	c.Check(deliverToCompID.Value, Equals, "JCD")

	deliverToSubID := new(fix.StringValue)
	err = builder.Header().GetField(tag.DeliverToSubID, deliverToSubID)
	c.Check(err, IsNil)
	c.Check(deliverToSubID.Value, Equals, "CS")

	deliverToLocationID := new(fix.StringValue)
	err = builder.Header().GetField(tag.DeliverToLocationID, deliverToLocationID)
	c.Check(err, IsNil)
	c.Check(deliverToLocationID.Value, Equals, "BB")

	onBehalfOfCompID := new(fix.StringValue)
	err = builder.Header().GetField(tag.OnBehalfOfCompID, onBehalfOfCompID)
	c.Check(err, IsNil)
	c.Check(onBehalfOfCompID.Value, Equals, "MG")

	onBehalfOfSubID := new(fix.StringValue)
	err = builder.Header().GetField(tag.OnBehalfOfSubID, onBehalfOfSubID)
	c.Check(err, IsNil)
	c.Check(onBehalfOfSubID.Value, Equals, "CB")

	onBehalfOfLocationID := new(fix.StringValue)
	err = builder.Header().GetField(tag.OnBehalfOfLocationID, onBehalfOfLocationID)
	c.Check(err, IsNil)
	c.Check(onBehalfOfLocationID.Value, Equals, "BH")
}

func (s *FieldMapTests) TestReverseRouteIgnoreEmpty(c *C) {
	msg, err := ParseMessage([]byte("8=FIX.4.09=12835=D34=249=TW52=20060102-15:04:0556=ISLD115=116=CS128=MG129=CB11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"))
	c.Check(err, IsNil)
	builder := msg.ReverseRoute()

	//don't reverse if empty
	deliverToCompID := new(fix.StringValue)
	err = builder.Header().GetField(tag.DeliverToCompID, deliverToCompID)
	c.Check(err, NotNil)

}

func (s *FieldMapTests) TestReverseRouteFIX40(c *C) {
	//onbehalfof/deliverto location id not supported in fix 4.0

	msg, err := ParseMessage([]byte("8=FIX.4.09=17135=D34=249=TW50=KK52=20060102-15:04:0556=ISLD57=AP144=BB115=JCD116=CS128=MG129=CB142=JV143=RY145=BH11=ID21=338=10040=w54=155=INTC60=20060102-15:04:0510=123"))

	c.Check(err, IsNil)
	builder := msg.ReverseRoute()

	deliverToLocationID := new(fix.StringValue)
	err = builder.Header().GetField(tag.DeliverToLocationID, deliverToLocationID)
	c.Check(err, NotNil)

	onBehalfOfLocationID := new(fix.StringValue)
	err = builder.Header().GetField(tag.OnBehalfOfLocationID, onBehalfOfLocationID)
	c.Check(err, NotNil)
}