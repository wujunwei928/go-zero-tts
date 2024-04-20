package logic

import (
	"context"
	"errors"
	"github.com/wujunwei928/edge-tts-go/edge_tts"

	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVoicesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVoicesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVoicesLogic {
	return &ListVoicesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVoicesLogic) ListVoices(req *types.ListVoicesRequest) (resp *types.ListVoicesResponse, err error) {
	switch req.Platform {
	case "edge-tts":
		voiceList, voiceListErr := edge_tts.ListVoices("")
		if voiceListErr != nil {
			err = voiceListErr
			return
		}

		formatVoiceList := make([]*types.Voice, 0, len(voiceList))
		for _, voice := range voiceList {
			formatVoiceList = append(formatVoiceList, &types.Voice{
				Name:               voice.Name,
				ShortName:          voice.ShortName,
				Gender:             voice.Gender,
				Locale:             voice.Locale,
				FriendlyName:       voice.FriendlyName,
				ContentCategories:  voice.VoiceTag.ContentCategories,
				VoicePersonalities: voice.VoiceTag.VoicePersonalities,
			})
		}
		logx.Infof("voice list: %v", formatVoiceList)

		return &types.ListVoicesResponse{
			Voices: formatVoiceList,
		}, nil
	default:
		err = errors.New("unsupported platform")
	}
	return
}
