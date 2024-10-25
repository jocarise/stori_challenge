flowchart TD
    A[Start: Call AddNewsletterWithFile_FUNC] --> B[Check if file type is valid]
    B --> C{Is valid?}
    C -- No --> D[Return error 'invalid file type']
    C -- Yes --> E[Save file and get file path]
    
    E --> F{Is saving successful?}
    F -- No --> G[Log error and return 'invalid file type']
    F -- Yes --> H[Validate NewsletterDTO]
    
    H --> I{Is validation successful?}
    I -- No --> J[Return validation error]
    I -- Yes --> K[Find category by ID]
    
    K --> L{Is category found?}
    L -- No --> M[Log error and return 'invalid category' ]
    L -- Yes --> N[Create Newsletter instance]

    N --> O[Save newsletter to repository]
    O --> P{Is saving successful?}
    P -- No --> Q[Return error]
    P -- Yes --> R[Prepare base URL for unsubscribe links]

    R --> S[Retrieve recipient emails from DTO]
    S --> T[Find existing recipients by emails]
    
    T --> U{Are existing recipients found?}
    U -- No --> V[Continue to process new recipients]
    U -- Yes --> W[Create a map of existing recipients]

    W --> X[Iterate over recipient emails]
    X --> Y{Is recipient found?}
    Y -- Yes --> Z[Associate existing recipient with newsletter]
    Y -- No --> AA[Create new recipient]
    
    AA --> AB[Save new recipient to repository]
    AB --> AC[Associate new recipient with newsletter]

    AC --> AD[Return created newsletter]