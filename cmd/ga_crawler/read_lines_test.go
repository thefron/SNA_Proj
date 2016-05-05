package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	utc, err := time.LoadLocation("UTC")
	assert.Nil(t, err)
	c := make(chan *Event, 10)
	err = readLines("./test_data/ten_line_sample.json", c)
	assert.Nil(t, err)

	event := <-c
	assert.Equal(t, "IssuesEvent", event.Type)
	assert.Equal(t, 793041, event.Actor.Id)
	assert.Equal(t, "OndraM", event.Actor.Login)
	assert.Equal(t, 35304655, event.Repo.Id)
	assert.Equal(t, "lmc-eu/steward", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 0, 0, utc), event.CreatedAt)
	assert.Equal(t, 429615, event.Org.Id)
	assert.Equal(t, "lmc-eu", event.Org.Login)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 793041, event.Actor.Id)
	assert.Equal(t, "OndraM", event.Actor.Login)
	assert.Equal(t, 35304655, event.Repo.Id)
	assert.Equal(t, "lmc-eu/steward", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 0, 0, utc), event.CreatedAt)
	assert.Equal(t, 429615, event.Org.Id)
	assert.Equal(t, "lmc-eu", event.Org.Login)

	event = <-c
	assert.Equal(t, "CreateEvent", event.Type)
	assert.Equal(t, 8417855, event.Actor.Id)
	assert.Equal(t, "dwtechgogators", event.Actor.Login)
	assert.Equal(t, 48892041, event.Repo.Id)
	assert.Equal(t, "dwtechgogators/MSProjectTaskEmail", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 1, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 4620546, event.Actor.Id)
	assert.Equal(t, "webbam46", event.Actor.Login)
	assert.Equal(t, 46007818, event.Repo.Id)
	assert.Equal(t, "webbam46/YBVisual", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 1, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)

	event = <-c
	assert.Equal(t, "CreateEvent", event.Type)
	assert.Equal(t, 11765156, event.Actor.Id)
	assert.Equal(t, "murraykj", event.Actor.Login)
	assert.Equal(t, 42083632, event.Repo.Id)
	assert.Equal(t, "murraykj/OnTheMap", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 2, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 11095732, event.Actor.Id)
	assert.Equal(t, "JackyChiu", event.Actor.Login)
	assert.Equal(t, 48559800, event.Repo.Id)
	assert.Equal(t, "JackyChiu/iOS_Calculator", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 2, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 2941391, event.Actor.Id)
	assert.Equal(t, "midfingr", event.Actor.Login)
	assert.Equal(t, 48193506, event.Repo.Id)
	assert.Equal(t, "midfingr/obx", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 2, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 1759081, event.Actor.Id)
	assert.Equal(t, "christianfredh", event.Actor.Login)
	assert.Equal(t, 30062860, event.Repo.Id)
	assert.Equal(t, "aptitud/aptitud.github.io", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 2, 0, utc), event.CreatedAt)
	assert.Equal(t, 2202349, event.Org.Id)
	assert.Equal(t, "aptitud", event.Org.Login)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 533180, event.Actor.Id)
	assert.Equal(t, "scriptzteam", event.Actor.Login)
	assert.Equal(t, 41249359, event.Repo.Id)
	assert.Equal(t, "scriptzteam/BitTorrent_DHT", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 4, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)

	event = <-c
	assert.Equal(t, "PushEvent", event.Type)
	assert.Equal(t, 6753598, event.Actor.Id)
	assert.Equal(t, "bradleyboehmke", event.Actor.Login)
	assert.Equal(t, 48668209, event.Repo.Id)
	assert.Equal(t, "bradleyboehmke/bradleyboehmke.github.io", event.Repo.Name)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 4, 0, utc), event.CreatedAt)
	assert.Nil(t, event.Org)
}
