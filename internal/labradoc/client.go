package labradoc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// Client is the Labradoc API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
	logger     *slog.Logger
}

// NewClient creates a new Labradoc API client.
func NewClient(apiKey, baseURL string, logger *slog.Logger) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: logger,
	}
}

// FileStatus represents the status of a file.
type FileStatus string

const (
	FileStatusActive   FileStatus = "active"
	FileStatusArchived FileStatus = "archived"
)

// File represents a Labradoc file.
type File struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      FileStatus `json:"status"`
	ContentType string    `json:"content_type,omitempty"`
	Size        int64     `json:"size,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// FilesResponse represents the response for listing files.
type FilesResponse struct {
	Items      []File `json:"items"`
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	TotalPages int    `json:"total_pages"`
	TotalItems int    `json:"total_items"`
}

// FilesListParams represents parameters for listing files.
type FilesListParams struct {
	Status     string
	PageSize   int
	PageNumber int
	Query      string
}

// FilesList returns a paginated list of files.
func (c *Client) FilesList(ctx context.Context, params FilesListParams) (*FilesResponse, error) {
	if params.PageSize <= 0 {
		params.PageSize = 50
	}
	if params.PageNumber <= 0 {
		params.PageNumber = 1
	}

	url := fmt.Sprintf("%s/api/v4/files?page_size=%d&page_number=%d", c.baseURL, params.PageSize, params.PageNumber)
	if params.Status != "" {
		url += "&status=" + params.Status
	}
	if params.Query != "" {
		url += "&query=" + params.Query
	}

	var result FilesResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FilesSearch searches for files.
func (c *Client) FilesSearch(ctx context.Context, query string) (*FilesResponse, error) {
	url := fmt.Sprintf("%s/api/v4/files/search?q=%s", c.baseURL, query)
	var result FilesResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FileGet returns a single file by ID.
func (c *Client) FileGet(ctx context.Context, fileID string) (*File, error) {
	url := fmt.Sprintf("%s/api/v4/files/%s", c.baseURL, fileID)
	var result File
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FilesDelete archives files by their IDs.
func (c *Client) FilesDelete(ctx context.Context, ids []string) error {
	body, err := json.Marshal(map[string][]string{"ids": ids})
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("%s/api/v4/files/archive", c.baseURL)
	return c.post(ctx, url, body, nil)
}

// EmailAddress represents an inbound email address.
type EmailAddress struct {
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	ForwardTo   string    `json:"forward_to,omitempty"`
}

// EmailAddressesResponse represents the response for listing email addresses.
type EmailAddressesResponse struct {
	Items []EmailAddress `json:"items"`
}

// EmailAddressesList returns all inbound email addresses.
func (c *Client) EmailAddressesList(ctx context.Context) (*EmailAddressesResponse, error) {
	url := fmt.Sprintf("%s/api/v4/email/addresses", c.baseURL)
	var result EmailAddressesResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EmailAddressCreateParams represents parameters for creating an email address.
type EmailAddressCreateParams struct {
	Description string
}

// EmailAddressCreate requests a new inbound email address.
func (c *Client) EmailAddressCreate(ctx context.Context, params EmailAddressCreateParams) (*EmailAddress, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("%s/api/v4/email/addresses", c.baseURL)
	var result EmailAddress
	if err := c.post(ctx, url, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Email represents an ingested email.
type Email struct {
	ID          string    `json:"id"`
	From        string    `json:"from,omitempty"`
	Subject     string    `json:"subject,omitempty"`
	To          string    `json:"to,omitempty"`
	Body        string    `json:"body,omitempty"`
	ReceivedAt  time.Time `json:"received_at,omitempty"`
	Attachments []File    `json:"attachments,omitempty"`
}

// EmailsResponse represents the response for listing emails.
type EmailsResponse struct {
	Items      []Email `json:"items"`
	PageSize   int     `json:"page_size"`
	PageNumber int     `json:"page_number"`
	TotalPages int     `json:"total_pages"`
	TotalItems int     `json:"total_items"`
}

// EmailsList returns a list of ingested emails.
func (c *Client) EmailsList(ctx context.Context) (*EmailsResponse, error) {
	url := fmt.Sprintf("%s/api/v4/emails", c.baseURL)
	var result EmailsResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Task represents a task extracted from documents.
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	DueDate     string    `json:"due_date,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

// TasksResponse represents the response for listing tasks.
type TasksResponse struct {
	Items []Task `json:"items"`
}

