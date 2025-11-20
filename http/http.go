package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/carlmjohnson/requests"
	"io"
	"net/http"
	"reflect"
	"tk-error-handle/model"
)

func checkBaseResp(response any) error {
	respValue := reflect.ValueOf(response)
	if respValue.Kind() != reflect.Ptr {
		return fmt.Errorf("响应体必须是一个指针")
	}

	structValue := respValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf("响应体必须指向一个结构体")
	}

	baseRespField := structValue.FieldByName("BaseResp")
	if !baseRespField.IsValid() {
		// 如果没有 BaseResp 字段，直接返回成功（向后兼容）
		return nil
	}

	if baseRespField.Type() != reflect.TypeOf(model.BaseResp{}) {
		return fmt.Errorf("基础响应不为对应的类型")
	}

	codeField := baseRespField.FieldByName("Code")
	messageField := baseRespField.FieldByName("Message")

	if !codeField.IsValid() || !messageField.IsValid() {
		return fmt.Errorf("基础响应丢失Code或Message参数")
	}

	code := codeField.Int()
	if code != 0 {
		message := messageField.String()
		return fmt.Errorf("响应错误: %d - %s", code, message)
	}

	return nil
}

func Request[T any, U any](
	method string,
	uri string,
	sessionId string,
	ctx context.Context,
	request *T,
	response *U,
) error {
	builder := requests.
		URL("https://api16-normal-sg.tiktokshopglobalselling.com").
		Pathf("/api/full-service/product-center%s", uri).
		Method(method).
		Header("Cookie", fmt.Sprintf("sessionid=%s", sessionId))

	if request != nil {
		builder = builder.BodyJSON(request)
	}

	return builder.
		Handle(func(res *http.Response) error {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return fmt.Errorf("读取响应失败: %w", err)
			}

			if err := json.Unmarshal(body, response); err != nil {
				return fmt.Errorf("解析响应失败: %w", err)
			}

			if err := checkBaseResp(response); err != nil {
				return err
			}

			return nil
		}).
		Fetch(ctx)
}
