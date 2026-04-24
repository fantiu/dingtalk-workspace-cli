// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const (
	// DefaultPageLimit is the maximum number of pages fetched with --page-all.
	DefaultPageLimit = 10

	// DefaultPageDelay is the delay between paginated requests in milliseconds.
	DefaultPageDelay = 200
)

// PaginationOptions controls automatic pagination behavior.
type PaginationOptions struct {
	PageLimit int // Maximum pages (0 = unlimited)
	PageDelay int // Delay between pages in milliseconds
}

// PaginateAll fetches all pages of a paginated API and merges the results.
// DingTalk APIs use two pagination patterns:
//   - cursor/next_cursor/has_more (in response body)
//   - next_token (in response body)
//
// The function auto-detects which pattern the API uses.
func (c *APIClient) PaginateAll(ctx context.Context, req RawAPIRequest, opts PaginationOptions) ([]any, error) {
	if opts.PageLimit == 0 {
		opts.PageLimit = DefaultPageLimit
	}
	if opts.PageDelay <= 0 {
		opts.PageDelay = DefaultPageDelay
	}

	var allResults []any
	pageCount := 0

	for {
		pageCount++
		if opts.PageLimit > 0 && pageCount > opts.PageLimit {
			break
		}

		resp, err := c.Do(ctx, req)
		if err != nil {
			if pageCount == 1 {
				return nil, err
			}
			// Non-first page error: return what we have so far.
			return allResults, fmt.Errorf("分页第 %d 页请求失败 (已获取 %d 页结果): %w", pageCount, pageCount-1, err)
		}

		result, hasMore, nextToken, parseErr := parsePaginatedResponse(resp)
		if parseErr != nil {
			if pageCount == 1 {
				return nil, parseErr
			}
			return allResults, nil
		}

		allResults = append(allResults, result)

		if !hasMore || nextToken == "" {
			break
		}

		// Inject the next page token into the request.
		req = injectPageToken(req, nextToken)

		// Delay between pages to prevent API throttling.
		select {
		case <-ctx.Done():
			return allResults, ctx.Err()
		case <-time.After(time.Duration(opts.PageDelay) * time.Millisecond):
		}
	}

	return allResults, nil
}

// parsePaginatedResponse extracts the response payload and pagination info.
// It auto-detects DingTalk's two pagination patterns.
func parsePaginatedResponse(resp *RawAPIResponse) (result any, hasMore bool, nextToken string, err error) {
	contentType := resp.Header.Get("Content-Type")
	if !isJSONContentType(contentType) {
		return nil, false, "", fmt.Errorf("分页响应非 JSON 格式 (Content-Type: %s)", contentType)
	}

	if len(resp.Body) == 0 {
		return nil, false, "", fmt.Errorf("分页响应体为空 (HTTP %d)", resp.StatusCode)
	}

	var payload map[string]any
	if unmarshalErr := jsonUnmarshal(resp.Body, &payload); unmarshalErr != nil {
		return nil, false, "", fmt.Errorf("解析分页 JSON 响应失败: %w", unmarshalErr)
	}

	// Check for DingTalk errors first.
	if apiErr := checkDingTalkError(payload, resp.StatusCode); apiErr != nil {
		return nil, false, "", apiErr
	}

	// Pattern 1: cursor/next_cursor/has_more (often nested in "result" or top-level)
	if resultObj, ok := payload["result"]; ok {
		if resultMap, isMap := resultObj.(map[string]any); isMap {
			hasMore, _ = resultMap["has_more"].(bool)
			if nc, ok := resultMap["next_cursor"].(float64); ok && nc > 0 {
				nextToken = fmt.Sprintf("%.0f", nc)
			}
			return payload, hasMore, nextToken, nil
		}
	}

	// Top-level has_more / next_cursor
	if hm, ok := payload["has_more"]; ok {
		hasMore, _ = hm.(bool)
	}
	if nc, ok := payload["next_cursor"]; ok {
		if ncf, isFloat := nc.(float64); isFloat && ncf > 0 {
			nextToken = fmt.Sprintf("%.0f", ncf)
		}
	}

	// Pattern 2: next_token
	if nt, ok := payload["next_token"]; ok {
		if nts, isStr := nt.(string); isStr && nts != "" {
			nextToken = nts
			hasMore = true
		}
	}

	return payload, hasMore, nextToken, nil
}

// injectPageToken injects the pagination token into the next request.
// For GET requests, it's added as a query param; for POST, it's in the body.
func injectPageToken(req RawAPIRequest, token string) RawAPIRequest {
	method := req.Method
	if method == "GET" {
		if req.Params == nil {
			req.Params = make(map[string]any)
		}
		// Try to detect which param name the API uses
		if _, ok := req.Params["cursor"]; ok {
			req.Params["cursor"] = token
		} else if _, ok := req.Params["next_token"]; ok {
			req.Params["next_token"] = token
		} else {
			// Default to next_token for GET requests
			req.Params["next_token"] = token
		}
	} else {
		// For POST/PUT requests, inject into the body
		if bodyMap, ok := req.Data.(map[string]any); ok {
			if _, hasCursor := bodyMap["cursor"]; hasCursor {
				bodyMap["cursor"] = token
			} else {
				bodyMap["next_token"] = token
			}
			req.Data = bodyMap
		}
	}
	return req
}

// jsonUnmarshal is a helper for JSON unmarshaling.
func jsonUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
