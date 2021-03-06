/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type RunningRace struct {
	// The unique identifier of this race.
	Id int64 `json:"id,omitempty"`
	// The name of this race.
	Name string `json:"name,omitempty"`
	// The type of this race.
	RunningRaceType int32 `json:"running_race_type,omitempty"`
	// The race's distance, in meters.
	Distance float32 `json:"distance,omitempty"`
	// The time at which the race begins started in the local timezone.
	StartDateLocal time.Time `json:"start_date_local,omitempty"`
	// The name of the city in which the race is taking place.
	City string `json:"city,omitempty"`
	// The name of the state or geographical region in which the race is taking place.
	State string `json:"state,omitempty"`
	// The name of the country in which the race is taking place.
	Country string `json:"country,omitempty"`
	// The set of routes that cover this race's course.
	RouteIds []int64 `json:"route_ids,omitempty"`
	// The unit system in which the race should be displayed.
	MeasurementPreference string `json:"measurement_preference,omitempty"`
	// The vanity URL of this race on Strava.
	Url string `json:"url,omitempty"`
	// The URL of this race's website.
	WebsiteUrl string `json:"website_url,omitempty"`
}
