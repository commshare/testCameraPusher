package rtspserver

import "../livemedia"

type StreamServerState struct {
	subsession  livemedia.IServerMediaSubsession
	streamToken *livemedia.StreamState
}
