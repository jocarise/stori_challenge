sequenceDiagram
    Client->>+UserService: POST /users/register
    UserService-->>Client: User

    Client->>+UserService: POST /users/authenticate
    UserService-->>Client: JWT Token

    Client->>+NewsletterService: POST /newsletter
    NewsletterService-->>Client: newsletterId

    Client->>+NewsletterService: POST /newsletter/{newsletterId}/send
    NewsletterService->>Email Server: 