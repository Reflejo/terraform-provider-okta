package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/okta/terraform-provider-okta/sdk/query"
)

type BrandResource resource

type Brand struct {
	Links                      interface{} `json:"_links,omitempty"`
	AgreeToCustomPrivacyPolicy *bool       `json:"agreeToCustomPrivacyPolicy,omitempty"`
	CustomPrivacyPolicyUrl     string      `json:"customPrivacyPolicyUrl,omitempty"`
	Id                         string      `json:"id,omitempty"`
	RemovePoweredByOkta        *bool       `json:"removePoweredByOkta,omitempty"`
}

// Fetches a brand by &#x60;brandId&#x60;
func (m *BrandResource) GetBrand(ctx context.Context, brandId string) (*Brand, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v", brandId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var brand *Brand

	resp, err := rq.Do(ctx, req, &brand)
	if err != nil {
		return nil, resp, err
	}

	return brand, resp, nil
}

// Updates a brand by &#x60;brandId&#x60;
func (m *BrandResource) UpdateBrand(ctx context.Context, brandId string, body Brand) (*Brand, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v", brandId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var brand *Brand

	resp, err := rq.Do(ctx, req, &brand)
	if err != nil {
		return nil, resp, err
	}

	return brand, resp, nil
}

// List all the brands in your org.
func (m *BrandResource) ListBrands(ctx context.Context) ([]*Brand, *Response, error) {
	url := "/api/v1/brands"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var brand []*Brand

	resp, err := rq.Do(ctx, req, &brand)
	if err != nil {
		return nil, resp, err
	}

	return brand, resp, nil
}

// List email templates in your organization with pagination.
func (m *BrandResource) ListEmailTemplates(ctx context.Context, brandId string, qp *query.Params) ([]*EmailTemplate, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email", brandId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplate []*EmailTemplate

	resp, err := rq.Do(ctx, req, &emailTemplate)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplate, resp, nil
}

// Fetch an email template by templateName
func (m *BrandResource) GetEmailTemplate(ctx context.Context, brandId, templateName string) (*EmailTemplate, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplate *EmailTemplate

	resp, err := rq.Do(ctx, req, &emailTemplate)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplate, resp, nil
}