// TasksList returns all tasks extracted from documents.
func (c *Client) TasksList(ctx context.Context) (*TasksResponse, error) {
	url := fmt.Sprintf("%s/api/v4/tasks", c.baseURL)
	var result TasksResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TasksClose closes/completes tasks by their IDs.
func (c *Client) TasksClose(ctx context.Context, ids []string) error {
	body, err := json.Marshal(map[string][]string{"ids": ids})
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("%s/api/v4/tasks/close", c.baseURL)
	return c.post(ctx, url, body, nil)
}

// UserStats represents user statistics.
type UserStats struct {
	CompletedPages  int   `json:"completed_pages"`
	UnlimitedPages  bool  `json:"unlimited_pages"`
	StorageUsed     int64 `json:"storage_used,omitempty"`
	StorageQuota    int64 `json:"storage_quota,omitempty"`
}

// UserStats returns user statistics.
func (c *Client) UserStats(ctx context.Context) (*UserStats, error) {
	url := fmt.Sprintf("%s/api/v4/users/me/stats", c.baseURL)
	var result UserStats
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BillingCheckoutResponse represents the Stripe checkout response.
type BillingCheckoutResponse struct {
	URL string `json:"url"`
}

// BillingCheckout creates a Stripe checkout session for AI credits.
func (c *Client) BillingCheckout(ctx context.Context) (*BillingCheckoutResponse, error) {
	url := fmt.Sprintf("%s/api/v4/billing/checkout", c.baseURL)
	var result BillingCheckoutResponse
	if err := c.post(ctx, url, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IntegrationStatus represents the status of an integration.
type IntegrationStatus struct {
	Connected bool   `json:"connected"`
	Email     string `json:"email,omitempty"`
}

// GoogleDriveStatus returns the status of Google Drive integration.
func (c *Client) GoogleDriveStatus(ctx context.Context) (*IntegrationStatus, error) {
	url := fmt.Sprintf("%s/api/v4/integrations/google/drive/status", c.baseURL)
	var result IntegrationStatus
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GoogleDriveConnect starts the Google Drive OAuth flow.
func (c *Client) GoogleDriveConnect(ctx context.Context) (*BillingCheckoutResponse, error) {
	url := fmt.Sprintf("%s/api/v4/integrations/google/drive/connect", c.baseURL)
	var result BillingCheckoutResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GoogleGmailStatus returns the status of Gmail integration.
func (c *Client) GoogleGmailStatus(ctx context.Context) (*IntegrationStatus, error) {
	url := fmt.Sprintf("%s/api/v4/integrations/google/gmail/status", c.baseURL)
	var result IntegrationStatus
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GoogleGmailConnect starts the Gmail OAuth flow.
func (c *Client) GoogleGmailConnect(ctx context.Context) (*BillingCheckoutResponse, error) {
	url := fmt.Sprintf("%s/api/v4/integrations/google/gmail/connect", c.baseURL)
	var result BillingCheckoutResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// MicrosoftOutlookStatus returns the status of Microsoft Outlook integration.
func (c *Client) MicrosoftOutlookStatus(ctx context.Context) (*IntegrationStatus, error) {
	url := fmt.Sprintf("%s/api/v4/integrations/microsoft/outlook/status", c.baseURL)
	var result IntegrationStatus
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// MicrosoftOutlookConnect starts the Microsoft Outlook OAuth flow.
func (c *Client) MicrosoftOutlookConnect(ctx context.Context) (*BillingCheckoutResponse, error) {
	url := fmt.Sprintf("%s/api/v4/integrations/microsoft/outlook/connect", c.baseURL)
	var result BillingCheckoutResponse
	if err := c.get(ctx, url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// APIError represents an API error response.
type APIError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("API error: code=%d", e.Code)
}

// get performs a GET request and unmarshals the response into v.
func (c *Client) get(ctx context.Context, url string, v interface{}) error {
	return c.doRequest(ctx, http.MethodGet, url, nil, v)
}

// post performs a POST request and unmarshals the response into v.
func (c *Client) post(ctx context.Context, url string, body []byte, v interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}
	return c.doRequest(ctx, http.MethodPost, url, bodyReader, v)
}

// doRequest performs an HTTP request and unmarshals the response into v.
func (c *Client) doRequest(ctx context.Context, method string, url string, body io.Reader, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	c.logger.Debug("Labradoc API request",
		slog.String("method", method),
		slog.String("url", url),
	)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	c.logger.Debug("Labradoc API response",
		slog.Int("status", resp.StatusCode),
		slog.String("body", string(respBody)),
	)

	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.Unmarshal(respBody, &apiErr); err == nil && apiErr.Message != "" {
			return fmt.Errorf("API error (%d): %s", resp.StatusCode, apiErr.Message)
		}
		return fmt.Errorf("API error: status=%d body=%s", resp.StatusCode, string(respBody))
	}

	// If v is nil, we're done (for operations with no response body)
	if v == nil {
		return nil
	}

	// Unmarshal the response
	if err := json.Unmarshal(respBody, v); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}
