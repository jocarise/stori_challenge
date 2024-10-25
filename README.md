# Stori Challenge

This Email Newsletter Service allows users to easily send newsletters to their subscribers.

Key Features:
Create Newsletters: Design and customize newsletters with ease.
Send Emails: Quickly send newsletters to your subscribers.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Architecture](#architecture)

## Prerequisites

Before running the project, ensure you have the following installed on your machine:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installation

Before proceeding with the installation, you need to configure the environment variables.

1. Clone the repo:
   ``git clone https://github.com/jocaris/stori_challenge.git``

2.  Set up the environment variables in the following files:
   	- ./challenge/newsletter-service/.env.example
	* ./challenge/user-service/.env.example
	+ ./challenge/client/.env.example
	- ./challenge/.env.example


	For details on where to obtain these environment variables or their functions, refer to:
	* ./challenge/doc/envs.md

4. Run docker services:
 ``docker compose build && docker compose up -d``

5. Go to http://localhost:3000 in you browser

	Be aware of the ports that are being forwarded from the docker containers: 
	- PORT = 3000 // Client
	* PORT = 4000 // User Service
	+ PORT = 4500 // Newsletter Service
	- PORT = 5432 // User Service DB
	* PORT = 5433 // Newsletter Service DB

6. To stop the containers run:
 ``docker compose down``

## Architecture

### Client UI
![image](https://github.com/user-attachments/assets/03841029-0955-47f2-b6ee-dc5bf64ec0cf)

### Initial Approach
```mermaid
   sequenceDiagram
    Client->>+NewsletterService: POST /newsletter
    NewsletterService->AWS_S3: attachment.ext
    AWS_S3-->>NewsletterService: attachmentUrl
    NewsletterService-->>Client: newsletterId
    Client->>+NewsletterService: POST /newsletter/{newsletterId}/send
    NewsletterService->>CommunicationsService: Emails Batch
    CommunicationsService->>AWS_SES: email
    AWS_SES->>Recipient: email
```
### Final Approach
```mermaid
  sequenceDiagram
    Client->>+NewsletterService: POST /newsletter
    NewsletterService->Server: attachment.ext
    Server-->>NewsletterService: attachmentPath
    NewsletterService-->>Client: newsletterId
    Client->>+NewsletterService: POST /newsletter/{newsletterId}/send
    NewsletterService->>GmailSmtp: Email
    GmailSmtp->>Recipient: email
```

### Happy Path
```mermaid
 sequenceDiagram
    Client->>+UserService: POST /users/register
    UserService-->>Client: User

    Client->>+UserService: POST /users/authenticate
    UserService-->>Client: JWT Token

    Client->>+NewsletterService: POST /newsletter
    NewsletterService-->>Client: newsletterId

    Client->>+NewsletterService: POST /newsletter/{newsletterId}/send
    NewsletterService->>Email Server: 
```

### Data Models
```mermaid
 erDiagram
    %% User Model
    USER {
        string ID PK "Primary Key"
        string Email "Unique email"
        string Role "User role"
        string Password "User password"
        time CreatedAt "Created timestamp"
        time UpdatedAt "Updated timestamp"
    }

    %% Newsletter Models
    CATEGORY {
        uint ID PK "Primary Key"
        string Title "Unique title"
        time CreatedAt "Created timestamp"
        time UpdatedAt "Updated timestamp"
    }

    NEWSLETTER {
        string ID PK "Primary Key"
        string Title "Newsletter title"
        string Attachment "Attachment file"
        text Html "HTML content"
        time CreatedAt "Created timestamp"
        time UpdatedAt "Updated timestamp"
        time ScheduledDate "Scheduled date"
        bool Scheduled "Is scheduled"
        uint CategoryID "Foreign key to Category"
    }

    RECIPIENT {
        string ID PK "Primary Key"
        string Email "Email address"
        string UnsuscribeUrl "URL to unsubscribe"
        time CreatedAt "Created timestamp"
        time UpdatedAt "Updated timestamp"
    }

    %% Relationships for Newsletter Service
    CATEGORY ||--o{ NEWSLETTER : has_many
    NEWSLETTER ||--|{ CATEGORY : belongs_to
    NEWSLETTER ||--o{ RECIPIENT : sends_to
    RECIPIENT ||--o{ NEWSLETTER : receives
```

###End
