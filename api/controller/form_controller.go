package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/handiism/kidum/domain"
	"github.com/handiism/kidum/internal/response"
	"github.com/handiism/kidum/usecase"
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
)

type FormController struct {
	formUsecase *usecase.FormUsecase
}

func (f *FormController) Insert() fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(err.Error()))
		}

		// agreement
		length := len(form.Value["agreement"])
		if length < 1 {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "agreement" is required`))
		}

		agreement, err := strconv.ParseBool(form.Value["agreement"][0])
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "agreement" must be either string "true" or "false"`))
		}

		if !agreement {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "agreement" must be a true boolean`))
		}

		// type
		length = len(form.Value["type"])
		if length < 1 {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "type" is required`))
		}

		formType, err := domain.NewFormType(form.Value["type"][0])
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(`Field "type" must be either "INDIVIDUAL" or "GROUP"`)
		}

		// destinationRouteId
		length = len(form.Value["destinationRouteId"])
		if length < 1 {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "destinationRouteId" is required`))
		}

		destinationRouteId, err := strconv.ParseInt(form.Value["destinationRouteId"][0], 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "destinationRouteId" must be a number`))
		}

		// email
		length = len(form.Value["email"])
		if length < 1 {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "email" is required`))
		}

		email := form.Value["email"][0]

		if e := validation.Validate(&email, validation.Required, is.Email); e != nil {
			err, ok := e.(validation.Error)
			if !ok {
				return c.Status(fiber.StatusInternalServerError).
					JSON(response.Error(e.Error()))
			}

			message := `Field "email" is invalid`
			switch err.Code() {
			case "validation_required":
				message = `Field "email" must not empty`
			case "validation_is_email":
				message = `Field "email" must be a valid email`
			}

			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(message))
		}

		// phoneNumber
		length = len(form.Value["phoneNumber"])
		if length < 1 {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "phoneNumber" is required`))
		}

		phoneNumber := form.Value["phoneNumber"][0]
		if e := validation.Validate(&phoneNumber, validation.Match(regexp.MustCompile(`^62\d{7,14}$`))); e != nil {
			numberLength := len(phoneNumber)

			if numberLength == 0 {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "phoneNumber" is required`))
			}

			validPrefix := strings.HasPrefix(phoneNumber, "62")
			if !validPrefix {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "phoneNumber" must starts with "62"`))
			}

			if numberLength < 9 {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "phoneNumber" must be minimum 9 characters length`))
			}

			if numberLength > 16 {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "phoneNumber" must be maximum 16 characters length`))
			}

			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "phoneNumber" must be a number`))
		}

		// emergencyPhoneNumber
		length = len(form.Value["emergencyPhoneNumber"])
		if length < 1 {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "emergencyPhoneNumber" is required`))
		}

		emergencyPhoneNumber := form.Value["emergencyPhoneNumber"][0]
		if e := validation.Validate(&emergencyPhoneNumber, validation.Match(regexp.MustCompile(`^62\d{7,14}$`))); e != nil {
			numberLength := len(emergencyPhoneNumber)

			if numberLength == 0 {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "emergencyPhoneNumber" is required`))
			}

			validPrefix := strings.HasPrefix(emergencyPhoneNumber, "62")
			if !validPrefix {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "emergencyPhoneNumber" must starts with "62"`))
			}

			if numberLength < 9 {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "emergencyPhoneNumber" must be minimum 9 characters length`))
			}

			if numberLength > 16 {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(`Field "emergencyPhoneNumber" must be maximum 16 characters length`))
			}

			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "emergencyPhoneNumber" must be a number`))
		}

		usersLength := len(form.Value["userNames"])

		users := make([]domain.UserRequest, usersLength)

		usersValue := map[string][]string{
			"userNames":       form.Value["userNames"],
			"userAges":        form.Value["userAges"],
			"userAddresses":   form.Value["userAddresses"],
			"userGenders":     form.Value["userGenders"],
			"userCredTypes":   form.Value["userCredTypes"],
			"userCredNumbers": form.Value["userCredNumbers"],
		}

		userCredImagesLength := len(form.File["userCredImages"])

		if userCredImagesLength != usersLength {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.Fail(`Field "userCredImages" have a different array length than other "user*" array`))
		}

		for key, value := range usersValue {
			if usersLength != len(value) {
				return c.Status(fiber.StatusBadRequest).
					JSON(response.Fail(fmt.Sprintf(`Field "%s" have a different array length than other "user*" array`, key)))
			}

			for i := 0; i < len(value); i++ {
				if key == "userNames" {
					userName := value[i]
					if len(userName) == 0 {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must not empty`, key, i)))
					}

					users[i].Name = userName
				}

				if key == "userAges" {
					if len(value[i]) == 0 {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must not empty`, key, i)))
					}

					age, err := strconv.ParseInt(value[i], 10, 64)
					if err != nil {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "userAges[%d]" must be a number`, i)))
					}

					users[i].Age = age
				}

				if key == "userAddresses" {
					if len(value[i]) == 0 {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must not empty`, key, i)))
					}

					users[i].Address = value[i]
				}

				if key == "userGenders" {
					if len(value[i]) == 0 {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must not empty`, key, i)))
					}

					gender, err := domain.NewUserGender(value[i])
					if err != nil {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must be either "MALE" or "FEMALE"`, key, i)))
					}
					users[i].Gender = gender
				}

				if key == "userCredTypes" {
					if len(value[i]) == 0 {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must not empty`, key, i)))
					}

					credType, err := domain.NewCredentialsType(value[i])
					if err != nil {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must be either "NATIONAL_ID" or "DRIVER_LICENSE"`, key, i)))
					}
					users[i].Credential.Type = credType
				}

				if key == "userCredNumbers" {
					if len(value[i]) == 0 {
						return c.Status(fiber.StatusBadRequest).
							JSON(response.Fail(fmt.Sprintf(`Field "%s[%d]" must not empty`, key, i)))
					}

					users[i].Credential.Number = value[i]
				}
			}
		}

		for i := 0; i < userCredImagesLength; i++ {
			users[i].Credential.Image = form.File["userCredImages"][i]
		}

		_ = domain.FormRequest{
			Agreement:          agreement,
			Type:               formType,
			Users:              users,
			DestinationRouteId: destinationRouteId,
			Contact: domain.ContactRequest{
				PhoneNumber:          phoneNumber,
				EmergencyPhoneNumber: emergencyPhoneNumber,
				Email:                email,
			},
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}
