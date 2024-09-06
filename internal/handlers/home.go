package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/internal/database"
	home "github.com/suchithsridhar/suchicodes/web/views/home"
)

const HUMAN_TEST_PHRASE = "i am human"

func (h *Handler) handleIndexShow(ctx echo.Context) error {
	indexData, err := loadJSONFromFile[home.IndexJSON](home.IndexJsonFile)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	return renderTemplate(ctx, home.IndexShow(indexData, navbarData, footerData))
}

func (h *Handler) handleAboutShow(ctx echo.Context) error {
	aboutData, err := loadJSONFromFile[home.AboutJSON](home.AboutJsonFile)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	return renderTemplate(ctx, home.AboutShow(aboutData, navbarData, footerData))
}

func (h *Handler) handleContactShow(ctx echo.Context) error {
	contactData, err := loadJSONFromFile[home.ContactJSON](home.ContactJsonFile)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	contactProps := &home.ContactProps{
	}

	return renderTemplate(ctx, home.ContactShow(contactProps, contactData, navbarData, footerData))
}

func (h *Handler) handleContactPost(ctx echo.Context) error {
	timestamp := time.Now().Unix()
	id := uuid.New().String()
	subject := ctx.FormValue("subject")
	message := ctx.FormValue("message")
	humanTest := ctx.FormValue("humantest")
	ipAddress := ctx.Get("clientIpAddress").(string)

	contactProps := &home.ContactProps{
	}

	if strings.ToLower(strings.TrimSpace(humanTest)) != HUMAN_TEST_PHRASE {
		return ctx.String(http.StatusForbidden, "You were identified as a bot. Please fill out the human test field.")
	}

	newContactMessage := database.Contact {
		Id: id,
		Subject: subject,
		Message: message,
		IpAddress: ipAddress,
		CreatedAt: timestamp,
	}

	result := h.App.Db.Create(&newContactMessage)
	if result.Error != nil {
		h.App.Logger.Error(fmt.Sprintf("Failed to create contact message at %v.", timestamp))
		contactProps.Notification = true
		contactProps.NotificationTitle = "Failed to send contact message"
		contactProps.NotificationBody = "The system was unable to process your contact message, please try one of the other contact options."
		contactProps.NotificationType = "failure"
	} else {
		contactProps.Notification = true
		contactProps.NotificationTitle = "Successfully sent contact message!"
		contactProps.NotificationBody = "The message was processed by the system. Suchith has been notified!"
		contactProps.NotificationType = "success"
	}

	contactData, err := loadJSONFromFile[home.ContactJSON](home.ContactJsonFile)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	return renderTemplate(ctx, home.ContactShow(contactProps, contactData, navbarData, footerData))
}

func (h *Handler) handleIndexApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.IndexJSON](home.IndexJsonFile)
	return ctx.JSON(status, data)
}

func (h *Handler) handleAboutApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.AboutJSON](home.AboutJsonFile)
	return ctx.JSON(status, data)
}

func (h *Handler) handleContactApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.ContactJSON](home.ContactJsonFile)
	return ctx.JSON(status, data)
}
