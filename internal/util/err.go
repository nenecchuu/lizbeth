package util

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	errlib "github.com/nenecchuu/arcana/err"
	"github.com/nenecchuu/arcana/response"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	"github.com/rs/zerolog/log"
)

func NewRequiredFieldErr(requiredField string) errlib.Error {
	err := fmt.Errorf(constants.ErrRequiredField, requiredField)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewDataNotFoundErr(entityName string) errlib.Error {
	err := fmt.Errorf(constants.ErrDataNotFound, entityName)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewInvalidFieldFormatErr(invalidField string) errlib.Error {
	err := fmt.Errorf(constants.ErrRequiredField, invalidField)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewInvalidEnumErr(value interface{}, enumName string) errlib.Error {
	err := fmt.Errorf(constants.ErrInvalidEnum, value, enumName)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewBadRequestErr(msg string) errlib.Error {
	err := fmt.Errorf(msg)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func ReturnErrorToFiberResponse(fc *fiber.Ctx, err error) error {
	e := errlib.GetError(err)
	res := response.NewJSONResponse().WithErrorString(e.Error()).WithStatusCode(e.HttpStatusCode())
	return fc.JSON(res)
}
