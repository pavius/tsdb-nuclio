// +build unit

/*
Copyright 2018 Iguazio Systems Ltd.

Licensed under the Apache License, Version 2.0 (the "License") with
an addition restriction as set forth herein. You may not use this
file except in compliance with the License. You may obtain a copy of
the License at http://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing
permissions and limitations under the License.

In addition, you may not use the software for any purposes that are
illegal under applicable law, and the grant of the foregoing license
under the Apache 2.0 license is conditioned upon your compliance with
such restriction.
*/

package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type testTimeSuite struct {
	suite.Suite
}

func (suite *testTimeSuite) TestStr2durationOneHour() {
	res, err := Str2duration("1h")
	suite.Require().Nil(err)
	suite.Require().Equal(int64(time.Hour/time.Millisecond), res)
}

func (suite *testTimeSuite) TestStr2durationFiveMinutes() {
	res, err := Str2duration("5m")
	suite.Require().Nil(err)
	suite.Require().Equal(int64(5*time.Minute/time.Millisecond), res)
}

func (suite *testTimeSuite) TestStr2durationZeroDays() {
	res, err := Str2duration("0d")
	suite.Require().Nil(err)
	suite.Require().Equal(int64(0), res)
}

func (suite *testTimeSuite) TestStr2durationSixSeconds() {
	res, err := Str2duration("6s")
	suite.Require().Nil(err)
	suite.Require().Equal(int64(6*time.Second/time.Millisecond), res)
}

func (suite *testTimeSuite) TestStr2durationZero() {
	res, err := Str2duration("0")
	suite.Require().Nil(err)
	suite.Require().Equal(int64(0), res)
}

func (suite *testTimeSuite) TestStr2durationNegative() {
	_, err := Str2duration("-1")
	suite.Require().Error(err)
}

func (suite *testTimeSuite) TestStr2durationBadFormat() {
	_, err := Str2duration("1dhm")
	suite.Require().Error(err)
}

func (suite *testTimeSuite) TestStr2unixTime() {
	expectedDuration, err := Str2duration("2m")
	suite.Require().Nil(err)
	endTime, err := Str2unixTime("now+1m")
	suite.Require().Nil(err)
	startTime, err := Str2unixTime("now-1m")
	suite.Require().Nil(err)
	actualDuration := endTime - startTime
	// "now" might slightly change from call to call, so we allow a safety range
	suite.Require().LessOrEqual(expectedDuration-actualDuration, int64(2))
	suite.Require().GreaterOrEqual(expectedDuration-actualDuration, int64(0))
}

func (suite *testTimeSuite) TestStr2unixTimeWithNow() {
	expectedDuration, err := Str2duration("1m")
	suite.Require().Nil(err)
	endTime, err := Str2unixTime("now")
	suite.Require().Nil(err)
	startTime, err := Str2unixTime("now-1m")
	suite.Require().Nil(err)
	actualDuration := endTime - startTime
	// "now" might slightly change from call to call, so we allow a safety range
	suite.Require().LessOrEqual(expectedDuration-actualDuration, int64(2))
	suite.Require().GreaterOrEqual(expectedDuration-actualDuration, int64(0))
}

func (suite *testTimeSuite) TestStr2unixTimeWithNowPlus() {
	expectedDuration, err := Str2duration("1m")
	suite.Require().Nil(err)
	endTime, err := Str2unixTime("now+")
	suite.Require().Nil(err)
	startTime, err := Str2unixTime("now-1m")
	suite.Require().Nil(err)
	actualDuration := endTime - startTime
	// "now" might slightly change from call to call, so we allow a safety range
	suite.Require().LessOrEqual(expectedDuration-actualDuration, int64(2))
	suite.Require().GreaterOrEqual(expectedDuration-actualDuration, int64(0))
}

func (suite *testTimeSuite) TestStr2unixTimeWithNowPlusMinus() {
	expectedDuration, err := Str2duration("0m")
	suite.Require().Nil(err)
	endTime, err := Str2unixTime("now+")
	suite.Require().Nil(err)
	startTime, err := Str2unixTime("now-")
	suite.Require().Nil(err)
	actualDuration := endTime - startTime
	// "now" might slightly change from call to call, so we allow a safety range
	suite.Require().LessOrEqual(expectedDuration-actualDuration, int64(2))
	suite.Require().GreaterOrEqual(expectedDuration-actualDuration, int64(0))
}

func (suite *testTimeSuite) TestRFCStr2unixTime() {
	timeStr := "2020-07-20T14:55:56.770808000Z"
	realTime, err := time.Parse(time.RFC3339, timeStr)
	suite.Require().NoError(err)
	timeInMillis, err := Str2unixTime(timeStr)
	suite.Require().NoError(err)
	suite.Require().Equal(realTime.UnixNano()/int64(time.Millisecond), timeInMillis)
}

func TestTimeSuite(t *testing.T) {
	suite.Run(t, new(testTimeSuite))
}
