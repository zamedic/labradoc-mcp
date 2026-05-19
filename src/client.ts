// Labradoc API client

export interface File {
  id: string;
  user_id: string;
  file_id: string;
  source: string;
  file_kind: string;
  file_name: string;
  status: string;
  original_mime_type: string;
  screenshots?: unknown;
  page_count: number;
  error: string;
  recommend_file_name: string;
  document_type: string;
  created: string;
  original_received: string;
  hash: string;
  hash_link?: string;
  archived: boolean;
}

export interface FilesResponse {
  files: File[];
  total: number;
}

export interface FilesListParams {
  status?: string;
  page_size?: number;
  page_number?: number;
  query?: string;
}

export interface EmailAddress {
  id: string;
  address: string;
  description?: string;
  created_at?: string;
  forward_to?: string;
}

export interface EmailAddressesResponse {
  email_addresses: EmailAddress[];
}

export interface EmailAddressCreateParams {
  description?: string;
}

export interface Email {
  id: string;
  from?: string;
  subject?: string;
  to?: string;
  body?: string;
  date?: string;
  files?: string[];
}

export interface EmailsResponse {
  emails: Email[];
}

export interface Task {
  id: string;
  description?: string;
  status?: string;
  dueDate?: string;
  userID?: string;
  fileID?: string;
  type?: string;
  createdAt?: string;
}

export interface TasksResponse {
  tasks: Task[];
}

export interface UserStats {
  completed_pages: number;
  unlimited_pages: boolean;
  storage_used?: number;
  storage_quota?: number;
}

export interface BillingCheckoutResponse {
  url: string;
}

export interface IntegrationStatus {
  connected: boolean;
  email?: string;
}

export class LabradocClient {
  private readonly baseURL: string;
  private readonly apiKey: string;
  private readonly fetch: typeof globalThis.fetch;
  private readonly log: (level: string, msg: string, meta?: Record<string, unknown>) => void;

  constructor(
    apiKey: string,
    baseURL: string = "https://labradoc.eu",
    log: (level: string, msg: string, meta?: Record<string, unknown>) => void = () => {},
  ) {
    this.apiKey = apiKey;
    this.baseURL = baseURL;
    this.fetch = globalThis.fetch;
    this.log = log;
  }

  private async request<T>(
    method: string,
    url: string,
    body?: unknown,
  ): Promise<T> {
    this.log("debug", `Labradoc API request`, { method, url });

    const resp = await this.fetch(url, {
      method,
      headers: {
        Authorization: `Bearer ${this.apiKey}`,
        "X-API-Key": this.apiKey,
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: body != null ? JSON.stringify(body) : undefined,
    });

    const text = await resp.text();
    this.log("debug", `Labradoc API response`, { status: resp.status, body: text });

    if (resp.status >= 400) {
      let msg = `API error: status=${resp.status} body=${text}`;
      try {
        const err = JSON.parse(text) as { message?: string };
        if (err.message) msg = `API error (${resp.status}): ${err.message}`;
      } catch {}
      throw new Error(msg);
    }

    if (!text) return {} as T;
    return JSON.parse(text) as T;
  }

  // Files
  async filesList(params: FilesListParams = {}): Promise<FilesResponse> {
    const ps = params.page_size ?? 50;
    const pn = params.page_number ?? 1;
    let url = `${this.baseURL}/api/user/files?limit=${ps}&page=${pn}`;
    if (params.status) url += `&status=${params.status}`;
    if (params.query) url += `&query=${encodeURIComponent(params.query)}`;
    return this.request("GET", url);
  }

  async filesSearch(query: string): Promise<FilesResponse> {
    const url = `${this.baseURL}/api/user/files?query=${encodeURIComponent(query)}`;
    return this.request("GET", url);
  }

  async fileGet(fileId: string): Promise<File> {
    const url = `${this.baseURL}/api/user/files/${encodeURIComponent(fileId)}`;
    return this.request("GET", url);
  }

  async filesArchive(ids: string[]): Promise<void> {
    const url = `${this.baseURL}/api/user/files/archive`;
    await this.request("POST", url, { ids });
  }

  // Email addresses
  async emailAddressesList(): Promise<EmailAddressesResponse> {
    const url = `${this.baseURL}/api/user/email-addresses`;
    return this.request("GET", url);
  }

  async emailAddressCreate(params: EmailAddressCreateParams = {}): Promise<EmailAddress> {
    const url = `${this.baseURL}/api/user/email-addresses`;
    return this.request("POST", url, params);
  }

  // Emails
  async emailsList(): Promise<EmailsResponse> {
    const url = `${this.baseURL}/api/emails`;
    return this.request("GET", url);
  }

  // Tasks
  async tasksList(): Promise<TasksResponse> {
    const url = `${this.baseURL}/api/tasks`;
    return this.request("GET", url);
  }

  async tasksClose(ids: string[]): Promise<void> {
    const url = `${this.baseURL}/api/tasks/close`;
    await this.request("POST", url, { ids });
  }

  // User
  async userStats(): Promise<UserStats> {
    const url = `${this.baseURL}/api/user/stats`;
    return this.request("GET", url);
  }

  async billingCheckout(): Promise<BillingCheckoutResponse> {
    const url = `${this.baseURL}/api/billing/checkout`;
    return this.request("POST", url, {});
  }

  // Integrations
  async googleDriveStatus(): Promise<IntegrationStatus> {
    const url = `${this.baseURL}/api/integrations/google/drive/status`;
    return this.request("GET", url);
  }

  async googleDriveConnect(): Promise<BillingCheckoutResponse> {
    const url = `${this.baseURL}/api/integrations/google/drive/connect`;
    return this.request("GET", url);
  }

  async googleGmailStatus(): Promise<IntegrationStatus> {
    const url = `${this.baseURL}/api/integrations/google/gmail/status`;
    return this.request("GET", url);
  }

  async googleGmailConnect(): Promise<BillingCheckoutResponse> {
    const url = `${this.baseURL}/api/integrations/google/gmail/connect`;
    return this.request("GET", url);
  }

  async microsoftOutlookStatus(): Promise<IntegrationStatus> {
    const url = `${this.baseURL}/api/integrations/microsoft/outlook/status`;
    return this.request("GET", url);
  }

  async microsoftOutlookConnect(): Promise<BillingCheckoutResponse> {
    const url = `${this.baseURL}/api/integrations/microsoft/outlook/connect`;
    return this.request("GET", url);
  }
}

export type Client = LabradocClient;
