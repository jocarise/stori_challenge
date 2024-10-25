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