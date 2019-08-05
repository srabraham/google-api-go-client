// Copyright 2019 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package fcm provides access to the Firebase Cloud Messaging API.
//
// For product documentation, see: https://firebase.google.com/docs/cloud-messaging
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/fcm/v1"
//   ...
//   ctx := context.Background()
//   fcmService, err := fcm.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   fcmService, err := fcm.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   fcmService, err := fcm.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package fcm // import "google.golang.org/api/fcm/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	option "google.golang.org/api/option"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "fcm:v1"
const apiName = "fcm"
const apiVersion = "v1"
const basePath = "https://fcm.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Messages = NewProjectsMessagesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Messages *ProjectsMessagesService
}

func NewProjectsMessagesService(s *Service) *ProjectsMessagesService {
	rs := &ProjectsMessagesService{s: s}
	return rs
}

type ProjectsMessagesService struct {
	s *Service
}

// AndroidConfig: Android specific options for messages sent
// through
// [FCM connection server](https://goo.gl/4GLdUl).
type AndroidConfig struct {
	// CollapseKey: An identifier of a group of messages that can be
	// collapsed, so that only
	// the last message gets sent when delivery can be resumed. A maximum of
	// 4
	// different collapse keys is allowed at any given time.
	CollapseKey string `json:"collapseKey,omitempty"`

	// Data: Arbitrary key/value payload. If present, it will
	// override
	// google.firebase.fcm.v1.Message.data.
	Data map[string]string `json:"data,omitempty"`

	// FcmOptions: Options for features provided by the FCM SDK for Android.
	FcmOptions *AndroidFcmOptions `json:"fcmOptions,omitempty"`

	// Notification: Notification to send to android devices.
	Notification *AndroidNotification `json:"notification,omitempty"`

	// Priority: Message priority. Can take "normal" and "high" values.
	// For more information, see [Setting the priority of
	// a
	// message](https://goo.gl/GjONJv).
	//
	// Possible values:
	//   "NORMAL" - Default priority for data messages. Normal priority
	// messages won't open
	// network connections on a sleeping device, and their delivery may
	// be
	// delayed to conserve the battery. For less time-sensitive messages,
	// such
	// as notifications of new email or other data to sync, choose
	// normal
	// delivery priority.
	//   "HIGH" - Default priority for notification messages. FCM attempts
	// to deliver high
	// priority messages immediately, allowing the FCM service to wake
	// a
	// sleeping device when possible and open a network connection to your
	// app
	// server. Apps with instant messaging, chat, or voice call alerts,
	// for
	// example, generally need to open a network connection and make sure
	// FCM
	// delivers the message to the device without delay. Set high priority
	// if
	// the message is time-critical and requires the user's
	// immediate
	// interaction, but beware that setting your messages to high
	// priority
	// contributes more to battery drain compared with normal priority
	// messages.
	Priority string `json:"priority,omitempty"`

	// RestrictedPackageName: Package name of the application where the
	// registration token must match in
	// order to receive the message.
	RestrictedPackageName string `json:"restrictedPackageName,omitempty"`

	// Ttl: How long (in seconds) the message should be kept in FCM storage
	// if the
	// device is offline. The maximum time to live supported is 4 weeks, and
	// the
	// default value is 4 weeks if not set. Set it to 0 if want to send
	// the
	// message immediately.
	// In JSON format, the Duration type is encoded as a string rather than
	// an
	// object, where the string ends in the suffix "s" (indicating seconds)
	// and
	// is preceded by the number of seconds, with nanoseconds expressed
	// as
	// fractional seconds. For example, 3 seconds with 0 nanoseconds should
	// be
	// encoded in JSON format as "3s", while 3 seconds and 1 nanosecond
	// should
	// be expressed in JSON format as "3.000000001s". The ttl will be
	// rounded down
	// to the nearest second.
	Ttl string `json:"ttl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollapseKey") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollapseKey") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidConfig) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidFcmOptions: Options for features provided by the FCM SDK for
// Android.
type AndroidFcmOptions struct {
	// AnalyticsLabel: Label associated with the message's analytics data.
	AnalyticsLabel string `json:"analyticsLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnalyticsLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnalyticsLabel") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AndroidFcmOptions) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidFcmOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AndroidNotification: Notification to send to android devices.
