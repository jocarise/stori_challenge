sequenceDiagram
    Client->>+NewsletterService: POST /newsletter
    NewsletterService->AWS_S3: attachment.ext
    AWS_S3-->>NewsletterService: attachmentUrl
    NewsletterService-->>Client: newsletterId
    Client->>+NewsletterService: POST /newsletter/{newsletterId}/send
    NewsletterService->>CommunicationsService: Emails Batch
    CommunicationsService->>AWS_SES: email
    AWS_SES->>Recipient: email