sequenceDiagram
    Client->>+NewsletterService: POST /newsletter
    NewsletterService->Server: attachment.ext
    Server-->>NewsletterService: attachmentPath
    NewsletterService-->>Client: newsletterId
    Client->>+NewsletterService: POST /newsletter/{newsletterId}/send
    NewsletterService->>GmailSmtp: Email
    GmailSmtp->>Recipient: email