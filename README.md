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
	1. Create a new .env file inside the root folder and copy .env.example into .env:
	```ruby
	# user-service
	POSTGRES_USERVICE_USER=admin
	POSTGRES_USERVICE_PASSWORD=root
	POSTGRES_USERVICE_DB=users_db
	# newsletter-service
	POSTGRES_NSERVICE_USER=admin
	POSTGRES_NSERVICE_PASSWORD=root
	POSTGRES_NSERVICE_DB=newsletters_db
	```

 	2.Create a new .env file inside the root/client folder and copy .env.example into .env:
	```ruby
	USER_SERVICE_API=http://localhost:4000
	NEWSLETTER_SERVICE_API=http://localhost:4500
	NEWSLETTER_DOCKER_SERVICE_API=http://newsletter_service:4500
	```

   	3.Create a new .env file inside the root/user-service folder and copy .env.example into .env:
	```ruby
	API_VERSION=/v1
	PORT=:4000
	DATABASE_URL=postgres://admin:root@user_service_db:5432/users_db
	JWT_SECRET=cab2e30b56e8618d6515370efa8acb934f70041ff05b72527ffb2142aefe088a
	```

	4.Create a new .env file inside the root/user-service folder and copy .env.example into .env:
	```ruby
	API_VERSION=/v1
	PORT=:4500
	DATABASE_URL=postgres://admin:root@newsletter_service_db:5432/newsletters_db
	JWT_SECRET=cab2e30b56e8618d6515370efa8acb934f70041ff05b72527ffb2142aefe088a
	FILES_PATH=/app/uploads
	BASE_URL=http://localhost:4500
	GMAIL_EMAIL=GMAIL_USER
	GMAIL_KEY=GMAIL_APP_KEY
	GMAIL_HOST=smtp.gmail.com
	GMAIL_PORT=587
	```
	If you need to test the smpt email please contact with me through email `jorgeocarisaa@gmail.dom` and I can give you the keys that I used to test.
	If you run the project without `GMAIL_APP_KEY` and `GMAIL_USER` the emails won't sended.

	For details on where to obtain these environment variables or their behavior, refer to:
	* ./doc/envs.txt

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
