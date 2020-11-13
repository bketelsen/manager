package main

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

type conferenceService struct {
	tracer opentracing.Tracer
}

func (s conferenceService) List(ctx context.Context, r ListConferenceRequest) (*ListConferenceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "conference.list")
	defer span.Finish()
	resp := &ListConferenceResponse{}
	return resp, nil
}

func (s conferenceService) Get(ctx context.Context, r GetConferenceRequest) (*GetConferenceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "conference.get")
	defer span.Finish()
	resp := &GetConferenceResponse{
		Conference: Conference{
			Name: "Gophercon",
		},
	}
	return resp, nil
}
