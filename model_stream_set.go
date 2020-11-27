/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type StreamSet struct {
	Time *TimeStream `json:"time,omitempty"`
	Distance *DistanceStream `json:"distance,omitempty"`
	Latlng *LatLngStream `json:"latlng,omitempty"`
	Altitude *AltitudeStream `json:"altitude,omitempty"`
	VelocitySmooth *SmoothVelocityStream `json:"velocity_smooth,omitempty"`
	Heartrate *HeartrateStream `json:"heartrate,omitempty"`
	Cadence *CadenceStream `json:"cadence,omitempty"`
	Watts *PowerStream `json:"watts,omitempty"`
	Temp *TemperatureStream `json:"temp,omitempty"`
	Moving *MovingStream `json:"moving,omitempty"`
	GradeSmooth *SmoothGradeStream `json:"grade_smooth,omitempty"`
}