type AndroidNotification struct {
	// Body: The notification's body text. If present, it will
	// override
	// google.firebase.fcm.v1.Notification.body.
	Body string `json:"body,omitempty"`

	// BodyLocArgs: Variable string values to be used in place of the format
	// specifiers in
	// body_loc_key to use to localize the body text to the user's
	// current
	// localization.
	// See [Formatting and Styling](https://goo.gl/MalYE3) for more
	// information.
	BodyLocArgs []string `json:"bodyLocArgs,omitempty"`

	// BodyLocKey: The key to the body string in the app's string resources
	// to use to localize
	// the body text to the user's current localization.
	// See [String Resources](https://goo.gl/NdFZGI) for more information.
	BodyLocKey string `json:"bodyLocKey,omitempty"`

	// ChannelId: The [notification's
	// channel
	// id](https://developer.android.com/guide/topics/ui/notifiers/no
	// tifications#ManageChannels)
	// (new in Android O). The app must create a channel with this channel
	// ID
	// before any notification with this channel ID is received. If you
	// don't send
	// this channel ID in the request, or if the channel ID provided has not
	// yet
	// been created by the app, FCM uses the channel ID specified in the
	// app
	// manifest.
	ChannelId string `json:"channelId,omitempty"`

	// ClickAction: The action associated with a user click on the
	// notification.
	// If specified, an activity with a matching intent filter is launched
	// when
	// a user clicks on the notification.
	ClickAction string `json:"clickAction,omitempty"`

	// Color: The notification's icon color, expressed in #rrggbb format.
	Color string `json:"color,omitempty"`

	// Icon: The notification's icon.
	// Sets the notification icon to myicon for drawable resource myicon.
	// If you don't send this key in the request, FCM displays the launcher
	// icon
	// specified in your app manifest.
	Icon string `json:"icon,omitempty"`

	// Image: Contains the URL of an image that is going to be displayed in
	// a
	// notification. If present, it will
	// override
	// google.firebase.fcm.v1.Notification.image.
	Image string `json:"image,omitempty"`

	// Sound: The sound to play when the device receives the
	// notification.
	// Supports "default" or the filename of a sound resource bundled in the
	// app.
	// Sound files must reside in /res/raw/.
	Sound string `json:"sound,omitempty"`

	// Tag: Identifier used to replace existing notifications in the
	// notification
	// drawer.
	// If not specified, each request creates a new notification.
	// If specified and a notification with the same tag is already being
	// shown,
	// the new notification replaces the existing one in the notification
	// drawer.
	Tag string `json:"tag,omitempty"`

	// Title: The notification's title. If present, it will
	// override
	// google.firebase.fcm.v1.Notification.title.
	Title string `json:"title,omitempty"`

	// TitleLocArgs: Variable string values to be used in place of the
	// format specifiers in
	// title_loc_key to use to localize the title text to the user's
	// current
	// localization.
	// See [Formatting and Styling](https://goo.gl/MalYE3) for more
	// information.
	TitleLocArgs []string `json:"titleLocArgs,omitempty"`

	// TitleLocKey: The key to the title string in the app's string
	// resources to use to
	// localize the title text to the user's current localization.
	// See [String Resources](https://goo.gl/NdFZGI) for more information.
	TitleLocKey string `json:"titleLocKey,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Body") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Body") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidNotification) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidNotification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ApnsConfig: [Apple Push Notification Service](https://goo.gl/MXRTPa)
// specific options.
type ApnsConfig struct {
	// FcmOptions: Options for features provided by the FCM SDK for iOS.
	FcmOptions *ApnsFcmOptions `json:"fcmOptions,omitempty"`

	// Headers: HTTP request headers defined in Apple Push Notification
	// Service. Refer to
	// [APNs request headers](https://goo.gl/C6Yhia) for
	// supported headers, e.g. "apns-priority": "10".
	Headers map[string]string `json:"headers,omitempty"`

	// Payload: APNs payload as a JSON object, including both `aps`
	// dictionary and custom
	// payload. See [Payload Key Reference](https://goo.gl/32Pl5W).
	// If present, it overrides
	// google.firebase.fcm.v1.Notification.title
	// and google.firebase.fcm.v1.Notification.body.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FcmOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FcmOptions") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ApnsConfig) MarshalJSON() ([]byte, error) {
	type NoMethod ApnsConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ApnsFcmOptions: Options for features provided by the FCM SDK for iOS.
type ApnsFcmOptions struct {
	// AnalyticsLabel: Label associated with the message's analytics data.
	AnalyticsLabel string `json:"analyticsLabel,omitempty"`

	// Image: Contains the URL of an image that is going to be displayed in
	// a
	// notification. If present, it will
	// override
	// google.firebase.fcm.v1.Notification.image.
	Image string `json:"image,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnalyticsLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnalyticsLabel") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ApnsFcmOptions) MarshalJSON() ([]byte, error) {
	type NoMethod ApnsFcmOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FcmOptions: Platform independent options for features provided by the
// FCM SDKs.
type FcmOptions struct {
	// AnalyticsLabel: Label associated with the message's analytics data.
	AnalyticsLabel string `json:"analyticsLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AnalyticsLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AnalyticsLabel") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FcmOptions) MarshalJSON() ([]byte, error) {
	type NoMethod FcmOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Message: Message to send by Firebase Cloud Messaging Service.
type Message struct {
	// Android: Input only. Android specific options for messages sent
	// through
	// [FCM connection server](https://goo.gl/4GLdUl).
	Android *AndroidConfig `json:"android,omitempty"`

	// Apns: Input only. [Apple Push Notification
	// Service](https://goo.gl/MXRTPa)
	// specific options.
	Apns *ApnsConfig `json:"apns,omitempty"`

	// Condition: Condition to send a message to,
	// e.g. "'foo' in topics && 'bar' in topics".
	Condition string `json:"condition,omitempty"`

	// Data: Input only. Arbitrary key/value payload.
	Data map[string]string `json:"data,omitempty"`

	// FcmOptions: Input only. Template for FCM SDK feature options to use
	// across all
	// platforms.
	FcmOptions *FcmOptions `json:"fcmOptions,omitempty"`

	// Name: Output Only. The identifier of the message sent, in the format
	// of
	// `projects/*/messages/{message_id}`.
	Name string `json:"name,omitempty"`

	// Notification: Input only. Basic notification template to use across
	// all platforms.
	Notification *Notification `json:"notification,omitempty"`

	// Token: Registration token to send a message to.
	Token string `json:"token,omitempty"`

	// Topic: Topic name to send a message to, e.g. "weather".
	// Note: "/topics/" prefix should not be provided.
	Topic string `json:"topic,omitempty"`

	// Webpush: Input only. [Webpush
	// protocol](https://tools.ietf.org/html/rfc8030)
	// options.
	Webpush *WebpushConfig `json:"webpush,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Android") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Android") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Message) MarshalJSON() ([]byte, error) {
	type NoMethod Message
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Notification: Basic notification template to use across all
// platforms.
type Notification struct {
	// Body: The notification's body text.
	Body string `json:"body,omitempty"`

	// Image: Contains the URL of an image that is going to be downloaded on
	// the device
	// and displayed in a notification.
	// JPEG, PNG, BMP have full support across platforms. Animated GIF and
	// video
	// only work on iOS. WebP and HEIF have varying levels of support
	// across
	// platforms and platform versions.
	// Android has 1MB image size limit.
	// Quota usage and implications/costs for hosting image on Firebase
	// Storage:
	// https://firebase.google.com/pricing
	Image string `json:"image,omitempty"`

	// Title: The notification's title.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Body") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Body") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Notification) MarshalJSON() ([]byte, error) {
	type NoMethod Notification
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SendMessageRequest: Request to send a message to specified target.
type SendMessageRequest struct {
	// Message: Required. Message to send.
	Message *Message `json:"message,omitempty"`

	// ValidateOnly: Flag for testing the request without actually
	// delivering the message.
	ValidateOnly bool `json:"validateOnly,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Message") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Message") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SendMessageRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SendMessageRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WebpushConfig: [Webpush
// protocol](https://tools.ietf.org/html/rfc8030) options.
type WebpushConfig struct {
	// Data: Arbitrary key/value payload. If present, it will
	// override
	// google.firebase.fcm.v1.Message.data.
	Data map[string]string `json:"data,omitempty"`

	// FcmOptions: Options for features provided by the FCM SDK for Web.
	FcmOptions *WebpushFcmOptions `json:"fcmOptions,omitempty"`

	// Headers: HTTP headers defined in webpush protocol. Refer to
	// [Webpush protocol](https://tools.ietf.org/html/rfc8030#section-5)
	// for
	// supported headers, e.g. "TTL": "15".
	Headers map[string]string `json:"headers,omitempty"`

	// Notification: Web Notification options as a JSON object. Supports
	// Notification instance
	// properties as defined in [Web
	// Notification
	// API](https://developer.mozilla.org/en-US/docs/Web/API/Not
	// ification). If
	// present, "title" and "body" fields
	// override
	// [google.firebase.fcm.v1.Notification.title]
	// and
	// [google.firebase.fcm.v1.Notification.body].
	Notification googleapi.RawMessage `json:"notification,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Data") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Data") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WebpushConfig) MarshalJSON() ([]byte, error) {
	type NoMethod WebpushConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WebpushFcmOptions: Options for features provided by the FCM SDK for
// Web.
type WebpushFcmOptions struct {
	// Link: The link to open when the user clicks on the notification.
	// For all URL values, HTTPS is required.
	Link string `json:"link,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Link") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Link") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WebpushFcmOptions) MarshalJSON() ([]byte, error) {
	type NoMethod WebpushFcmOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "fcm.projects.messages.send":

type ProjectsMessagesSendCall struct {
	s                  *Service
	parentid           string
	sendmessagerequest *SendMessageRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Send: Send a message to specified target (a registration token,
// topic
// or condition).
func (r *ProjectsMessagesService) Send(parentid string, sendmessagerequest *SendMessageRequest) *ProjectsMessagesSendCall {
	c := &ProjectsMessagesSendCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parentid = parentid
	c.sendmessagerequest = sendmessagerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMessagesSendCall) Fields(s ...googleapi.Field) *ProjectsMessagesSendCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMessagesSendCall) Context(ctx context.Context) *ProjectsMessagesSendCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMessagesSendCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMessagesSendCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/1.12.5 gdcl/20190802")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sendmessagerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/messages:send")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parentid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "fcm.projects.messages.send" call.
// Exactly one of *Message or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Message.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsMessagesSendCall) Do(opts ...googleapi.CallOption) (*Message, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Message{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Send a message to specified target (a registration token, topic\nor condition).",
	//   "flatPath": "v1/projects/{projectsId}/messages:send",
	//   "httpMethod": "POST",
	//   "id": "fcm.projects.messages.send",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. It contains the Firebase project id (i.e. the unique identifier\nfor your Firebase project), in the format of `projects/{project_id}`.\nFor legacy support, the numeric project number with no padding is also\nsupported in the format of `projects/{project_number}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/messages:send",
	//   "request": {
	//     "$ref": "SendMessageRequest"
	//   },
	//   "response": {
	//     "$ref": "Message"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
