package logic

import (
	"context"
	"github.com/wujunwei928/edge-tts-go/edge_tts"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TextToSpeechLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTextToSpeechLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TextToSpeechLogic {
	return &TextToSpeechLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TextToSpeechLogic) TextToSpeech(req *types.TextToSpeechRequest) (resp *types.TextToSpeechResponse, err error) {
	connOptions := []edge_tts.CommunicateOption{
		edge_tts.SetVoice(req.Voice),
		//edge_tts.SetRate(rate),
		//edge_tts.SetVolume(volume),
		//edge_tts.SetPitch(pitch),
		//edge_tts.SetReceiveTimeout(20),
	}

	conn, err := edge_tts.NewCommunicate(
		req.Text,
		connOptions...,
	)
	if err != nil {
		return
	}
	audioData, err := conn.Stream()
	if err != nil {
		return
	}

	return &types.TextToSpeechResponse{
		Audio: audioData,
	}, nil
}
