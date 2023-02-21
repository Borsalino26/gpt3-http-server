package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	gogpt "github.com/sashabaranov/go-gpt3"
	log "github.com/sirupsen/logrus"
	"gpt3-http-server/biz/constdef"
	"gpt3-http-server/biz/dal"
)

type GPTCompletionArgs struct {
	Model    string `query:"model"`
	MaxToken int    `query:"max_token"`
	Prompt   string `json:"prompt"`
}

// GPTCompletion .
func GPTCompletion(ctx context.Context, c *app.RequestContext) {
	req, err := MakeGPTCompletionReq(ctx, c)
	if err != nil {
		log.Errorf("MakeGPTCompletionReq BindAndValidate err:%v", err)
		return
	}
	log.Infof("GPTCompletion request. maxToken:%d, model:%s, prompt:%s", req.MaxTokens, req.Model, req.Prompt)
	resp, err := dal.GPTClient.CreateCompletion(ctx, req)
	if err != nil || len(resp.Choices) == 0 {
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"text": resp.Choices[0].Text,
	})
}

func MakeGPTCompletionReq(ctx context.Context, c *app.RequestContext) (gogpt.CompletionRequest, error) {
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 30,
	}
	var arg GPTCompletionArgs
	err := c.BindAndValidate(&arg)
	if err != nil {
		return req, err
	}
	if arg.MaxToken > constdef.GPTMaxToken {
		return req, fmt.Errorf("arg.maxToken illegal. arg.maxToken:%d", arg.MaxToken)
	}
	req.MaxTokens = arg.MaxToken
	req.Prompt = arg.Prompt
	req.Model = chooseModel(arg.Model)
	return req, nil
}

func chooseModel(level string) string {
	if level == "high" {
		return gogpt.GPT3TextDavinci003
	}
	if level == "medium" {
		return gogpt.GPT3TextCurie001
	}
	if level == "low" {
		return gogpt.GPT3TextBabbage001
	}
	return gogpt.GPT3TextDavinci003
}
