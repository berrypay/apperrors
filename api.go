/*
 * Project: Application Error Library
 * Filename: /api.go
 * Created Date: Thursday April 13th 2023 13:10:09 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Thursday April 13th 2023 15:46:22 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package apperrors

import "fmt"

const (
	ERR_API_CALL_FAILURE               = "ERR_API_CALL_FAILURE" // general error
	ERR_API_FAILURE_PAYLOAD_MISSING    = "ERR_API_FAILURE_PAYLOAD_MISSING"
	ERR_API_FAILURE_SIGNATURE_MISMATCH = "ERR_API_FAILURE_SIGNATURE_MISMATCH"
)

type APICallError struct {
	ErrorCode      string `json:"errorCode"`
	Message        string `json:"message"`
	StatusCode     int    `json:"statusCode"`
	ResponseCode   string `json:"responseCode"`
	ResponseDetail string `json:"responseDetail"`
	SourceError    error  `json:"sourceError"`
}

func (e APICallError) Error() string {
	return fmt.Sprintf("[%s] %s - StatusCode{%d} ResponseCode{%s} ResponseDetail{%s}", e.ErrorCode, e.Message, e.StatusCode, e.ResponseCode, e.ResponseDetail)
}

func NewAPICallFailedError(message string, statusCode int, responseCode string, responseDetail string, err error) APICallError {
	return APICallError{
		ErrorCode:      ERR_API_CALL_FAILURE,
		Message:        message,
		StatusCode:     statusCode,
		ResponseCode:   responseCode,
		ResponseDetail: responseDetail,
		SourceError:    err,
	}
}

func NewAPICallFailedPayloadMissingError(message string, statusCode int, responseCode string, responseDetail string, err error) APICallError {
	return APICallError{
		ErrorCode:      ERR_API_FAILURE_PAYLOAD_MISSING,
		Message:        message,
		StatusCode:     statusCode,
		ResponseCode:   responseCode,
		ResponseDetail: responseDetail,
		SourceError:    err,
	}
}

func NewAPICallFailedSignatureMismatchError(message string, statusCode int, responseCode string, responseDetail string, err error) APICallError {
	return APICallError{
		ErrorCode:      ERR_API_FAILURE_SIGNATURE_MISMATCH,
		Message:        message,
		StatusCode:     statusCode,
		ResponseCode:   responseCode,
		ResponseDetail: responseDetail,
		SourceError:    err,
	}
}
