package logic

import (
	"context"
	"fmt"
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

func (l *TextToSpeechLogic) floatToPercent(f float64) string {
	percent := int((f - 1) * 100)
	prefix := ""
	if percent > 0 {
		prefix = "+"
	}
	strPercent := fmt.Sprintf("%s%d%%", prefix, percent)
	return strPercent
}

func (l *TextToSpeechLogic) TextToSpeech(req *types.TextToSpeechRequest) (resp *types.TextToSpeechResponse, err error) {
	connOptions := []edge_tts.CommunicateOption{
		edge_tts.SetVoice(req.Voice),
	}
	if req.Volume != 1 {
		strVolume := l.floatToPercent(req.Volume)
		connOptions = append(connOptions, edge_tts.SetVolume(strVolume))
	}
	if req.Pitch != 1 {
		strPitch := l.floatToPercent(req.Pitch)
		connOptions = append(connOptions, edge_tts.SetPitch(strPitch))
	}
	if req.Rate != 1 {
		strRate := l.floatToPercent(req.Rate)
		connOptions = append(connOptions, edge_tts.SetRate(strRate))
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