// Delete all customizations for an email template. Also known as “Reset to Default”.
func (m *BrandResource) DeleteEmailTemplateCustomizations(ctx context.Context, brandId, templateName string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List all email customizations for an email template
func (m *BrandResource) ListEmailTemplateCustomizations(ctx context.Context, brandId, templateName string) ([]*EmailTemplateCustomization, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateCustomization []*EmailTemplateCustomization

	resp, err := rq.Do(ctx, req, &emailTemplateCustomization)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateCustomization, resp, nil
}

// Create an email customization
func (m *BrandResource) CreateEmailTemplateCustomization(ctx context.Context, brandId, templateName string, body EmailTemplateCustomizationRequest) (*EmailTemplateCustomization, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateCustomization *EmailTemplateCustomization

	resp, err := rq.Do(ctx, req, &emailTemplateCustomization)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateCustomization, resp, nil
}

// Delete an email customization
func (m *BrandResource) DeleteEmailTemplateCustomization(ctx context.Context, brandId, templateName, customizationId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations/%v", brandId, templateName, customizationId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetch an email customization by id.
func (m *BrandResource) GetEmailTemplateCustomization(ctx context.Context, brandId, templateName, customizationId string) (*EmailTemplateCustomization, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations/%v", brandId, templateName, customizationId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateCustomization *EmailTemplateCustomization

	resp, err := rq.Do(ctx, req, &emailTemplateCustomization)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateCustomization, resp, nil
}

// Update an email customization
func (m *BrandResource) UpdateEmailTemplateCustomization(ctx context.Context, brandId, templateName, customizationId string, body EmailTemplateCustomizationRequest) (*EmailTemplateCustomization, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations/%v", brandId, templateName, customizationId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateCustomization *EmailTemplateCustomization

	resp, err := rq.Do(ctx, req, &emailTemplateCustomization)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateCustomization, resp, nil
}

// Get a preview of an email template customization.
func (m *BrandResource) GetEmailTemplateCustomizationPreview(ctx context.Context, brandId, templateName, customizationId string) (*EmailTemplateContent, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/customizations/%v/preview", brandId, templateName, customizationId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateContent *EmailTemplateContent

	resp, err := rq.Do(ctx, req, &emailTemplateContent)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateContent, resp, nil
}

// Fetch the default content for an email template.
func (m *BrandResource) GetEmailTemplateDefaultContent(ctx context.Context, brandId, templateName string) (*EmailTemplateContent, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/default-content", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateContent *EmailTemplateContent

	resp, err := rq.Do(ctx, req, &emailTemplateContent)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateContent, resp, nil
}

// Fetch a preview of an email template&#x27;s default content by populating velocity references with the current user&#x27;s environment.
func (m *BrandResource) GetEmailTemplateDefaultContentPreview(ctx context.Context, brandId, templateName string) (*EmailTemplateContent, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/default-content/preview", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var emailTemplateContent *EmailTemplateContent

	resp, err := rq.Do(ctx, req, &emailTemplateContent)
	if err != nil {
		return nil, resp, err
	}

	return emailTemplateContent, resp, nil
}

// Send a test email to the current users primary and secondary email addresses. The email content is selected based on the following priority: An email customization specifically for the users locale. The default language of email customizations. The email templates default content.
func (m *BrandResource) SendTestEmail(ctx context.Context, brandId, templateName string, body EmailTemplateTestRequest) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/templates/email/%v/test", brandId, templateName)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// List all the themes in your brand
func (m *BrandResource) ListBrandThemes(ctx context.Context, brandId string) ([]*ThemeResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes", brandId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var themeResponse []*ThemeResponse

	resp, err := rq.Do(ctx, req, &themeResponse)
	if err != nil {
		return nil, resp, err
	}

	return themeResponse, resp, nil
}

// Fetches a theme for a brand
func (m *BrandResource) GetBrandTheme(ctx context.Context, brandId, themeId string) (*ThemeResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var themeResponse *ThemeResponse

	resp, err := rq.Do(ctx, req, &themeResponse)
	if err != nil {
		return nil, resp, err
	}

	return themeResponse, resp, nil
}

// Updates a theme for a brand
func (m *BrandResource) UpdateBrandTheme(ctx context.Context, brandId, themeId string, body Theme) (*ThemeResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var themeResponse *ThemeResponse

	resp, err := rq.Do(ctx, req, &themeResponse)
	if err != nil {
		return nil, resp, err
	}

	return themeResponse, resp, nil
}

// Deletes a Theme background image
func (m *BrandResource) DeleteBrandThemeBackgroundImage(ctx context.Context, brandId, themeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/background-image", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates the background image for your Theme
func (m *BrandResource) UploadBrandThemeBackgroundImage(ctx context.Context, brandId, themeId, file string) (*ImageUploadResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/background-image", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType(writer.FormDataContentType()).NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var imageUploadResponse *ImageUploadResponse

	resp, err := rq.Do(ctx, req, &imageUploadResponse)
	if err != nil {
		return nil, resp, err
	}

	return imageUploadResponse, resp, nil
}

// Deletes a Theme favicon. The org then uses the Okta default favicon.
func (m *BrandResource) DeleteBrandThemeFavicon(ctx context.Context, brandId, themeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/favicon", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates the favicon for your theme
func (m *BrandResource) UploadBrandThemeFavicon(ctx context.Context, brandId, themeId, file string) (*ImageUploadResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/favicon", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType(writer.FormDataContentType()).NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var imageUploadResponse *ImageUploadResponse

	resp, err := rq.Do(ctx, req, &imageUploadResponse)
	if err != nil {
		return nil, resp, err
	}

	return imageUploadResponse, resp, nil
}

// Deletes a Theme logo. The org then uses the Okta default logo.
func (m *BrandResource) DeleteBrandThemeLogo(ctx context.Context, brandId, themeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/logo", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates the logo for your Theme
func (m *BrandResource) UploadBrandThemeLogo(ctx context.Context, brandId, themeId, file string) (*ImageUploadResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/logo", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType(writer.FormDataContentType()).NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var imageUploadResponse *ImageUploadResponse

	resp, err := rq.Do(ctx, req, &imageUploadResponse)
	if err != nil {
		return nil, resp, err
	}

	return imageUploadResponse, resp, nil
}
